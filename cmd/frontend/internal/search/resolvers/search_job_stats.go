package resolvers

import (
	"github.com/khulnasoft/khulnasoft/cmd/frontend/graphqlbackend"
	"github.com/khulnasoft/khulnasoft/internal/search/exhaustive/types"
)

var _ graphqlbackend.SearchJobStatsResolver = &searchJobStatsResolver{}

type searchJobStatsResolver struct {
	*types.RepoRevJobStats
}

func (e *searchJobStatsResolver) Total() int32 {
	return e.RepoRevJobStats.Total
}

func (e *searchJobStatsResolver) Completed() int32 {
	return e.RepoRevJobStats.Completed
}

func (e *searchJobStatsResolver) Failed() int32 {
	return e.RepoRevJobStats.Failed
}

func (e *searchJobStatsResolver) InProgress() int32 {
	return e.RepoRevJobStats.InProgress
}
