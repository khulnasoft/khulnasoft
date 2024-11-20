// Copyright 2016-2024, Pulumi Corporation.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

type MyComponent struct {
	khulnasoft.ResourceState
}

func NewMyComponent(ctx *khulnasoft.Context, name string) (*MyComponent, error) {
	component := &MyComponent{}

	err := ctx.RegisterComponentResource("test:index:MyComponent", name, component)
	if err != nil {
		return nil, err
	}

	return component, nil
}

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		err := ctx.Log.Debug("A debug message", nil)
		if err != nil {
			return err
		}

		_, err = NewMyComponent(ctx, "mycomponent")
		return err
	})
}
