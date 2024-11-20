package main

import (
	"example.com/khulnasoft-primitive-ref/sdk/go/v11/primitiveref"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		_, err := primitiveref.NewResource(ctx, "res", &primitiveref.ResourceArgs{
			Data: &primitiveref.DataArgs{
				Boolean: khulnasoft.Bool(false),
				Float:   khulnasoft.Float64(2.17),
				Integer: khulnasoft.Int(-12),
				String:  khulnasoft.String("Goodbye"),
				BoolArray: khulnasoft.BoolArray{
					khulnasoft.Bool(false),
					khulnasoft.Bool(true),
				},
				StringMap: khulnasoft.StringMap{
					"two":   khulnasoft.String("turtle doves"),
					"three": khulnasoft.String("french hens"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
