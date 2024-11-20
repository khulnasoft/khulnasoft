//go:build !all
// +build !all

package main

import (
	"time"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		time.Sleep(5 * time.Second)
		ctx.Export("exp_static", khulnasoft.String("foo"))
		return nil
	})
}
