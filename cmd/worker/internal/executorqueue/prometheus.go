package executorqueue

import (
	"github.com/khulnasoft/khulnasoft/internal/observation"
	"github.com/khulnasoft/khulnasoft/internal/workerutil"
	"github.com/khulnasoft/khulnasoft/internal/workerutil/dbworker"
	"github.com/khulnasoft/khulnasoft/internal/workerutil/dbworker/store"
)

func initPrometheusMetric[T workerutil.Record](observationCtx *observation.Context, queueName string, store store.Store[T]) {
	dbworker.InitPrometheusMetric(observationCtx, store, "", "executor", map[string]string{"queue": queueName})
}
