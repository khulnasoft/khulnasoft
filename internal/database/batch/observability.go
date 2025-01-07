package batch

import (
	"fmt"
	"sync"

	"github.com/sourcegraph/log"

	"github.com/khulnasoft/khulnasoft/internal/metrics"
	"github.com/khulnasoft/khulnasoft/internal/observation"
)

type operations struct {
	flush *observation.Operation
}

func newOperations(observationCtx *observation.Context) *operations {
	redMetrics := metrics.NewREDMetrics(
		observationCtx.Registerer,
		"database_batch",
		metrics.WithLabels("op"),
		metrics.WithCountHelp("Total number of method invocations."),
	)

	op := func(name string) *observation.Operation {
		return observationCtx.Operation(observation.Op{
			Name:              fmt.Sprintf("database.batch.%s", name),
			MetricLabelValues: []string{name},
			Metrics:           redMetrics,
		})
	}

	return &operations{
		flush: op("Flush"),
	}
}

var (
	ops     *operations
	opsOnce sync.Once
)

func getOperations(logger log.Logger) *operations {
	opsOnce.Do(func() {
		observationCtx := observation.NewContext(logger)

		ops = newOperations(observationCtx)
	})

	return ops
}
