package own

import (
	"context"

	"github.com/khulnasoft/khulnasoft/cmd/frontend/enterprise"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/own/resolvers"
	"github.com/khulnasoft/khulnasoft/internal/codeintel"
	"github.com/khulnasoft/khulnasoft/internal/conf/conftypes"
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/gitserver"
	"github.com/khulnasoft/khulnasoft/internal/observation"
	"github.com/khulnasoft/khulnasoft/internal/own"
)

// Init initializes the given enterpriseServices to include the required
// resolvers for Khulnasoft Own.
func Init(
	_ context.Context,
	observationCtx *observation.Context,
	db database.DB,
	_ codeintel.Services,
	_ conftypes.UnifiedWatchable,
	enterpriseServices *enterprise.Services,
) error {
	if !own.IsEnabled() {
		return nil
	}
	g := gitserver.NewClient("graphql.own")
	enterpriseServices.OwnResolver = resolvers.New(db, g, observationCtx.Logger.Scoped("own"))
	return nil
}
