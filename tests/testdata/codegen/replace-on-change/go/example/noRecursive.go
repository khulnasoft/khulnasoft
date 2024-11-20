// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"context"
	"reflect"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"replace-on-change/example/internal"
)

type NoRecursive struct {
	khulnasoft.CustomResourceState

	Rec       RecPtrOutput           `khulnasoft:"rec"`
	ReplaceMe khulnasoft.StringPtrOutput `khulnasoft:"replaceMe"`
}

// NewNoRecursive registers a new resource with the given unique name, arguments, and options.
func NewNoRecursive(ctx *khulnasoft.Context,
	name string, args *NoRecursiveArgs, opts ...khulnasoft.ResourceOption) (*NoRecursive, error) {
	if args == nil {
		args = &NoRecursiveArgs{}
	}

	replaceOnChanges := khulnasoft.ReplaceOnChanges([]string{
		"replaceMe",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NoRecursive
	err := ctx.RegisterResource("example::NoRecursive", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNoRecursive gets an existing NoRecursive resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNoRecursive(ctx *khulnasoft.Context,
	name string, id khulnasoft.IDInput, state *NoRecursiveState, opts ...khulnasoft.ResourceOption) (*NoRecursive, error) {
	var resource NoRecursive
	err := ctx.ReadResource("example::NoRecursive", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NoRecursive resources.
type noRecursiveState struct {
}

type NoRecursiveState struct {
}

func (NoRecursiveState) ElementType() reflect.Type {
	return reflect.TypeOf((*noRecursiveState)(nil)).Elem()
}

type noRecursiveArgs struct {
}

// The set of arguments for constructing a NoRecursive resource.
type NoRecursiveArgs struct {
}

func (NoRecursiveArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*noRecursiveArgs)(nil)).Elem()
}

type NoRecursiveInput interface {
	khulnasoft.Input

	ToNoRecursiveOutput() NoRecursiveOutput
	ToNoRecursiveOutputWithContext(ctx context.Context) NoRecursiveOutput
}

func (*NoRecursive) ElementType() reflect.Type {
	return reflect.TypeOf((**NoRecursive)(nil)).Elem()
}

func (i *NoRecursive) ToNoRecursiveOutput() NoRecursiveOutput {
	return i.ToNoRecursiveOutputWithContext(context.Background())
}

func (i *NoRecursive) ToNoRecursiveOutputWithContext(ctx context.Context) NoRecursiveOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(NoRecursiveOutput)
}

// NoRecursiveArrayInput is an input type that accepts NoRecursiveArray and NoRecursiveArrayOutput values.
// You can construct a concrete instance of `NoRecursiveArrayInput` via:
//
//	NoRecursiveArray{ NoRecursiveArgs{...} }
type NoRecursiveArrayInput interface {
	khulnasoft.Input

	ToNoRecursiveArrayOutput() NoRecursiveArrayOutput
	ToNoRecursiveArrayOutputWithContext(context.Context) NoRecursiveArrayOutput
}

type NoRecursiveArray []NoRecursiveInput

func (NoRecursiveArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NoRecursive)(nil)).Elem()
}

func (i NoRecursiveArray) ToNoRecursiveArrayOutput() NoRecursiveArrayOutput {
	return i.ToNoRecursiveArrayOutputWithContext(context.Background())
}

func (i NoRecursiveArray) ToNoRecursiveArrayOutputWithContext(ctx context.Context) NoRecursiveArrayOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(NoRecursiveArrayOutput)
}

// NoRecursiveMapInput is an input type that accepts NoRecursiveMap and NoRecursiveMapOutput values.
// You can construct a concrete instance of `NoRecursiveMapInput` via:
//
//	NoRecursiveMap{ "key": NoRecursiveArgs{...} }
type NoRecursiveMapInput interface {
	khulnasoft.Input

	ToNoRecursiveMapOutput() NoRecursiveMapOutput
	ToNoRecursiveMapOutputWithContext(context.Context) NoRecursiveMapOutput
}

type NoRecursiveMap map[string]NoRecursiveInput

func (NoRecursiveMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NoRecursive)(nil)).Elem()
}

func (i NoRecursiveMap) ToNoRecursiveMapOutput() NoRecursiveMapOutput {
	return i.ToNoRecursiveMapOutputWithContext(context.Background())
}

func (i NoRecursiveMap) ToNoRecursiveMapOutputWithContext(ctx context.Context) NoRecursiveMapOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(NoRecursiveMapOutput)
}

type NoRecursiveOutput struct{ *khulnasoft.OutputState }

func (NoRecursiveOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NoRecursive)(nil)).Elem()
}

func (o NoRecursiveOutput) ToNoRecursiveOutput() NoRecursiveOutput {
	return o
}

func (o NoRecursiveOutput) ToNoRecursiveOutputWithContext(ctx context.Context) NoRecursiveOutput {
	return o
}

func (o NoRecursiveOutput) Rec() RecPtrOutput {
	return o.ApplyT(func(v *NoRecursive) RecPtrOutput { return v.Rec }).(RecPtrOutput)
}

func (o NoRecursiveOutput) ReplaceMe() khulnasoft.StringPtrOutput {
	return o.ApplyT(func(v *NoRecursive) khulnasoft.StringPtrOutput { return v.ReplaceMe }).(khulnasoft.StringPtrOutput)
}

type NoRecursiveArrayOutput struct{ *khulnasoft.OutputState }

func (NoRecursiveArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NoRecursive)(nil)).Elem()
}

func (o NoRecursiveArrayOutput) ToNoRecursiveArrayOutput() NoRecursiveArrayOutput {
	return o
}

func (o NoRecursiveArrayOutput) ToNoRecursiveArrayOutputWithContext(ctx context.Context) NoRecursiveArrayOutput {
	return o
}

func (o NoRecursiveArrayOutput) Index(i khulnasoft.IntInput) NoRecursiveOutput {
	return khulnasoft.All(o, i).ApplyT(func(vs []interface{}) *NoRecursive {
		return vs[0].([]*NoRecursive)[vs[1].(int)]
	}).(NoRecursiveOutput)
}

type NoRecursiveMapOutput struct{ *khulnasoft.OutputState }

func (NoRecursiveMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NoRecursive)(nil)).Elem()
}

func (o NoRecursiveMapOutput) ToNoRecursiveMapOutput() NoRecursiveMapOutput {
	return o
}

func (o NoRecursiveMapOutput) ToNoRecursiveMapOutputWithContext(ctx context.Context) NoRecursiveMapOutput {
	return o
}

func (o NoRecursiveMapOutput) MapIndex(k khulnasoft.StringInput) NoRecursiveOutput {
	return khulnasoft.All(o, k).ApplyT(func(vs []interface{}) *NoRecursive {
		return vs[0].(map[string]*NoRecursive)[vs[1].(string)]
	}).(NoRecursiveOutput)
}

func init() {
	khulnasoft.RegisterOutputType(NoRecursiveOutput{})
	khulnasoft.RegisterOutputType(NoRecursiveArrayOutput{})
	khulnasoft.RegisterOutputType(NoRecursiveMapOutput{})
}