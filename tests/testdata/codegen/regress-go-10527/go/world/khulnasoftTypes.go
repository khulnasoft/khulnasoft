// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package world

import (
	"context"
	"reflect"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"regress-go-10527/world/internal"
)

var _ = internal.GetEnvOrDefault

type World struct {
	Name *string `khulnasoft:"name"`
}

// WorldInput is an input type that accepts WorldArgs and WorldOutput values.
// You can construct a concrete instance of `WorldInput` via:
//
//	WorldArgs{...}
type WorldInput interface {
	khulnasoft.Input

	ToWorldOutput() WorldOutput
	ToWorldOutputWithContext(context.Context) WorldOutput
}

type WorldArgs struct {
	Name khulnasoft.StringPtrInput `khulnasoft:"name"`
}

func (WorldArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*World)(nil)).Elem()
}

func (i WorldArgs) ToWorldOutput() WorldOutput {
	return i.ToWorldOutputWithContext(context.Background())
}

func (i WorldArgs) ToWorldOutputWithContext(ctx context.Context) WorldOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(WorldOutput)
}

type WorldOutput struct{ *khulnasoft.OutputState }

func (WorldOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*World)(nil)).Elem()
}

func (o WorldOutput) ToWorldOutput() WorldOutput {
	return o
}

func (o WorldOutput) ToWorldOutputWithContext(ctx context.Context) WorldOutput {
	return o
}

func (o WorldOutput) Name() khulnasoft.StringPtrOutput {
	return o.ApplyT(func(v World) *string { return v.Name }).(khulnasoft.StringPtrOutput)
}

func init() {
	khulnasoft.RegisterInputType(reflect.TypeOf((*WorldInput)(nil)).Elem(), WorldArgs{})
	khulnasoft.RegisterOutputType(WorldOutput{})
	khulnasoft.RegisterOutputType(WorldMapOutput{})
}