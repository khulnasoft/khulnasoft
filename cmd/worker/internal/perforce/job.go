package perforce

import (
	"context"

	"github.com/khulnasoft/khulnasoft/cmd/worker/job"
	workerdb "github.com/khulnasoft/khulnasoft/cmd/worker/shared/init/db"
	"github.com/khulnasoft/khulnasoft/internal/actor"
	"github.com/khulnasoft/khulnasoft/internal/env"
	"github.com/khulnasoft/khulnasoft/internal/gitserver"
	"github.com/khulnasoft/khulnasoft/internal/goroutine"
	"github.com/khulnasoft/khulnasoft/internal/metrics"
	"github.com/khulnasoft/khulnasoft/internal/observation"
)

type perforceChangelistMappingJob struct {
	cfg *Config
}

func NewPerforceChangelistMappingJob() job.Job {
	return &perforceChangelistMappingJob{cfg: &Config{}}
}

func (j *perforceChangelistMappingJob) Description() string {
	return "Background job indexing mapping data from Perforce changelist numbers to imported Git commits"
}

func (j *perforceChangelistMappingJob) Config() []env.Config {
	return []env.Config{j.cfg}
}

func (j *perforceChangelistMappingJob) Routines(startupCtx context.Context, observationCtx *observation.Context) ([]goroutine.BackgroundRoutine, error) {
	db, err := workerdb.InitDB(observationCtx)
	if err != nil {
		return nil, err
	}

	logger := observationCtx.Logger.Scoped("perforce-changelist-mapper")

	gs := gitserver.NewClient("perforce-changelist-mapper")

	// make sure to use an internal actor. The changelist mapper has to be able to
	// access any repo to index changelist IDs.
	ctx := actor.WithInternalActor(context.Background())

	redMetrics := metrics.NewREDMetrics(
		observationCtx.Registerer,
		"perforce_changelist_mapper",
		metrics.WithLabels("op"),
		metrics.WithCountHelp("Total number of method invocations."),
	)

	op := observationCtx.Operation(observation.Op{
		Name:              "perforce.changelist-mapper",
		MetricLabelValues: []string{"processRepo"},
		Metrics:           redMetrics,
	})

	return []goroutine.BackgroundRoutine{
		goroutine.NewPeriodicGoroutine(
			ctx,
			&perforceChangelistMapper{
				cfg:    j.cfg,
				db:     db,
				gs:     gs,
				logger: logger,
			},
			goroutine.WithName("perforce-changelist-mapper"),
			goroutine.WithDescription("Background job indexing mapping data from Perforce changelist numbers to imported Git commits"),
			goroutine.WithInterval(j.cfg.Interval),
			goroutine.WithOperation(op),
		),
	}, nil
}
