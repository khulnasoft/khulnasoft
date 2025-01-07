package telemetry

import (
	"context"

	"github.com/khulnasoft/khulnasoft/cmd/frontend/enterprise"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/graphqlbackend"
	"github.com/khulnasoft/khulnasoft/internal/codeintel"
	"github.com/khulnasoft/khulnasoft/internal/conf/conftypes"
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/observation"

	resolvers "github.com/khulnasoft/khulnasoft/cmd/frontend/internal/telemetry/resolvers"
)

// Init initializes the given enterpriseServices to include the required resolvers for telemetry.
func Init(
	ctx context.Context,
	observationCtx *observation.Context,
	db database.DB,
	_ codeintel.Services,
	_ conftypes.UnifiedWatchable,
	enterpriseServices *enterprise.Services,
) error {
	enterpriseServices.TelemetryRootResolver = &graphqlbackend.TelemetryRootResolver{
		Resolver: resolvers.New(
			observationCtx.Logger.Scoped("telemetry"),
			db),
	}

	return nil
}
