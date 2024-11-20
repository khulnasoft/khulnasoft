package main

import (
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		ctx.Export("zero", khulnasoft.Float64(0))
		ctx.Export("one", khulnasoft.Float64(1))
		ctx.Export("e", khulnasoft.Float64(2.718))
		ctx.Export("minInt32", khulnasoft.Float64(-2147483648))
		ctx.Export("max", khulnasoft.Float64(1.7976931348623157e+308))
		ctx.Export("min", khulnasoft.Float64(5e-324))
		return nil
	})
}
