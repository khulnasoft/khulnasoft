package codeintel

import (
	"context"

	"github.com/khulnasoft/khulnasoft/cmd/worker/job"
	"github.com/khulnasoft/khulnasoft/cmd/worker/shared/init/codeintel"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/autoindexing"
	"github.com/khulnasoft/khulnasoft/internal/env"
	"github.com/khulnasoft/khulnasoft/internal/goroutine"

	"github.com/khulnasoft/khulnasoft/internal/observation"
)

type autoindexingSummaryBuilder struct{}

func NewAutoindexingSummaryBuilder() job.Job {
	return &autoindexingSummaryBuilder{}
}

func (j *autoindexingSummaryBuilder) Description() string {
	return ""
}

func (j *autoindexingSummaryBuilder) Config() []env.Config {
	return []env.Config{
		autoindexing.SummaryConfigInst,
	}
}

func (j *autoindexingSummaryBuilder) Routines(_ context.Context, observationCtx *observation.Context) ([]goroutine.BackgroundRoutine, error) {
	services, err := codeintel.InitServices(observationCtx)
	if err != nil {
		return nil, err
	}

	return autoindexing.NewSummaryBuilder(
		observationCtx,
		services.AutoIndexingService,
		services.UploadsService,
	), nil
}
