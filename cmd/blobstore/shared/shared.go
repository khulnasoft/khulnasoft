// Package shared is the shared main entrypoint for blobstore, a simple service which exposes
// an S3-compatible API for object storage. See the blobstore package for more information.
package shared

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/sourcegraph/log"

	"github.com/khulnasoft/khulnasoft/cmd/blobstore/internal/blobstore"
	"github.com/khulnasoft/khulnasoft/internal/actor"
	"github.com/khulnasoft/khulnasoft/internal/conf/deploy"
	"github.com/khulnasoft/khulnasoft/internal/goroutine"
	"github.com/khulnasoft/khulnasoft/internal/instrumentation"
	"github.com/khulnasoft/khulnasoft/internal/observation"
	"github.com/khulnasoft/khulnasoft/internal/service"
	"github.com/khulnasoft/khulnasoft/internal/tenant"
	"github.com/khulnasoft/khulnasoft/internal/trace"
)

func shutdownOnSignal(ctx context.Context, server *http.Server) error {
	// Listen for shutdown signals. When we receive one attempt to clean up,
	// but do an insta-shutdown if we receive more than one signal.
	c := make(chan os.Signal, 2)
	signal.Notify(c, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM)

	// Once we receive one of the signals from above, continues with the shutdown
	// process.
	select {
	case <-c:
	case <-ctx.Done(): // still call shutdown below
	}

	go func() {
		// If a second signal is received, exit immediately.
		<-c
		os.Exit(1)
	}()

	// Wait for at most for the configured shutdown timeout.
	ctx, cancel := context.WithTimeout(ctx, goroutine.GracefulShutdownTimeout)
	defer cancel()
	// Stop accepting requests.
	return server.Shutdown(ctx)
}

func Start(ctx context.Context, observationCtx *observation.Context, config *Config, ready service.ReadyFunc) error {
	logger := observationCtx.Logger

	// Ready immediately
	ready()

	bsService := &blobstore.Service{
		DataDir:        config.DataDir,
		Log:            logger,
		ObservationCtx: observation.NewContext(logger),
	}

	// Set up handler middleware
	handler := actor.HTTPMiddleware(logger, bsService)
	handler = trace.HTTPMiddleware(logger, handler)
	handler = tenant.InternalHTTPMiddleware(logger, handler)
	handler = instrumentation.HTTPMiddleware("", handler)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	g, ctx := errgroup.WithContext(ctx)

	host, port := deploy.BlobstoreHostPort()
	addr := net.JoinHostPort(host, port)
	server := &http.Server{
		ReadTimeout:  75 * time.Second,
		WriteTimeout: 10 * time.Minute,
		Addr:         addr,
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// For cluster liveness and readiness probes
			if r.URL.Path == "/healthz" {
				w.WriteHeader(200)
				_, _ = w.Write([]byte("ok"))
				return
			}
			handler.ServeHTTP(w, r)
		}),
	}

	// Listen
	g.Go(func() error {
		logger.Info("listening", log.String("addr", server.Addr))
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	// Shutdown
	g.Go(func() error {
		return shutdownOnSignal(ctx, server)
	})

	return g.Wait()
}
