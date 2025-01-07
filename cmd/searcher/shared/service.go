package shared

import (
	"context"

	"github.com/khulnasoft/khulnasoft/internal/debugserver"
	"github.com/khulnasoft/khulnasoft/internal/env"
	"github.com/khulnasoft/khulnasoft/internal/observation"
	"github.com/khulnasoft/khulnasoft/internal/service"
)

type svc struct{}

func (svc) Name() string { return "searcher" }

func (svc) Configure() (env.Config, []debugserver.Endpoint) {
	c := LoadConfig()

	return c, []debugserver.Endpoint{GRPCWebUIDebugEndpoint(c.ListenAddress)}
}

func (svc) Start(ctx context.Context, observationCtx *observation.Context, ready service.ReadyFunc, c env.Config) error {
	return Start(ctx, observationCtx, ready, c.(*Config))
}

var Service service.Service = svc{}
