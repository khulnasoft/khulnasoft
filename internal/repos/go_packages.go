package repos

import (
	"context"

	"github.com/khulnasoft/khulnasoft/internal/api"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/dependencies"
	"github.com/khulnasoft/khulnasoft/internal/conf/reposource"
	"github.com/khulnasoft/khulnasoft/internal/extsvc/gomodproxy"
	"github.com/khulnasoft/khulnasoft/internal/httpcli"
	"github.com/khulnasoft/khulnasoft/internal/jsonc"
	"github.com/khulnasoft/khulnasoft/internal/types"
	"github.com/khulnasoft/khulnasoft/lib/errors"
	"github.com/khulnasoft/khulnasoft/schema"
)

// NewGoPackagesSource returns a new GoModulesSource from the given external service.
func NewGoPackagesSource(ctx context.Context, svc *types.ExternalService, cf *httpcli.Factory) (*PackagesSource, error) {
	rawConfig, err := svc.Config.Decrypt(ctx)
	if err != nil {
		return nil, errors.Errorf("external service id=%d config error: %s", svc.ID, err)
	}
	var c schema.GoModulesConnection
	if err := jsonc.Unmarshal(rawConfig, &c); err != nil {
		return nil, errors.Errorf("external service id=%d config error: %s", svc.ID, err)
	}

	return &PackagesSource{
		svc:        svc,
		configDeps: c.Dependencies,
		scheme:     dependencies.GoPackagesScheme,
		src: &goPackagesSource{
			client: gomodproxy.NewClient(svc.URN(), c.Urls, cf),
		},
	}, nil
}

type goPackagesSource struct {
	client *gomodproxy.Client
}

var _ packagesSource = &goPackagesSource{}

func (goPackagesSource) ParseVersionedPackageFromConfiguration(dep string) (reposource.VersionedPackage, error) {
	return reposource.ParseGoVersionedPackage(dep)
}

func (goPackagesSource) ParsePackageFromName(name reposource.PackageName) (reposource.Package, error) {
	return reposource.ParseGoDependencyFromName(name)
}

func (goPackagesSource) ParsePackageFromRepoName(repoName api.RepoName) (reposource.Package, error) {
	return reposource.ParseGoDependencyFromRepoName(repoName)
}
