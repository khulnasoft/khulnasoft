// Copyright 2016-2024, Pulumi Corporation.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"reflect"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

type componentArgs struct {
	Foo string `khulnasoft:"foo"`
}

type ComponentArgs struct {
	Foo khulnasoft.StringInput
}

func (ComponentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentArgs)(nil)).Elem()
}

type Component struct {
	khulnasoft.ResourceState

	Foo khulnasoft.StringOutput `khulnasoft:"foo"`
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
