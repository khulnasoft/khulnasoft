package batches

import (
	"context"

	sglog "github.com/sourcegraph/log"

	"github.com/khulnasoft/khulnasoft/cmd/frontend/enterprise"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/batches/httpapi"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/batches/resolvers"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/batches/webhooks"
	"github.com/khulnasoft/khulnasoft/internal/batches"
	"github.com/khulnasoft/khulnasoft/internal/batches/store"
	"github.com/khulnasoft/khulnasoft/internal/codeintel"
	"github.com/khulnasoft/khulnasoft/internal/conf/conftypes"
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/encryption/keyring"
	"github.com/khulnasoft/khulnasoft/internal/gitserver"
	"github.com/khulnasoft/khulnasoft/internal/observation"
)

// Init initializes the given enterpriseServices to include the required
// resolvers for Batch Changes and sets up webhook handlers for changeset
// events.
func Init(
	ctx context.Context,
	observationCtx *observation.Context,
	db database.DB,
	_ codeintel.Services,
	_ conftypes.UnifiedWatchable,
	enterpriseServices *enterprise.Services,
) error {
	if !batches.IsEnabled() {
		return nil
	}

	// Initialize store.
	bstore := store.New(db, observationCtx, keyring.Default().BatchChangesCredentialKey)

	// Register enterprise services.
	logger := sglog.Scoped("Batches")
	enterpriseServices.BatchChangesResolver = resolvers.New(db, bstore, gitserver.NewClient("graphql.batches"), logger)
	gitserverClient := gitserver.NewClient("http.batches.webhook")
	enterpriseServices.BatchesGitHubWebhook = webhooks.NewGitHubWebhook(bstore, gitserverClient.Scoped("github"), logger)
	enterpriseServices.BatchesBitbucketServerWebhook = webhooks.NewBitbucketServerWebhook(bstore, gitserverClient.Scoped("bitbucketserver"), logger)
	enterpriseServices.BatchesBitbucketCloudWebhook = webhooks.NewBitbucketCloudWebhook(bstore, gitserverClient.Scoped("bitbucketcloud"), logger)
	enterpriseServices.BatchesGitLabWebhook = webhooks.NewGitLabWebhook(bstore, gitserverClient.Scoped("gitlab"), logger)
	enterpriseServices.BatchesAzureDevOpsWebhook = webhooks.NewAzureDevOpsWebhook(bstore, gitserverClient.Scoped("azure"), logger)

	operations := httpapi.NewOperations(observationCtx)
	fileHandler := httpapi.NewFileHandler(db, bstore, operations)
	enterpriseServices.BatchesChangesFileGetHandler = fileHandler.Get()
	enterpriseServices.BatchesChangesFileExistsHandler = fileHandler.Exists()
	enterpriseServices.BatchesChangesFileUploadHandler = fileHandler.Upload()

	return nil
}
