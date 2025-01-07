package dependencies

import (
	"github.com/khulnasoft/khulnasoft/internal/codeintel/dependencies/internal/background"
	dependenciesstore "github.com/khulnasoft/khulnasoft/internal/codeintel/dependencies/internal/store"
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/goroutine"
	"github.com/khulnasoft/khulnasoft/internal/observation"
)

func NewService(observationCtx *observation.Context, db database.DB) *Service {
	return newService(scopedContext("service", observationCtx), dependenciesstore.New(scopedContext("store", observationCtx), db))
}

// TestService creates a new dependencies service with noop observation contexts.
func TestService(db database.DB) *Service {
	store := dependenciesstore.New(&observation.TestContext, db)

	return newService(&observation.TestContext, store)
}

func scopedContext(component string, parent *observation.Context) *observation.Context {
	return observation.ScopedContext("codeintel", "dependencies", component, parent)
}

func PackageFiltersJob(
	obsctx *observation.Context,
	db database.DB,
) goroutine.CombinedRoutine {
	return []goroutine.BackgroundRoutine{
		background.NewPackagesFilterApplicator(obsctx, db),
	}
}
