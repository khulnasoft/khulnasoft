package main

import "github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		panic("great sadness")
	})
}
