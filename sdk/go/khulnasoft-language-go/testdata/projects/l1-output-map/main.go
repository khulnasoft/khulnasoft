package main

import (
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		ctx.Export("empty", khulnasoft.Map{})
		ctx.Export("strings", khulnasoft.StringMap{
			"greeting": khulnasoft.String("Hello, world!"),
			"farewell": khulnasoft.String("Goodbye, world!"),
		})
		ctx.Export("numbers", khulnasoft.Float64Map{
			"1": khulnasoft.Float64(1),
			"2": khulnasoft.Float64(2),
		})
		ctx.Export("keys", khulnasoft.Float64Map{
			"my.key": khulnasoft.Float64(1),
			"my-key": khulnasoft.Float64(2),
			"my_key": khulnasoft.Float64(3),
			"MY_KEY": khulnasoft.Float64(4),
			"mykey":  khulnasoft.Float64(5),
			"MYKEY":  khulnasoft.Float64(6),
		})
		return nil
	})
}
