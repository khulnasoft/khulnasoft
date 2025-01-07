package symbol

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/khulnasoft/khulnasoft/internal/actor"
	"github.com/khulnasoft/khulnasoft/internal/api"
	"github.com/khulnasoft/khulnasoft/internal/authz"
	srp "github.com/khulnasoft/khulnasoft/internal/authz/subrepoperms"
	"github.com/khulnasoft/khulnasoft/internal/conf"
	"github.com/khulnasoft/khulnasoft/internal/search/result"
	"github.com/khulnasoft/khulnasoft/internal/types"
	"github.com/khulnasoft/khulnasoft/lib/errors"
	"github.com/khulnasoft/khulnasoft/schema"
)

func TestSearchZoektDoesntPanicWithNilQuery(t *testing.T) {
	// As soon as we reach Streamer.Search function, we can consider test successful,
	// that's why we can just mock it.
	mockStreamer := NewMockStreamer()
	expectedErr := errors.New("short circuit")
	mockStreamer.SearchFunc.SetDefaultReturn(nil, expectedErr)

	_, err := searchZoekt(context.Background(), mockStreamer, types.MinimalRepo{ID: 1}, "commitID", nil, "branch", nil, nil, nil)
	assert.ErrorIs(t, err, expectedErr)
}

func TestFilterZoektResults(t *testing.T) {
	conf.Mock(&conf.Unified{
		SiteConfiguration: schema.SiteConfiguration{
			ExperimentalFeatures: &schema.ExperimentalFeatures{
				SubRepoPermissions: &schema.SubRepoPermissions{
					Enabled: true,
				},
			},
		},
	})
	t.Cleanup(func() { conf.Mock(nil) })

	repoName := api.RepoName("foo")
	ctx := context.Background()
	ctx = actor.WithActor(ctx, &actor.Actor{
		UID: 1,
	})
	checker := srp.NewSimpleChecker(repoName, []authz.PathWithIP{
		{
			Path: "/**",
			IP:   "*",
		},
		{
			Path: "-/*_test.go",
			IP:   "*",
		},
	})

	results := []*result.SymbolMatch{
		{
			Symbol: result.Symbol{},
			File: &result.File{
				Path: "foo.go",
			},
		},
		{
			Symbol: result.Symbol{},
			File: &result.File{
				Path: "foo_test.go",
			},
		},
	}
	filtered, err := filterZoektResults(ctx, checker, repoName, results)
	if err != nil {
		t.Fatal(err)
	}
	assert.Len(t, filtered, 1)
	r := filtered[0]
	assert.Equal(t, r.File.Path, "foo.go")
}
