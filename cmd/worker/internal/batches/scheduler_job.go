package batches

import (
	"context"

	"github.com/khulnasoft/khulnasoft/cmd/worker/job"
	"github.com/khulnasoft/khulnasoft/internal/actor"
	"github.com/khulnasoft/khulnasoft/internal/batches"
	"github.com/khulnasoft/khulnasoft/internal/batches/scheduler"
	"github.com/khulnasoft/khulnasoft/internal/env"
	"github.com/khulnasoft/khulnasoft/internal/goroutine"
	"github.com/khulnasoft/khulnasoft/internal/observation"
)

type schedulerJob struct{}

func NewSchedulerJob() job.Job {
	return &schedulerJob{}
}

func (j *schedulerJob) Description() string {
	return ""
}

func (j *schedulerJob) Config() []env.Config {
	return []env.Config{}
}

func (j *schedulerJob) Routines(_ context.Context, observationCtx *observation.Context) ([]goroutine.BackgroundRoutine, error) {
	if !batches.IsEnabled() {
		return nil, nil
	}
	workCtx := actor.WithInternalActor(context.Background())

	bstore, err := InitStore()
	if err != nil {
		return nil, err
	}

	routines := []goroutine.BackgroundRoutine{
		scheduler.NewScheduler(workCtx, bstore),
	}

	return routines, nil
}
