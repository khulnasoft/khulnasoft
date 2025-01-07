package gitcli

import (
	"github.com/sourcegraph/log"

	"github.com/hashicorp/golang-lru/v2"

	"github.com/khulnasoft/khulnasoft/cmd/gitserver/internal/common"
	"github.com/khulnasoft/khulnasoft/cmd/gitserver/internal/git"
	"github.com/khulnasoft/khulnasoft/internal/api"
	"github.com/khulnasoft/khulnasoft/internal/wrexec"
)

func NewBackend(logger log.Logger, rcf *wrexec.RecordingCommandFactory, dir common.GitDir, repoName api.RepoName) git.GitBackend {
	return &gitCLIBackend{
		logger:         logger,
		rcf:            rcf,
		dir:            dir,
		repoName:       repoName,
		revAtTimeCache: globalRevAtTimeCache,
	}
}

type gitCLIBackend struct {
	logger         log.Logger
	rcf            *wrexec.RecordingCommandFactory
	dir            common.GitDir
	repoName       api.RepoName
	revAtTimeCache *lru.Cache[revAtTimeCacheKey, api.CommitID]
}
