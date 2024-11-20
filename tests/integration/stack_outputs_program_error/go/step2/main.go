// Copyright 2016-2024, Pulumi Corporation.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"errors"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		ctx.Export("xyz", khulnasoft.String("DEF"))
		return errors.New("program error")
	})
}
