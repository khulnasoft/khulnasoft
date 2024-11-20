// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"reflect"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

type Component struct {
	khulnasoft.ResourceState
}

func NewComponent(ctx *khulnasoft.Context, name string, args ComponentArgs,
	opts ...khulnasoft.ResourceOption,
) (*Component, error) {
	var resource Component
	err := ctx.RegisterRemoteComponentResource("testcomponent:index:Component", name, &args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type componentArgs struct {
	Children *int `khulnasoft:"children"`
}

type ComponentArgs struct {
	Children *int
}

func (ComponentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentArgs)(nil)).Elem()
}

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		children := 5
		if _, err := NewComponent(ctx, "component", ComponentArgs{Children: &children}); err != nil {
			return err
		}
		return nil
	})
}
