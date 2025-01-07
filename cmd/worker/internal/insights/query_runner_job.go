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

type insightsQueryRunnerJob struct {
	env.BaseConfig
}

func (s *insightsQueryRunnerJob) Description() string {
	return ""
}

func (s *insightsQueryRunnerJob) Config() []env.Config {
	return nil
}

func (s *insightsQueryRunnerJob) Routines(_ context.Context, observationCtx *observation.Context) ([]goroutine.BackgroundRoutine, error) {
	if !insights.IsEnabled() {
		observationCtx.Logger.Debug("Code Insights disabled. Disabling query runner.")
		return nil, nil
	}
	observationCtx.Logger.Debug("Code Insights enabled. Enabling query runner.")

	db, err := workerdb.InitDB(observationCtx)
	if err != nil {
		return nil, err
	}

	insightsDB, err := workerinsightsdb.InitDB(observationCtx)
	if err != nil {
		return nil, err
	}

	return background.GetBackgroundQueryRunnerJob(context.Background(), observationCtx.Logger, db, insightsDB), nil
}

func NewInsightsQueryRunnerJob() job.Job {
	return &insightsQueryRunnerJob{}
}
