package reposource

import (
	"strings"

	"github.com/khulnasoft/khulnasoft/internal/api"
)

func PerforceRepoName(repositoryPathPattern, depot string) api.RepoName {
	if repositoryPathPattern == "" {
		repositoryPathPattern = "{depot}"
	}

	return api.RepoName(strings.NewReplacer(
		"{depot}", depot,
	).Replace(repositoryPathPattern))
}
