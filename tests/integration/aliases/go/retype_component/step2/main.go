// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

type FooResource struct {
	khulnasoft.ResourceState
}

type FooComponent struct {
	khulnasoft.ResourceState
}

func NewFooResource(ctx *khulnasoft.Context, name string, opts ...khulnasoft.ResourceOption) (*FooResource, error) {
	fooRes := &FooResource{}
	err := ctx.RegisterComponentResource("my:module:FooResource", name, fooRes, opts...)
	if err != nil {
		return nil, err
	}
	return fooRes, nil
}

// Scenario #4 - change the type of a component
func NewFooComponent(ctx *khulnasoft.Context, name string) (*FooComponent, error) {
	fooComp := &FooComponent{}
	alias := &khulnasoft.Alias{
		Type: khulnasoft.StringInput(khulnasoft.String("my:module:FooComponent44")),
	}
	aliasOpt := khulnasoft.Aliases([]khulnasoft.Alias{*alias})
	err := ctx.RegisterComponentResource("my:diffmodule:FooComponent55DiffType", name, fooComp, aliasOpt)
	if err != nil {
		return nil, err
	}
	parentOpt := khulnasoft.Parent(fooComp)
	_, err = NewFooResource(ctx, "otherchild", parentOpt)
	if err != nil {
		return nil, err
	}
	return fooComp, nil
}

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		_, err := NewFooComponent(ctx, "comp4")
		if err != nil {
			return err
		}

		return nil
	})
}
