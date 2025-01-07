package search

import (
	"context"

	"github.com/khulnasoft/khulnasoft/cmd/frontend/enterprise"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/search/httpapi"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/search/resolvers"
	"github.com/khulnasoft/khulnasoft/internal/codeintel"
	"github.com/khulnasoft/khulnasoft/internal/conf"
	"github.com/khulnasoft/khulnasoft/internal/conf/conftypes"
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/gitserver"
	"github.com/khulnasoft/khulnasoft/internal/observation"
	"github.com/khulnasoft/khulnasoft/internal/search"
	"github.com/khulnasoft/khulnasoft/internal/search/client"
	"github.com/khulnasoft/khulnasoft/internal/search/exhaustive"
	"github.com/khulnasoft/khulnasoft/internal/search/exhaustive/service"
	"github.com/khulnasoft/khulnasoft/internal/search/exhaustive/store"
)

func LoadConfig() {
	search.ObjectStorageConfigInst.Load()
}

// Init initializes the given enterpriseServices to include the required resolvers for search.
func Init(
	ctx context.Context,
	observationCtx *observation.Context,
	db database.DB,
	_ codeintel.Services,
	_ conftypes.UnifiedWatchable,
	enterpriseServices *enterprise.Services,
) error {
	if !exhaustive.IsEnabled(conf.Get()) {
		return nil
	}
	logger := observationCtx.Logger
	store := store.New(db, observationCtx)

	uploadStore, err := search.NewObjectStorage(ctx, observationCtx, search.ObjectStorageConfigInst)
	if err != nil {
		return err
	}

	searchClient := client.New(logger, db, gitserver.NewClient("http.search"))
	newSearcher := service.FromSearchClient(searchClient)

	svc := service.New(observationCtx, store, uploadStore, newSearcher)

	enterpriseServices.SearchJobsResolver = resolvers.New(logger, db, svc)
	enterpriseServices.SearchJobsDataExportHandler = httpapi.ServeSearchJobDownload(logger, svc)
	enterpriseServices.SearchJobsLogsHandler = httpapi.ServeSearchJobLogs(logger, svc)

	return nil
}
