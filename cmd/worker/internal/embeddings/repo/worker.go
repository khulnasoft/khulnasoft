package repo

import (
	"context"
	"time"

	"github.com/khulnasoft/khulnasoft/cmd/worker/job"
	"github.com/khulnasoft/khulnasoft/cmd/worker/shared/init/codeintel"
	workerdb "github.com/khulnasoft/khulnasoft/cmd/worker/shared/init/db"
	"github.com/khulnasoft/khulnasoft/internal/actor"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/ranking"
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/embeddings"
	repoembeddingsbg "github.com/khulnasoft/khulnasoft/internal/embeddings/background/repo"
	"github.com/khulnasoft/khulnasoft/internal/embeddings/embed"
	"github.com/khulnasoft/khulnasoft/internal/env"
	"github.com/khulnasoft/khulnasoft/internal/gitserver"
	"github.com/khulnasoft/khulnasoft/internal/goroutine"
	"github.com/khulnasoft/khulnasoft/internal/object"
	"github.com/khulnasoft/khulnasoft/internal/observation"
	"github.com/khulnasoft/khulnasoft/internal/workerutil"
	"github.com/khulnasoft/khulnasoft/internal/workerutil/dbworker"
	dbworkerstore "github.com/khulnasoft/khulnasoft/internal/workerutil/dbworker/store"
)

type repoEmbeddingJob struct{}

func NewRepoEmbeddingJob() job.Job {
	return &repoEmbeddingJob{}
}

func (s *repoEmbeddingJob) Description() string {
	return ""
}

func (s *repoEmbeddingJob) Config() []env.Config {
	return []env.Config{embeddings.ObjectStorageConfigInst}
}

func (s *repoEmbeddingJob) Routines(_ context.Context, observationCtx *observation.Context) ([]goroutine.BackgroundRoutine, error) {
	db, err := workerdb.InitDB(observationCtx)
	if err != nil {
		return nil, err
	}

	uploadStore, err := embeddings.NewObjectStorage(context.Background(), observationCtx, embeddings.ObjectStorageConfigInst)
	if err != nil {
		return nil, err
	}

	services, err := codeintel.InitServices(observationCtx)
	if err != nil {
		return nil, err
	}

	workCtx := actor.WithInternalActor(context.Background())
	return []goroutine.BackgroundRoutine{
		newRepoEmbeddingJobWorker(
			workCtx,
			observationCtx,
			repoembeddingsbg.NewRepoEmbeddingJobWorkerStore(observationCtx, db.Handle()),
			db,
			uploadStore,
			gitserver.NewClient("embeddings.worker"),
			services.ContextService,
			repoembeddingsbg.NewRepoEmbeddingJobsStore(db),
			services.RankingService,
		),
	}, nil
}

func newRepoEmbeddingJobWorker(
	ctx context.Context,
	observationCtx *observation.Context,
	workerStore dbworkerstore.Store[*repoembeddingsbg.RepoEmbeddingJob],
	db database.DB,
	uploadStore object.Storage,
	gitserverClient gitserver.Client,
	contextService embed.ContextService,
	repoEmbeddingJobsStore repoembeddingsbg.RepoEmbeddingJobsStore,
	rankingService *ranking.Service,
) *workerutil.Worker[*repoembeddingsbg.RepoEmbeddingJob] {
	handler := &handler{
		db:                     db,
		uploadStore:            uploadStore,
		gitserverClient:        gitserverClient,
		contextService:         contextService,
		repoEmbeddingJobsStore: repoEmbeddingJobsStore,
		rankingService:         rankingService,
	}
	return dbworker.NewWorker[*repoembeddingsbg.RepoEmbeddingJob](ctx, workerStore, handler, workerutil.WorkerOptions{
		Name:              "repo_embedding_job_worker",
		Interval:          10 * time.Second, // Poll for a job once every 10 seconds
		NumHandlers:       1,                // Process only one job at a time (per instance)
		HeartbeatInterval: 10 * time.Second,
		Metrics:           workerutil.NewMetrics(observationCtx, "repo_embedding_job_worker"),
	})
}
