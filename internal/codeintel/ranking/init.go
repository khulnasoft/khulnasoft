package ranking

import (
	"github.com/khulnasoft/khulnasoft/internal/codeintel/ranking/internal/background"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/ranking/internal/background/coordinator"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/ranking/internal/background/exporter"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/ranking/internal/background/janitor"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/ranking/internal/background/mapper"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/ranking/internal/background/reducer"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/ranking/internal/lsifstore"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/ranking/internal/store"
	codeintelshared "github.com/khulnasoft/khulnasoft/internal/codeintel/shared"
	"github.com/khulnasoft/khulnasoft/internal/conf"
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/goroutine"
	"github.com/khulnasoft/khulnasoft/internal/observation"
)

func NewService(
	observationCtx *observation.Context,
	db database.DB,
	codeIntelDB codeintelshared.CodeIntelDB,
) *Service {
	return newService(
		scopedContext("service", observationCtx),
		store.New(scopedContext("store", observationCtx), db),
		lsifstore.New(scopedContext("lsifstore", observationCtx), codeIntelDB),
		conf.DefaultClient(),
	)
}

var (
	ExporterConfigInst    = &exporter.Config{}
	CoordinatorConfigInst = &coordinator.Config{}
	MapperConfigInst      = &mapper.Config{}
	ReducerConfigInst     = &reducer.Config{}
	JanitorConfigInst     = &janitor.Config{}
)

func NewSymbolExporter(observationCtx *observation.Context, rankingService *Service) goroutine.BackgroundRoutine {
	return background.NewSymbolExporter(
		scopedContext("exporter", observationCtx),
		rankingService.store,
		rankingService.lsifstore,
		ExporterConfigInst,
	)
}

func NewCoordinator(observationCtx *observation.Context, rankingService *Service) goroutine.BackgroundRoutine {
	return background.NewCoordinator(
		scopedContext("coordinator", observationCtx),
		rankingService.store,
		CoordinatorConfigInst,
	)
}

func NewMapper(observationCtx *observation.Context, rankingService *Service) []goroutine.BackgroundRoutine {
	return background.NewMapper(
		scopedContext("mapper", observationCtx),
		rankingService.store,
		MapperConfigInst,
	)
}

func NewReducer(observationCtx *observation.Context, rankingService *Service) goroutine.BackgroundRoutine {
	return background.NewReducer(
		scopedContext("reducer", observationCtx),
		rankingService.store,
		ReducerConfigInst,
	)
}

func NewSymbolJanitor(observationCtx *observation.Context, rankingService *Service) []goroutine.BackgroundRoutine {
	return background.NewSymbolJanitor(
		scopedContext("janitor", observationCtx),
		rankingService.store,
		JanitorConfigInst,
	)
}

func scopedContext(component string, observationCtx *observation.Context) *observation.Context {
	return observation.ScopedContext("codeintel", "ranking", component, observationCtx)
}
