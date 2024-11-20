// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"reflect"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft/internals"
)

type componentArgs struct {
	First  string `khulnasoft:"first"`
	Second string `khulnasoft:"second"`
}

type ComponentArgs struct {
	First  khulnasoft.StringInput
	Second khulnasoft.StringInput
}

func (ComponentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentArgs)(nil)).Elem()
}

type Component struct {
	khulnasoft.ResourceState
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

func (c *Component) GetMessage(ctx *khulnasoft.Context, args *ComponentGetMessageArgs) (ComponentGetMessageResultOutput, error) {
	out, err := ctx.Call("testcomponent:index:Component/getMessage", args, ComponentGetMessageResultOutput{}, c)
	if err != nil {
		return ComponentGetMessageResultOutput{}, err
	}
	return out.(ComponentGetMessageResultOutput), nil
}

type componentGetMessageArgs struct {
	Name string `khulnasoft:"name"`
}

type ComponentGetMessageArgs struct {
	Name khulnasoft.StringInput
}

func (ComponentGetMessageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentGetMessageArgs)(nil)).Elem()
}

type ComponentGetMessageResult struct {
	Message string `khulnasoft:"message"`
}

type ComponentGetMessageResultOutput struct{ *khulnasoft.OutputState }

func (ComponentGetMessageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComponentGetMessageResult)(nil)).Elem()
}

func (o ComponentGetMessageResultOutput) Message() khulnasoft.StringOutput {
	return o.ApplyT(func(v ComponentGetMessageResult) string { return v.Message }).(khulnasoft.StringOutput)
}

func (*Component) ElementType() reflect.Type {
	return reflect.TypeOf((*Component)(nil))
}

func init() {
	khulnasoft.RegisterOutputType(ComponentGetMessageResultOutput{})
}

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		component, err := NewComponent(ctx, "component", &ComponentArgs{
			First:  khulnasoft.String("Hello"),
			Second: khulnasoft.String("World"),
		})
		if err != nil {
			return err
		}
		result, err := component.GetMessage(ctx, &ComponentGetMessageArgs{
			Name: khulnasoft.String("Alice"),
		})
		if err != nil {
			return err
		}
		message := result.Message()
		ctx.Export("message", message)
		ctx.Export("messagedeps", awaitDependencies(ctx, message))

		return nil
	})
}

func awaitDependencies(ctx *khulnasoft.Context, o khulnasoft.Output) khulnasoft.URNArray {
	r, err := internals.UnsafeAwaitOutput(ctx.Context(), o)
	if err != nil {
		panic(err)
	}
	var deps khulnasoft.URNArray
	for _, dep := range r.Dependencies {
		deps = append(deps, dep.URN())
	}
	return deps
}
