package inference

import (
	"context"
	"io"
	"io/fs"
	"sort"
	"testing"

	"golang.org/x/time/rate"

	"github.com/khulnasoft/khulnasoft/internal/api"
	"github.com/khulnasoft/khulnasoft/internal/fileutil"
	"github.com/khulnasoft/khulnasoft/internal/gitserver"
	"github.com/khulnasoft/khulnasoft/internal/luasandbox"
	"github.com/khulnasoft/khulnasoft/internal/observation"
	"github.com/khulnasoft/khulnasoft/internal/ratelimit"
	"github.com/khulnasoft/khulnasoft/internal/unpack/unpacktest"
)

func testService(t *testing.T, repositoryContents map[string]string) *Service {
	repositoryPaths := make([]string, 0, len(repositoryContents))
	for path := range repositoryContents {
		repositoryPaths = append(repositoryPaths, path)
	}
	sort.Strings(repositoryPaths)

	// Real deal
	sandboxService := luasandbox.NewService()

	// Fake deal
	gitService := NewMockGitService()
	gitService.ReadDirFunc.SetDefaultHook(func(ctx context.Context, _ api.RepoName, _ api.CommitID, path string, recurse bool) ([]fs.FileInfo, error) {
		var fds []fs.FileInfo
		for _, repositoryPath := range repositoryPaths {
			fds = append(fds, &fileutil.FileInfo{
				Name_: repositoryPath,
			})
		}
		return fds, nil
	})
	gitService.ArchiveFunc.SetDefaultHook(func(ctx context.Context, repoName api.RepoName, opts gitserver.ArchiveOptions) (io.ReadCloser, error) {
		files := map[string]string{}
		for _, path := range opts.Paths {
			if contents, ok := repositoryContents[path]; ok {
				files[path] = contents
			}
		}

		return unpacktest.CreateTarArchive(t, files), nil
	})

	return newService(observation.TestContextTB(t), sandboxService, gitService, ratelimit.NewInstrumentedLimiter("TestInference", rate.NewLimiter(rate.Limit(100), 1)), 100, 1024*1024)
}
