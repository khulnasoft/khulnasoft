package batches

import (
	"github.com/khulnasoft/khulnasoft/cmd/worker/internal/executorqueue"
	"github.com/khulnasoft/khulnasoft/internal/env"
	"github.com/khulnasoft/khulnasoft/lib/errors"
)

type janitorConfig struct {
	env.BaseConfig

	MetricsConfig *executorqueue.Config
}

var janitorConfigInst = &janitorConfig{}

func (c *janitorConfig) Load() {
	c.MetricsConfig = executorqueue.InitMetricsConfig()
	c.MetricsConfig.Load()
}

func (c *janitorConfig) Validate() error {
	var errs error
	errs = errors.Append(errs, c.BaseConfig.Validate())
	errs = errors.Append(errs, c.MetricsConfig.Validate())
	return errs
}
