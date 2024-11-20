// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"context"
	"reflect"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"naming-collisions/example/internal"
)

type ResourceInputResource struct {
	khulnasoft.CustomResourceState

	Bar khulnasoft.StringPtrOutput `khulnasoft:"bar"`
}

// NewResourceInputResource registers a new resource with the given unique name, arguments, and options.
func NewResourceInputResource(ctx *khulnasoft.Context,
	name string, args *ResourceInputResourceArgs, opts ...khulnasoft.ResourceOption) (*ResourceInputResource, error) {
	if args == nil {
		args = &ResourceInputResourceArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ResourceInputResource
	err := ctx.RegisterResource("example::ResourceInput", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceInputResource gets an existing ResourceInputResource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceInputResource(ctx *khulnasoft.Context,
	name string, id khulnasoft.IDInput, state *ResourceInputResourceState, opts ...khulnasoft.ResourceOption) (*ResourceInputResource, error) {
	var resource ResourceInputResource
	err := ctx.ReadResource("example::ResourceInput", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceInputResource resources.
type resourceInputResourceState struct {
}

type ResourceInputResourceState struct {
}

func (ResourceInputResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceInputResourceState)(nil)).Elem()
}

type resourceInputResourceArgs struct {
}

// The set of arguments for constructing a ResourceInputResource resource.
type ResourceInputResourceArgs struct {
}

func (ResourceInputResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceInputResourceArgs)(nil)).Elem()
}

type ResourceInputResourceInput interface {
	khulnasoft.Input

	ToResourceInputResourceOutput() ResourceInputResourceOutput
	ToResourceInputResourceOutputWithContext(ctx context.Context) ResourceInputResourceOutput
}

func (*ResourceInputResource) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceInputResource)(nil)).Elem()
}

func (i *ResourceInputResource) ToResourceInputResourceOutput() ResourceInputResourceOutput {
	return i.ToResourceInputResourceOutputWithContext(context.Background())
}

func (i *ResourceInputResource) ToResourceInputResourceOutputWithContext(ctx context.Context) ResourceInputResourceOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(ResourceInputResourceOutput)
}

type ResourceInputResourceOutput struct{ *khulnasoft.OutputState }

func (ResourceInputResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceInputResource)(nil)).Elem()
}

func (o ResourceInputResourceOutput) ToResourceInputResourceOutput() ResourceInputResourceOutput {
	return o
}

func (o ResourceInputResourceOutput) ToResourceInputResourceOutputWithContext(ctx context.Context) ResourceInputResourceOutput {
	return o
}

func (o ResourceInputResourceOutput) Bar() khulnasoft.StringPtrOutput {
	return o.ApplyT(func(v *ResourceInputResource) khulnasoft.StringPtrOutput { return v.Bar }).(khulnasoft.StringPtrOutput)
}

func init() {
	khulnasoft.RegisterInputType(reflect.TypeOf((*ResourceInputResourceInput)(nil)).Elem(), &ResourceInputResource{})
	khulnasoft.RegisterOutputType(ResourceInputResourceOutput{})
}
