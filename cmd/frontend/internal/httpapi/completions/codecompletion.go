package completions

import (
	"net/http"

	"github.com/sourcegraph/log"

	"github.com/khulnasoft/khulnasoft/internal/completions/types"
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/guardrails"
	"github.com/khulnasoft/khulnasoft/internal/redispool"
	"github.com/khulnasoft/khulnasoft/internal/telemetry/telemetryrecorder"
)

// NewCodeCompletionsHandler is an http handler which sends back code completion results.
func NewCodeCompletionsHandler(logger log.Logger, db database.DB, test guardrails.AttributionTest) http.Handler {
	logger = logger.Scoped("code")
	rl := NewRateLimiter(db, redispool.Store, types.CompletionsFeatureCode)
	return newCompletionsHandler(
		logger,
		db,
		db.Users(),
		db.AccessTokens(),
		telemetryrecorder.New(db),
		test,
		types.CompletionsFeatureCode,
		rl,
		"code",
		getCodeCompletionModelFn())
}
