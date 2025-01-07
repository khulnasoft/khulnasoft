package adminanalytics

import (
	"context"
	"time"

	"github.com/khulnasoft/khulnasoft/cmd/worker/job"
	workerdb "github.com/khulnasoft/khulnasoft/cmd/worker/shared/init/db"
	"github.com/khulnasoft/khulnasoft/internal/adminanalytics"
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/env"
	"github.com/khulnasoft/khulnasoft/internal/goroutine"
	"github.com/khulnasoft/khulnasoft/internal/metrics"
	"github.com/khulnasoft/khulnasoft/internal/observation"
	"github.com/khulnasoft/khulnasoft/internal/redispool"
)

type refreshAnalyticsCacheJob struct{}

func NewRefreshAnalyticsCacheJob() job.Job {
	return &refreshAnalyticsCacheJob{}
}

func (e refreshAnalyticsCacheJob) Description() string {
	return "refreshes the admin analytics cache"
}

func (e refreshAnalyticsCacheJob) Config() []env.Config {
	return nil
}

func (e refreshAnalyticsCacheJob) Routines(_ context.Context, observationCtx *observation.Context) ([]goroutine.BackgroundRoutine, error) {
	db, err := workerdb.InitDB(observationCtx)
	if err != nil {
		return nil, err
	}

	return []goroutine.BackgroundRoutine{
			newRefreshAnalyticsCacheJob(observationCtx, redispool.Store, db),
		},
		nil
}

func newRefreshAnalyticsCacheJob(observationCtx *observation.Context, cache adminanalytics.KeyValue, db database.DB) goroutine.BackgroundRoutine {
	handler := goroutine.HandlerFunc(func(ctx context.Context) error {
		return refreshAnalyticsCache(ctx, cache, db)
	})

	operation := observationCtx.Operation(observation.Op{
		Name: "analytics.cache.update",
		Metrics: metrics.NewREDMetrics(
			observationCtx.Registerer,
			"refresh_analytics_cache",
			metrics.WithCountHelp("Total number of refresh_analytics_cache executions"),
		),
	})

	return goroutine.NewPeriodicGoroutine(
		context.Background(),
		handler,
		goroutine.WithName("refresh_analytics_cache"),
		goroutine.WithDescription("refresh analytics cache"),
		goroutine.WithInterval(24*time.Hour),
		goroutine.WithOperation(operation),
	)
}
