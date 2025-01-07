package main

import (
	"github.com/getsentry/sentry-go"
	"github.com/sourcegraph/log"

	"github.com/khulnasoft/khulnasoft/cmd/cody-gateway/shared"
	"github.com/khulnasoft/khulnasoft/internal/conf"
	"github.com/khulnasoft/khulnasoft/internal/env"
	"github.com/khulnasoft/khulnasoft/internal/sanitycheck"
	"github.com/khulnasoft/khulnasoft/internal/service/svcmain"
)

var sentryDSN = env.Get("CODY_GATEWAY_SENTRY_DSN", "", "Sentry DSN")

func main() {
	sanitycheck.Pass()
	svcmain.SingleServiceMainWithoutConf(shared.Service, nil, svcmain.OutOfBandConfiguration{
		Logging: func() conf.LogSinksSource {
			if sentryDSN == "" {
				return nil
			}

			return conf.NewStaticLogsSinksSource(log.SinksConfig{
				Sentry: &log.SentrySink{
					ClientOptions: sentry.ClientOptions{
						Dsn: sentryDSN,
					},
				},
			})
		}(),
		Tracing: nil,
	})
}
