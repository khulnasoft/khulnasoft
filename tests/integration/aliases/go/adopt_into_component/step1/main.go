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

type FooComponent2 struct {
	khulnasoft.ResourceState
}

type FooComponent3 struct {
	khulnasoft.ResourceState
}

type FooComponent4 struct {
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

func NewFooComponent(ctx *khulnasoft.Context, name string, opts ...khulnasoft.ResourceOption) (*FooComponent, error) {
	fooComp := &FooComponent{}
	err := ctx.RegisterComponentResource("my:module:FooComponent", name, fooComp, opts...)
	if err != nil {
		return nil, err
	}
	return fooComp, nil
}

func NewFooComponent2(ctx *khulnasoft.Context, name string, opts ...khulnasoft.ResourceOption) (*FooComponent2, error) {
	fooComp := &FooComponent2{}
	err := ctx.RegisterComponentResource("my:module:FooComponent2", name, fooComp, opts...)
	if err != nil {
		return nil, err
	}
	return fooComp, nil
}

func NewFooComponent3(ctx *khulnasoft.Context, name string, opts ...khulnasoft.ResourceOption) (*FooComponent3, error) {
	fooComp := &FooComponent3{}
	err := ctx.RegisterComponentResource("my:module:FooComponent3", name, fooComp, opts...)
	if err != nil {
		return nil, err
	}
	_, err = NewFooComponent2(ctx, name+"-child", opts...)
	if err != nil {
		return nil, err
	}
	return fooComp, nil
}

func NewFooComponent4(ctx *khulnasoft.Context, name string, opts ...khulnasoft.ResourceOption) (*FooComponent4, error) {
	fooComp := &FooComponent4{}
	err := ctx.RegisterComponentResource("my:module:FooComponent4", name, fooComp, opts...)
	if err != nil {
		return nil, err
	}
	return fooComp, nil
}

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		_, err := NewFooResource(ctx, "res2")
		if err != nil {
			return err
		}
		comp2, err := NewFooComponent(ctx, "comp2")
		if err != nil {
			return err
		}
		_, err = NewFooComponent2(ctx, "unparented")
		if err != nil {
			return err
		}
		_, err = NewFooComponent3(ctx, "parentedbystack")
		if err != nil {
			return err
		}
		pbcOpt := khulnasoft.Parent(comp2)
		_, err = NewFooComponent3(ctx, "parentedbycomponent", pbcOpt)
		if err != nil {
			return err
		}
		dupeOpt := khulnasoft.Parent(comp2)
		_, err = NewFooComponent4(ctx, "duplicateAliases", dupeOpt)
		if err != nil {
			return err
		}
		return nil
	})
}
