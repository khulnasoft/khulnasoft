// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"context"
	"reflect"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"regress-go-12971/example/internal"
)

type Provider struct {
	khulnasoft.ProviderResourceState
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *khulnasoft.Context,
	name string, args *ProviderArgs, opts ...khulnasoft.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}

	if args.Name == nil {
		if d := internal.GetEnvOrDefault(nil, nil, "WORLD_NAME"); d != nil {
			args.Name = khulnasoft.StringPtr(d.(string))
		}
	}
	if args.Populated == nil {
		if d := internal.GetEnvOrDefault(nil, internal.ParseEnvBool, "WORLD_POPULATED"); d != nil {
			args.Populated = khulnasoft.BoolPtr(d.(bool))
		}
	}
	if args.RadiusKm == nil {
		if d := internal.GetEnvOrDefault(nil, internal.ParseEnvFloat, "WORLD_RADIUS_KM"); d != nil {
			args.RadiusKm = khulnasoft.Float64Ptr(d.(float64))
		}
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Provider
	err := ctx.RegisterResource("khulnasoft:providers:world", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	Name      *string  `khulnasoft:"name"`
	Populated *bool    `khulnasoft:"populated"`
	RadiusKm  *float64 `khulnasoft:"radiusKm"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	Name      khulnasoft.StringPtrInput
	Populated khulnasoft.BoolPtrInput
	RadiusKm  khulnasoft.Float64PtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	khulnasoft.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(ProviderOutput)
}

type ProviderOutput struct{ *khulnasoft.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

func init() {
	khulnasoft.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	khulnasoft.RegisterOutputType(ProviderOutput{})
}
