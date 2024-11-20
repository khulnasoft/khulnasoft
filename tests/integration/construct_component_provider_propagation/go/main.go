//go:build !all
// +build !all

package main

import (
	"reflect"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		provider, err := NewRandomProvider(ctx, "explicit")
		if err != nil {
			return err
		}

		if _, err := NewComponent(ctx, "uses_default", nil); err != nil {
			return err
		}

		if _, err := NewComponent(ctx, "uses_provider", nil, khulnasoft.Provider(provider)); err != nil {
			return err
		}

		if _, err := NewComponent(ctx, "uses_providers", nil, khulnasoft.Providers(provider)); err != nil {
			return err
		}

		providerMap := map[string]khulnasoft.ProviderResource{
			"testprovider": provider,
		}
		if _, err := NewComponent(ctx, "uses_providers_map", nil, khulnasoft.ProviderMap(providerMap)); err != nil {
			return err
		}

		return nil
	})
}

type RandomProvider struct {
	khulnasoft.ProviderResourceState
}

func NewRandomProvider(ctx *khulnasoft.Context, name string) (*RandomProvider, error) {
	var provider RandomProvider
	err := ctx.RegisterResource("khulnasoft:providers:testprovider", "explicit", nil, &provider)
	return &provider, err
}

type Component struct {
	khulnasoft.ResourceState

	Result khulnasoft.StringOutput `khulnasoft:"result"`
}

func NewComponent(ctx *khulnasoft.Context, name string, args *ComponentArgs, opts ...khulnasoft.ResourceOption) (*Component, error) {
	var resource Component
	err := ctx.RegisterRemoteComponentResource("testcomponent:index:Component", name, args, &resource, opts...)
	return &resource, err
}

type ComponentArgs struct {
	Result khulnasoft.StringInput `khulnasoft:"result"`
}

func (ComponentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fooComponentArgs)(nil)).Elem()
}

type fooComponentArgs struct {
	Result khulnasoft.StringInput `khulnasoft:"result"`
}
