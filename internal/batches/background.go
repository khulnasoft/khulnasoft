package batches

import (
	"context"

	"github.com/sourcegraph/log"

	"github.com/khulnasoft/khulnasoft/internal/actor"
	"github.com/khulnasoft/khulnasoft/internal/batches/store"
	"github.com/khulnasoft/khulnasoft/internal/batches/syncer"
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/encryption"
	"github.com/khulnasoft/khulnasoft/internal/httpcli"
	"github.com/khulnasoft/khulnasoft/internal/observation"
)

// InitBackgroundJobs starts all jobs required to run batches. Currently, it is called from
// repo-updater and in the future will be the main entry point for the batch changes worker.
func InitBackgroundJobs(
	ctx context.Context,
	db database.DB,
	key encryption.Key,
	cf *httpcli.Factory,
) syncer.ChangesetSyncRegistry {
	// We use an internal actor so that we can freely load dependencies from
	// the database without repository permissions being enforced.
	// We do check for repository permissions consciously in the Rewirer when
	// creating new changesets and in the executor, when talking to the code
	// host, we manually check for BatchChangesCredentials.
	ctx = actor.WithInternalActor(ctx)

	observationCtx := observation.NewContext(log.Scoped("batches.background"))
	bstore := store.New(db, observationCtx, key)

	syncRegistry := syncer.NewSyncRegistry(ctx, observationCtx, bstore, cf)

	return syncRegistry
}
