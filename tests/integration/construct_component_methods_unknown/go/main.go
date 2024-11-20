// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		r, err := NewRandom(ctx, "resource", &RandomArgs{Length: khulnasoft.Int(10)})
		if err != nil {
			return err
		}

		component, err := NewComponent(ctx, "component")
		if err != nil {
			return err
		}
		result, err := component.GetMessage(ctx, &ComponentGetMessageArgs{
			Echo: r.ID().ToStringOutput(),
		})
		if err != nil {
			return err
		}

		ctx.Export("result", result.ApplyT(func(val ComponentGetMessageResult) string {
			panic("should not run (result)")
		}))
		return nil
	})
}
