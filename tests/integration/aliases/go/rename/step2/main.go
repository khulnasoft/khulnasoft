// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

// FooComponent is a component resource
type FooComponent struct {
	khulnasoft.ResourceState
}

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		fooComponent := &FooComponent{}
		alias := &khulnasoft.Alias{
			Name: khulnasoft.String("foo"),
		}
		opts := khulnasoft.Aliases([]khulnasoft.Alias{*alias})
		return ctx.RegisterComponentResource("foo:component", "newfoo", fooComponent, opts)
	})
}
