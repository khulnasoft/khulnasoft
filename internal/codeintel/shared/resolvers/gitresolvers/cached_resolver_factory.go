package gitresolvers

import (
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/gitserver"
)

type CachedLocationResolverFactory struct {
	repoStore       database.RepoStore
	gitserverClient gitserver.Client
}

func NewCachedLocationResolverFactory(repoStore database.RepoStore, gitserverClient gitserver.Client) *CachedLocationResolverFactory {
	return &CachedLocationResolverFactory{
		repoStore:       repoStore,
		gitserverClient: gitserverClient,
	}
}

func (f *CachedLocationResolverFactory) Create() *CachedLocationResolver {
	return newCachedLocationResolver(f.repoStore, f.gitserverClient)
}
