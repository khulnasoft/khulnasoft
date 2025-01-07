package repo

import (
	"context"
	"time"

	"github.com/khulnasoft/khulnasoft/cmd/worker/job"
	workerdb "github.com/khulnasoft/khulnasoft/cmd/worker/shared/init/db"
	repoembeddingsbg "github.com/khulnasoft/khulnasoft/internal/embeddings/background/repo"
	"github.com/khulnasoft/khulnasoft/internal/env"
	"github.com/khulnasoft/khulnasoft/internal/goroutine"
	"github.com/khulnasoft/khulnasoft/internal/observation"
	"github.com/khulnasoft/khulnasoft/internal/workerutil/dbworker"
	dbworkerstore "github.com/khulnasoft/khulnasoft/internal/workerutil/dbworker/store"
)

type repoEmbeddingJanitorJob struct{}

func NewRepoEmbeddingJanitorJob() job.Job {
	return &repoEmbeddingJanitorJob{}
}

func (j *repoEmbeddingJanitorJob) Description() string {
	return ""
}

func (j *repoEmbeddingJanitorJob) Config() []env.Config {
	return []env.Config{}
}

func (j *repoEmbeddingJanitorJob) Routines(_ context.Context, observationCtx *observation.Context) ([]goroutine.BackgroundRoutine, error) {
	db, err := workerdb.InitDB(observationCtx)
	if err != nil {
		return nil, err
	}
	store := repoembeddingsbg.NewRepoEmbeddingJobWorkerStore(observationCtx, db.Handle())
	return []goroutine.BackgroundRoutine{newRepoEmbeddingJobResetter(observationCtx, store)}, nil
}

func newRepoEmbeddingJobResetter(observationCtx *observation.Context, workerStore dbworkerstore.Store[*repoembeddingsbg.RepoEmbeddingJob]) *dbworker.Resetter[*repoembeddingsbg.RepoEmbeddingJob] {
	return dbworker.NewResetter(observationCtx.Logger, workerStore, dbworker.ResetterOptions{
		Name:     "repo_embedding_job_worker_resetter",
		Interval: time.Minute, // Check for orphaned jobs every minute
		Metrics:  dbworker.NewResetterMetrics(observationCtx, "repo_embedding_job_worker"),
	})
}
