package migrations

import (
	"context"
	"os"

	"github.com/khulnasoft/khulnasoft/cmd/worker/job"
	workerdb "github.com/khulnasoft/khulnasoft/cmd/worker/shared/init/db"
	"github.com/khulnasoft/khulnasoft/internal/env"
	"github.com/khulnasoft/khulnasoft/internal/goroutine"
	"github.com/khulnasoft/khulnasoft/internal/observation"
	"github.com/khulnasoft/khulnasoft/internal/oobmigration"
	"github.com/khulnasoft/khulnasoft/internal/version/upgradestore"
	"github.com/khulnasoft/khulnasoft/lib/errors"
)

// migrator configures an out of band migration runner process to execute in the background.
type migrator struct {
	registerMigrators oobmigration.RegisterMigratorsFunc
}

var _ job.Job = &migrator{}

func NewMigrator(registerMigrators oobmigration.RegisterMigratorsFunc) job.Job {
	return &migrator{registerMigrators}
}

func (m *migrator) Description() string {
	return ""
}

func (m *migrator) Config() []env.Config {
	return nil
}

func (m *migrator) Routines(startupCtx context.Context, observationCtx *observation.Context) ([]goroutine.BackgroundRoutine, error) {
	db, err := workerdb.InitDB(observationCtx)
	if err != nil {
		return nil, err
	}

	outOfBandMigrationRunner := oobmigration.NewRunnerWithDB(observationCtx, db, oobmigration.RefreshInterval)

	if err := outOfBandMigrationRunner.SynchronizeMetadata(startupCtx); err != nil {
		return nil, errors.Wrap(err, "failed to synchronize out of band migration metadata")
	}

	if err := m.registerMigrators(startupCtx, db, outOfBandMigrationRunner); err != nil {
		return nil, err
	}

	if os.Getenv("SRC_DISABLE_OOBMIGRATION_VALIDATION") != "" {
		observationCtx.Logger.Warn("Skipping out-of-band migrations check")
	} else {
		if err := oobmigration.ValidateOutOfBandMigrationRunner(startupCtx, db, outOfBandMigrationRunner); err != nil {
			return nil, err
		}
	}

	currentVersion, err := currentVersion(observationCtx.Logger)
	if err != nil {
		return nil, err
	}

	firstVersionString, _, err := upgradestore.New(db).GetFirstServiceVersion(startupCtx)
	if err != nil {
		return nil, err
	}

	firstVersion, ok := oobmigration.NewVersionFromString(firstVersionString)
	if !ok {
		return nil, err
	}

	return []goroutine.BackgroundRoutine{
		&outOfBandMigrationRunnerWrapper{
			Runner:         outOfBandMigrationRunner,
			currentVersion: currentVersion,
			firstVersion:   firstVersion,
		},
	}, nil
}

type outOfBandMigrationRunnerWrapper struct {
	*oobmigration.Runner
	currentVersion oobmigration.Version
	firstVersion   oobmigration.Version
}

func (w *outOfBandMigrationRunnerWrapper) Start() {
	w.Runner.Start(w.currentVersion, w.firstVersion)
}
