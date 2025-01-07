package http

import (
	"net/http"
	"sync"

	"github.com/sourcegraph/log"

	"github.com/khulnasoft/khulnasoft/internal/codeintel/uploads"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/uploads/transport/http/auth"
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/gitserver"
	"github.com/khulnasoft/khulnasoft/internal/object"
	"github.com/khulnasoft/khulnasoft/internal/observation"
	"github.com/khulnasoft/khulnasoft/internal/uploadhandler"
)

var (
	handler         http.Handler
	handlerWithAuth http.Handler
	handlerOnce     sync.Once
)

func GetHandler(svc *uploads.Service, db database.DB, gitserverClient gitserver.Client, uploadStore object.Storage, withCodeHostAuthAuth bool) http.Handler {
	handlerOnce.Do(func() {
		logger := log.Scoped(
			"uploads.handler",
		)

		observationCtx := observation.NewContext(logger)

		operations := newOperations(observationCtx)
		uploadHandlerOperations := uploadhandler.NewOperations(observationCtx, "codeintel")

		userStore := db.Users()
		repoStore := db.Repos()

		// Construct base handler, used in internal routes and as internal handler wrapped
		// in the auth middleware defined on the next few lines
		handler = newHandler(observationCtx, repoStore, gitserverClient, uploadStore, svc.UploadHandlerStore(), uploadHandlerOperations)

		// 🚨 SECURITY: Non-internal installations of this handler will require a user/repo
		// visibility check with the remote code host (if enabled via site configuration).
		handlerWithAuth = auth.AuthMiddleware(
			handler,
			userStore,
			repoStore,
			auth.DefaultValidatorByCodeHost,
			operations.authMiddleware,
		)
	})

	if withCodeHostAuthAuth {
		return handlerWithAuth
	}
	return handler
}
