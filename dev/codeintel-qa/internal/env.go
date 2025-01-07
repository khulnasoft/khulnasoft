package internal

import "github.com/khulnasoft/khulnasoft/internal/env"

var (
	KhulnasoftEndpoint    = env.Get("KHULNASOFT_BASE_URL", "http://127.0.0.1:3080", "Khulnasoft frontend endpoint")
	KhulnasoftAccessToken = env.Get("KHULNASOFT_SUDO_TOKEN", "", "Khulnasoft access token with sudo privileges")
)
