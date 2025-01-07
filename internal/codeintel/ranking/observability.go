package ranking

import (
	"fmt"

	"github.com/khulnasoft/khulnasoft/internal/metrics"
	"github.com/khulnasoft/khulnasoft/internal/observation"
)

type operations struct {
	getRepoRank      *observation.Operation
	getDocumentRanks *observation.Operation
}

var (
	m = new(metrics.SingletonREDMetrics)
)

func newOperations(observationCtx *observation.Context) *operations {
	m := m.Get(func() *metrics.REDMetrics {
		return metrics.NewREDMetrics(
			observationCtx.Registerer,
			"codeintel_ranking",
			metrics.WithLabels("op"),
			metrics.WithCountHelp("Total number of method invocations."),
		)
	})

	op := func(name string) *observation.Operation {
		return observationCtx.Operation(observation.Op{
			Name:              fmt.Sprintf("codeintel.ranking.%s", name),
			MetricLabelValues: []string{name},
			Metrics:           m,
		})
	}

	return &operations{
		getRepoRank:      op("GetRepoRank"),
		getDocumentRanks: op("GetDocumentRanks"),
	}
}
