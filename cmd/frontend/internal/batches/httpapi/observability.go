package httpapi

import (
	"fmt"

	"github.com/khulnasoft/khulnasoft/internal/metrics"
	"github.com/khulnasoft/khulnasoft/internal/observation"
)

type Operations struct {
	get    *observation.Operation
	exists *observation.Operation
	upload *observation.Operation
}

func NewOperations(observationCtx *observation.Context) *Operations {
	m := metrics.NewREDMetrics(
		observationCtx.Registerer,
		"batches_httpapi",
		metrics.WithLabels("op"),
		metrics.WithCountHelp("Total number of method invocations."),
	)

	op := func(name string) *observation.Operation {
		return observationCtx.Operation(observation.Op{
			Name:              fmt.Sprintf("batches.httpapi.%s", name),
			MetricLabelValues: []string{name},
			Metrics:           m,
		})
	}

	return &Operations{
		get:    op("get"),
		exists: op("exists"),
		upload: op("upload"),
	}
}
