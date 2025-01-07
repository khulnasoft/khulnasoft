package executorqueue

import (
	"github.com/sourcegraph/log"

	"github.com/khulnasoft/khulnasoft/internal/conf/conftypes"
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/observation"

	"github.com/khulnasoft/khulnasoft/cmd/frontend/enterprise"
)

// Init initializes the executor endpoints required for use with the executor service.
func Init(
	observationCtx *observation.Context,
	db database.DB,
	conf conftypes.UnifiedWatchable,
	enterpriseServices *enterprise.Services,
) error {
	codeintelUploadHandler := enterpriseServices.NewCodeIntelUploadHandler(false)
	batchesWorkspaceFileGetHandler := enterpriseServices.BatchesChangesFileGetHandler
	batchesWorkspaceFileExistsHandler := enterpriseServices.BatchesChangesFileGetHandler

	accessToken := func() string {
		return conf.SiteConfig().ExecutorsAccessToken
	}

	logger := log.Scoped("executorqueue")

	queueHandler := newExecutorQueuesHandler(
		observationCtx,
		db,
		logger,
		accessToken,
		codeintelUploadHandler,
		batchesWorkspaceFileGetHandler,
		batchesWorkspaceFileExistsHandler,
	)

	enterpriseServices.NewExecutorProxyHandler = queueHandler
	return nil
}
