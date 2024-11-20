// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft/config"
	"using-shared-types-in-config/credentials/internal"
)

var _ = internal.GetEnvOrDefault

// The (entirely uncryptographic) hash function used to encode the "password".
func GetHash(ctx *khulnasoft.Context) string {
	return config.Get(ctx, "credentials:hash")
}

// The password. It is very secret.
func GetPassword(ctx *khulnasoft.Context) string {
	v, err := config.Try(ctx, "credentials:password")
	if err == nil {
		return v
	}
	var value string
	if d := internal.GetEnvOrDefault("", nil, "FOO"); d != nil {
		value = d.(string)
	}
	return value
}
func GetShared(ctx *khulnasoft.Context) string {
	return config.Get(ctx, "credentials:shared")
}

// The username. Its important but not secret.
func GetUser(ctx *khulnasoft.Context) string {
	return config.Get(ctx, "credentials:user")
}
