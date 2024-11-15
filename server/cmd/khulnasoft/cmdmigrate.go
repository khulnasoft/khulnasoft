// Copyright 2020 - KhulnaSoft Authors <admin@khulnasoft.com>
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"
	"fmt"

	"github.com/khulnasoft/khulnasoft/server/persistence"
	"github.com/khulnasoft/khulnasoft/server/persistence/relational"
)

var migrateUsage = `
"migrate" applies all pending database migrations to the connected database.
Only run this command when you run KhulnaSoft as a horizontally scaling service as
the default installation will handle this routine by itself.

Usage of "migrate":
`

func cmdMigrate(subcommand string, flags []string) {
	cmd := flag.NewFlagSet(subcommand, flag.ExitOnError)
	cmd.Usage = func() {
		fmt.Fprint(
			flag.CommandLine.Output(), migrateUsage)
		cmd.PrintDefaults()
	}
	var (
		envFile = cmd.String("envfile", "", "the env file to use")
	)
	cmd.Parse(flags)
	a := newApp(false, true, *envFile)

	gormDB, dbErr := newDB(a.config, a.logger)
	if dbErr != nil {
		a.logger.WithError(dbErr).Fatal("Error establishing database connection")
	}

	db, err := persistence.New(
		relational.NewRelationalDAL(gormDB),
	)
	if err != nil {
		a.logger.WithError(err).Fatal("Error creating persistence layer")
	}

	if err := db.Migrate(); err != nil {
		a.logger.WithError(err).Fatal("Error applying database migrations")
	}
	a.logger.Info("Successfully ran database migrations")
}
