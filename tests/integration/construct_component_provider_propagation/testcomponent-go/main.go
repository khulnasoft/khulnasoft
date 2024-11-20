// Copyright 2016-2023, Pulumi Corporation.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/khulnasoft/khulnasoft/pkg/v3/resource/provider"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/common/util/cmdutil"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	khulnasoftprovider "github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft/provider"
)

func main() {
	if err := provider.ComponentMain("testcomponent", "0.0.1", nil /* schema */, construct); err != nil {
		cmdutil.Exit(err)
	}
}

func construct(
	ctx *khulnasoft.Context,
	typ, name string,
	inputs khulnasoftprovider.ConstructInputs,
	options khulnasoft.ResourceOption,
) (*khulnasoftprovider.ConstructResult, error) {
	if typ != "testcomponent:index:Component" {
		return nil, fmt.Errorf("unknown resource type %q", typ)
	}

	comp, err := NewComponent(ctx, name, options)
	if err != nil {
		return nil, err
	}

	return khulnasoftprovider.NewConstructResult(comp)
}

// Component is a component resource.
//
// It's exposed to other SDKs from 'construct' above.
type Component struct {
	khulnasoft.ResourceState

	Result khulnasoft.StringOutput `khulnasoft:"result"`
}

// NewComponent builds a new component resource with the given name.
//
// It will instantiate a random resource as a child of the component
// with the same name.
func NewComponent(ctx *khulnasoft.Context, name string, opts ...khulnasoft.ResourceOption) (*Component, error) {
	var comp Component
	if err := ctx.RegisterComponentResource("testcomponent:index:Component", name, &comp, opts...); err != nil {
		return nil, err
	}

	r, err := NewRandom(ctx, name, &RandomArgs{Length: khulnasoft.Int(10)}, khulnasoft.Parent(&comp))
	if err != nil {
		return nil, err
	}

	comp.Result = r.Result
	return &comp, ctx.RegisterResourceOutputs(&comp, khulnasoft.Map{
		"result": comp.Result,
	})
}

// Random is a custom resource that generates a random string.
//
// It's implemented in the tests/testprovider directory.
// This is a Go-level reference to that resource.
type Random struct {
	khulnasoft.CustomResourceState

	Length khulnasoft.IntOutput    `khulnasoft:"length"`
	Result khulnasoft.StringOutput `khulnasoft:"result"`
}

// NewRandom builds a new random resource with the given name.
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

// RandomArgs specifies the parameters for a Random resource.
type RandomArgs struct {
	// Length of the random string to generate.
	Length khulnasoft.IntInput
}

// ElementType implements the khulnasoft.Input interface.
func (RandomArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*randomArgs)(nil)).Elem()
}

type randomArgs struct {
	Length int `khulnasoft:"length"`
}
