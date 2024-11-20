//go:build !all
// +build !all

package main

import (
	"fmt"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		return fmt.Errorf("hello world")
	})
}
