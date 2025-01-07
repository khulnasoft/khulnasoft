package background

import (
	"github.com/khulnasoft/khulnasoft/internal/codeintel/autoindexing/internal/background/dependencies"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/autoindexing/internal/background/scheduler"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/autoindexing/internal/background/summary"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/autoindexing/internal/jobselector"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/autoindexing/internal/store"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/reposcheduler"
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/goroutine"
	"github.com/khulnasoft/khulnasoft/internal/observation"
	"github.com/khulnasoft/khulnasoft/internal/workerutil"
	dbworkerstore "github.com/khulnasoft/khulnasoft/internal/workerutil/dbworker/store"
)

var (
	IndexWorkerStoreOptions                 = dependencies.IndexWorkerStoreOptions
	DependencySyncingJobWorkerStoreOptions  = dependencies.DependencySyncingJobWorkerStoreOptions
	DependencyIndexingJobWorkerStoreOptions = dependencies.DependencyIndexingJobWorkerStoreOptions
)

func NewIndexSchedulers(
	observationCtx *observation.Context,
	policiesSvc scheduler.PoliciesService,
	policyMatcher scheduler.PolicyMatcher,
	repoSchedulingSvc reposcheduler.RepositorySchedulingService,
	indexEnqueuer scheduler.IndexEnqueuer,
	repoStore database.RepoStore,
	store store.Store,
	config *scheduler.Config,
) []goroutine.BackgroundRoutine {
	return []goroutine.BackgroundRoutine{
		scheduler.NewScheduler(
			observationCtx,
			repoSchedulingSvc,
			policiesSvc,
			policyMatcher,
			indexEnqueuer,
			repoStore,
			config,
		),

		scheduler.NewOnDemandScheduler(
			store,
			indexEnqueuer,
			config,
		),
	}
}

func NewDependencyIndexSchedulers(
	observationCtx *observation.Context,
	db database.DB,
	uploadSvc dependencies.UploadService,
	depsSvc dependencies.DependenciesService,
	store store.Store,
	indexEnqueuer dependencies.IndexEnqueuer,
	config *dependencies.Config,
) []goroutine.BackgroundRoutine {
	metrics := dependencies.NewResetterMetrics(observationCtx)
	indexStore := dbworkerstore.New(observationCtx, db.Handle(), dependencies.IndexWorkerStoreOptions)
	dependencySyncStore := dbworkerstore.New(observationCtx, db.Handle(), DependencySyncingJobWorkerStoreOptions)
	dependencyIndexingStore := dbworkerstore.New(observationCtx, db.Handle(), dependencies.DependencyIndexingJobWorkerStoreOptions)

	externalServiceStore := db.ExternalServices()
	repoStore := db.Repos()
	gitserverRepoStore := db.GitserverRepos()

	return []goroutine.BackgroundRoutine{
		dependencies.NewDependencySyncScheduler(
			dependencySyncStore,
			uploadSvc,
			depsSvc,
			store,
			externalServiceStore,
			workerutil.NewMetrics(observationCtx, "codeintel_dependency_index_processor"),
			config,
		),
		dependencies.NewDependencyIndexingScheduler(
			dependencyIndexingStore,
			uploadSvc,
			repoStore,
			externalServiceStore,
			gitserverRepoStore,
			indexEnqueuer,
			workerutil.NewMetrics(observationCtx, "codeintel_dependency_index_queueing"),
			config,
		),

		dependencies.NewIndexResetter(observationCtx.Logger.Scoped("indexResetter"), config.ResetterInterval, indexStore, metrics),
		dependencies.NewDependencyIndexResetter(observationCtx.Logger.Scoped("dependencyIndexResetter"), config.ResetterInterval, dependencyIndexingStore, metrics),
	}
}

func NewSummaryBuilder(
	observationCtx *observation.Context,
	store store.Store,
	jobSelector *jobselector.JobSelector,
	uploadSvc summary.UploadService,
	config *summary.Config,
) []goroutine.BackgroundRoutine {
	return []goroutine.BackgroundRoutine{
		summary.NewSummaryBuilder(
			observationCtx,
			store,
			jobSelector,
			uploadSvc,
			config,
		),
	}
}
