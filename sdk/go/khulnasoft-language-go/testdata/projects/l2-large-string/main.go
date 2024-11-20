package main

import (
	"example.com/khulnasoft-large/sdk/go/v4/large"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		res, err := large.NewString(ctx, "res", &large.StringArgs{
			Value: khulnasoft.String("hello world"),
		})
		if err != nil {
			return err
		}
		ctx.Export("output", res.Value)
		return nil
	})
}
