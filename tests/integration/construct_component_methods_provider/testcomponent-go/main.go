// Copyright 2016-2023, Pulumi Corporation.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"errors"
	"fmt"

	"github.com/blang/semver"

	"github.com/khulnasoft/khulnasoft/pkg/v3/resource/provider"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/common/util/cmdutil"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	khulnasoftprovider "github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft/provider"
)

type Component struct {
	khulnasoft.ResourceState
	First  khulnasoft.StringOutput `khulnasoft:"first"`
	Second khulnasoft.StringOutput `khulnasoft:"second"`
}

type ComponentArgs struct {
	First  khulnasoft.StringInput `khulnasoft:"first"`
	Second khulnasoft.StringInput `khulnasoft:"second"`
}

func NewComponent(ctx *khulnasoft.Context, name string, args *ComponentArgs,
	opts ...khulnasoft.ResourceOption,
) (*Component, error) {
	if args == nil {
		return nil, errors.New("args is required")
	}

	component := &Component{}
	err := ctx.RegisterComponentResource("testcomponent:index:Component", name, component, opts...)
	if err != nil {
		return nil, err
	}

	component.First = args.First.ToStringOutput()
	component.Second = args.Second.ToStringOutput()

	if err := ctx.RegisterResourceOutputs(component, khulnasoft.Map{
		"first":  args.First,
		"second": args.Second,
	}); err != nil {
		return nil, err
	}

	return component, nil
}

type ComponentGetMessageArgs struct {
	Name khulnasoft.StringInput `khulnasoft:"name"`
}

type ComponentGetMessageResult struct {
	Message khulnasoft.StringOutput `khulnasoft:"message"`
}

func (c *Component) GetMessage(args *ComponentGetMessageArgs) (*ComponentGetMessageResult, error) {
	return &ComponentGetMessageResult{
		Message: khulnasoft.Sprintf("%s %s, %s!", c.First, c.Second, args.Name),
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

			args := &ComponentArgs{}
			if err := inputs.CopyTo(args); err != nil {
				return nil, fmt.Errorf("setting args: %w", err)
			}

			component, err := NewComponent(ctx, name, args, options)
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
