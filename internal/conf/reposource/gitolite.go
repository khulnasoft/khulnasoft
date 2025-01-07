package reposource

import (
	"strings"

	"github.com/khulnasoft/khulnasoft/internal/api"
	"github.com/khulnasoft/khulnasoft/schema"
)

type Gitolite struct {
	*schema.GitoliteConnection
}

var _ RepoSource = Gitolite{}

func (c Gitolite) CloneURLToRepoName(cloneURL string) (repoName api.RepoName, err error) {
	parsedCloneURL, err := parseCloneURL(cloneURL)
	if err != nil {
		return "", err
	}
	parsedHostURL, err := parseCloneURL(c.Host)
	if err != nil {
		return "", err
	}
	if parsedHostURL.Hostname() != parsedCloneURL.Hostname() {
		return "", nil
	}
	return GitoliteRepoName(c.Prefix, strings.TrimPrefix(strings.TrimSuffix(parsedCloneURL.Path, ".git"), "/")), nil
}

// GitoliteRepoName returns the Khulnasoft name for a repository given the Gitolite prefix (defined
// in the Gitolite external service config) and the Gitolite repository name. This is normally just
// the prefix concatenated with the Gitolite name. Gitolite permits the "@" character, but
// Khulnasoft does not, so "@" characters are rewritten to be "-".
func GitoliteRepoName(prefix, gitoliteName string) api.RepoName {
	gitoliteNameWithNoIllegalChars := strings.ReplaceAll(gitoliteName, "@", "-")
	return api.RepoName(strings.NewReplacer(
		"{prefix}", prefix,
		"{gitoliteName}", gitoliteNameWithNoIllegalChars,
	).Replace("{prefix}{gitoliteName}"))
}
