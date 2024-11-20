// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"go-plain-ref-repro/repro/internal"
)

type FargateTaskDefinition struct {
	khulnasoft.ResourceState

	LoadBalancers khulnasoft.StringArrayOutput `khulnasoft:"loadBalancers"`
}

// NewFargateTaskDefinition registers a new resource with the given unique name, arguments, and options.
func NewFargateTaskDefinition(ctx *khulnasoft.Context,
	name string, args *FargateTaskDefinitionArgs, opts ...khulnasoft.ResourceOption) (*FargateTaskDefinition, error) {
	if args == nil {
		args = &FargateTaskDefinitionArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FargateTaskDefinition
	err := ctx.RegisterRemoteComponentResource("repro:ecs:FargateTaskDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type fargateTaskDefinitionArgs struct {
	Container  *TaskDefinitionContainerDefinition           `khulnasoft:"container"`
	Containers map[string]TaskDefinitionContainerDefinition `khulnasoft:"containers"`
}

// The set of arguments for constructing a FargateTaskDefinition resource.
type FargateTaskDefinitionArgs struct {
	Container  *TaskDefinitionContainerDefinitionArgs
	Containers map[string]TaskDefinitionContainerDefinitionArgs
}

func (FargateTaskDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fargateTaskDefinitionArgs)(nil)).Elem()
}

type FargateTaskDefinitionInput interface {
	khulnasoft.Input

	ToFargateTaskDefinitionOutput() FargateTaskDefinitionOutput
	ToFargateTaskDefinitionOutputWithContext(ctx context.Context) FargateTaskDefinitionOutput
}

func (*FargateTaskDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((**FargateTaskDefinition)(nil)).Elem()
}

func (i *FargateTaskDefinition) ToFargateTaskDefinitionOutput() FargateTaskDefinitionOutput {
	return i.ToFargateTaskDefinitionOutputWithContext(context.Background())
}

func (i *FargateTaskDefinition) ToFargateTaskDefinitionOutputWithContext(ctx context.Context) FargateTaskDefinitionOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(FargateTaskDefinitionOutput)
}

// FargateTaskDefinitionArrayInput is an input type that accepts FargateTaskDefinitionArray and FargateTaskDefinitionArrayOutput values.
// You can construct a concrete instance of `FargateTaskDefinitionArrayInput` via:
//
//	FargateTaskDefinitionArray{ FargateTaskDefinitionArgs{...} }
type FargateTaskDefinitionArrayInput interface {
	khulnasoft.Input

	ToFargateTaskDefinitionArrayOutput() FargateTaskDefinitionArrayOutput
	ToFargateTaskDefinitionArrayOutputWithContext(context.Context) FargateTaskDefinitionArrayOutput
}

type FargateTaskDefinitionArray []FargateTaskDefinitionInput

func (FargateTaskDefinitionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FargateTaskDefinition)(nil)).Elem()
}

func (i FargateTaskDefinitionArray) ToFargateTaskDefinitionArrayOutput() FargateTaskDefinitionArrayOutput {
	return i.ToFargateTaskDefinitionArrayOutputWithContext(context.Background())
}

func (i FargateTaskDefinitionArray) ToFargateTaskDefinitionArrayOutputWithContext(ctx context.Context) FargateTaskDefinitionArrayOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(FargateTaskDefinitionArrayOutput)
}

// FargateTaskDefinitionMapInput is an input type that accepts FargateTaskDefinitionMap and FargateTaskDefinitionMapOutput values.
// You can construct a concrete instance of `FargateTaskDefinitionMapInput` via:
//
//	FargateTaskDefinitionMap{ "key": FargateTaskDefinitionArgs{...} }
type FargateTaskDefinitionMapInput interface {
	khulnasoft.Input

	ToFargateTaskDefinitionMapOutput() FargateTaskDefinitionMapOutput
	ToFargateTaskDefinitionMapOutputWithContext(context.Context) FargateTaskDefinitionMapOutput
}

type FargateTaskDefinitionMap map[string]FargateTaskDefinitionInput

func (FargateTaskDefinitionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FargateTaskDefinition)(nil)).Elem()
}

func (i FargateTaskDefinitionMap) ToFargateTaskDefinitionMapOutput() FargateTaskDefinitionMapOutput {
	return i.ToFargateTaskDefinitionMapOutputWithContext(context.Background())
}

func (i FargateTaskDefinitionMap) ToFargateTaskDefinitionMapOutputWithContext(ctx context.Context) FargateTaskDefinitionMapOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(FargateTaskDefinitionMapOutput)
}

type FargateTaskDefinitionOutput struct{ *khulnasoft.OutputState }

func (FargateTaskDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FargateTaskDefinition)(nil)).Elem()
}

func (o FargateTaskDefinitionOutput) ToFargateTaskDefinitionOutput() FargateTaskDefinitionOutput {
	return o
}

func (o FargateTaskDefinitionOutput) ToFargateTaskDefinitionOutputWithContext(ctx context.Context) FargateTaskDefinitionOutput {
	return o
}

func (o FargateTaskDefinitionOutput) LoadBalancers() khulnasoft.StringArrayOutput {
	return o.ApplyT(func(v *FargateTaskDefinition) khulnasoft.StringArrayOutput { return v.LoadBalancers }).(khulnasoft.StringArrayOutput)
}

type FargateTaskDefinitionArrayOutput struct{ *khulnasoft.OutputState }

func (FargateTaskDefinitionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FargateTaskDefinition)(nil)).Elem()
}

func (o FargateTaskDefinitionArrayOutput) ToFargateTaskDefinitionArrayOutput() FargateTaskDefinitionArrayOutput {
	return o
}

func (o FargateTaskDefinitionArrayOutput) ToFargateTaskDefinitionArrayOutputWithContext(ctx context.Context) FargateTaskDefinitionArrayOutput {
	return o
}

func (o FargateTaskDefinitionArrayOutput) Index(i khulnasoft.IntInput) FargateTaskDefinitionOutput {
	return khulnasoft.All(o, i).ApplyT(func(vs []interface{}) *FargateTaskDefinition {
		return vs[0].([]*FargateTaskDefinition)[vs[1].(int)]
	}).(FargateTaskDefinitionOutput)
}

type FargateTaskDefinitionMapOutput struct{ *khulnasoft.OutputState }

func (FargateTaskDefinitionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FargateTaskDefinition)(nil)).Elem()
}

func (o FargateTaskDefinitionMapOutput) ToFargateTaskDefinitionMapOutput() FargateTaskDefinitionMapOutput {
	return o
}

func (o FargateTaskDefinitionMapOutput) ToFargateTaskDefinitionMapOutputWithContext(ctx context.Context) FargateTaskDefinitionMapOutput {
	return o
}

func (o FargateTaskDefinitionMapOutput) MapIndex(k khulnasoft.StringInput) FargateTaskDefinitionOutput {
	return khulnasoft.All(o, k).ApplyT(func(vs []interface{}) *FargateTaskDefinition {
		return vs[0].(map[string]*FargateTaskDefinition)[vs[1].(string)]
	}).(FargateTaskDefinitionOutput)
}

func init() {
	khulnasoft.RegisterInputType(reflect.TypeOf((*FargateTaskDefinitionInput)(nil)).Elem(), &FargateTaskDefinition{})
	khulnasoft.RegisterInputType(reflect.TypeOf((*FargateTaskDefinitionArrayInput)(nil)).Elem(), FargateTaskDefinitionArray{})
	khulnasoft.RegisterInputType(reflect.TypeOf((*FargateTaskDefinitionMapInput)(nil)).Elem(), FargateTaskDefinitionMap{})
	khulnasoft.RegisterOutputType(FargateTaskDefinitionOutput{})
	khulnasoft.RegisterOutputType(FargateTaskDefinitionArrayOutput{})
	khulnasoft.RegisterOutputType(FargateTaskDefinitionMapOutput{})
}
