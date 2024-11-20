package main

import (
	"github.com/khulnasoft/khulnasoft-unknown/sdk/go/unknown"
	"github.com/khulnasoft/khulnasoft-unknown/sdk/go/unknown/eks"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		data, err := unknown.GetData(ctx, map[string]interface{}{
			"input": "hello",
		}, nil)
		if err != nil {
			return err
		}
		_, err = eks.ModuleValues(ctx, map[string]interface{}{}, nil)
		if err != nil {
			return err
		}
		ctx.Export("content", data.Content)
		return nil
	})
}
