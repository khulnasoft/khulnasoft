package store

import (
	"context"
	"database/sql"
	"testing"

	"github.com/sourcegraph/log/logtest"

	bt "github.com/khulnasoft/khulnasoft/internal/batches/testing"
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/database/basestore"
	"github.com/khulnasoft/khulnasoft/internal/database/dbtest"
	"github.com/khulnasoft/khulnasoft/internal/encryption"
	"github.com/khulnasoft/khulnasoft/internal/observation"
	"github.com/khulnasoft/khulnasoft/internal/timeutil"
)

type storeTestFunc func(*testing.T, context.Context, *Store, bt.Clock)

// storeTest converts a storeTestFunc into a func(*testing.T) in which all
// dependencies are set up and injected into the storeTestFunc.
func storeTest(db *sql.DB, key encryption.Key, f storeTestFunc) func(*testing.T) {
	return func(t *testing.T) {
		logger := logtest.Scoped(t)
		c := &bt.TestClock{Time: timeutil.Now()}

		// Store tests all run in a transaction that's rolled back at the end
		// of the tests, so that foreign key constraints can be deferred and we
		// don't need to insert a lot of dependencies into the DB (users,
		// repos, ...) to setup the tests.
		tx := database.NewDBWith(logger, basestore.NewWithHandle(basestore.NewHandleWithTx(dbtest.NewTx(t, db), sql.TxOptions{})))
		s := NewWithClock(tx, observation.TestContextTB(t), key, c.Now)

		f(t, context.Background(), s, c)
	}
}
