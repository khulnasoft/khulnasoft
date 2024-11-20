// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"fmt"

	"github.com/khulnasoft/khulnasoft/pkg/v3/resource/provider"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/common/util/cmdutil"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	khulnasoftprovider "github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft/provider"
)

type Component struct {
	khulnasoft.ResourceState
}

func NewComponent(ctx *khulnasoft.Context, name string, opts ...khulnasoft.ResourceOption) (*Component, error) {
	component := &Component{}
	err := ctx.RegisterComponentResource("testcomponent:index:Component", name, component, opts...)
	if err != nil {
		return nil, err
	}

	if err := ctx.RegisterResourceOutputs(component, khulnasoft.Map{}); err != nil {
		return nil, err
	}

	return component, nil
}

const (
	providerName = "testcomponent"
	version      = "0.0.1"
)

func main() {
	if err := provider.MainWithOptions(provider.Options{
		Name:    providerName,
		Version: version,
		Construct: func(ctx *khulnasoft.Context, typ, name string, inputs khulnasoftprovider.ConstructInputs,
			options khulnasoft.ResourceOption,
		) (*khulnasoftprovider.ConstructResult, error) {
			if typ != "testcomponent:index:Component" {
				return nil, fmt.Errorf("unknown resource type %s", typ)
			}

			component, err := NewComponent(ctx, name, options)
			if err != nil {
				return nil, fmt.Errorf("creating component: %w", err)
			}

			return khulnasoftprovider.NewConstructResult(component)
		},
		Call: func(ctx *khulnasoft.Context, tok string, args khulnasoftprovider.CallArgs) (*khulnasoftprovider.CallResult, error) {
			if tok != "testcomponent:index:Component/getMessage" {
				return nil, fmt.Errorf("unknown method %s", tok)
			}

			return &khulnasoftprovider.CallResult{
				Failures: []khulnasoftprovider.CallFailure{
					{
						Property: "the failure property",
						Reason:   "the failure reason",
					},
				},
			}, nil
		},
	}); err != nil {
		cmdutil.ExitError(err.Error())
	}
}
