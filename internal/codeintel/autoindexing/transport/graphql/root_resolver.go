package graphql

import (
	resolverstubs "github.com/khulnasoft/khulnasoft/internal/codeintel/resolvers"
	sharedresolvers "github.com/khulnasoft/khulnasoft/internal/codeintel/shared/resolvers"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/shared/resolvers/gitresolvers"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/uploads/transport/graphql"
	"github.com/khulnasoft/khulnasoft/internal/conf"
	"github.com/khulnasoft/khulnasoft/internal/observation"
	"github.com/khulnasoft/khulnasoft/lib/errors"
)

type rootResolver struct {
	autoindexSvc                AutoIndexingService
	siteAdminChecker            sharedresolvers.SiteAdminChecker
	uploadLoaderFactory         graphql.UploadLoaderFactory
	autoIndexJobLoaderFactory   graphql.AutoIndexJobLoaderFactory
	locationResolverFactory     *gitresolvers.CachedLocationResolverFactory
	preciseIndexResolverFactory *graphql.PreciseIndexResolverFactory
	operations                  *operations
}

func NewRootResolver(
	observationCtx *observation.Context,
	autoindexSvc AutoIndexingService,
	siteAdminChecker sharedresolvers.SiteAdminChecker,
	uploadLoaderFactory graphql.UploadLoaderFactory,
	autoIndexJobLoaderFactory graphql.AutoIndexJobLoaderFactory,
	locationResolverFactory *gitresolvers.CachedLocationResolverFactory,
	preciseIndexResolverFactory *graphql.PreciseIndexResolverFactory,
) resolverstubs.AutoindexingServiceResolver {
	return &rootResolver{
		autoindexSvc:                autoindexSvc,
		siteAdminChecker:            siteAdminChecker,
		uploadLoaderFactory:         uploadLoaderFactory,
		autoIndexJobLoaderFactory:   autoIndexJobLoaderFactory,
		locationResolverFactory:     locationResolverFactory,
		preciseIndexResolverFactory: preciseIndexResolverFactory,
		operations:                  newOperations(observationCtx),
	}
}

var (
	autoIndexingEnabled       = conf.CodeIntelAutoIndexingEnabled
	errAutoIndexingNotEnabled = errors.New("precise code intelligence auto-indexing is not enabled")
)
