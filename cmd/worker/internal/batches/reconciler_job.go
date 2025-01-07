package batches

import (
	"context"

	"github.com/khulnasoft/khulnasoft/cmd/worker/internal/batches/workers"
	"github.com/khulnasoft/khulnasoft/cmd/worker/job"
	"github.com/khulnasoft/khulnasoft/internal/actor"
	"github.com/khulnasoft/khulnasoft/internal/batches"
	"github.com/khulnasoft/khulnasoft/internal/batches/sources"
	"github.com/khulnasoft/khulnasoft/internal/env"
	"github.com/khulnasoft/khulnasoft/internal/gitserver"
	"github.com/khulnasoft/khulnasoft/internal/goroutine"
	"github.com/khulnasoft/khulnasoft/internal/httpcli"
	"github.com/khulnasoft/khulnasoft/internal/observation"
)

type reconcilerJob struct{}

func NewReconcilerJob() job.Job {
	return &reconcilerJob{}
}

func (j *reconcilerJob) Description() string {
	return ""
}

func (j *reconcilerJob) Config() []env.Config {
	return []env.Config{}
}

func (j *reconcilerJob) Routines(_ context.Context, observationCtx *observation.Context) ([]goroutine.BackgroundRoutine, error) {
	if !batches.IsEnabled() {
		return nil, nil
	}
	observationCtx = observation.NewContext(observationCtx.Logger.Scoped("routines"))
	workCtx := actor.WithInternalActor(context.Background())

	bstore, err := InitStore()
	if err != nil {
		return nil, err
	}

	reconcilerStore, err := InitReconcilerWorkerStore()
	if err != nil {
		return nil, err
	}

	reconcilerWorker := workers.NewReconcilerWorker(
		workCtx,
		observationCtx,
		bstore,
		reconcilerStore,
		gitserver.NewClient("batches.reconciler"),
		sources.NewSourcer(httpcli.NewExternalClientFactory(
			httpcli.NewLoggingMiddleware(observationCtx.Logger.Scoped("sourcer")),
		)),
	)

	routines := []goroutine.BackgroundRoutine{
		reconcilerWorker,
	}

	return routines, nil
}
