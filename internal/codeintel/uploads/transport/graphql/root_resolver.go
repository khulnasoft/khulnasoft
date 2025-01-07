package graphql

import (
	resolverstubs "github.com/khulnasoft/khulnasoft/internal/codeintel/resolvers"
	sharedresolvers "github.com/khulnasoft/khulnasoft/internal/codeintel/shared/resolvers"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/shared/resolvers/gitresolvers"
	"github.com/khulnasoft/khulnasoft/internal/observation"
)

type rootResolver struct {
	uploadSvc                   UploadsService
	autoindexSvc                AutoIndexingService
	siteAdminChecker            sharedresolvers.SiteAdminChecker
	uploadLoaderFactory         UploadLoaderFactory
	autoIndexJobLoaderFactory   AutoIndexJobLoaderFactory
	locationResolverFactory     *gitresolvers.CachedLocationResolverFactory
	preciseIndexResolverFactory *PreciseIndexResolverFactory
	operations                  *operations
}

func NewRootResolver(
	observationCtx *observation.Context,
	uploadSvc UploadsService,
	autoindexSvc AutoIndexingService,
	siteAdminChecker sharedresolvers.SiteAdminChecker,
	uploadLoaderFactory UploadLoaderFactory,
	autoIndexJobLoaderFactory AutoIndexJobLoaderFactory,
	locationResolverFactory *gitresolvers.CachedLocationResolverFactory,
	preciseIndexResolverFactory *PreciseIndexResolverFactory,
) resolverstubs.UploadsServiceResolver {
	return &rootResolver{
		uploadSvc:                   uploadSvc,
		autoindexSvc:                autoindexSvc,
		siteAdminChecker:            siteAdminChecker,
		uploadLoaderFactory:         uploadLoaderFactory,
		autoIndexJobLoaderFactory:   autoIndexJobLoaderFactory,
		locationResolverFactory:     locationResolverFactory,
		preciseIndexResolverFactory: preciseIndexResolverFactory,
		operations:                  newOperations(observationCtx),
	}
}
