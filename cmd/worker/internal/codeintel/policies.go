package codeintel

import (
	"context"

	"github.com/khulnasoft/khulnasoft/cmd/worker/job"
	"github.com/khulnasoft/khulnasoft/cmd/worker/shared/init/codeintel"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/policies"
	"github.com/khulnasoft/khulnasoft/internal/env"
	"github.com/khulnasoft/khulnasoft/internal/goroutine"
	"github.com/khulnasoft/khulnasoft/internal/observation"
)

type policiesRepositoryMatcherJob struct{}

func NewPoliciesRepositoryMatcherJob() job.Job {
	return &policiesRepositoryMatcherJob{}
}

func (j *policiesRepositoryMatcherJob) Description() string {
	return "code-intel policies repository matcher"
}

func (j *policiesRepositoryMatcherJob) Config() []env.Config {
	return []env.Config{
		policies.RepositoryMatcherConfigInst,
	}
}

func (j *policiesRepositoryMatcherJob) Routines(_ context.Context, observationCtx *observation.Context) ([]goroutine.BackgroundRoutine, error) {
	services, err := codeintel.InitServices(observationCtx)
	if err != nil {
		return nil, err
	}

	// TODO(nsc): https://github.com/khulnasoft/khulnasoft/pull/42765
	return policies.NewRepositoryMatcherRoutines(observationCtx, services.PoliciesService), nil
}
