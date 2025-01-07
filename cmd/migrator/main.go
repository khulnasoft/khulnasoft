package main

import (
	"github.com/sourcegraph/log"

	"github.com/khulnasoft/khulnasoft/cmd/migrator/shared"
	"github.com/khulnasoft/khulnasoft/internal/env"
	"github.com/khulnasoft/khulnasoft/internal/oobmigration/migrations/register"
	"github.com/khulnasoft/khulnasoft/internal/sanitycheck"
	"github.com/khulnasoft/khulnasoft/internal/version"
)

func main() {
	sanitycheck.Pass()
	liblog := log.Init(log.Resource{
		Name:    env.MyName,
		Version: version.Version(),
	})
	defer liblog.Sync()

	logger := log.Scoped("migrator")

	if err := shared.Start(logger, register.RegisterEnterpriseMigratorsUsingConfAndStoreFactory); err != nil {
		logger.Fatal(err.Error())
	}
}
