package streaming

import (
	"context"

	"github.com/sourcegraph/log"

	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/gitserver"
	"github.com/khulnasoft/khulnasoft/internal/search"
	"github.com/khulnasoft/khulnasoft/internal/search/client"
	"github.com/khulnasoft/khulnasoft/internal/search/streaming"
	"github.com/khulnasoft/khulnasoft/lib/pointers"
)

type SearchClient interface {
	Search(ctx context.Context, query string, patternType *string, sender streaming.Sender) (*search.Alert, error)
}

func NewInsightsSearchClient(db database.DB) SearchClient {
	logger := log.Scoped("insightsSearchClient")
	return &insightsSearchClient{
		db:           db,
		searchClient: client.New(logger, db, gitserver.NewClient("insights.search")),
	}
}

type insightsSearchClient struct {
	db           database.DB
	searchClient client.SearchClient
}

func (r *insightsSearchClient) Search(ctx context.Context, query string, patternType *string, sender streaming.Sender) (*search.Alert, error) {
	inputs, err := r.searchClient.Plan(
		ctx,
		"V3",
		patternType,
		query,
		search.Precise,
		search.Streaming,
		pointers.Ptr(int32(0)),
	)
	if err != nil {
		return nil, err
	}
	return r.searchClient.Execute(ctx, sender, inputs)
}
