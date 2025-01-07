package migrations

import (
	"context"

	"github.com/khulnasoft/khulnasoft/internal/database/basestore"
)

type StoreFactory interface {
	Store(ctx context.Context, schemaName string) (*basestore.Store, error)
}
