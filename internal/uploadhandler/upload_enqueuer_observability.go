package uploadhandler

import (
	"fmt"
	"syscall"

	"github.com/khulnasoft/khulnasoft/internal/metrics"
	"github.com/khulnasoft/khulnasoft/internal/observation"
	"github.com/khulnasoft/khulnasoft/lib/errors"
)

type EnqueuerOperations struct {
	enqueueSinglePayload *observation.Operation
}

func NewEnqueuerOperations(observationCtx *observation.Context) *EnqueuerOperations {
	redMetrics := metrics.NewREDMetrics(
		observationCtx.Registerer,
		"upload_enqueuer",
		metrics.WithLabels("op"),
		metrics.WithCountHelp("Total number of method invocations."),
	)

	op := func(name string) *observation.Operation {
		return observationCtx.Operation(observation.Op{
			Name:              fmt.Sprintf("upload_enqueuer.%s", name),
			MetricLabelValues: []string{name},
			Metrics:           redMetrics,
			ErrorFilter: func(err error) observation.ErrorFilterBehaviour {
				var errno syscall.Errno
				if errors.As(err, &errno) && errno == syscall.ECONNREFUSED {
					return observation.EmitForDefault ^ observation.EmitForSentry
				}
				return observation.EmitForDefault
			},
		})
	}

	return &EnqueuerOperations{
		enqueueSinglePayload: op("enqueueSinglePayload"),
	}
}
