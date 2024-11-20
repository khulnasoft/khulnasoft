// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

// FooComponent is a component resource
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

// Scenario #3 - rename a component (and all it's children)
func NewFooComponent(ctx *khulnasoft.Context, name string, opts ...khulnasoft.ResourceOption) (*FooComponent, error) {
	fooComp := &FooComponent{}
	err := ctx.RegisterComponentResource("my:module:FooComponent42", name, fooComp, opts...)
	if err != nil {
		return nil, err
	}
	// Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
	// alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
	parentOpt := khulnasoft.Parent(fooComp)
	_, err = NewFooResource(ctx, name+"-child", parentOpt)
	if err != nil {
		return nil, err
	}
	_, err = NewFooResource(ctx, "otherchild", parentOpt)
	if err != nil {
		return nil, err
	}
	return fooComp, nil
}

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		_, err := NewFooComponent(ctx, "comp3")
		if err != nil {
			return err
		}

		return nil
	})
}
