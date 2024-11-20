// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"context"
	"reflect"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"simple-resource-schema/example/internal"
)

type BarResource struct {
	khulnasoft.ResourceState

	Foo ResourceOutput `khulnasoft:"foo"`
}

// NewBarResource registers a new resource with the given unique name, arguments, and options.
func NewBarResource(ctx *khulnasoft.Context,
	name string, args *BarResourceArgs, opts ...khulnasoft.ResourceOption) (*BarResource, error) {
	if args == nil {
		args = &BarResourceArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BarResource
	err := ctx.RegisterRemoteComponentResource("bar::BarResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type barResourceArgs struct {
	Foo *Resource `khulnasoft:"foo"`
}

// The set of arguments for constructing a BarResource resource.
type BarResourceArgs struct {
	Foo ResourceInput
}

func (BarResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*barResourceArgs)(nil)).Elem()
}

type BarResourceInput interface {
	khulnasoft.Input

	ToBarResourceOutput() BarResourceOutput
	ToBarResourceOutputWithContext(ctx context.Context) BarResourceOutput
}

func (*BarResource) ElementType() reflect.Type {
	return reflect.TypeOf((**BarResource)(nil)).Elem()
}

func (i *BarResource) ToBarResourceOutput() BarResourceOutput {
	return i.ToBarResourceOutputWithContext(context.Background())
}

func (i *BarResource) ToBarResourceOutputWithContext(ctx context.Context) BarResourceOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(BarResourceOutput)
}

type BarResourceOutput struct{ *khulnasoft.OutputState }

func (BarResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BarResource)(nil)).Elem()
}

func (o BarResourceOutput) ToBarResourceOutput() BarResourceOutput {
	return o
}

func (o BarResourceOutput) ToBarResourceOutputWithContext(ctx context.Context) BarResourceOutput {
	return o
}

func (o BarResourceOutput) Foo() ResourceOutput {
	return o.ApplyT(func(v *BarResource) ResourceOutput { return v.Foo }).(ResourceOutput)
}

func init() {
	khulnasoft.RegisterInputType(reflect.TypeOf((*BarResourceInput)(nil)).Elem(), &BarResource{})
	khulnasoft.RegisterOutputType(BarResourceOutput{})
}
