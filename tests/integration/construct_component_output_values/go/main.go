// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		if _, err := NewComponent(ctx, "component", &ComponentArgs{
			Foo: &FooArgs{
				Something: khulnasoft.String("hello"),
			},
			Bar: &BarArgs{
				Tags: khulnasoft.StringMap{
					"a": khulnasoft.String("world"),
					"b": khulnasoft.ToSecret("shh").(khulnasoft.StringOutput),
				},
			},
		}); err != nil {
			return err
		}
		return nil
	})
}
