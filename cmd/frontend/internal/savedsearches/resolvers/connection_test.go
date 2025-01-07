package resolvers

import (
	"context"
	"fmt"
	"testing"

	"github.com/sourcegraph/log/logtest"

	"github.com/khulnasoft/khulnasoft/internal/actor"
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/database/dbtest"
	"github.com/khulnasoft/khulnasoft/internal/gqlutil"
	"github.com/khulnasoft/khulnasoft/internal/types"
	"github.com/khulnasoft/khulnasoft/lib/pointers"
)

func TestSavedSearchesConnectionStore(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	ctx := context.Background()
	db := database.NewDB(logtest.Scoped(t), dbtest.NewDB(t))

	user, err := db.Users().Create(ctx, database.NewUser{
		Email:           "test@khulnasoft.com",
		Username:        "test",
		EmailIsVerified: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	ctx = actor.WithActor(ctx, &actor.Actor{UID: user.ID})

	for i := range 10 {
		created, err := db.SavedSearches().Create(ctx, &types.SavedSearch{
			Description: fmt.Sprintf("Test Search %d", i),
			Query:       "r:src-cli",
			Owner:       types.NamespaceUser(user.ID),
		})
		if err != nil {
			t.Fatal(err)
		}

		// Adjust so each one has a different updated_at value (which is rounded to the second).
		if _, err := db.ExecContext(ctx, `UPDATE saved_searches SET created_at = '2024-07-04 12:34:56.123456', updated_at = '2024-07-05 19:46:03.515814'::timestamp with time zone - (INTERVAL '100 milliseconds' * $1) WHERE id = $2`, i, created.ID); err != nil {
			t.Fatal(err)
		}
	}

	owner := types.NamespaceUser(user.ID)
	connectionStore := &savedSearchesConnectionStore{
		db:       db,
		listArgs: database.SavedSearchListArgs{Owner: &owner},
	}

	t.Run("no orderBy", func(t *testing.T) {
		gqlutil.TestConnectionResolverStoreSuite(t, connectionStore, nil)
	})

	t.Run("orderBy updated_at", func(t *testing.T) {
		var pgArgs gqlutil.TestPaginationArgs
		pgArgs.OrderBy, pgArgs.Ascending = database.SavedSearchesOrderByUpdatedAt.ToOptions()
		gqlutil.TestConnectionResolverStoreSuite(t, connectionStore, &pgArgs)
	})

	t.Run("orderBy description", func(t *testing.T) {
		var pgArgs gqlutil.TestPaginationArgs
		pgArgs.OrderBy, pgArgs.Ascending = database.SavedSearchesOrderByDescription.ToOptions()
		gqlutil.TestConnectionResolverStoreSuite(t, connectionStore, &pgArgs)
	})
}

var dummyConnectionResolverArgs = gqlutil.ConnectionResolverArgs{First: pointers.Ptr[int32](1)}
