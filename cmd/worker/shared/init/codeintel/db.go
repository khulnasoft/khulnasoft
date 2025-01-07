package codeintel

import (
	"database/sql"

	codeintelshared "github.com/khulnasoft/khulnasoft/internal/codeintel/shared"
	"github.com/khulnasoft/khulnasoft/internal/conf"
	"github.com/khulnasoft/khulnasoft/internal/conf/conftypes"
	connections "github.com/khulnasoft/khulnasoft/internal/database/connections/live"
	"github.com/khulnasoft/khulnasoft/internal/memo"
	"github.com/khulnasoft/khulnasoft/internal/observation"
	"github.com/khulnasoft/khulnasoft/lib/errors"
)

func InitDB(observationCtx *observation.Context) (codeintelshared.CodeIntelDB, error) {
	rawDB, err := initDBMemo.Init(observationCtx)
	if err != nil {
		return nil, err
	}

	return codeintelshared.NewCodeIntelDB(observationCtx.Logger, rawDB), nil
}

var initDBMemo = memo.NewMemoizedConstructorWithArg(func(observationCtx *observation.Context) (*sql.DB, error) {
	dsn := conf.GetServiceConnectionValueAndRestartOnChange(func(serviceConnections conftypes.ServiceConnections) string {
		return serviceConnections.CodeIntelPostgresDSN
	})

	db, err := connections.EnsureNewCodeIntelDB(observationCtx, dsn, "worker")
	if err != nil {
		return nil, errors.Errorf("failed to connect to codeintel database: %s", err)
	}

	return db, nil
})
