package shared

import (
	"context"

	"github.com/khulnasoft/khulnasoft/internal/debugserver"
	"github.com/khulnasoft/khulnasoft/internal/env"
	"github.com/khulnasoft/khulnasoft/internal/observation"
	"github.com/khulnasoft/khulnasoft/internal/service"
)

type svc struct{}

func (svc) Name() string { return "symbols" }

func (svc) Configure() (env.Config, []debugserver.Endpoint) {
	LoadConfig()
	config := loadRockskipConfig(env.BaseConfig{}, CtagsConfig, config.NumCtagsProcesses, config.RequestBufferSize, RepositoryFetcherConfig)
	return &config, []debugserver.Endpoint{GRPCWebUIDebugEndpoint()}
}

func (svc) Start(ctx context.Context, observationCtx *observation.Context, ready service.ReadyFunc, config env.Config) error {
	return Main(ctx, observationCtx, ready, CreateSetup(*config.(*rockskipConfig)))
}

var Service service.Service = svc{}
