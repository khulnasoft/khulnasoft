package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/khulnasoft/khulnasoft/cmd/enterprise-portal/internal/dotcomdb"
	"github.com/khulnasoft/khulnasoft/lib/errors"
	"github.com/khulnasoft/khulnasoft/lib/managedservicesplatform/cloudsql"
)

func newDotComDBConn(ctx context.Context, config Config) (*dotcomdb.Reader, error) {
	readerOpts := dotcomdb.ReaderOptions{
		DevOnly: !config.DotComDB.IncludeProductionLicenses,
	}

	if override := config.DotComDB.PGDSNOverride; override != nil {
		db, err := pgxpool.New(ctx, *override)
		if err != nil {
			return nil, errors.Wrapf(err, "pgx.ConnectConfig %q", *override)
		}
		return dotcomdb.NewReader(db, readerOpts), nil
	}

	// Use IAM auth to connect to the Cloud SQL database.
	db, err := cloudsql.GetConnectionPool(ctx, config.DotComDB.ConnConfig)
	if err != nil {
		return nil, errors.Wrap(err, "contract.GetPostgreSQLDB")
	}
	return dotcomdb.NewReader(db, readerOpts), nil
}
