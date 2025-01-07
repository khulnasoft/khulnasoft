package policies

import (
	"fmt"

	"github.com/khulnasoft/khulnasoft/internal/metrics"
	"github.com/khulnasoft/khulnasoft/internal/observation"
)

type operations struct {
	updateConfigurationPolicy  *observation.Operation
	getRetentionPolicyOverview *observation.Operation
	getPreviewRepositoryFilter *observation.Operation
	getPreviewGitObjectFilter  *observation.Operation
}

var m = new(metrics.SingletonREDMetrics)

func newOperations(observationCtx *observation.Context) *operations {
	redMetrics := m.Get(func() *metrics.REDMetrics {
		return metrics.NewREDMetrics(
			observationCtx.Registerer,
			"codeintel_policies",
			metrics.WithLabels("op"),
			metrics.WithCountHelp("Total number of method invocations."),
		)
	})

	op := func(name string) *observation.Operation {
		return observationCtx.Operation(observation.Op{
			Name:              fmt.Sprintf("codeintel.policies.%s", name),
			MetricLabelValues: []string{name},
			Metrics:           redMetrics,
		})
	}

	return &operations{
		updateConfigurationPolicy:  op("UpdateConfigurationPolicy"),
		getRetentionPolicyOverview: op("GetRetentionPolicyOverview"),
		getPreviewRepositoryFilter: op("GetPreviewRepositoryFilter"),
		getPreviewGitObjectFilter:  op("GetPreviewGitObjectFilter"),
	}
}
