// Copyright 2016-2023, Pulumi Corporation.  All rights reserved.
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

// A remote component resource.
type Component struct {
	khulnasoft.ResourceState

	Message khulnasoft.StringOutput `khulnasoft:"message"`
}

// Creates a remote component resource.
func NewComponent(ctx *khulnasoft.Context, name string, opts ...khulnasoft.ResourceOption) (*Component, error) {
	var resource Component
	err := ctx.RegisterRemoteComponentResource("testcomponent:index:Component", name, nil, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// A local component resource.
type LocalComponent struct {
	khulnasoft.ResourceState

	Message khulnasoft.StringOutput
}

// Creates a regular local component resource, which creates a child remote component resource.
func NewLocalComponent(ctx *khulnasoft.Context, name string, opts ...khulnasoft.ResourceOption) (*LocalComponent, error) {
	var resource LocalComponent
	err := ctx.RegisterComponentResource("my:index:LocalComponent", name, &resource, opts...)
	if err != nil {
		return nil, err
	}

	component, err := NewComponent(ctx, name+"-mycomponent", khulnasoft.Parent(&resource))
	if err != nil {
		return nil, err
	}
	resource.Message = component.Message

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

		component, err := NewComponent(ctx, "mycomponent", khulnasoft.Provider(provider))
		if err != nil {
			return err
		}

		localComponent, err := NewLocalComponent(ctx, "mylocalcomponent", khulnasoft.Providers(provider))
		if err != nil {
			return err
		}

		ctx.Export("message", component.Message)
		ctx.Export("nestedMessage", localComponent.Message)

		return nil
	})
}
