package repos

import (
	"context"
	"testing"

	mockrequire "github.com/derision-test/go-mockgen/v2/testutil/require"

	"github.com/khulnasoft/khulnasoft/internal/extsvc"
	"github.com/khulnasoft/khulnasoft/internal/gitserver"
	"github.com/khulnasoft/khulnasoft/internal/types/typestest"
	"github.com/khulnasoft/khulnasoft/schema"
)

func TestGitoliteSource(t *testing.T) {
	gc := gitserver.NewMockClient()
	gc.ScopedFunc.SetDefaultReturn(gc)
	svc := typestest.MakeExternalService(t, extsvc.VariantGitolite, &schema.GitoliteConnection{})

	ctx := context.Background()
	s, err := NewGitoliteSource(ctx, svc, gc)
	if err != nil {
		t.Fatal(err)
	}

	res := make(chan SourceResult)
	go func() {
		s.ListRepos(ctx, res)
		close(res)
	}()

	for range res {
	}

	mockrequire.Called(t, gc.ListGitoliteReposFunc)
}
