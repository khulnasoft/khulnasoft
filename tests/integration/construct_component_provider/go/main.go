// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"reflect"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

type Provider struct {
	khulnasoft.ProviderResourceState

	Message khulnasoft.StringOutput `khulnasoft:"message"`
}

func NewProvider(ctx *khulnasoft.Context,
	name string, args *ProviderArgs, opts ...khulnasoft.ResourceOption,
) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}
	var resource Provider
	err := ctx.RegisterResource("khulnasoft:providers:testcomponent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	Message string `khulnasoft:"message"`
}

type ProviderArgs struct {
	Message khulnasoft.StringInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type Component struct {
	khulnasoft.ResourceState

	Message khulnasoft.StringOutput `khulnasoft:"message"`
}

func NewComponent(ctx *khulnasoft.Context, name string, opts ...khulnasoft.ResourceOption) (*Component, error) {
	var resource Component
	err := ctx.RegisterRemoteComponentResource("testcomponent:index:Component", name, nil, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		provider, err := NewProvider(ctx, "myprovider", &ProviderArgs{
			Message: khulnasoft.String("hello world"),
		})
		if err != nil {
			return err
		}

		component, err := NewComponent(ctx, "mycomponent", khulnasoft.Providers(provider))
		if err != nil {
			return err
		}

		ctx.Export("message", component.Message)

		return nil
	})
}
