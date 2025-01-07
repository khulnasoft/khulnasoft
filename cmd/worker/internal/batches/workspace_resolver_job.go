package batches

import (
	"context"

	"github.com/khulnasoft/khulnasoft/cmd/worker/internal/batches/workers"
	"github.com/khulnasoft/khulnasoft/cmd/worker/job"
	"github.com/khulnasoft/khulnasoft/internal/actor"
	"github.com/khulnasoft/khulnasoft/internal/batches"
	"github.com/khulnasoft/khulnasoft/internal/env"
	"github.com/khulnasoft/khulnasoft/internal/goroutine"
	"github.com/khulnasoft/khulnasoft/internal/observation"
)

type workspaceResolverJob struct{}

func NewWorkspaceResolverJob() job.Job {
	return &workspaceResolverJob{}
}

func (j *workspaceResolverJob) Description() string {
	return ""
}

func (j *workspaceResolverJob) Config() []env.Config {
	return []env.Config{}
}

func (j *workspaceResolverJob) Routines(_ context.Context, observationCtx *observation.Context) ([]goroutine.BackgroundRoutine, error) {
	if !batches.IsEnabled() {
		return nil, nil
	}
	observationCtx = observation.NewContext(observationCtx.Logger.Scoped("routines"))
	workCtx := actor.WithInternalActor(context.Background())

	bstore, err := InitStore()
	if err != nil {
		return nil, err
	}

	resStore, err := InitBatchSpecResolutionWorkerStore()
	if err != nil {
		return nil, err
	}

	resolverWorker := workers.NewBatchSpecResolutionWorker(
		workCtx,
		observationCtx,
		bstore,
		resStore,
	)

	routines := []goroutine.BackgroundRoutine{
		resolverWorker,
	}

	return routines, nil
}
