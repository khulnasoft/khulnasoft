package workers

import (
	"context"
	"time"

	"github.com/khulnasoft/khulnasoft/internal/batches/reconciler"
	"github.com/khulnasoft/khulnasoft/internal/batches/sources"
	"github.com/khulnasoft/khulnasoft/internal/batches/store"
	btypes "github.com/khulnasoft/khulnasoft/internal/batches/types"
	"github.com/khulnasoft/khulnasoft/internal/gitserver"
	"github.com/khulnasoft/khulnasoft/internal/observation"
	"github.com/khulnasoft/khulnasoft/internal/workerutil"
	"github.com/khulnasoft/khulnasoft/internal/workerutil/dbworker"
	dbworkerstore "github.com/khulnasoft/khulnasoft/internal/workerutil/dbworker/store"
)

// NewReconcilerWorker creates a dbworker.newWorker that fetches enqueued changesets
// from the database and passes them to the changeset reconciler for
// processing.
func NewReconcilerWorker(
	ctx context.Context,
	observationCtx *observation.Context,
	s *store.Store,
	workerStore dbworkerstore.Store[*btypes.Changeset],
	gitClient gitserver.Client,
	sourcer sources.Sourcer,
) *workerutil.Worker[*btypes.Changeset] {
	r := reconciler.New(gitClient, sourcer, s)

	options := workerutil.WorkerOptions{
		Name:              "batches_reconciler_worker",
		Description:       "changeset reconciler that publishes, modifies and closes changesets on the code host",
		NumHandlers:       5,
		Interval:          5 * time.Second,
		HeartbeatInterval: 15 * time.Second,
		Metrics:           workerutil.NewMetrics(observationCtx, "batch_changes_reconciler"),
	}

	worker := dbworker.NewWorker[*btypes.Changeset](ctx, workerStore, r.HandlerFunc(), options)
	return worker
}
