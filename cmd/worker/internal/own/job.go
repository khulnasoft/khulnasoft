package own

import (
	"context"

	"github.com/khulnasoft/khulnasoft/cmd/worker/job"
	workerdb "github.com/khulnasoft/khulnasoft/cmd/worker/shared/init/db"
	"github.com/khulnasoft/khulnasoft/internal/env"
	"github.com/khulnasoft/khulnasoft/internal/goroutine"
	"github.com/khulnasoft/khulnasoft/internal/observation"
	"github.com/khulnasoft/khulnasoft/internal/own"
	"github.com/khulnasoft/khulnasoft/internal/own/background"
)

type ownRepoIndexingQueue struct{}

func NewOwnRepoIndexingQueue() job.Job {
	return &ownRepoIndexingQueue{}
}

func (o *ownRepoIndexingQueue) Description() string {
	return "Queue used to index ownership data partitioned per repository"
}

func (o *ownRepoIndexingQueue) Config() []env.Config {
	return nil
}

func (o *ownRepoIndexingQueue) Routines(startupCtx context.Context, observationCtx *observation.Context) ([]goroutine.BackgroundRoutine, error) {
	if !own.IsEnabled() {
		return nil, nil
	}
	db, err := workerdb.InitDB(observationCtx)
	if err != nil {
		return nil, err
	}

	var routines []goroutine.BackgroundRoutine
	routines = append(routines, background.NewOwnBackgroundWorker(context.Background(), db, observationCtx)...)
	routines = append(routines, background.GetOwnIndexSchedulerRoutines(db, observationCtx)...)

	return routines, nil
}
