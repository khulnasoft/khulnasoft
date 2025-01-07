package testing

import (
	"testing"

	"github.com/khulnasoft/khulnasoft/internal/batches/types/scheduler/config"
	"github.com/khulnasoft/khulnasoft/internal/conf"
)

func MockConfig(t testing.TB, mockery *conf.Unified) {
	t.Helper()

	conf.Mock(mockery)
	t.Cleanup(func() { conf.Mock(nil) })
	config.Reset()
}
