package httpapi

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/sourcegraph/log"

	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/handlerutil"
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/repoupdater"
)

func serveRepoRefresh(db database.DB) func(http.ResponseWriter, *http.Request) error {
	logger := log.Scoped("serveRepoRefresh")
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		repo, err := handlerutil.GetRepo(ctx, logger, db, mux.Vars(r))
		if err != nil {
			return err
		}

		_, err = repoupdater.DefaultClient.EnqueueRepoUpdate(ctx, repo.Name)
		return err
	}
}
