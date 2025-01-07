package policies

import (
	"github.com/khulnasoft/khulnasoft/internal/codeintel/policies/internal/background"
	repomatcher "github.com/khulnasoft/khulnasoft/internal/codeintel/policies/internal/background/repository_matcher"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/policies/internal/store"
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/gitserver"
	"github.com/khulnasoft/khulnasoft/internal/goroutine"
	"github.com/khulnasoft/khulnasoft/internal/observation"
)

func NewService(
	observationCtx *observation.Context,
	db database.DB,
	uploadSvc UploadService,
	gitserverClient gitserver.Client,
) *Service {
	return newService(
		scopedContext("service", observationCtx),
		store.New(scopedContext("store", observationCtx), db),
		db.Repos(),
		uploadSvc,
		gitserverClient,
	)
}

var RepositoryMatcherConfigInst = &repomatcher.Config{}

func NewRepositoryMatcherRoutines(observationCtx *observation.Context, service *Service) []goroutine.BackgroundRoutine {
	return background.PolicyMatcherJobs(
		scopedContext("repository-matcher", observationCtx),
		service.store,
		RepositoryMatcherConfigInst,
	)
}

func scopedContext(component string, parent *observation.Context) *observation.Context {
	return observation.ScopedContext("codeintel", "policies", component, parent)
}
