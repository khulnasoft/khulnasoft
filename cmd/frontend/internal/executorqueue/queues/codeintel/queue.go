package codeintel

import (
	"context"

	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/executorqueue/handler"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/autoindexing"
	uploadsshared "github.com/khulnasoft/khulnasoft/internal/codeintel/uploads/shared"
	"github.com/khulnasoft/khulnasoft/internal/database"
	apiclient "github.com/khulnasoft/khulnasoft/internal/executor/types"
	"github.com/khulnasoft/khulnasoft/internal/observation"
	dbworkerstore "github.com/khulnasoft/khulnasoft/internal/workerutil/dbworker/store"
)

func QueueHandler(observationCtx *observation.Context, db database.DB, accessToken func() string) handler.QueueHandler[uploadsshared.AutoIndexJob] {
	recordTransformer := func(ctx context.Context, _ string, record uploadsshared.AutoIndexJob, resourceMetadata handler.ResourceMetadata) (apiclient.Job, error) {
		return transformRecord(ctx, db, record, resourceMetadata, accessToken())
	}

	store := dbworkerstore.New(observationCtx, db.Handle(), autoindexing.IndexWorkerStoreOptions)

	return handler.QueueHandler[uploadsshared.AutoIndexJob]{
		Name:              "codeintel",
		Store:             store,
		RecordTransformer: recordTransformer,
	}
}
