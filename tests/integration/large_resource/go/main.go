//go:build !all
// +build !all

package main

import (
	"strings"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		// Create and export a very long string (>4mb)
		ctx.Export("longString", khulnasoft.String(strings.Repeat("a", 5*1024*1024)))
		return nil
	})
}
