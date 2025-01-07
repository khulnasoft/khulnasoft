package background

import (
	repomatcher "github.com/khulnasoft/khulnasoft/internal/codeintel/policies/internal/background/repository_matcher"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/policies/internal/store"
	"github.com/khulnasoft/khulnasoft/internal/goroutine"
	"github.com/khulnasoft/khulnasoft/internal/observation"
)

func PolicyMatcherJobs(observationCtx *observation.Context, store store.Store, config *repomatcher.Config) []goroutine.BackgroundRoutine {
	return []goroutine.BackgroundRoutine{
		repomatcher.NewRepositoryMatcher(
			store,
			observationCtx,
			config.Interval,
			config.ConfigurationPolicyMembershipBatchSize,
		),
	}
}
