package store

import (
	"fmt"

	"github.com/khulnasoft/khulnasoft/internal/metrics"
	"github.com/khulnasoft/khulnasoft/internal/observation"
)

type operations struct {
	listPackageRepoRefs              *observation.Operation
	insertPackageRepoRefs            *observation.Operation
	deletePackageRepoRefsByID        *observation.Operation
	deletePackageRepoRefVersionsByID *observation.Operation

	listPackageRepoFilters  *observation.Operation
	createPackageRepoFilter *observation.Operation
	updatePackageRepoFilter *observation.Operation
	deletePackageRepoFilter *observation.Operation

	shouldRefilterPackageRepoRefs *observation.Operation
	updateAllBlockedStatuses      *observation.Operation
}

var m = new(metrics.SingletonREDMetrics)

func newOperations(observationCtx *observation.Context) *operations {
	redMetrics := m.Get(func() *metrics.REDMetrics {
		return metrics.NewREDMetrics(
			observationCtx.Registerer,
			"codeintel_dependencies_store",
			metrics.WithLabels("op"),
			metrics.WithCountHelp("Total number of method invocations."),
		)
	})

	op := func(name string) *observation.Operation {
		return observationCtx.Operation(observation.Op{
			Name:              fmt.Sprintf("codeintel.dependencies.store.%s", name),
			MetricLabelValues: []string{name},
			Metrics:           redMetrics,
		})
	}

	return &operations{
		listPackageRepoRefs:              op("ListDependencyRepos"),
		insertPackageRepoRefs:            op("InsertDependencyRepos"),
		deletePackageRepoRefsByID:        op("DeleteDependencyRepoRefsByID"),
		deletePackageRepoRefVersionsByID: op("DeletePackageRepoRefVersionsByID"),

		listPackageRepoFilters:  op("ListPackageRepoFilters"),
		createPackageRepoFilter: op("CreatePackageRepoFilter"),
		updatePackageRepoFilter: op("UpdatePackageRepoFilter"),
		deletePackageRepoFilter: op("DeletePackageRepoFilter"),

		shouldRefilterPackageRepoRefs: op("ShouldRefilterPackageRepoRefs"),
		updateAllBlockedStatuses:      op("UpdateAllBlockedStatuses"),
	}
}
