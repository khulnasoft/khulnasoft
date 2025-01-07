package background

import (
	"github.com/khulnasoft/khulnasoft/internal/codeintel/ranking/internal/background/coordinator"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/ranking/internal/background/exporter"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/ranking/internal/background/janitor"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/ranking/internal/background/mapper"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/ranking/internal/background/reducer"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/ranking/internal/lsifstore"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/ranking/internal/store"
	"github.com/khulnasoft/khulnasoft/internal/goroutine"
	"github.com/khulnasoft/khulnasoft/internal/observation"
)

func NewSymbolExporter(observationCtx *observation.Context, store store.Store, lsifstore lsifstore.Store, config *exporter.Config) goroutine.BackgroundRoutine {
	return exporter.NewSymbolExporter(observationCtx, store, lsifstore, config)
}

func NewCoordinator(observationCtx *observation.Context, store store.Store, config *coordinator.Config) goroutine.BackgroundRoutine {
	return coordinator.NewCoordinator(observationCtx, store, config)
}

func NewMapper(observationCtx *observation.Context, store store.Store, config *mapper.Config) []goroutine.BackgroundRoutine {
	return []goroutine.BackgroundRoutine{
		mapper.NewMapper(observationCtx, store, config),
		mapper.NewSeedMapper(observationCtx, store, config),
	}
}

func NewReducer(observationCtx *observation.Context, store store.Store, config *reducer.Config) goroutine.BackgroundRoutine {
	return reducer.NewReducer(observationCtx, store, config)
}

func NewSymbolJanitor(observationCtx *observation.Context, store store.Store, config *janitor.Config) []goroutine.BackgroundRoutine {
	return []goroutine.BackgroundRoutine{
		janitor.NewExportedUploadsJanitor(observationCtx, store, config),
		janitor.NewDeletedUploadsJanitor(observationCtx, store, config),
		janitor.NewAbandonedExportedUploadsJanitor(observationCtx, store, config),
		janitor.NewProcessedReferencesJanitor(observationCtx, store, config),
		janitor.NewProcessedPathsJanitor(observationCtx, store, config),
		janitor.NewRankCountsJanitor(observationCtx, store, config),
		janitor.NewRankJanitor(observationCtx, store, config),
	}
}
