package codenav

import (
	"github.com/sourcegraph/log"

	"github.com/khulnasoft/khulnasoft/internal/codeintel/codenav/internal/lsifstore"
	codeintelshared "github.com/khulnasoft/khulnasoft/internal/codeintel/shared"
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/gitserver"
	"github.com/khulnasoft/khulnasoft/internal/observation"
	searchClient "github.com/khulnasoft/khulnasoft/internal/search/client"
)

func NewService(
	observationCtx *observation.Context,
	db database.DB,
	codeIntelDB codeintelshared.CodeIntelDB,
	uploadSvc UploadService,
	gitserver gitserver.Client,
) *Service {
	lsifStore := lsifstore.New(scopedContext("lsifstore", observationCtx), codeIntelDB)
	logger := log.Scoped("codenav")
	searcher := searchClient.New(logger, db, gitserver)

	return newService(
		observationCtx,
		db.Repos(),
		lsifStore,
		uploadSvc,
		gitserver,
		searcher,
		logger,
	)
}

func scopedContext(component string, parent *observation.Context) *observation.Context {
	return observation.ScopedContext("codeintel", "codenav", component, parent)
}
