// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

type BucketComponent struct {
	khulnasoft.ResourceState
}

func NewBucketComponent(ctx *khulnasoft.Context, name string, opts ...khulnasoft.ResourceOption) (*BucketComponent, error) {
	component := &BucketComponent{}
	err := ctx.RegisterRemoteComponentResource("wibble:index:BucketComponent", name, nil, component, opts...)
	if err != nil {
		return nil, err
	}
	return component, nil
}

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		_, err := NewBucketComponent(ctx, "main-bucket")
		return err
	})
}
