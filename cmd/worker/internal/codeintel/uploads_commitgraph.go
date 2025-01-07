package codeintel

import (
	"context"

	"github.com/khulnasoft/khulnasoft/cmd/worker/job"
	"github.com/khulnasoft/khulnasoft/cmd/worker/shared/init/codeintel"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/uploads"
	"github.com/khulnasoft/khulnasoft/internal/env"
	"github.com/khulnasoft/khulnasoft/internal/goroutine"
	"github.com/khulnasoft/khulnasoft/internal/observation"
)

type commitGraphUpdaterJob struct{}

func NewCommitGraphUpdaterJob() job.Job {
	return &commitGraphUpdaterJob{}
}

func (j *commitGraphUpdaterJob) Description() string {
	return ""
}

func (j *commitGraphUpdaterJob) Config() []env.Config {
	return []env.Config{
		uploads.CommitGraphConfigInst,
	}
}

func (j *commitGraphUpdaterJob) Routines(_ context.Context, observationCtx *observation.Context) ([]goroutine.BackgroundRoutine, error) {
	services, err := codeintel.InitServices(observationCtx)
	if err != nil {
		return nil, err
	}

	return uploads.NewCommitGraphUpdater(services.UploadsService, services.GitserverClient), nil
}
