package encryption

import (
	"context"

	"github.com/khulnasoft/khulnasoft/cmd/worker/job"
	workerdb "github.com/khulnasoft/khulnasoft/cmd/worker/shared/init/db"
	"github.com/khulnasoft/khulnasoft/internal/database/basestore"
	"github.com/khulnasoft/khulnasoft/internal/env"
	"github.com/khulnasoft/khulnasoft/internal/goroutine"
	"github.com/khulnasoft/khulnasoft/internal/observation"
)

type recordEncrypterJob struct{}

func NewRecordEncrypterJob() job.Job {
	return &recordEncrypterJob{}
}

func (j *recordEncrypterJob) Description() string {
	return "encrypter routines"
}

func (j *recordEncrypterJob) Config() []env.Config {
	return []env.Config{
		ConfigInst,
	}
}

func (j *recordEncrypterJob) Routines(_ context.Context, observationCtx *observation.Context) ([]goroutine.BackgroundRoutine, error) {
	metrics := newMetrics(observationCtx)

	db, err := workerdb.InitDB(observationCtx)
	if err != nil {
		return nil, err
	}
	store := newRecordEncrypter(basestore.NewWithHandle(db.Handle()))

	return []goroutine.BackgroundRoutine{
		goroutine.NewPeriodicGoroutine(
			context.Background(),
			&recordEncrypterRoutine{
				store:   store,
				decrypt: ConfigInst.Decrypt,
				metrics: metrics,
				logger:  observationCtx.Logger,
			},
			goroutine.WithName("encryption.record-encrypter"),
			goroutine.WithDescription("encrypts/decrypts existing data when a key is provided/removed"),
			goroutine.WithInterval(ConfigInst.EncryptionInterval),
		),
		goroutine.NewPeriodicGoroutine(
			context.Background(),
			&recordCounter{
				store:   store,
				metrics: metrics,
				logger:  observationCtx.Logger,
			},
			goroutine.WithName("encryption.operation-metrics"),
			goroutine.WithDescription("tracks number of encrypted vs unencrypted records"),
			goroutine.WithInterval(ConfigInst.MetricsInterval),
		),
	}, nil
}
