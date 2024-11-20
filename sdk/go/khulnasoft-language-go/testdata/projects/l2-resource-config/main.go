package main

import (
	"example.com/khulnasoft-config/sdk/go/v9/config"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		prov, err := config.NewProvider(ctx, "prov", &config.ProviderArgs{
			Name:              khulnasoft.String("my config"),
			PluginDownloadURL: khulnasoft.String("not the same as the khulnasoft resource option"),
		})
		if err != nil {
			return err
		}
		// Note this isn't _using_ the explicit provider, it's just grabbing a value from it.
		_, err = config.NewResource(ctx, "res", &config.ResourceArgs{
			Text: prov.Version,
		})
		if err != nil {
			return err
		}
		ctx.Export("pluginDownloadURL", prov.PluginDownloadURL)
		return nil
	})
}
