package embeddings

import (
	"context"

	"github.com/khulnasoft/khulnasoft/cmd/frontend/enterprise"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/embeddings/resolvers"
	"github.com/khulnasoft/khulnasoft/internal/codeintel"
	"github.com/khulnasoft/khulnasoft/internal/conf/conftypes"
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/embeddings"
	"github.com/khulnasoft/khulnasoft/internal/embeddings/background/repo"
	"github.com/khulnasoft/khulnasoft/internal/gitserver"
	"github.com/khulnasoft/khulnasoft/internal/observation"
)

func Init(
	_ context.Context,
	observationCtx *observation.Context,
	db database.DB,
	_ codeintel.Services,
	_ conftypes.UnifiedWatchable,
	enterpriseServices *enterprise.Services,
) error {
	repoEmbeddingsStore := repo.NewRepoEmbeddingJobsStore(db)
	gitserverClient := gitserver.NewClient("graphql.embeddings")
	embeddingsClient := embeddings.NewDefaultClient()
	enterpriseServices.EmbeddingsResolver = resolvers.NewResolver(
		db,
		observationCtx.Logger,
		gitserverClient,
		embeddingsClient,
		repoEmbeddingsStore,
	)

	return nil
}
