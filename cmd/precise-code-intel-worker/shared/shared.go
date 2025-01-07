package shared

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/sourcegraph/log"

	"github.com/khulnasoft/khulnasoft/internal/authz"
	srp "github.com/khulnasoft/khulnasoft/internal/authz/subrepoperms"
	"github.com/khulnasoft/khulnasoft/internal/codeintel"
	codeintelshared "github.com/khulnasoft/khulnasoft/internal/codeintel/shared"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/shared/lsifuploadstore"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/uploads"
	"github.com/khulnasoft/khulnasoft/internal/conf"
	"github.com/khulnasoft/khulnasoft/internal/conf/conftypes"
	"github.com/khulnasoft/khulnasoft/internal/database"
	connections "github.com/khulnasoft/khulnasoft/internal/database/connections/live"
	"github.com/khulnasoft/khulnasoft/internal/encryption/keyring"
	"github.com/khulnasoft/khulnasoft/internal/goroutine"
	"github.com/khulnasoft/khulnasoft/internal/honey"
	"github.com/khulnasoft/khulnasoft/internal/httpserver"
	"github.com/khulnasoft/khulnasoft/internal/object"
	"github.com/khulnasoft/khulnasoft/internal/observation"
	"github.com/khulnasoft/khulnasoft/internal/service"
	"github.com/khulnasoft/khulnasoft/lib/errors"
)

const addr = ":3188"

func Main(ctx context.Context, observationCtx *observation.Context, ready service.ReadyFunc, config Config) error {
	logger := observationCtx.Logger

	// Initialize tracing/metrics
	observationCtx = observation.NewContext(logger, observation.Honeycomb(&honey.Dataset{
		Name: "codeintel-worker",
	}))

	if err := keyring.Init(ctx); err != nil {
		return errors.Wrap(err, "initializing keyring")
	}

	// Connect to databases
	db := database.NewDB(logger, mustInitializeDB(observationCtx))
	codeIntelDB := mustInitializeCodeIntelDB(observationCtx)

	// Migrations may take a while, but after they're done we'll immediately
	// spin up a server and can accept traffic. Inform external clients we'll
	// be ready for traffic.
	ready()

	// Initialize sub-repo permissions client
	authz.DefaultSubRepoPermsChecker = srp.NewSubRepoPermsClient(db.SubRepoPerms())

	services, err := codeintel.NewServices(codeintel.ServiceDependencies{
		DB:             db,
		CodeIntelDB:    codeIntelDB,
		ObservationCtx: observationCtx,
	})
	if err != nil {
		return errors.Wrap(err, "creating codeintel services")
	}

	// Initialize stores
	uploadStore, err := lsifuploadstore.New(ctx, observationCtx, config.LSIFUploadStoreConfig)
	if err != nil {
		return errors.Wrap(err, "creating upload store")
	}
	if err := initializeUploadStore(ctx, uploadStore); err != nil {
		return errors.Wrap(err, "initializing upload store")
	}

	// Initialize worker
	worker := uploads.NewUploadProcessorJob(
		observationCtx,
		services.UploadsService,
		db,
		uploadStore,
		config.WorkerConcurrency,
		config.WorkerBudget,
		config.WorkerPollInterval,
		config.MaximumRuntimePerJob,
	)

	// Initialize health server
	server := httpserver.NewFromAddr(addr, &http.Server{
		ReadTimeout:  75 * time.Second,
		WriteTimeout: 10 * time.Minute,
		Handler:      httpserver.NewHandler(nil),
	})

	// Go!
	return goroutine.MonitorBackgroundRoutines(ctx, append(worker, server)...)
}

func mustInitializeDB(observationCtx *observation.Context) *sql.DB {
	dsn := conf.GetServiceConnectionValueAndRestartOnChange(func(serviceConnections conftypes.ServiceConnections) string {
		return serviceConnections.PostgresDSN
	})
	sqlDB, err := connections.EnsureNewFrontendDB(observationCtx, dsn, "precise-code-intel-worker")
	if err != nil {
		log.Scoped("init db").Fatal("Failed to connect to frontend database", log.Error(err))
	}

	return sqlDB
}

func mustInitializeCodeIntelDB(observationCtx *observation.Context) codeintelshared.CodeIntelDB {
	dsn := conf.GetServiceConnectionValueAndRestartOnChange(func(serviceConnections conftypes.ServiceConnections) string {
		return serviceConnections.CodeIntelPostgresDSN
	})
	db, err := connections.EnsureNewCodeIntelDB(observationCtx, dsn, "precise-code-intel-worker")
	if err != nil {
		log.Scoped("init db").Fatal("Failed to connect to codeintel database", log.Error(err))
	}

	return codeintelshared.NewCodeIntelDB(observationCtx.Logger, db)
}

func initializeUploadStore(ctx context.Context, uploadStore object.Storage) error {
	for {
		if err := uploadStore.Init(ctx); err == nil || !isRequestError(err) {
			return err
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(250 * time.Millisecond):
		}
	}
}

func isRequestError(err error) bool {
	return errors.HasType[*smithyhttp.RequestSendError](err)
}
