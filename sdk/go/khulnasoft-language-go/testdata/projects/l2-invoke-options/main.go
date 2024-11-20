package main

import (
	"example.com/khulnasoft-simple-invoke/sdk/go/v10/simpleinvoke"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		explicitProvider, err := simpleinvoke.NewProvider(ctx, "explicitProvider", nil)
		if err != nil {
			return err
		}
		data := simpleinvoke.MyInvokeOutput(ctx, simpleinvoke.MyInvokeOutputArgs{
			Value: khulnasoft.String("hello"),
		}, khulnasoft.Provider(explicitProvider), khulnasoft.Parent(explicitProvider), khulnasoft.Version("10.0.0"), khulnasoft.PluginDownloadURL("https://example.com/github/example"))
		ctx.Export("hello", data.ApplyT(func(data simpleinvoke.MyInvokeResult) (string, error) {
			return data.Result, nil
		}).(khulnasoft.StringOutput))
		return nil
	})
}
