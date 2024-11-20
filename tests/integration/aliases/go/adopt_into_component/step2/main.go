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
	var nilInput khulnasoft.StringInput
	aliasURN := khulnasoft.CreateURN(
		khulnasoft.StringInput(khulnasoft.String("res2")),
		khulnasoft.StringInput(khulnasoft.String("my:module:FooResource")),
		nilInput,
		khulnasoft.StringInput(khulnasoft.String(ctx.Project())),
		khulnasoft.StringInput(khulnasoft.String(ctx.Stack())))
	alias := &khulnasoft.Alias{
		URN: aliasURN,
	}
	aliasOpt := khulnasoft.Aliases([]khulnasoft.Alias{*alias})
	parentOpt := khulnasoft.Parent(fooComp)
	_, err = NewFooResource(ctx, name+"-child", aliasOpt, parentOpt)
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

func NewFooComponent3(ctx *khulnasoft.Context,
	name string,
	childAliasParent khulnasoft.Resource,
	opts ...khulnasoft.ResourceOption,
) (*FooComponent3, error) {
	fooComp := &FooComponent3{}
	err := ctx.RegisterComponentResource("my:module:FooComponent3", name, fooComp, opts...)
	if err != nil {
		return nil, err
	}

	alias := &khulnasoft.Alias{}
	if childAliasParent != nil {
		alias.Parent = childAliasParent
	} else {
		alias.NoParent = khulnasoft.Bool(true)
	}

	aliasOpt := khulnasoft.Aliases([]khulnasoft.Alias{*alias})
	parentOpt := khulnasoft.Parent(fooComp)
	_, err = NewFooComponent2(ctx, name+"-child", aliasOpt, parentOpt)
	if err != nil {
		return nil, err
	}
	return fooComp, nil
}

func NewFooComponent4(ctx *khulnasoft.Context, name string, opts ...khulnasoft.ResourceOption) (*FooComponent4, error) {
	fooComp := &FooComponent4{}
	alias := &khulnasoft.Alias{
		Parent: nil,
	}
	aliasOpt := khulnasoft.Aliases([]khulnasoft.Alias{*alias, *alias})
	o := []khulnasoft.ResourceOption{aliasOpt}
	o = append(o, opts...)
	err := ctx.RegisterComponentResource("my:module:FooComponent4", name, fooComp, o...)
	if err != nil {
		return nil, err
	}
	return fooComp, nil
}

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		comp2, err := NewFooComponent(ctx, "comp2")
		if err != nil {
			return err
		}
		alias := &khulnasoft.Alias{
			NoParent: khulnasoft.Bool(true),
		}
		aliasOpt := khulnasoft.Aliases([]khulnasoft.Alias{*alias})
		parentOpt := khulnasoft.Parent(comp2)
		_, err = NewFooComponent2(ctx, "unparented", aliasOpt, parentOpt)
		if err != nil {
			return err
		}
		_, err = NewFooComponent3(ctx, "parentedbystack", nil)
		if err != nil {
			return err
		}
		pbcOpt := khulnasoft.Parent(comp2)
		_, err = NewFooComponent3(ctx, "parentedbycomponent", comp2, pbcOpt)
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
