package codeintel

import (
	"github.com/khulnasoft/khulnasoft/internal/codeintel/autoindexing"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/codenav"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/context"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/dependencies"
	ossdependencies "github.com/khulnasoft/khulnasoft/internal/codeintel/dependencies"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/policies"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/ranking"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/reposcheduler"
	codeintelshared "github.com/khulnasoft/khulnasoft/internal/codeintel/shared"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/uploads"
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/gitserver"
	"github.com/khulnasoft/khulnasoft/internal/observation"
)

type Services struct {
	AutoIndexingService          *autoindexing.Service
	PreciseRepoSchedulingService reposcheduler.RepositorySchedulingService
	CodenavService               *codenav.Service
	DependenciesService          *ossdependencies.Service
	PoliciesService              *policies.Service
	RankingService               *ranking.Service
	UploadsService               *uploads.Service
	ContextService               *context.Service
	GitserverClient              gitserver.Client
}

type ServiceDependencies struct {
	DB             database.DB
	CodeIntelDB    codeintelshared.CodeIntelDB
	ObservationCtx *observation.Context
}

func NewServices(deps ServiceDependencies) (Services, error) {
	db, codeIntelDB := deps.DB, deps.CodeIntelDB
	gitserverClient := gitserver.NewClient("codeintel")

	uploadsSvc := uploads.NewService(deps.ObservationCtx, db, codeIntelDB, gitserverClient.Scoped("uploads"))
	dependenciesSvc := dependencies.NewService(deps.ObservationCtx, db)
	policiesSvc := policies.NewService(deps.ObservationCtx, db, uploadsSvc, gitserverClient.Scoped("policies"))
	autoIndexingSvc := autoindexing.NewService(deps.ObservationCtx, db, dependenciesSvc, policiesSvc, gitserverClient.Scoped("autoindexing"))
	codenavSvc := codenav.NewService(deps.ObservationCtx, db, codeIntelDB, uploadsSvc, gitserverClient.Scoped("codenav"))
	rankingSvc := ranking.NewService(deps.ObservationCtx, db, codeIntelDB)
	contextService := context.NewService(deps.ObservationCtx, db)
	reposchedulingService := reposcheduler.NewService(reposcheduler.NewPreciseStore(deps.ObservationCtx, db))

	return Services{
		AutoIndexingService:          autoIndexingSvc,
		PreciseRepoSchedulingService: reposchedulingService,
		CodenavService:               codenavSvc,
		DependenciesService:          dependenciesSvc,
		PoliciesService:              policiesSvc,
		RankingService:               rankingSvc,
		UploadsService:               uploadsSvc,
		ContextService:               contextService,
		GitserverClient:              gitserverClient,
	}, nil
}
