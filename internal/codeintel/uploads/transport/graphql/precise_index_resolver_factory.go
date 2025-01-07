package graphql

import (
	"context"

	resolverstubs "github.com/khulnasoft/khulnasoft/internal/codeintel/resolvers"
	sharedresolvers "github.com/khulnasoft/khulnasoft/internal/codeintel/shared/resolvers"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/shared/resolvers/gitresolvers"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/uploads/shared"
	uploadsshared "github.com/khulnasoft/khulnasoft/internal/codeintel/uploads/shared"
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/gitserver"
	"github.com/khulnasoft/khulnasoft/internal/observation"
)

type PreciseIndexResolverFactory struct {
	uploadsSvc       UploadsService
	policySvc        PolicyService
	gitserverClient  gitserver.Client
	siteAdminChecker sharedresolvers.SiteAdminChecker
	repoStore        database.RepoStore
}

func NewPreciseIndexResolverFactory(
	uploadsSvc UploadsService,
	policySvc PolicyService,
	gitserverClient gitserver.Client,
	siteAdminChecker sharedresolvers.SiteAdminChecker,
	repoStore database.RepoStore,
) *PreciseIndexResolverFactory {
	return &PreciseIndexResolverFactory{
		uploadsSvc:       uploadsSvc,
		policySvc:        policySvc,
		gitserverClient:  gitserverClient,
		siteAdminChecker: siteAdminChecker,
		repoStore:        repoStore,
	}
}

func (f *PreciseIndexResolverFactory) Create(
	ctx context.Context,
	uploadLoader UploadLoader,
	autoIndexJobLoader AutoIndexJobLoader,
	locationResolver *gitresolvers.CachedLocationResolver,
	traceErrs *observation.ErrCollector,
	upload *shared.Upload,
	index *uploadsshared.AutoIndexJob,
) (resolverstubs.PreciseIndexResolver, error) {
	return newPreciseIndexResolver(
		ctx,
		f.uploadsSvc,
		f.policySvc,
		f.gitserverClient,
		uploadLoader,
		autoIndexJobLoader,
		f.siteAdminChecker,
		f.repoStore,
		locationResolver,
		traceErrs,
		upload,
		index,
	)
}
