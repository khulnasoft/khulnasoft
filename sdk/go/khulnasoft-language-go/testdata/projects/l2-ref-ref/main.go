package main

import (
	"example.com/khulnasoft-ref-ref/sdk/go/v12/refref"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		_, err := refref.NewResource(ctx, "res", &refref.ResourceArgs{
			Data: &refref.DataArgs{
				InnerData: &refref.InnerDataArgs{
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
				Boolean:   khulnasoft.Bool(true),
				Float:     khulnasoft.Float64(4.5),
				Integer:   khulnasoft.Int(1024),
				String:    khulnasoft.String("Hello"),
				BoolArray: khulnasoft.BoolArray{},
				StringMap: khulnasoft.StringMap{
					"x": khulnasoft.String("100"),
					"y": khulnasoft.String("200"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
