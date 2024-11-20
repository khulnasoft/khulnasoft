// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

type BucketComponentV2 struct {
	khulnasoft.ResourceState
}

func NewBucketComponentV2(ctx *khulnasoft.Context, name string, opts ...khulnasoft.ResourceOption) (*BucketComponentV2, error) {
	component := &BucketComponentV2{}
	err := ctx.RegisterRemoteComponentResource("wibble:index:BucketComponentV2", name, nil, component, opts...)
	if err != nil {
		return nil, err
	}
	return component, nil
}

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		_, err := NewBucketComponentV2(ctx, "main-bucket")
		return err
	})
}
