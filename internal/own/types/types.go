package types

import (
	"time"

	"github.com/khulnasoft/khulnasoft/internal/api"
	codeownerspb "github.com/khulnasoft/khulnasoft/internal/own/codeowners/v1"
)

type CodeownersFile struct {
	CreatedAt time.Time
	UpdatedAt time.Time

	RepoID   api.RepoID
	Contents string
	Proto    *codeownerspb.File
}

// These signal constants should match the names in the `own_signal_configurations` table
const (
	SignalRecentContributors = "recent-contributors"
	SignalRecentViews        = "recent-views"
	Analytics                = "analytics"
)
