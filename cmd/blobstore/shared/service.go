package shared

import (
	"context"

	"github.com/khulnasoft/khulnasoft/internal/debugserver"
	"github.com/khulnasoft/khulnasoft/internal/env"
	"github.com/khulnasoft/khulnasoft/internal/observation"
	"github.com/khulnasoft/khulnasoft/internal/service"
)

type svc struct{}

func (svc) Name() string { return "blobstore" }

func (svc) Configure() (env.Config, []debugserver.Endpoint) {
	return LoadConfig(), nil
}

func (svc) Start(ctx context.Context, observationCtx *observation.Context, ready service.ReadyFunc, config env.Config) error {
	return Start(ctx, observationCtx, config.(*Config), ready)
}

var Service service.Service = svc{}

type Config struct {
	env.BaseConfig

	DataDir string
}

func (c *Config) Load() {
	c.DataDir = c.Get("BLOBSTORE_DATA_DIR", "/data", "directory to store blobstore buckets and objects")
}

func LoadConfig() *Config {
	var config Config
	config.Load()
	return &config
}
