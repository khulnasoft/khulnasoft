//go:build !all
// +build !all

package main

import (
	"errors"
	"reflect"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		r, err := NewRandom(ctx, "default", &RandomArgs{
			Length: khulnasoft.Int(10),
		}, khulnasoft.PluginDownloadURL("get.example.test"))
		if err != nil {
			return err
		}

		provider, err := NewProvider(ctx, "explicit",
			khulnasoft.PluginDownloadURL("get.khulnasoft.test/providers"))
		e, err := NewRandom(ctx, "explicit", &RandomArgs{
			Length: khulnasoft.Int(8),
		}, khulnasoft.Provider(provider))
		ctx.Export("default provider", r.Result)
		ctx.Export("explicit provider", e.Result)
		return nil
	})
}

type Random struct {
	khulnasoft.CustomResourceState

	Length khulnasoft.IntOutput    `khulnasoft:"length"`
	Result khulnasoft.StringOutput `khulnasoft:"result"`
}

func NewProvider(ctx *khulnasoft.Context, name string,
	opts ...khulnasoft.ResourceOption,
) (khulnasoft.ProviderResource, error) {
	provider := Provider{}
	err := ctx.RegisterResource("khulnasoft:providers:testprovider",
		"provider", nil, &provider, opts...)
	if err != nil {
		return nil, err
	}
	return &provider, nil
}

type Provider struct {
	khulnasoft.ProviderResourceState
}

func NewRandom(ctx *khulnasoft.Context,
	name string, args *RandomArgs, opts ...khulnasoft.ResourceOption,
) (*Random, error) {
	if args == nil || args.Length == nil {
		return nil, errors.New("missing required argument 'Length'")
	}
	if args == nil {
		args = &RandomArgs{}
	}
	var resource Random
	err := ctx.RegisterResource("testprovider:index:Random", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type randomArgs struct {
	Length int `khulnasoft:"length"`
}

type RandomArgs struct {
	Length khulnasoft.IntInput
}

func (RandomArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*randomArgs)(nil)).Elem()
}
