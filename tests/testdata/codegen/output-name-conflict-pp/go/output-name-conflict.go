package main

import (
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft/config"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		cfg := config.New(ctx, "")
		cidrBlock := "Test config variable"
		if param := cfg.Get("cidrBlock"); param != "" {
			cidrBlock = param
		}
		ctx.Export("cidrBlock", cidrBlock)
		return nil
	})
}
