// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"context"
	"reflect"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"simple-yaml-schema/example/internal"
)

type Resource struct {
	khulnasoft.CustomResourceState

	Bar khulnasoft.StringPtrOutput `khulnasoft:"bar"`
}

// NewResource registers a new resource with the given unique name, arguments, and options.
func NewResource(ctx *khulnasoft.Context,
	name string, args *ResourceArgs, opts ...khulnasoft.ResourceOption) (*Resource, error) {
	if args == nil {
		args = &ResourceArgs{}
	}

	if args.Bar != nil {
		args.Bar = khulnasoft.ToSecret(args.Bar).(khulnasoft.StringPtrInput)
	}
	secrets := khulnasoft.AdditionalSecretOutputs([]string{
		"bar",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Resource
	err := ctx.RegisterResource("example::Resource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResource gets an existing Resource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResource(ctx *khulnasoft.Context,
	name string, id khulnasoft.IDInput, state *ResourceState, opts ...khulnasoft.ResourceOption) (*Resource, error) {
	var resource Resource
	err := ctx.ReadResource("example::Resource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Resource resources.
type resourceState struct {
}

type ResourceState struct {
}

func (ResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceState)(nil)).Elem()
}

type resourceArgs struct {
	Bar *string `khulnasoft:"bar"`
}

// The set of arguments for constructing a Resource resource.
type ResourceArgs struct {
	Bar khulnasoft.StringPtrInput
}

func (ResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceArgs)(nil)).Elem()
}

type ResourceInput interface {
	khulnasoft.Input

	ToResourceOutput() ResourceOutput
	ToResourceOutputWithContext(ctx context.Context) ResourceOutput
}

func (*Resource) ElementType() reflect.Type {
	return reflect.TypeOf((**Resource)(nil)).Elem()
}

func (i *Resource) ToResourceOutput() ResourceOutput {
	return i.ToResourceOutputWithContext(context.Background())
}

func (i *Resource) ToResourceOutputWithContext(ctx context.Context) ResourceOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(ResourceOutput)
}

type ResourceOutput struct{ *khulnasoft.OutputState }

func (ResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Resource)(nil)).Elem()
}

func (o ResourceOutput) ToResourceOutput() ResourceOutput {
	return o
}

func (o ResourceOutput) ToResourceOutputWithContext(ctx context.Context) ResourceOutput {
	return o
}

func (o ResourceOutput) Bar() khulnasoft.StringPtrOutput {
	return o.ApplyT(func(v *Resource) khulnasoft.StringPtrOutput { return v.Bar }).(khulnasoft.StringPtrOutput)
}

func init() {
	khulnasoft.RegisterInputType(reflect.TypeOf((*ResourceInput)(nil)).Elem(), &Resource{})
	khulnasoft.RegisterOutputType(ResourceOutput{})
}
