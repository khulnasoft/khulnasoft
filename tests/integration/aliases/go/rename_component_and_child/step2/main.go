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

// Scenario #5 - composing #1 and #3 and making both changes at the same time
func NewFooComponent(ctx *khulnasoft.Context, name string, opts ...khulnasoft.ResourceOption) (*FooComponent, error) {
	fooComp := &FooComponent{}
	err := ctx.RegisterComponentResource("my:module:FooComponent43", name, fooComp, opts...)
	if err != nil {
		return nil, err
	}
	parentOpt := khulnasoft.Parent(fooComp)
	alias := &khulnasoft.Alias{
		Name:   khulnasoft.StringInput(khulnasoft.String("otherchild")),
		Parent: fooComp,
	}
	aliasOpt := khulnasoft.Aliases([]khulnasoft.Alias{*alias})
	_, err = NewFooResource(ctx, "otherchildrenamed", parentOpt, aliasOpt)
	if err != nil {
		return nil, err
	}
	return fooComp, nil
}

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		alias := &khulnasoft.Alias{Name: khulnasoft.StringInput(khulnasoft.String("comp5"))}
		aliasOpt := khulnasoft.Aliases([]khulnasoft.Alias{*alias})
		_, err := NewFooComponent(ctx, "newcomp5", aliasOpt)
		if err != nil {
			return err
		}

		return nil
	})
}
