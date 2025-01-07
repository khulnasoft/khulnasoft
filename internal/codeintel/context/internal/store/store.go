package store

import (
	logger "github.com/sourcegraph/log"

	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/database/basestore"
	"github.com/khulnasoft/khulnasoft/internal/observation"
)

type Store interface {
	// TODO
}

type store struct {
	db         *basestore.Store
	logger     logger.Logger
	operations *operations
}

func New(observationCtx *observation.Context, db database.DB) Store {
	return &store{
		db:         basestore.NewWithHandle(db.Handle()),
		logger:     logger.Scoped("context.store"),
		operations: newOperations(observationCtx),
	}
}
