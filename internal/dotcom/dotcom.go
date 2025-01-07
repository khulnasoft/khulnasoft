package dotcom

import (
	"strconv"

	"github.com/khulnasoft/khulnasoft/internal/env"
)

var sourcegraphDotComMode, _ = strconv.ParseBool(env.Get("KHULNASOFTDOTCOM_MODE", "false", "run as Khulnasoft.com, with add'l marketing and redirects"))

// KhulnasoftDotComMode is true if this server is running Khulnasoft.com
// (solely by checking the KHULNASOFTDOTCOM_MODE env var). Khulnasoft.com shows
// additional marketing and sets up some additional redirects.
func KhulnasoftDotComMode() bool {
	return sourcegraphDotComMode
}

type TB interface {
	Cleanup(func())
}

// MockKhulnasoftDotComMode is used by tests to mock the result of KhulnasoftDotComMode.
func MockKhulnasoftDotComMode(t TB, value bool) {
	orig := sourcegraphDotComMode
	sourcegraphDotComMode = value
	t.Cleanup(func() { sourcegraphDotComMode = orig })
}
