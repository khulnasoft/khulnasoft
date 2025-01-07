package context

import (
	"github.com/khulnasoft/khulnasoft/internal/codeintel/context/internal/store"
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/observation"
)

func NewService(
	observationCtx *observation.Context,
	db database.DB,
) *Service {
	store := store.New(scopedContext("store", observationCtx), db)

	return newService(
		observationCtx,
		store,
	)
}

func scopedContext(component string, parent *observation.Context) *observation.Context {
	return observation.ScopedContext("codeintel", "context", component, parent)
}
