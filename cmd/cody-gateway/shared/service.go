package shared

import (
	"context"

	"github.com/khulnasoft/khulnasoft/cmd/cody-gateway/shared/config"
	"github.com/khulnasoft/khulnasoft/internal/debugserver"
	"github.com/khulnasoft/khulnasoft/internal/env"
	"github.com/khulnasoft/khulnasoft/internal/observation"
	"github.com/khulnasoft/khulnasoft/internal/service"
)

// Service is the shared cody-gateway service.
var Service service.Service = svc{}

type svc struct{}

func (svc) Name() string { return "cody-gateway" }

func (svc) Configure() (env.Config, []debugserver.Endpoint) {
	c := &config.Config{}
	c.Load()
	return c, []debugserver.Endpoint{}
}

func (svc) Start(ctx context.Context, observationCtx *observation.Context, ready service.ReadyFunc, envCfg env.Config) error {
	return Main(ctx, observationCtx, ready, envCfg.(*config.Config))
}
