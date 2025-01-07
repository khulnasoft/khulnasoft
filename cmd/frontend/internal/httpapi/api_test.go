package httpapi

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/throttled/throttled/v2/store/memstore"

	"github.com/sourcegraph/log/logtest"

	"github.com/khulnasoft/khulnasoft/cmd/frontend/enterprise"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/graphqlbackend"
	"github.com/khulnasoft/khulnasoft/internal/conf"
	"github.com/khulnasoft/khulnasoft/internal/database/dbmocks"
	"github.com/khulnasoft/khulnasoft/internal/httptestutil"
	"github.com/khulnasoft/khulnasoft/internal/txemail"
)

func init() {
	txemail.DisableSilently()
}

func newTest(t *testing.T) *httptestutil.Client {
	conf.Mock(&conf.Unified{})
	t.Cleanup(func() {
		conf.Mock(nil)
	})

	logger := logtest.Scoped(t)
	enterpriseServices := enterprise.DefaultServices()
	rateLimitStore, _ := memstore.NewCtx(1024)
	rateLimiter := graphqlbackend.NewBasicLimitWatcher(logger, rateLimitStore)

	db := dbmocks.NewMockDB()

	handler, err := NewHandler(db,
		nil,
		rateLimiter,
		&Handlers{
			BatchesGitHubWebhook:            enterpriseServices.BatchesGitHubWebhook,
			BatchesGitLabWebhook:            enterpriseServices.BatchesGitLabWebhook,
			GitHubSyncWebhook:               enterpriseServices.ReposGithubWebhook,
			GitLabSyncWebhook:               enterpriseServices.ReposGitLabWebhook,
			BitbucketServerSyncWebhook:      enterpriseServices.ReposBitbucketServerWebhook,
			BitbucketCloudSyncWebhook:       enterpriseServices.ReposBitbucketCloudWebhook,
			BatchesBitbucketServerWebhook:   enterpriseServices.BatchesBitbucketServerWebhook,
			BatchesBitbucketCloudWebhook:    enterpriseServices.BatchesBitbucketCloudWebhook,
			BatchesAzureDevOpsWebhook:       enterpriseServices.BatchesAzureDevOpsWebhook,
			SCIMHandler:                     enterpriseServices.SCIMHandler,
			NewCodeIntelUploadHandler:       enterpriseServices.NewCodeIntelUploadHandler,
			NewComputeStreamHandler:         enterpriseServices.NewComputeStreamHandler,
			PermissionsGitHubWebhook:        enterpriseServices.PermissionsGitHubWebhook,
			NewChatCompletionsStreamHandler: enterpriseServices.NewChatCompletionsStreamHandler,
			NewCodeCompletionsHandler:       enterpriseServices.NewCodeCompletionsHandler,
		},
	)
	require.NoError(t, err)
	return httptestutil.NewTest(handler)
}
