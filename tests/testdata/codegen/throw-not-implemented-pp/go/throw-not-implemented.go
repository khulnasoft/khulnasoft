package main

import (
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func notImplemented(message string) khulnasoft.AnyOutput {
	panic(message)
}

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		ctx.Export("result", khulnasoft.Any(notImplemented("expression here is not implemented yet")))
		return nil
	})
}
