package githubapps

import (
	"context"
	"time"

	"github.com/sourcegraph/log"

	"github.com/khulnasoft/khulnasoft/cmd/worker/internal/githubapps/worker"
	"github.com/khulnasoft/khulnasoft/cmd/worker/job"
	workerdb "github.com/khulnasoft/khulnasoft/cmd/worker/shared/init/db"
	"github.com/khulnasoft/khulnasoft/internal/env"
	"github.com/khulnasoft/khulnasoft/internal/goroutine"
	"github.com/khulnasoft/khulnasoft/internal/observation"
	"github.com/khulnasoft/khulnasoft/lib/errors"
)

type githubAppsInstallationJob struct{}

func NewGitHubApsInstallationJob() job.Job {
	return &githubAppsInstallationJob{}
}

func (gh *githubAppsInstallationJob) Description() string {
	return "Job to validate and backfill github app installations"
}

func (gh *githubAppsInstallationJob) Config() []env.Config {
	return nil
}

func (gh *githubAppsInstallationJob) Routines(ctx context.Context, observationCtx *observation.Context) ([]goroutine.BackgroundRoutine, error) {
	db, err := workerdb.InitDB(observationCtx)
	if err != nil {
		return nil, errors.Wrap(err, "init DB")
	}

	logger := log.Scoped("github_apps_installation")
	return []goroutine.BackgroundRoutine{
		goroutine.NewPeriodicGoroutine(
			context.Background(),
			worker.NewGitHubInstallationWorker(db, logger),
			goroutine.WithName("github_apps.installation_backfill"),
			goroutine.WithDescription("backfills github apps installation ids and removes deleted github app installations"),
			goroutine.WithInterval(24*time.Hour),
		),
	}, nil
}
