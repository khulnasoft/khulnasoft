package migration

import (
	"database/sql"
	"strings"

	connections "github.com/khulnasoft/khulnasoft/internal/database/connections/live"
	"github.com/khulnasoft/khulnasoft/internal/database/migration/runner"
	"github.com/khulnasoft/khulnasoft/internal/database/migration/schemas"
	"github.com/khulnasoft/khulnasoft/internal/database/migration/store"
	"github.com/khulnasoft/khulnasoft/internal/database/postgresdsn"
	"github.com/khulnasoft/khulnasoft/internal/env"
	"github.com/khulnasoft/khulnasoft/internal/observation"
	"github.com/khulnasoft/khulnasoft/lib/output"
)

func NewRunnerWithSchemas(
	observationCtx *observation.Context,
	out *output.Output,
	appName string,
	schemaNames []string,
	schemas []*schemas.Schema,
) (*runner.Runner, error) {
	dsns, err := postgresdsn.DSNsBySchema(schemaNames)
	if err != nil {
		return nil, err
	}
	var verbose = env.LogLevel == "dbug"

	var dsnsStrings []string
	for schema, dsn := range dsns {
		dsnsStrings = append(dsnsStrings, schema+" => "+dsn)
	}
	if verbose {
		out.WriteLine(output.Linef(output.EmojiInfo, output.StyleGrey, " Connection DSNs used: %s", strings.Join(dsnsStrings, ", ")))
	}

	storeFactory := func(db *sql.DB, migrationsTable string) connections.Store {
		return connections.NewStoreShim(store.NewWithDB(observationCtx, db, migrationsTable))
	}
	r, err := connections.RunnerFromDSNsWithSchemas(out, observationCtx.Logger, dsns, appName, storeFactory, schemas)
	if err != nil {
		return nil, err
	}

	return r, nil
}
