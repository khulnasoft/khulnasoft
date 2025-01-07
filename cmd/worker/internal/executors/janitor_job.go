package executors

import (
	"context"

	"github.com/khulnasoft/khulnasoft/cmd/worker/job"
	workerdb "github.com/khulnasoft/khulnasoft/cmd/worker/shared/init/db"
	"github.com/khulnasoft/khulnasoft/internal/env"
	executortypes "github.com/khulnasoft/khulnasoft/internal/executor/types"
	"github.com/khulnasoft/khulnasoft/internal/goroutine"
	"github.com/khulnasoft/khulnasoft/internal/observation"
	"github.com/khulnasoft/khulnasoft/internal/rcache"
	"github.com/khulnasoft/khulnasoft/internal/redispool"
)

type janitorJob struct{}

func NewJanitorJob() job.Job {
	return &janitorJob{}
}

func (j *janitorJob) Description() string {
	return ""
}

func (j *janitorJob) Config() []env.Config {
	return []env.Config{janitorConfigInst}
}

func (j *janitorJob) Routines(_ context.Context, observationCtx *observation.Context) ([]goroutine.BackgroundRoutine, error) {
	db, err := workerdb.InitDB(observationCtx)
	if err != nil {
		return nil, err
	}

	dequeueCache := rcache.New(redispool.Cache, executortypes.DequeueCachePrefix)

	routines := []goroutine.BackgroundRoutine{
		goroutine.NewPeriodicGoroutine(
			context.Background(),
			goroutine.HandlerFunc(func(ctx context.Context) error {
				return db.Executors().DeleteInactiveHeartbeats(ctx, janitorConfigInst.HeartbeatRecordsMaxAge)
			}),
			goroutine.WithName("executor.heartbeat-janitor"),
			goroutine.WithDescription("clean up executor heartbeat records for presumed dead executors"),
			goroutine.WithInterval(janitorConfigInst.CleanupTaskInterval),
		),
		NewMultiqueueCacheCleaner(executortypes.ValidQueueNames, dequeueCache, janitorConfigInst.CacheDequeueTtl, janitorConfigInst.CacheCleanupInterval),
	}

	return routines, nil
}
