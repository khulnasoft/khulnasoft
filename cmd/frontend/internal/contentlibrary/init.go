package contentlibrary

import (
	"context"

	logger "github.com/sourcegraph/log"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/enterprise"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/graphqlbackend"
	"github.com/khulnasoft/khulnasoft/internal/codeintel"
	"github.com/khulnasoft/khulnasoft/internal/conf/conftypes"
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/observation"
)

// Init initializes the given enterpriseServices to include the required
// resolvers for the search content library.
func Init(
	_ context.Context,
	observationCtx *observation.Context,
	db database.DB,
	_ codeintel.Services,
	_ conftypes.UnifiedWatchable,
	enterpriseServices *enterprise.Services,
) error {
	logger := logger.Scoped("contentlibrary")
	enterpriseServices.ContentLibraryResolver = graphqlbackend.NewContentLibraryResolver(db, logger)
	return nil
}
