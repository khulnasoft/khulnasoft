package scheduler

import (
	"context"
	"time"

	"github.com/khulnasoft/khulnasoft/internal/api"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/dependencies"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/policies"
	policiesshared "github.com/khulnasoft/khulnasoft/internal/codeintel/policies/shared"
	uploadsshared "github.com/khulnasoft/khulnasoft/internal/codeintel/uploads/shared"
)

type PolicyMatcher interface {
	CommitsDescribedByPolicy(ctx context.Context, repositoryID int, repoName api.RepoName, policies []policiesshared.ConfigurationPolicy, now time.Time, filterCommits ...string) (map[string][]policies.PolicyMatch, error)
}

type PoliciesService interface {
	GetConfigurationPolicies(ctx context.Context, opts policiesshared.GetConfigurationPoliciesOptions) ([]policiesshared.ConfigurationPolicy, int, error)
}

type IndexEnqueuer interface {
	QueueAutoIndexJobs(ctx context.Context, repositoryID int, rev, configuration string, force, bypassLimit bool) (_ []uploadsshared.AutoIndexJob, err error)
	QueueAutoIndexJobsForPackage(ctx context.Context, pkg dependencies.MinimialVersionedPackageRepo) (err error)
}
