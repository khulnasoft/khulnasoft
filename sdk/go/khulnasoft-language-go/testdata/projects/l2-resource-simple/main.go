package main

import (
	"example.com/khulnasoft-simple/sdk/go/v2/simple"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		_, err := simple.NewResource(ctx, "res", &simple.ResourceArgs{
			Value: khulnasoft.Bool(true),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
