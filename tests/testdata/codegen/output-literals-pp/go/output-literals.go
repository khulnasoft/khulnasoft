package main

import (
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		ctx.Export("output_true", khulnasoft.Bool(true))
		ctx.Export("output_false", khulnasoft.Bool(false))
		ctx.Export("output_number", khulnasoft.Float64(4))
		ctx.Export("output_string", khulnasoft.String("hello"))
		return nil
	})
}
