package codeintel

import (
	"context"

	"github.com/khulnasoft/khulnasoft/cmd/worker/internal/executorqueue"
	"github.com/khulnasoft/khulnasoft/cmd/worker/job"
	"github.com/khulnasoft/khulnasoft/cmd/worker/shared/init/codeintel"

	workerdb "github.com/khulnasoft/khulnasoft/cmd/worker/shared/init/db"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/autoindexing"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/uploads"
	"github.com/khulnasoft/khulnasoft/internal/env"
	"github.com/khulnasoft/khulnasoft/internal/goroutine"
	"github.com/khulnasoft/khulnasoft/internal/observation"
	"github.com/khulnasoft/khulnasoft/internal/workerutil/dbworker"
	dbworkerstore "github.com/khulnasoft/khulnasoft/internal/workerutil/dbworker/store"
)

type metricsReporterJob struct{}

func NewMetricsReporterJob() job.Job {
	return &metricsReporterJob{}
}

func (j *metricsReporterJob) Description() string {
	return "executor push-based metrics reporting routines"
}

func (j *metricsReporterJob) Config() []env.Config {
	return []env.Config{
		configInst,
	}
}

func (j *metricsReporterJob) Routines(_ context.Context, observationCtx *observation.Context) ([]goroutine.BackgroundRoutine, error) {
	services, err := codeintel.InitServices(observationCtx)
	if err != nil {
		return nil, err
	}

	db, err := workerdb.InitDB(observationCtx)
	if err != nil {
		return nil, err
	}

	// TODO: move this and dependency {sync,index} metrics back to their respective jobs and keep for executor reporting only
	uploads.MetricReporters(observationCtx, services.UploadsService)

	dependencySyncStore := dbworkerstore.New(observationCtx, db.Handle(), autoindexing.DependencySyncingJobWorkerStoreOptions)
	dependencyIndexingStore := dbworkerstore.New(observationCtx, db.Handle(), autoindexing.DependencyIndexingJobWorkerStoreOptions)
	dbworker.InitPrometheusMetric(observationCtx, dependencySyncStore, "codeintel", "dependency_sync", nil)
	dbworker.InitPrometheusMetric(observationCtx, dependencyIndexingStore, "codeintel", "dependency_index", nil)

	executorMetricsReporter, err := executorqueue.NewMetricReporter(
		observationCtx,
		"codeintel",
		dbworkerstore.New(observationCtx, db.Handle(), autoindexing.IndexWorkerStoreOptions),
		configInst.MetricsConfig,
	)
	if err != nil {
		return nil, err
	}

	return []goroutine.BackgroundRoutine{executorMetricsReporter}, nil
}

type janitorConfig struct {
	MetricsConfig *executorqueue.Config
}

var configInst = &janitorConfig{}

func (c *janitorConfig) Load() {
	metricsConfig := executorqueue.InitMetricsConfig()
	metricsConfig.Load()
	c.MetricsConfig = metricsConfig
}

func (c *janitorConfig) Validate() error {
	return c.MetricsConfig.Validate()
}
