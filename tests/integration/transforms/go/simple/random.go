// Copyright 2016-2024, Pulumi Corporation.  All rights reserved.
//go:build !all
// +build !all

// Exposes the Random resource from the testprovider.

package main

import (
	"errors"
	"reflect"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

type Random struct {
	khulnasoft.CustomResourceState

	Length khulnasoft.IntOutput    `khulnasoft:"length"`
	Result khulnasoft.StringOutput `khulnasoft:"result"`
}

func NewRandom(ctx *khulnasoft.Context,
	name string, args *RandomArgs, opts ...khulnasoft.ResourceOption,
) (*Random, error) {
	if args == nil || args.Length == nil {
		return nil, errors.New("missing required argument 'Length'")
	}
	var resource Random
	err := ctx.RegisterResource("testprovider:index:Random", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

func (r *Random) RandomInvoke(ctx *khulnasoft.Context, args map[string]interface{}) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := ctx.Invoke("testprovider:index:returnArgs", args, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type randomArgs struct {
	Length int    `khulnasoft:"length"`
	Prefix string `khulnasoft:"prefix"`
}

type RandomArgs struct {
	Length khulnasoft.IntInput
	Prefix khulnasoft.StringInput
}

func (RandomArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*randomArgs)(nil)).Elem()
}

type Component struct {
	khulnasoft.ResourceState

	Length  khulnasoft.IntOutput    `khulnasoft:"length"`
	ChildID khulnasoft.StringOutput `khulnasoft:"childId"`
}

func NewComponent(ctx *khulnasoft.Context,
	name string, args *ComponentArgs, opts ...khulnasoft.ResourceOption,
) (*Random, error) {
	if args == nil || args.Length == nil {
		return nil, errors.New("missing required argument 'Length'")
	}
	var resource Random
	err := ctx.RegisterRemoteComponentResource("testprovider:index:Component", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type componentArgs struct {
	Length int `khulnasoft:"length"`
}

type ComponentArgs struct {
	Length khulnasoft.IntInput
}

func (ComponentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentArgs)(nil)).Elem()
}

type Provider struct {
	khulnasoft.ProviderResourceState
}

func NewProvider(ctx *khulnasoft.Context,
	name string, opts ...khulnasoft.ResourceOption,
) (*Provider, error) {
	var resource Provider
	err := ctx.RegisterResource("khulnasoft:providers:testprovider", name, nil, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}
