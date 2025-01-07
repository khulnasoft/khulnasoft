package embeddings

import (
	"context"

	"github.com/khulnasoft/khulnasoft/cmd/frontend/enterprise"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/codycontext"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/context/resolvers"
	"github.com/khulnasoft/khulnasoft/internal/codeintel"
	"github.com/khulnasoft/khulnasoft/internal/conf/conftypes"
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/embeddings"
	"github.com/khulnasoft/khulnasoft/internal/gitserver"
	"github.com/khulnasoft/khulnasoft/internal/observation"
	"github.com/khulnasoft/khulnasoft/internal/search/client"
)

func Init(
	_ context.Context,
	observationCtx *observation.Context,
	db database.DB,
	services codeintel.Services,
	_ conftypes.UnifiedWatchable,
	enterpriseServices *enterprise.Services,
) error {
	observationCtx = observationCtx.Clone()
	observationCtx.Logger = observationCtx.Logger.Scoped("codycontext")

	embeddingsClient := embeddings.NewDefaultClient()
	searchClient := client.New(observationCtx.Logger, db, gitserver.NewClient("graphql.context.search"))

	contextClient := codycontext.NewCodyContextClient(
		observationCtx,
		db,
		embeddingsClient,
		searchClient,
		services.GitserverClient.Scoped("codycontext.client"),
	)
	enterpriseServices.CodyContextResolver = resolvers.NewResolver(
		db,
		services.GitserverClient,
		contextClient,
		observationCtx.Logger,
	)

	return nil
}
