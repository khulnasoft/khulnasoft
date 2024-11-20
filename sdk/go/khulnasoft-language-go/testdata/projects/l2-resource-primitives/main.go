package main

import (
	"example.com/khulnasoft-primitive/sdk/go/v7/primitive"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		_, err := primitive.NewResource(ctx, "res", &primitive.ResourceArgs{
			Boolean: khulnasoft.Bool(true),
			Float:   khulnasoft.Float64(3.14),
			Integer: khulnasoft.Int(42),
			String:  khulnasoft.String("hello"),
			NumberArray: khulnasoft.Float64Array{
				khulnasoft.Float64(-1),
				khulnasoft.Float64(0),
				khulnasoft.Float64(1),
			},
			BooleanMap: khulnasoft.BoolMap{
				"t": khulnasoft.Bool(true),
				"f": khulnasoft.Bool(false),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
