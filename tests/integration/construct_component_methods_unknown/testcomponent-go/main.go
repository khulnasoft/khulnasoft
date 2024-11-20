// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"fmt"

	"github.com/blang/semver"

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

type ComponentGetMessageArgs struct {
	Echo khulnasoft.StringInput `khulnasoft:"echo"`
}

type ComponentGetMessageResult struct {
	Message khulnasoft.StringOutput `khulnasoft:"message"`
}

func (c *Component) GetMessage(args *ComponentGetMessageArgs) (*ComponentGetMessageResult, error) {
	message := args.Echo.ToStringOutput().ApplyT(func(val string) string {
		panic("should not run (echo)")
	}).(khulnasoft.StringOutput)

	return &ComponentGetMessageResult{
		Message: message,
	}, nil
}

const (
	providerName = "testcomponent"
	version      = "0.0.1"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *khulnasoft.Context, name, typ, urn string) (r khulnasoft.Resource, err error) {
	switch typ {
	case "testcomponent:index:Component":
		r = &Component{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, khulnasoft.URN_(urn))
	return
}

func main() {
	// Register any resources that can come back as resource references that need to be rehydrated.
	khulnasoft.RegisterResourceModule("testcomponent", "index", &module{semver.MustParse(version)})

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

			methodArgs := &ComponentGetMessageArgs{}
			res, err := args.CopyTo(methodArgs)
			if err != nil {
				return nil, fmt.Errorf("setting args: %w", err)
			}
			component := res.(*Component)

			result, err := component.GetMessage(methodArgs)
			if err != nil {
				return nil, fmt.Errorf("calling method: %w", err)
			}

			return khulnasoftprovider.NewCallResult(result)
		},
	}); err != nil {
		cmdutil.ExitError(err.Error())
	}
}
