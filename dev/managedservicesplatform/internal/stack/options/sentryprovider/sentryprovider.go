package sentryprovider

import (
	sentry "github.com/sourcegraph/managed-services-platform-cdktf/gen/sentry/provider"

	"github.com/khulnasoft/khulnasoft/dev/managedservicesplatform/internal/resource/gsmsecret"
	"github.com/khulnasoft/khulnasoft/dev/managedservicesplatform/internal/resourceid"
	"github.com/khulnasoft/khulnasoft/dev/managedservicesplatform/internal/stack"
	"github.com/khulnasoft/khulnasoft/lib/pointers"
)

// With configures a stack to be able to use Sentry resources.
func With(sentryToken gsmsecret.DataConfig) stack.NewStackOption {
	return func(s stack.Stack) error {
		_ = sentry.NewSentryProvider(s.Stack, pointers.Ptr("sentry"),
			&sentry.SentryProviderConfig{
				Token: &gsmsecret.Get(s.Stack, resourceid.New("sentry-provider-token"), sentryToken).Value,
			})
		return nil
	}
}
