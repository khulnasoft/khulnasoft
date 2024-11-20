// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		component, err := NewComponent(ctx, "component")
		if err != nil {
			return err
		}
		result, err := component.CreateRandom(ctx, &ComponentCreateRandomArgs{
			Length: khulnasoft.Int(10),
		})
		if err != nil {
			return err
		}
		ctx.Export("result", result.Result())
		return nil
	})
}
