// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.
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
	if args == nil {
		args = &RandomArgs{}
	}
	var resource Random
	err := ctx.RegisterResource("testprovider:index:Random", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type randomArgs struct {
	Length int `khulnasoft:"length"`
}

type RandomArgs struct {
	Length khulnasoft.IntInput
}

func (RandomArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*randomArgs)(nil)).Elem()
}
