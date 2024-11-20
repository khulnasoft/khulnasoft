// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
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
		return nil
	})
}
