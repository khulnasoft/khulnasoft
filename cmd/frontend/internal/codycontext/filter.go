package codycontext

import (
	"context"

	"github.com/sourcegraph/log"

	"github.com/khulnasoft/khulnasoft/internal/api"
	"github.com/khulnasoft/khulnasoft/internal/dotcom"
	"github.com/khulnasoft/khulnasoft/internal/gitserver"
	"github.com/khulnasoft/khulnasoft/internal/types"
)

// fileMatcher returns true if the given repo and path are allowed to be returned. It is used to filter out
// file matches that don't satisfy Cody ignore rules
type fileMatcher func(repoID api.RepoID, path string) bool
type repoContentFilter interface {
	// getMatcher returns a matcher to filter out file matches, and returns a filtered down list of
	// repositories containing only the ones that are allowed to be searched.
	getMatcher(ctx context.Context, repos []types.RepoIDName) ([]types.RepoIDName, fileMatcher, error)
}

func newRepoContentFilter(logger log.Logger, client gitserver.Client) repoContentFilter {
	if dotcom.KhulnasoftDotComMode() {
		return newDotcomFilter(logger, client)
	}
	return newEnterpriseFilter(logger)
}
