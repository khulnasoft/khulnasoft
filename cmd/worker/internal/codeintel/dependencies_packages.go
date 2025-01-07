package codeintel

import (
	"context"

	"github.com/khulnasoft/khulnasoft/cmd/worker/job"
	workerdb "github.com/khulnasoft/khulnasoft/cmd/worker/shared/init/db"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/dependencies"
	"github.com/khulnasoft/khulnasoft/internal/env"
	"github.com/khulnasoft/khulnasoft/internal/goroutine"
	"github.com/khulnasoft/khulnasoft/internal/observation"
)

type packageFilterApplicatorJob struct{}

func NewPackagesFilterApplicatorJob() job.Job {
	return &packageFilterApplicatorJob{}
}

func (j *packageFilterApplicatorJob) Description() string {
	return "package repo filters applicator"
}

func (j *packageFilterApplicatorJob) Config() []env.Config {
	return nil
}

func (j *packageFilterApplicatorJob) Routines(_ context.Context, observationCtx *observation.Context) ([]goroutine.BackgroundRoutine, error) {
	db, err := workerdb.InitDB(observationCtx)
	if err != nil {
		return nil, err
	}

	return dependencies.PackageFiltersJob(observationCtx, db), nil
}
