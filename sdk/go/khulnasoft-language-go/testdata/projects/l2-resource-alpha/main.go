package main

import (
	"example.com/khulnasoft-alpha/sdk/go/v3/alpha"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		_, err := alpha.NewResource(ctx, "res", &alpha.ResourceArgs{
			Value: khulnasoft.Bool(true),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
