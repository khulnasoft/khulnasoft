package outboundwebhooks

import (
	"time"

	"github.com/keegancsmith/sqlf"

	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/database/basestore"
	"github.com/khulnasoft/khulnasoft/internal/database/dbutil"
	"github.com/khulnasoft/khulnasoft/internal/encryption"
	"github.com/khulnasoft/khulnasoft/internal/observation"
	"github.com/khulnasoft/khulnasoft/internal/types"
	"github.com/khulnasoft/khulnasoft/internal/workerutil/dbworker/store"
)

func makeStore(observationCtx *observation.Context, db basestore.TransactableHandle, key encryption.Key) store.Store[*types.OutboundWebhookJob] {
	return store.New(observationCtx, db, store.Options[*types.OutboundWebhookJob]{
		Name:              "outbound_webhooks_worker_store",
		TableName:         "outbound_webhook_jobs",
		ColumnExpressions: database.OutboundWebhookJobColumns,
		Scan: store.BuildWorkerScan(func(sc dbutil.Scanner) (*types.OutboundWebhookJob, error) {
			return database.ScanOutboundWebhookJob(key, sc)
		}),
		OrderByExpression: sqlf.Sprintf("id"),
		MaxNumResets:      5,
		StalledMaxAge:     10 * time.Second,
	})
}
