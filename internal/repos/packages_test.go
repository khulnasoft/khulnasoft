package repos

import (
	"context"
	"testing"

	"github.com/khulnasoft/khulnasoft/internal/api"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/dependencies"
	"github.com/khulnasoft/khulnasoft/internal/conf/reposource"
	"github.com/khulnasoft/khulnasoft/internal/extsvc"
	"github.com/khulnasoft/khulnasoft/internal/types"
)

func TestPackagesSource_GetRepo(t *testing.T) {
	ctx := context.Background()
	svc := testDependenciesService(ctx, t, []dependencies.MinimalPackageRepoRef{
		{
			Scheme:   "go",
			Name:     "github.com/sourcegraph-testing/go-repo-a",
			Versions: []dependencies.MinimalPackageRepoRefVersion{{Version: "1.0.0"}},
		},
	})

	dummySrc := &dummyPackagesSource{}
	src := &PackagesSource{src: dummySrc, svc: &types.ExternalService{
		ID:     1,
		Kind:   extsvc.KindGoPackages,
		Config: extsvc.NewEmptyConfig(),
	}, depsSvc: svc}

	src.GetRepo(ctx, "go/github.com/sourcegraph-testing/go-repo-a")

	if !dummySrc.parsePackageFromRepoNameCalled {
		t.Fatalf("expected ParsePackageFromRepoName to be called, was not")
	}

	// Flip the condition below after https://github.com/khulnasoft/khulnasoft/issues/39653 has been fixed.
	if dummySrc.getPackageCalled {
		t.Fatalf("expected GetPackage to not be called, but it was called")
	}
}

var _ packagesSource = &dummyPackagesSource{}

// dummyPackagesSource is a tiny shim around Go-specific methods to track when they're called.
type dummyPackagesSource struct {
	parseVersionedPackageFromConfiguration bool
	parsePackageFromRepoNameCalled         bool
	parsePackageFromNameCalled             bool
	getPackageCalled                       bool
}

// GetPackage implements packagesDownloadSource
func (d *dummyPackagesSource) GetPackage(ctx context.Context, name reposource.PackageName) (reposource.Package, error) {
	d.getPackageCalled = true
	return reposource.ParseGoDependencyFromName(name)
}

func (d *dummyPackagesSource) ParseVersionedPackageFromConfiguration(dep string) (reposource.VersionedPackage, error) {
	d.parseVersionedPackageFromConfiguration = true
	return reposource.ParseGoVersionedPackage(dep)
}

func (d *dummyPackagesSource) ParsePackageFromName(name reposource.PackageName) (reposource.Package, error) {
	d.parsePackageFromNameCalled = true
	return reposource.ParseGoDependencyFromName(name)
}

func (d *dummyPackagesSource) ParsePackageFromRepoName(repoName api.RepoName) (reposource.Package, error) {
	d.parsePackageFromRepoNameCalled = true
	return reposource.ParseGoDependencyFromRepoName(repoName)
}
