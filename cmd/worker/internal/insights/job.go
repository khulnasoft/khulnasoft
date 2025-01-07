package insights

import (
	"context"

	"github.com/khulnasoft/khulnasoft/cmd/worker/job"
	workerinsightsdb "github.com/khulnasoft/khulnasoft/cmd/worker/shared/init/codeinsights"
	workerdb "github.com/khulnasoft/khulnasoft/cmd/worker/shared/init/db"
	"github.com/khulnasoft/khulnasoft/internal/env"
	"github.com/khulnasoft/khulnasoft/internal/goroutine"
	"github.com/khulnasoft/khulnasoft/internal/insights"
	"github.com/khulnasoft/khulnasoft/internal/insights/background"
	"github.com/khulnasoft/khulnasoft/internal/observation"
)

type insightsJob struct{}

func (s *insightsJob) Description() string {
	return ""
}

func (s *insightsJob) Config() []env.Config {
	return nil
}

func (s *insightsJob) Routines(_ context.Context, observationCtx *observation.Context) ([]goroutine.BackgroundRoutine, error) {
	if !insights.IsEnabled() {
		observationCtx.Logger.Debug("Code Insights disabled. Disabling background jobs.")
		return nil, nil
	}
	observationCtx.Logger.Debug("Code Insights enabled. Enabling background jobs.")

	db, err := workerdb.InitDB(observationCtx)
	if err != nil {
		return nil, err
	}

	insightsDB, err := workerinsightsdb.InitDB(observationCtx)
	if err != nil {
		return nil, err
	}

	return background.GetBackgroundJobs(context.Background(), observationCtx.Logger, db, insightsDB), nil
}

func NewInsightsJob() job.Job {
	return &insightsJob{}
}
