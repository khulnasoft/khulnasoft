package compute

import (
	"context"
	"net/http"

	"github.com/sourcegraph/log"

	"github.com/khulnasoft/khulnasoft/cmd/frontend/enterprise"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/compute/resolvers"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/compute/streaming"
	"github.com/khulnasoft/khulnasoft/internal/codeintel"
	"github.com/khulnasoft/khulnasoft/internal/conf/conftypes"
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/observation"
)

func Init(
	ctx context.Context,
	observationCtx *observation.Context,
	db database.DB,
	_ codeintel.Services,
	_ conftypes.UnifiedWatchable,
	enterpriseServices *enterprise.Services,
) error {
	logger := log.Scoped("compute")
	enterpriseServices.ComputeResolver = resolvers.NewResolver(logger, db)
	enterpriseServices.NewComputeStreamHandler = func() http.Handler {
		return streaming.NewComputeStreamHandler(logger, db)
	}
	return nil
}
