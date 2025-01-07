package phabricator

import (
	"context"

	"github.com/khulnasoft/khulnasoft/cmd/worker/job"
	workerdb "github.com/khulnasoft/khulnasoft/cmd/worker/shared/init/db"
	"github.com/khulnasoft/khulnasoft/internal/env"
	"github.com/khulnasoft/khulnasoft/internal/goroutine"
	"github.com/khulnasoft/khulnasoft/internal/observation"
	"github.com/khulnasoft/khulnasoft/internal/repos"
)

type phabricatorRepoSyncerJob struct{}

func NewPhabricatorRepoSyncerJob() job.Job {
	return &phabricatorRepoSyncerJob{}
}

func (o *phabricatorRepoSyncerJob) Description() string {
	return "Periodically syncs repositories from Phabricator to Sourcegraph"
}

func (o *phabricatorRepoSyncerJob) Config() []env.Config {
	return nil
}

func (o *phabricatorRepoSyncerJob) Routines(startupCtx context.Context, observationCtx *observation.Context) ([]goroutine.BackgroundRoutine, error) {
	db, err := workerdb.InitDB(observationCtx)
	if err != nil {
		return nil, err
	}

	store := repos.NewStore(observationCtx.Logger.Scoped("store"), db)

	return []goroutine.BackgroundRoutine{
		NewRepositorySyncWorker(context.Background(), db, observationCtx.Logger, store),
	}, nil
}
