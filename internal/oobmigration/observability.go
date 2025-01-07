package oobmigration

import (
	"fmt"
	"strconv"

	"github.com/khulnasoft/khulnasoft/internal/metrics"
	"github.com/khulnasoft/khulnasoft/internal/observation"
)

type operations struct {
	upForMigration   func(migrationID int) *observation.Operation
	downForMigration func(migrationID int) *observation.Operation
}

func newOperations(observationCtx *observation.Context) *operations {
	redMetrics := metrics.NewREDMetrics(
		observationCtx.Registerer,
		"oobmigration",
		metrics.WithLabels("op", "migration"),
		metrics.WithCountHelp("Total number of migrator invocations."),
	)

	opForMigration := func(name string) func(migrationID int) *observation.Operation {
		return func(migrationID int) *observation.Operation {
			return observationCtx.Operation(observation.Op{
				Name:              fmt.Sprintf("oobmigration.%s", name),
				MetricLabelValues: []string{name, strconv.Itoa(migrationID)},
				Metrics:           redMetrics,
			})
		}
	}

	return &operations{
		upForMigration:   opForMigration("up"),
		downForMigration: opForMigration("down"),
	}
}
