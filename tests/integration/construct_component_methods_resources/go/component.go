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

func NewComponent(ctx *khulnasoft.Context, name string, opts ...khulnasoft.ResourceOption) (*Component, error) {
	var resource Component
	err := ctx.RegisterRemoteComponentResource("testcomponent:index:Component", name, nil, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

func (c *Component) CreateRandom(ctx *khulnasoft.Context, args *ComponentCreateRandomArgs) (ComponentCreateRandomResultOutput, error) {
	out, err := ctx.Call("testcomponent:index:Component/createRandom", args, ComponentCreateRandomResultOutput{}, c)
	if err != nil {
		return ComponentCreateRandomResultOutput{}, err
	}
	return out.(ComponentCreateRandomResultOutput), nil
}

type componentCreateRandomArgs struct {
	Length int `khulnasoft:"length"`
}

type ComponentCreateRandomArgs struct {
	Length khulnasoft.IntInput
}

func (ComponentCreateRandomArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentCreateRandomArgs)(nil)).Elem()
}

type ComponentCreateRandomResult struct {
	Result string `khulnasoft:"result"`
}

type ComponentCreateRandomResultOutput struct{ *khulnasoft.OutputState }

func (ComponentCreateRandomResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComponentCreateRandomResult)(nil)).Elem()
}

func (o ComponentCreateRandomResultOutput) Result() khulnasoft.StringOutput {
	return o.ApplyT(func(v ComponentCreateRandomResult) string { return v.Result }).(khulnasoft.StringOutput)
}

func init() {
	khulnasoft.RegisterOutputType(ComponentCreateRandomResultOutput{})
}
