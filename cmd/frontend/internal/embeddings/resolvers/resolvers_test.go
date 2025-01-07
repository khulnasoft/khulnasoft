package resolvers

import (
	"bytes"
	"context"
	"io"
	"os"
	"testing"
	"time"

	"github.com/graph-gophers/graphql-go"
	"github.com/sourcegraph/log/logtest"
	"github.com/stretchr/testify/require"

	"github.com/khulnasoft/khulnasoft/cmd/frontend/graphqlbackend"
	"github.com/khulnasoft/khulnasoft/internal/actor"
	"github.com/khulnasoft/khulnasoft/internal/api"
	"github.com/khulnasoft/khulnasoft/internal/conf"
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/database/dbmocks"
	"github.com/khulnasoft/khulnasoft/internal/embeddings"
	"github.com/khulnasoft/khulnasoft/internal/embeddings/background/repo"
	"github.com/khulnasoft/khulnasoft/internal/featureflag"
	"github.com/khulnasoft/khulnasoft/internal/gitserver"
	"github.com/khulnasoft/khulnasoft/internal/licensing"
	rtypes "github.com/khulnasoft/khulnasoft/internal/rbac/types"
	"github.com/khulnasoft/khulnasoft/internal/types"
	"github.com/khulnasoft/khulnasoft/lib/pointers"
	"github.com/khulnasoft/khulnasoft/schema"
)

func TestEmbeddingSearchResolver(t *testing.T) {
	logger := logtest.Scoped(t)
	conf.MockForceAllowEmbeddings(t, true)

	oldMock := licensing.MockCheckFeature
	licensing.MockCheckFeature = func(feature licensing.Feature) error {
		return nil
	}
	t.Cleanup(func() {
		licensing.MockCheckFeature = oldMock
	})

	mockDB := dbmocks.NewMockDB()
	mockRepos := dbmocks.NewMockRepoStore()
	mockRepos.GetByIDsFunc.SetDefaultReturn([]*types.Repo{{ID: 1, Name: "repo1"}}, nil)
	mockDB.ReposFunc.SetDefaultReturn(mockRepos)
	mockUsers := dbmocks.NewMockUserStore()
	mockUsers.GetByCurrentAuthUserFunc.SetDefaultReturn(&types.User{ID: 1, SiteAdmin: true}, nil)
	mockDB.UsersFunc.SetDefaultReturn(mockUsers)

	type Perm struct {
		namespace rtypes.PermissionNamespace
		action    rtypes.NamespaceAction
	}
	defaultUserPerms := map[Perm]bool{
		{rtypes.CodyNamespace, rtypes.CodyAccessAction}: true, // Cody access
	}
	users := dbmocks.NewMockUserStore()
	users.GetByCurrentAuthUserFunc.SetDefaultHook(func(ctx context.Context) (*types.User, error) {
		return &types.User{ID: 1, SiteAdmin: true}, nil
	})
	mockDB.UsersFunc.SetDefaultReturn(users)
	permissions := dbmocks.NewMockPermissionStore()
	permissions.GetPermissionForUserFunc.SetDefaultHook(func(ctx context.Context, opt database.GetPermissionForUserOpts) (*types.Permission, error) {
		if hasPermission, ok := defaultUserPerms[Perm{opt.Namespace, opt.Action}]; ok && hasPermission {
			return &types.Permission{ID: 1, Namespace: opt.Namespace, Action: opt.Action, CreatedAt: time.Now()}, nil
		}
		return nil, nil
	})
	mockDB.PermissionsFunc.SetDefaultReturn(permissions)

	mockGitserver := gitserver.NewMockClient()
	mockGitserver.NewFileReaderFunc.SetDefaultHook(func(ctx context.Context, rn api.RepoName, ci api.CommitID, fileName string) (io.ReadCloser, error) {
		if fileName == "testfile" {
			return io.NopCloser(bytes.NewReader([]byte("test\nfirst\nfour\nlines\nplus\nsome\nmore"))), nil
		}
		return nil, os.ErrNotExist
	})

	mockEmbeddingsClient := embeddings.NewMockClient()
	mockEmbeddingsClient.SearchFunc.SetDefaultReturn(&embeddings.EmbeddingCombinedSearchResults{
		CodeResults: embeddings.EmbeddingSearchResults{{
			FileName:  "testfile",
			StartLine: 0,
			EndLine:   4,
		}, {
			FileName:  "censored",
			StartLine: 0,
			EndLine:   4,
		}},
	}, nil)

	repoEmbeddingJobsStore := repo.NewMockRepoEmbeddingJobsStore()

	resolver := NewResolver(
		mockDB,
		logger,
		mockGitserver,
		mockEmbeddingsClient,
		repoEmbeddingJobsStore,
	)

	conf.Mock(&conf.Unified{
		SiteConfiguration: schema.SiteConfiguration{
			CodyEnabled: pointers.Ptr(true),
			LicenseKey:  "asdf",
			Embeddings:  &schema.Embeddings{},
		},
	})

	ctx := actor.WithActor(context.Background(), actor.FromMockUser(1))
	ffs := featureflag.NewMemoryStore(map[string]bool{"cody": true}, nil, nil)
	ctx = featureflag.WithFlags(ctx, ffs)

	results, err := resolver.EmbeddingsMultiSearch(ctx, graphqlbackend.EmbeddingsMultiSearchInputArgs{
		Repos:            []graphql.ID{graphqlbackend.MarshalRepositoryID(3)},
		Query:            "test",
		CodeResultsCount: 1,
		TextResultsCount: 1,
	})
	require.NoError(t, err)

	codeResults, err := results.CodeResults(ctx)
	require.NoError(t, err)
	require.Len(t, codeResults, 1)
	require.Equal(t, "test\nfirst\nfour\nlines", codeResults[0].Content(ctx))
}

func Test_extractLineRange(t *testing.T) {
	cases := []struct {
		input      []byte
		start, end int
		output     []byte
	}{{
		[]byte("zero\none\ntwo\nthree\n"),
		0, 2,
		[]byte("zero\none"),
	}, {
		[]byte("zero\none\ntwo\nthree\n"),
		1, 2,
		[]byte("one"),
	}, {
		[]byte("zero\none\ntwo\nthree\n"),
		1, 2,
		[]byte("one"),
	}, {
		[]byte(""),
		1, 2,
		[]byte(""),
	}}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := extractLineRange(tc.input, tc.start, tc.end)
			require.Equal(t, tc.output, got)
		})
	}
}
