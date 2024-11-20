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

type Provider struct {
	khulnasoft.ProviderResourceState

	Message khulnasoft.StringOutput `khulnasoft:"message"`
}

type Component struct {
	khulnasoft.ResourceState

	Message khulnasoft.StringOutput `khulnasoft:"message"`
}

func NewComponent(ctx *khulnasoft.Context, name string, opts ...khulnasoft.ResourceOption) (*Component, error) {
	component := &Component{}
	err := ctx.RegisterComponentResource("testcomponent:index:Component", name, component, opts...)
	if err != nil {
		return nil, err
	}

	// Test that we're indeed getting back an instance of `Provider` with its state.
	provider := component.GetProvider("testcomponent::").(*Provider)
	component.Message = provider.Message

	if err := ctx.RegisterResourceOutputs(component, khulnasoft.Map{
		"message": component.Message,
	}); err != nil {
		return nil, err
	}

	return component, nil
}

const (
	providerName = "testcomponent"
	version      = "0.0.1"
)

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *khulnasoft.Context, name, typ, urn string) (khulnasoft.ProviderResource, error) {
	if typ != "khulnasoft:providers:testcomponent" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, khulnasoft.URN_(urn))
	return r, err
}

func main() {
	khulnasoft.RegisterResourcePackage(providerName, &pkg{semver.MustParse(version)})

	if err := provider.ComponentMain(providerName, version, nil, func(ctx *khulnasoft.Context, typ, name string,
		inputs khulnasoftprovider.ConstructInputs, options khulnasoft.ResourceOption,
	) (*khulnasoftprovider.ConstructResult, error) {
		if typ != "testcomponent:index:Component" {
			return nil, fmt.Errorf("unknown resource type %s", typ)
		}

		component, err := NewComponent(ctx, name, options)
		if err != nil {
			return nil, fmt.Errorf("creating component: %w", err)
		}

		return khulnasoftprovider.NewConstructResult(component)
	}); err != nil {
		cmdutil.ExitError(err.Error())
	}
}
