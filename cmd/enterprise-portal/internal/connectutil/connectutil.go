package connectutil

import (
	"context"

	"connectrpc.com/connect"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/sourcegraph/log"

	sgtrace "github.com/khulnasoft/khulnasoft/internal/trace"
	"github.com/khulnasoft/khulnasoft/lib/errors"
)

// InternalError logs an error, adds it to the trace, and returns a connect
// error with a safe message.
func InternalError(ctx context.Context, logger log.Logger, err error, safeMsg string) error {
	trace.SpanFromContext(ctx).
		SetAttributes(
			attribute.String("full_error", err.Error()),
		)
	sgtrace.Logger(ctx, logger).
		AddCallerSkip(1).
		Error(safeMsg,
			log.String("code", connect.CodeInternal.String()),
			log.Error(err),
		)
	return connect.NewError(connect.CodeInternal, errors.New(safeMsg))
}
