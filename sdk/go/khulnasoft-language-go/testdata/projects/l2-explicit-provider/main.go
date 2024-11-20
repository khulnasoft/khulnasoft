package main

import (
	"example.com/khulnasoft-simple/sdk/go/v2/simple"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		prov, err := simple.NewProvider(ctx, "prov", nil)
		if err != nil {
			return err
		}
		_, err = simple.NewResource(ctx, "res", &simple.ResourceArgs{
			Value: khulnasoft.Bool(true),
		}, khulnasoft.Provider(prov))
		if err != nil {
			return err
		}
		return nil
	})
}
