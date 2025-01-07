package outboundwebhooks

import (
	"context"
	"net/http"
	"time"

	"github.com/khulnasoft/khulnasoft/cmd/worker/job"
	workerdb "github.com/khulnasoft/khulnasoft/cmd/worker/shared/init/db"
	"github.com/khulnasoft/khulnasoft/internal/actor"
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/encryption/keyring"
	"github.com/khulnasoft/khulnasoft/internal/env"
	"github.com/khulnasoft/khulnasoft/internal/goroutine"
	"github.com/khulnasoft/khulnasoft/internal/httpcli"
	"github.com/khulnasoft/khulnasoft/internal/observation"
	"github.com/khulnasoft/khulnasoft/internal/types"
	"github.com/khulnasoft/khulnasoft/internal/workerutil"
	"github.com/khulnasoft/khulnasoft/internal/workerutil/dbworker"
	"github.com/khulnasoft/khulnasoft/internal/workerutil/dbworker/store"
	"github.com/khulnasoft/khulnasoft/lib/errors"
)

type sender struct{}

func NewSender() job.Job {
	return &sender{}
}

func (s *sender) Description() string {
	return "Outbound webhook sender"
}

func (*sender) Config() []env.Config {
	return nil
}

func (s *sender) Routines(_ context.Context, observationCtx *observation.Context) ([]goroutine.BackgroundRoutine, error) {
	observationCtx = observation.NewContext(observationCtx.Logger.Scoped("sender"))
	ctx := actor.WithInternalActor(context.Background())

	db, err := workerdb.InitDB(observationCtx)
	if err != nil {
		return nil, errors.Wrap(err, "initialising database")
	}

	client := httpcli.ExternalClient
	key := keyring.Default().OutboundWebhookKey
	workerStore := makeStore(observationCtx, db.Handle(), key)

	return []goroutine.BackgroundRoutine{
		makeWorker(
			ctx, observationCtx, workerStore, client,
			database.OutboundWebhooksWith(db, key),
			database.OutboundWebhookLogsWith(db, key),
		),
		makeResetter(observationCtx, workerStore),
		makeJanitor(observationCtx, db.OutboundWebhookJobs(key)),
	}, nil
}

func makeWorker(
	ctx context.Context,
	observationCtx *observation.Context,
	workerStore store.Store[*types.OutboundWebhookJob],
	client *http.Client,
	webhookStore database.OutboundWebhookStore,
	logStore database.OutboundWebhookLogStore,
) *workerutil.Worker[*types.OutboundWebhookJob] {
	handler := &handler{
		client:   client,
		store:    webhookStore,
		logStore: logStore,
	}

	return dbworker.NewWorker[*types.OutboundWebhookJob](
		ctx, workerStore, handler, workerutil.WorkerOptions{
			Name:              "outbound_webhook_job_worker",
			Interval:          time.Second,
			NumHandlers:       1,
			HeartbeatInterval: 10 * time.Second,
			Metrics:           workerutil.NewMetrics(observationCtx, "outbound_webhook_job_worker"),
		},
	)
}

func makeResetter(
	observationCtx *observation.Context,
	workerStore store.Store[*types.OutboundWebhookJob],
) *dbworker.Resetter[*types.OutboundWebhookJob] {
	return dbworker.NewResetter(
		observationCtx.Logger, workerStore, dbworker.ResetterOptions{
			Name:     "outbound_webhook_job_resetter",
			Interval: 5 * time.Minute,
			Metrics:  dbworker.NewResetterMetrics(observationCtx, "outbound_webhook_job_resetter"),
		},
	)
}
