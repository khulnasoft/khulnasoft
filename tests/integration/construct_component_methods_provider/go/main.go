// Copyright 2016-2023, Pulumi Corporation.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"reflect"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

type TestProvider struct {
	khulnasoft.ProviderResourceState
}

func NewTestProvider(ctx *khulnasoft.Context, name string) (*TestProvider, error) {
	var resource TestProvider
	err := ctx.RegisterResource("khulnasoft:providers:testprovider", name, nil, &resource)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

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
		testProvider, err := NewTestProvider(ctx, "testProvider")
		if err != nil {
			return err
		}

		component1, err := NewComponent(ctx, "component1", &ComponentArgs{
			First:  khulnasoft.String("Hello"),
			Second: khulnasoft.String("World"),
		}, khulnasoft.Provider(testProvider))
		if err != nil {
			return err
		}
		result1, err := component1.GetMessage(ctx, &ComponentGetMessageArgs{
			Name: khulnasoft.String("Alice"),
		})
		if err != nil {
			return err
		}

		component2, err := NewComponent(ctx, "component2", &ComponentArgs{
			First:  khulnasoft.String("Hi"),
			Second: khulnasoft.String("There"),
		}, khulnasoft.Providers(testProvider))
		if err != nil {
			return err
		}
		result2, err := component2.GetMessage(ctx, &ComponentGetMessageArgs{
			Name: khulnasoft.String("Bob"),
		})
		if err != nil {
			return err
		}

		ctx.Export("message1", result1.Message())
		ctx.Export("message2", result2.Message())

		return nil
	})
}
