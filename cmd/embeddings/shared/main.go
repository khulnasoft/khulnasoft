package shared

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sourcegraph/log"

	"github.com/khulnasoft/khulnasoft/internal/actor"
	"github.com/khulnasoft/khulnasoft/internal/api"
	"github.com/khulnasoft/khulnasoft/internal/authz"
	srp "github.com/khulnasoft/khulnasoft/internal/authz/subrepoperms"
	"github.com/khulnasoft/khulnasoft/internal/conf"
	"github.com/khulnasoft/khulnasoft/internal/conf/conftypes"
	"github.com/khulnasoft/khulnasoft/internal/database"
	connections "github.com/khulnasoft/khulnasoft/internal/database/connections/live"
	"github.com/khulnasoft/khulnasoft/internal/embeddings"
	"github.com/khulnasoft/khulnasoft/internal/embeddings/background/repo"
	"github.com/khulnasoft/khulnasoft/internal/embeddings/embed"
	"github.com/khulnasoft/khulnasoft/internal/embeddings/embed/client"
	"github.com/khulnasoft/khulnasoft/internal/errcode"
	"github.com/khulnasoft/khulnasoft/internal/featureflag"
	"github.com/khulnasoft/khulnasoft/internal/goroutine"
	"github.com/khulnasoft/khulnasoft/internal/honey"
	"github.com/khulnasoft/khulnasoft/internal/httpserver"
	"github.com/khulnasoft/khulnasoft/internal/instrumentation"
	"github.com/khulnasoft/khulnasoft/internal/observation"
	"github.com/khulnasoft/khulnasoft/internal/service"
	"github.com/khulnasoft/khulnasoft/internal/tenant"
	"github.com/khulnasoft/khulnasoft/internal/trace"
	"github.com/khulnasoft/khulnasoft/lib/errors"
)

const addr = ":9991"

func Main(ctx context.Context, observationCtx *observation.Context, ready service.ReadyFunc, config *Config) error {
	logger := observationCtx.Logger

	// Initialize tracing/metrics
	observationCtx = observation.NewContext(logger, observation.Honeycomb(&honey.Dataset{
		Name:       "embeddings",
		SampleRate: 20,
	}))

	// Initialize main DB connection.
	sqlDB := mustInitializeFrontendDB(observationCtx)
	db := database.NewDB(logger, sqlDB)

	repoStore := db.Repos()
	repoEmbeddingJobsStore := repo.NewRepoEmbeddingJobsStore(db)

	// Run setup
	uploadStore, err := embeddings.NewObjectStorage(ctx, observationCtx, config.EmbeddingsUploadStoreConfig)
	if err != nil {
		return err
	}

	authz.DefaultSubRepoPermsChecker = srp.NewSubRepoPermsClient(db.SubRepoPerms())

	indexGetter, err := NewCachedEmbeddingIndexGetter(
		repoStore,
		repoEmbeddingJobsStore,
		func(ctx context.Context, repoID api.RepoID, repoName api.RepoName) (*embeddings.RepoEmbeddingIndex, error) {
			return embeddings.DownloadRepoEmbeddingIndex(ctx, uploadStore, repoID, repoName)
		},
		config.EmbeddingsCacheSize,
	)
	if err != nil {
		return err
	}

	// Create HTTP server
	handler := NewHandler(logger, indexGetter.Get, getQueryEmbedding)
	handler = handlePanic(logger, handler)
	handler = featureflag.Middleware(db.FeatureFlags(), handler)
	handler = trace.HTTPMiddleware(logger, handler)
	handler = instrumentation.HTTPMiddleware("", handler)
	handler = actor.HTTPMiddleware(logger, handler)
	handler = tenant.InternalHTTPMiddleware(logger, handler)
	server := httpserver.NewFromAddr(addr, &http.Server{
		ReadTimeout:  75 * time.Second,
		WriteTimeout: 10 * time.Minute,
		Handler:      handler,
	})

	// Mark health server as ready and go!
	ready()

	return goroutine.MonitorBackgroundRoutines(ctx, server)
}

func NewHandler(
	logger log.Logger,
	getRepoEmbeddingIndex getRepoEmbeddingIndexFn,
	getQueryEmbedding getQueryEmbeddingFn,
) http.Handler {
	// Initialize the legacy JSON API server
	mux := http.NewServeMux()
	mux.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, fmt.Sprintf("unsupported method %s", r.Method), http.StatusBadRequest)
			return
		}

		var args embeddings.EmbeddingsSearchParameters
		err := json.NewDecoder(r.Body).Decode(&args)
		if err != nil {
			http.Error(w, "could not parse request body: "+err.Error(), http.StatusBadRequest)
			return
		}

		res, err := searchRepoEmbeddingIndexes(r.Context(), args, getRepoEmbeddingIndex, getQueryEmbedding)
		if errcode.IsNotFound(err) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err != nil {
			logger.Error("error searching embedding index", log.Error(err))
			if errors.Is(err, &client.RateLimitExceededError{}) {
				http.Error(w, fmt.Sprintf("error searching embedding index: %s", err.Error()), http.StatusTooManyRequests)
				return
			}
			http.Error(w, fmt.Sprintf("error searching embedding index: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	})

	return mux
}

func getQueryEmbedding(ctx context.Context, query string) ([]float32, string, error) {
	c := conf.GetEmbeddingsConfig(conf.Get().SiteConfig())
	if c == nil {
		return nil, "", errors.New("embeddings not configured or disabled")
	}
	client, err := embed.NewEmbeddingsClient(c)
	if err != nil {
		return nil, "", errors.Wrap(err, "getting embeddings client")
	}

	embeddings, err := client.GetQueryEmbedding(ctx, query)
	if err != nil {
		return nil, "", errors.Wrap(err, "getting query embedding")
	}
	if len(embeddings.Failed) > 0 {
		return nil, "", errors.Newf("failed to get embeddings for query %s", query)
	}

	return embeddings.Embeddings, client.GetModelIdentifier(), nil
}

func mustInitializeFrontendDB(observationCtx *observation.Context) *sql.DB {
	dsn := conf.GetServiceConnectionValueAndRestartOnChange(func(serviceConnections conftypes.ServiceConnections) string {
		return serviceConnections.PostgresDSN
	})

	db, err := connections.EnsureNewFrontendDB(observationCtx, dsn, "embeddings")
	if err != nil {
		observationCtx.Logger.Fatal("failed to connect to database", log.Error(err))
	}

	return db
}

func handlePanic(logger log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				err := fmt.Sprintf("%v", rec)
				http.Error(w, fmt.Sprintf("%v", rec), http.StatusInternalServerError)
				logger.Error("recovered from panic", log.String("err", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
