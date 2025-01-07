package resolvers

import (
	"github.com/sourcegraph/log"

	"github.com/khulnasoft/khulnasoft/cmd/frontend/graphqlbackend"
	"github.com/khulnasoft/khulnasoft/internal/batches/store"
	btypes "github.com/khulnasoft/khulnasoft/internal/batches/types"
	"github.com/khulnasoft/khulnasoft/internal/gitserver"
	"github.com/khulnasoft/khulnasoft/internal/types"
)

type changesetJobErrorResolver struct {
	store           *store.Store
	logger          log.Logger
	changeset       *btypes.Changeset
	repo            *types.Repo
	error           string
	gitserverClient gitserver.Client
}

var _ graphqlbackend.ChangesetJobErrorResolver = &changesetJobErrorResolver{}

func (r *changesetJobErrorResolver) Changeset() graphqlbackend.ChangesetResolver {
	return NewChangesetResolver(r.store, r.gitserverClient, r.logger, r.changeset, r.repo)
}

func (r *changesetJobErrorResolver) Error() *string {
	// We only show the error when the changeset is visible to the requesting user.
	if r.repo == nil {
		return nil
	}
	return &r.error
}
