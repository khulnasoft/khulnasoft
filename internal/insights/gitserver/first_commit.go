package gitserver

import (
	"context"
	"sync"

	"github.com/khulnasoft/khulnasoft/internal/api"
	"github.com/khulnasoft/khulnasoft/internal/gitserver"
	"github.com/khulnasoft/khulnasoft/internal/gitserver/gitdomain"
	"github.com/khulnasoft/khulnasoft/lib/errors"
)

var (
	EmptyRepoErr = errors.New("empty repository")
)

func isFirstCommitEmptyRepoError(err error) bool {
	if gitdomain.IsRevisionNotFoundError(err) {
		return true
	}

	unwrappedErr := errors.Unwrap(err)
	if unwrappedErr != nil {
		return isFirstCommitEmptyRepoError(unwrappedErr)
	}
	return false
}

func GitFirstEverCommit(ctx context.Context, gitserverClient gitserver.Client, repoName api.RepoName) (*gitdomain.Commit, error) {
	return gitserverClient.FirstEverCommit(ctx, repoName)
}

func NewCachedGitFirstEverCommit() *CachedGitFirstEverCommit {
	return &CachedGitFirstEverCommit{
		impl: GitFirstEverCommit,
	}
}

// CachedGitFirstEverCommit is a simple in-memory cache for gitFirstEverCommit calls. It does so
// using a map, and entries are never evicted because they are expected to be small and in general
// unchanging.
type CachedGitFirstEverCommit struct {
	impl func(ctx context.Context, gitserverClient gitserver.Client, repoName api.RepoName) (*gitdomain.Commit, error)

	mu    sync.Mutex
	cache map[api.RepoName]*gitdomain.Commit
}

func (c *CachedGitFirstEverCommit) GitFirstEverCommit(ctx context.Context, gitserverClient gitserver.Client, repoName api.RepoName) (*gitdomain.Commit, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.cache == nil {
		c.cache = map[api.RepoName]*gitdomain.Commit{}
	}
	if cached, ok := c.cache[repoName]; ok {
		return cached, nil
	}
	entry, err := c.impl(ctx, gitserverClient, repoName)
	if err != nil {
		return nil, err
	}
	c.cache[repoName] = entry
	return entry, nil
}
