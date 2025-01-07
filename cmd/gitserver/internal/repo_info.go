package internal

import (
	"github.com/khulnasoft/khulnasoft/cmd/gitserver/internal/gitserverfs"
	"github.com/khulnasoft/khulnasoft/internal/api"
	"github.com/khulnasoft/khulnasoft/internal/gitserver/protocol"
	"github.com/khulnasoft/khulnasoft/lib/errors"
)

func repoCloneProgress(fs gitserverfs.FS, locker RepositoryLocker, repo api.RepoName) (*protocol.RepoCloneProgress, error) {
	cloned, err := fs.RepoCloned(repo)
	if err != nil {
		return nil, errors.Wrap(err, "determine clone status")
	}

	resp := protocol.RepoCloneProgress{
		Cloned: cloned,
	}
	cloneProgress, locked := locker.Status(repo)
	if isAlwaysCloningTest(repo) {
		resp.CloneInProgress = true
		resp.CloneProgress = "This will never finish cloning"
	}
	if !cloned && locked {
		resp.CloneInProgress = true
		resp.CloneProgress = cloneProgress
	}
	return &resp, nil
}
