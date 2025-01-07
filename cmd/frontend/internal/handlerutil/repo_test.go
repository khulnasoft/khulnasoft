package handlerutil

import (
	"context"
	"testing"

	"github.com/sourcegraph/log/logtest"

	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/backend"
	"github.com/khulnasoft/khulnasoft/internal/api"
	"github.com/khulnasoft/khulnasoft/internal/database/dbmocks"
	"github.com/khulnasoft/khulnasoft/internal/types"
	"github.com/khulnasoft/khulnasoft/lib/errors"
)

func TestGetRepo(t *testing.T) {
	logger := logtest.Scoped(t)
	t.Run("URLMovedError", func(t *testing.T) {
		backend.Mocks.Repos.GetByName = func(ctx context.Context, name api.RepoName) (*types.Repo, error) {
			return &types.Repo{Name: name + name}, nil
		}
		t.Cleanup(func() {
			backend.Mocks.Repos = backend.MockRepos{}
		})

		_, err := GetRepo(context.Background(), logger, dbmocks.NewMockDB(), map[string]string{"Repo": "repo1"})
		if !errors.HasType[*URLMovedError](err) {
			t.Fatalf("err: want type *URLMovedError but got %T", err)
		}
	})
}
