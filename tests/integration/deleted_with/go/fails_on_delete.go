// Copyright 2016-2023, Pulumi Corporation.  All rights reserved.
//go:build !all
// +build !all

// Exposes the FailsOnDelete resource from the testprovider.

package main

import (
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

type FailsOnDelete struct {
	khulnasoft.CustomResourceState
}

func NewFailsOnDelete(ctx *khulnasoft.Context, name string, opts ...khulnasoft.ResourceOption) (*FailsOnDelete, error) {
	var resource FailsOnDelete
	err := ctx.RegisterResource("testprovider:index:FailsOnDelete", name, nil, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}
