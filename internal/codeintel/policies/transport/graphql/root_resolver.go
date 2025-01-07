package graphql

import (
	"github.com/khulnasoft/khulnasoft/internal/codeintel/policies"
	resolverstubs "github.com/khulnasoft/khulnasoft/internal/codeintel/resolvers"
	sharedresolvers "github.com/khulnasoft/khulnasoft/internal/codeintel/shared/resolvers"
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/observation"
)

type rootResolver struct {
	policySvc        PoliciesService
	repoStore        database.RepoStore
	siteAdminChecker sharedresolvers.SiteAdminChecker
	operations       *operations
}

func NewRootResolver(
	observationCtx *observation.Context,
	policySvc *policies.Service,
	repoStore database.RepoStore,
	siteAdminChecker sharedresolvers.SiteAdminChecker,
) resolverstubs.PoliciesServiceResolver {
	return &rootResolver{
		policySvc:        policySvc,
		repoStore:        repoStore,
		siteAdminChecker: siteAdminChecker,
		operations:       newOperations(observationCtx),
	}
}
