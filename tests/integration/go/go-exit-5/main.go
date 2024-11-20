//go:build !all
// +build !all

package main

import (
	"os"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		os.Exit(5)
		return nil
	})
}
