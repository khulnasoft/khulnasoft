// Copyright 2016-2023, Pulumi Corporation.  All rights reserved.
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
	aliasOpt := khulnasoft.Aliases([]khulnasoft.Alias{{
		Type: khulnasoft.String("my:module:FooResource"),
	}})
	opts = append(opts, aliasOpt)
	err := ctx.RegisterComponentResource("my:module:FooResourceNew", name, fooRes, opts...)
	if err != nil {
		return nil, err
	}
	return fooRes, nil
}

func NewFooComponent(ctx *khulnasoft.Context, name string, opts ...khulnasoft.ResourceOption) (*FooComponent, error) {
	fooComp := &FooComponent{}
	err := ctx.RegisterComponentResource("my:module:FooComponent", name, fooComp, opts...)
	if err != nil {
		return nil, err
	}
	parentOpt := khulnasoft.Parent(fooComp)
	_, err = NewFooResource(ctx, "child", parentOpt)
	if err != nil {
		return nil, err
	}
	return fooComp, nil
}

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		_, err := NewFooComponent(ctx, "foo")
		if err != nil {
			return err
		}

		return nil
	})
}
