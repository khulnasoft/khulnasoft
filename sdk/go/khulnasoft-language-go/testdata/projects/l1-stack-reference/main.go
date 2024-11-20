package main

import (
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		ref, err := khulnasoft.NewStackReference(ctx, "ref", &khulnasoft.StackReferenceArgs{
			Name: khulnasoft.String("organization/other/dev"),
		})
		if err != nil {
			return err
		}
		ctx.Export("plain", ref.GetOutput(khulnasoft.String("plain")))
		ctx.Export("secret", ref.GetOutput(khulnasoft.String("secret")))
		return nil
	})
}
