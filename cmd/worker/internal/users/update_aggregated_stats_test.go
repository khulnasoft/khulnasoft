package users

import (
	"context"
	"testing"

	"github.com/sourcegraph/log/logtest"
	"github.com/stretchr/testify/require"

	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/database/dbtest"
)

func TestUpdateAggregatedUsersStatisticsTable(t *testing.T) {
	db := database.NewDB(logtest.NoOp(t), dbtest.NewDB(t))
	err := updateAggregatedUsersStatisticsTable(context.Background(), db)
	require.NoError(t, err)
}
