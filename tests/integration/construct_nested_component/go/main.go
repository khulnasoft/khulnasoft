// Copyright 2016-2022, Pulumi Corporation.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"reflect"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

type componentArgs struct {
	Echo interface{} `khulnasoft:"echo"`
}

type ComponentArgs struct {
	Echo khulnasoft.Input
}

func (ComponentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentArgs)(nil)).Elem()
}

type Component struct {
	khulnasoft.ResourceState

	Echo    khulnasoft.AnyOutput    `khulnasoft:"echo"`
	ChildID khulnasoft.StringOutput `khulnasoft:"childId"`
	Secret  khulnasoft.StringOutput `khulnasoft:"secret"`
}

func NewComponent(
	ctx *khulnasoft.Context, name string, args *ComponentArgs, opts ...khulnasoft.ResourceOption,
) (*Component, error) {
	var resource Component
	err := ctx.RegisterRemoteComponentResource("testcomponent:index:Component", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}

	return &resource, nil
}

func NewSecondComponent(
	ctx *khulnasoft.Context, name string, args *ComponentArgs, opts ...khulnasoft.ResourceOption,
) (*Component, error) {
	var resource Component
	err := ctx.RegisterRemoteComponentResource("secondtestcomponent:index:Component", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}

	return &resource, nil
}

func NewComponentComponent(
	ctx *khulnasoft.Context, name string, opts ...khulnasoft.ResourceOption,
) (*Component, error) {
	var resource Component
	err := ctx.RegisterRemoteComponentResource("secondtestcomponent:index:ComponentComponent", name, khulnasoft.Map{}, &resource, opts...)
	if err != nil {
		return nil, err
	}

	return &resource, nil
}

type provider struct {
	khulnasoft.ProviderResourceState
	expectResourceArg khulnasoft.Bool
}

type LocalComponent struct{ khulnasoft.ResourceState }

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		componentA, err := NewComponent(ctx, "a", &ComponentArgs{Echo: khulnasoft.Int(42)})
		if err != nil {
			return err
		}
		_, err = NewComponent(ctx, "b", &ComponentArgs{Echo: componentA.Echo})
		if err != nil {
			return err
		}
		_, err = NewComponent(ctx, "C", &ComponentArgs{Echo: componentA.ChildID})
		if err != nil {
			return err
		}

		provider := &provider{}
		err = ctx.RegisterResource("khulnasoft:providers:testcomponent", "provider", khulnasoft.Map{
			"expectResourceArg": khulnasoft.Bool(true),
		}, provider)
		if err != nil {
			return err
		}
		localComponent := &LocalComponent{}
		err = ctx.RegisterComponentResource("pkg:index:LocalComponent", "localComponent", localComponent, khulnasoft.Providers(provider))
		if err != nil {
			return err
		}
		parentProvider := khulnasoft.Parent(localComponent)
		_, err = NewComponent(ctx, "checkProvider1",
			&ComponentArgs{Echo: khulnasoft.String("checkExpected")}, parentProvider)
		if err != nil {
			return err
		}
		_, err = NewSecondComponent(ctx, "checkProvider2",
			&ComponentArgs{Echo: khulnasoft.String("checkExpected")}, parentProvider)
		if err != nil {
			return err
		}
		_, err = NewComponentComponent(ctx, "checkProvider12", parentProvider)
		if err != nil {
			return err
		}
		return nil
	})
}
