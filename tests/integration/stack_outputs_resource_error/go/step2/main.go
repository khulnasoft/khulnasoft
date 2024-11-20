// Copyright 2016-2024, Pulumi Corporation.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

type FailsOnCreate struct {
	khulnasoft.CustomResourceState

	Value khulnasoft.Float64Output `khulnasoft:"value"`
}

func NewFailsOnCreate(ctx *khulnasoft.Context, name string, opts ...khulnasoft.ResourceOption) (*FailsOnCreate, error) {
	var resource FailsOnCreate
	err := ctx.RegisterResource("testprovider:index:FailsOnCreate", name, nil, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		ctx.Export("xyz", khulnasoft.String("DEF"))
		res, _ := NewFailsOnCreate(ctx, "test")
		ctx.Export("foo", res.Value)
		return nil
	})
}
