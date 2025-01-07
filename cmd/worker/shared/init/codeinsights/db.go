package codeintel

import (
	"database/sql"

	"github.com/khulnasoft/khulnasoft/internal/conf"
	"github.com/khulnasoft/khulnasoft/internal/conf/conftypes"
	"github.com/khulnasoft/khulnasoft/internal/database"
	connections "github.com/khulnasoft/khulnasoft/internal/database/connections/live"
	"github.com/khulnasoft/khulnasoft/internal/memo"
	"github.com/khulnasoft/khulnasoft/internal/observation"
	"github.com/khulnasoft/khulnasoft/lib/errors"
)

// InitRawDB initializes and returns a connection to the codeinsights db.
func InitRawDB(observationCtx *observation.Context) (*sql.DB, error) {
	return initDBMemo.Init(observationCtx)
}

func InitDB(observationCtx *observation.Context) (database.InsightsDB, error) {
	rawDB, err := InitRawDB(observationCtx)
	if err != nil {
		return nil, err
	}

	return database.NewInsightsDB(rawDB, observationCtx.Logger), nil
}

var initDBMemo = memo.NewMemoizedConstructorWithArg(func(observationCtx *observation.Context) (*sql.DB, error) {
	dsn := conf.GetServiceConnectionValueAndRestartOnChange(func(serviceConnections conftypes.ServiceConnections) string {
		return serviceConnections.CodeInsightsDSN
	})

	db, err := connections.EnsureNewCodeInsightsDB(observationCtx, dsn, "worker")
	if err != nil {
		return nil, errors.Errorf("failed to connect to codeinsights database: %s", err)
	}

	return db, nil
})
