// Copyright 2016-2022, Pulumi Corporation.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

type Resource struct {
	khulnasoft.ResourceState
}

type ComponentSix struct {
	khulnasoft.ResourceState
}

type ComponentSixParent struct {
	khulnasoft.ResourceState
}

func NewResource(ctx *khulnasoft.Context, name string, opts ...khulnasoft.ResourceOption) (*Resource, error) {
	comp := &Resource{}
	err := ctx.RegisterComponentResource("my:module:Resource", name, comp, opts...)
	if err != nil {
		return nil, err
	}
	return comp, nil
}

// Scenario #6 - Nested parents changing types
func NewComponentSix(ctx *khulnasoft.Context, name string, opts ...khulnasoft.ResourceOption) (*ComponentSix, error) {
	comp := &ComponentSix{}
	err := ctx.RegisterComponentResource("my:module:ComponentSix-v0", name, comp, opts...)
	if err != nil {
		return nil, err
	}
	parentOpt := khulnasoft.Parent(comp)
	_, err = NewResource(ctx, "otherchild", parentOpt)
	if err != nil {
		return nil, err
	}
	return comp, nil
}

func NewComponentSixParent(ctx *khulnasoft.Context, name string, opts ...khulnasoft.ResourceOption) (*ComponentSixParent, error) {
	comp := &ComponentSixParent{}
	err := ctx.RegisterComponentResource("my:module:ComponentSixParent-v0", name, comp, opts...)
	if err != nil {
		return nil, err
	}
	parentOpt := khulnasoft.Parent(comp)
	_, err = NewComponentSix(ctx, "child", parentOpt)
	if err != nil {
		return nil, err
	}
	return comp, nil
}

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		_, err := NewComponentSixParent(ctx, "comp6")
		if err != nil {
			return err
		}

		return nil
	})
}