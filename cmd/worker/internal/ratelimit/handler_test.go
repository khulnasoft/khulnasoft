package ratelimit

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/sourcegraph/log/logtest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/khulnasoft/khulnasoft/internal/conf"
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/database/dbtest"
	"github.com/khulnasoft/khulnasoft/internal/extsvc"
	"github.com/khulnasoft/khulnasoft/internal/ratelimit"
	"github.com/khulnasoft/khulnasoft/internal/redispool"
	"github.com/khulnasoft/khulnasoft/internal/types"
	"github.com/khulnasoft/khulnasoft/lib/pointers"
	"github.com/khulnasoft/khulnasoft/schema"
)

func TestHandler_Handle(t *testing.T) {
	logger := logtest.Scoped(t)
	ctx := context.Background()
	db := database.NewDB(logger, dbtest.NewDB(t))

	prefix := "__test__" + t.Name()
	kv := redispool.NewTestKeyValue()

	t.Cleanup(func() {
		if err := redispool.DeleteAllKeysWithPrefix(kv, prefix); err != nil {
			t.Logf("Failed to clear redis: %+v\n", err)
		}
	})

	conf.Mock(&conf.Unified{
		SiteConfiguration: schema.SiteConfiguration{
			GitMaxCodehostRequestsPerSecond: pointers.Ptr(1),
		},
	})
	defer conf.Mock(nil)

	// Create the external service so that the first code host appears when the handler calls GetByURL.
	confGet := func() *conf.Unified { return &conf.Unified{} }
	extsvcConfig := extsvc.NewUnencryptedConfig(`{"url": "https://github.com/", "token":"abc", "repositoryQuery": ["none"], "rateLimit": {"enabled": true, "requestsPerHour": 150}}`)
	svc := &types.ExternalService{
		Kind:   extsvc.KindGitHub,
		Config: extsvcConfig,
	}
	err := db.ExternalServices().Create(ctx, confGet, svc)
	require.NoError(t, err)

	// Create the handler to start the test
	h := handler{
		externalServiceStore: db.ExternalServices(),
		newRateLimiterFunc: func(bucketName string) ratelimit.GlobalLimiter {
			return ratelimit.NewTestGlobalRateLimiter(kv.Pool(), prefix, bucketName)
		},
		logger: logger,
	}
	err = h.Handle(ctx)
	assert.NoError(t, err)

	info, err := ratelimit.GetGlobalLimiterStateFromStore(kv, prefix)
	require.NoError(t, err)

	if diff := cmp.Diff(map[string]ratelimit.GlobalLimiterInfo{
		svc.URN(): {
			Burst:             10,
			Limit:             150,
			Interval:          time.Hour,
			LastReplenishment: time.Unix(0, 0),
		},
		ratelimit.GitRPSLimiterBucketName: {
			Burst:             10,
			Limit:             1,
			Interval:          time.Second,
			LastReplenishment: time.Unix(0, 0),
		},
	}, info); diff != "" {
		t.Fatal(diff)
	}
}
