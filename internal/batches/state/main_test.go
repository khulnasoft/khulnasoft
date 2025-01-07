package state

import (
	"time"

	btypes "github.com/khulnasoft/khulnasoft/internal/batches/types"
	"github.com/khulnasoft/khulnasoft/internal/extsvc/github"
	"github.com/khulnasoft/khulnasoft/internal/extsvc/gitlab"
)

func setDraft(c *btypes.Changeset) *btypes.Changeset {
	switch m := c.Metadata.(type) {
	case *github.PullRequest:
		m.IsDraft = true
	case *gitlab.MergeRequest:
		m.WorkInProgress = true
	}
	return c
}

func timeToUnixMilli(t time.Time) int {
	return int(t.UnixNano()) / int(time.Millisecond)
}
