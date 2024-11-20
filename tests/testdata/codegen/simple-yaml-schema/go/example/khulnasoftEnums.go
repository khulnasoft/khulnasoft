// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"context"
	"reflect"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

type OutputOnlyEnumType string

const (
	OutputOnlyEnumTypeFoo = OutputOnlyEnumType("foo")
	OutputOnlyEnumTypeBar = OutputOnlyEnumType("bar")
)

type OutputOnlyEnumTypeOutput struct{ *khulnasoft.OutputState }

func (OutputOnlyEnumTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OutputOnlyEnumType)(nil)).Elem()
}

func (o OutputOnlyEnumTypeOutput) ToOutputOnlyEnumTypeOutput() OutputOnlyEnumTypeOutput {
	return o
}

func (o OutputOnlyEnumTypeOutput) ToOutputOnlyEnumTypeOutputWithContext(ctx context.Context) OutputOnlyEnumTypeOutput {
	return o
}

func (o OutputOnlyEnumTypeOutput) ToOutputOnlyEnumTypePtrOutput() OutputOnlyEnumTypePtrOutput {
	return o.ToOutputOnlyEnumTypePtrOutputWithContext(context.Background())
}

func (o OutputOnlyEnumTypeOutput) ToOutputOnlyEnumTypePtrOutputWithContext(ctx context.Context) OutputOnlyEnumTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OutputOnlyEnumType) *OutputOnlyEnumType {
		return &v
	}).(OutputOnlyEnumTypePtrOutput)
}

func (o OutputOnlyEnumTypeOutput) ToStringOutput() khulnasoft.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OutputOnlyEnumTypeOutput) ToStringOutputWithContext(ctx context.Context) khulnasoft.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OutputOnlyEnumType) string {
		return string(e)
	}).(khulnasoft.StringOutput)
}

func (o OutputOnlyEnumTypeOutput) ToStringPtrOutput() khulnasoft.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OutputOnlyEnumTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) khulnasoft.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OutputOnlyEnumType) *string {
		v := string(e)
		return &v
	}).(khulnasoft.StringPtrOutput)
}

type OutputOnlyEnumTypePtrOutput struct{ *khulnasoft.OutputState }

func (OutputOnlyEnumTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OutputOnlyEnumType)(nil)).Elem()
}

func (o OutputOnlyEnumTypePtrOutput) ToOutputOnlyEnumTypePtrOutput() OutputOnlyEnumTypePtrOutput {
	return o
}

func (o OutputOnlyEnumTypePtrOutput) ToOutputOnlyEnumTypePtrOutputWithContext(ctx context.Context) OutputOnlyEnumTypePtrOutput {
	return o
}

func (o OutputOnlyEnumTypePtrOutput) Elem() OutputOnlyEnumTypeOutput {
	return o.ApplyT(func(v *OutputOnlyEnumType) OutputOnlyEnumType {
		if v != nil {
			return *v
		}
		var ret OutputOnlyEnumType
		return ret
	}).(OutputOnlyEnumTypeOutput)
}

func (o OutputOnlyEnumTypePtrOutput) ToStringPtrOutput() khulnasoft.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OutputOnlyEnumTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) khulnasoft.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OutputOnlyEnumType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(khulnasoft.StringPtrOutput)
}

type OutputOnlyEnumTypeMapOutput struct{ *khulnasoft.OutputState }

func (OutputOnlyEnumTypeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]OutputOnlyEnumType)(nil)).Elem()
}

func (o OutputOnlyEnumTypeMapOutput) ToOutputOnlyEnumTypeMapOutput() OutputOnlyEnumTypeMapOutput {
	return o
}

func (o OutputOnlyEnumTypeMapOutput) ToOutputOnlyEnumTypeMapOutputWithContext(ctx context.Context) OutputOnlyEnumTypeMapOutput {
	return o
}

func (o OutputOnlyEnumTypeMapOutput) MapIndex(k khulnasoft.StringInput) OutputOnlyEnumTypeOutput {
	return khulnasoft.All(o, k).ApplyT(func(vs []interface{}) OutputOnlyEnumType {
		return vs[0].(map[string]OutputOnlyEnumType)[vs[1].(string)]
	}).(OutputOnlyEnumTypeOutput)
}

// types of rubber trees
type RubberTreeVariety string

const (
	// A burgundy rubber tree.
	RubberTreeVarietyBurgundy = RubberTreeVariety("Burgundy")
	// A ruby rubber tree.
	RubberTreeVarietyRuby = RubberTreeVariety("Ruby")
	// A tineke rubber tree.
	RubberTreeVarietyTineke = RubberTreeVariety("Tineke")
)

func (RubberTreeVariety) ElementType() reflect.Type {
	return reflect.TypeOf((*RubberTreeVariety)(nil)).Elem()
}

func (e RubberTreeVariety) ToRubberTreeVarietyOutput() RubberTreeVarietyOutput {
	return khulnasoft.ToOutput(e).(RubberTreeVarietyOutput)
}

func (e RubberTreeVariety) ToRubberTreeVarietyOutputWithContext(ctx context.Context) RubberTreeVarietyOutput {
	return khulnasoft.ToOutputWithContext(ctx, e).(RubberTreeVarietyOutput)
}

func (e RubberTreeVariety) ToRubberTreeVarietyPtrOutput() RubberTreeVarietyPtrOutput {
	return e.ToRubberTreeVarietyPtrOutputWithContext(context.Background())
}

func (e RubberTreeVariety) ToRubberTreeVarietyPtrOutputWithContext(ctx context.Context) RubberTreeVarietyPtrOutput {
	return RubberTreeVariety(e).ToRubberTreeVarietyOutputWithContext(ctx).ToRubberTreeVarietyPtrOutputWithContext(ctx)
}

func (e RubberTreeVariety) ToStringOutput() khulnasoft.StringOutput {
	return khulnasoft.ToOutput(khulnasoft.String(e)).(khulnasoft.StringOutput)
}

func (e RubberTreeVariety) ToStringOutputWithContext(ctx context.Context) khulnasoft.StringOutput {
	return khulnasoft.ToOutputWithContext(ctx, khulnasoft.String(e)).(khulnasoft.StringOutput)
}

func (e RubberTreeVariety) ToStringPtrOutput() khulnasoft.StringPtrOutput {
	return khulnasoft.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RubberTreeVariety) ToStringPtrOutputWithContext(ctx context.Context) khulnasoft.StringPtrOutput {
	return khulnasoft.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RubberTreeVarietyOutput struct{ *khulnasoft.OutputState }

func (RubberTreeVarietyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RubberTreeVariety)(nil)).Elem()
}

func (o RubberTreeVarietyOutput) ToRubberTreeVarietyOutput() RubberTreeVarietyOutput {
	return o
}

func (o RubberTreeVarietyOutput) ToRubberTreeVarietyOutputWithContext(ctx context.Context) RubberTreeVarietyOutput {
	return o
}

func (o RubberTreeVarietyOutput) ToRubberTreeVarietyPtrOutput() RubberTreeVarietyPtrOutput {
	return o.ToRubberTreeVarietyPtrOutputWithContext(context.Background())
}

func (o RubberTreeVarietyOutput) ToRubberTreeVarietyPtrOutputWithContext(ctx context.Context) RubberTreeVarietyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RubberTreeVariety) *RubberTreeVariety {
		return &v
	}).(RubberTreeVarietyPtrOutput)
}

func (o RubberTreeVarietyOutput) ToStringOutput() khulnasoft.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RubberTreeVarietyOutput) ToStringOutputWithContext(ctx context.Context) khulnasoft.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RubberTreeVariety) string {
		return string(e)
	}).(khulnasoft.StringOutput)
}

func (o RubberTreeVarietyOutput) ToStringPtrOutput() khulnasoft.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RubberTreeVarietyOutput) ToStringPtrOutputWithContext(ctx context.Context) khulnasoft.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RubberTreeVariety) *string {
		v := string(e)
		return &v
	}).(khulnasoft.StringPtrOutput)
}

type RubberTreeVarietyPtrOutput struct{ *khulnasoft.OutputState }

func (RubberTreeVarietyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RubberTreeVariety)(nil)).Elem()
}

func (o RubberTreeVarietyPtrOutput) ToRubberTreeVarietyPtrOutput() RubberTreeVarietyPtrOutput {
	return o
}

func (o RubberTreeVarietyPtrOutput) ToRubberTreeVarietyPtrOutputWithContext(ctx context.Context) RubberTreeVarietyPtrOutput {
	return o
}

func (o RubberTreeVarietyPtrOutput) Elem() RubberTreeVarietyOutput {
	return o.ApplyT(func(v *RubberTreeVariety) RubberTreeVariety {
		if v != nil {
			return *v
		}
		var ret RubberTreeVariety
		return ret
	}).(RubberTreeVarietyOutput)
}

func (o RubberTreeVarietyPtrOutput) ToStringPtrOutput() khulnasoft.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RubberTreeVarietyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) khulnasoft.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RubberTreeVariety) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(khulnasoft.StringPtrOutput)
}

// RubberTreeVarietyInput is an input type that accepts values of the RubberTreeVariety enum
// A concrete instance of `RubberTreeVarietyInput` can be one of the following:
//
//	RubberTreeVarietyBurgundy
//	RubberTreeVarietyRuby
//	RubberTreeVarietyTineke
type RubberTreeVarietyInput interface {
	khulnasoft.Input

	ToRubberTreeVarietyOutput() RubberTreeVarietyOutput
	ToRubberTreeVarietyOutputWithContext(context.Context) RubberTreeVarietyOutput
}

var rubberTreeVarietyPtrType = reflect.TypeOf((**RubberTreeVariety)(nil)).Elem()

type RubberTreeVarietyPtrInput interface {
	khulnasoft.Input

	ToRubberTreeVarietyPtrOutput() RubberTreeVarietyPtrOutput
	ToRubberTreeVarietyPtrOutputWithContext(context.Context) RubberTreeVarietyPtrOutput
}

type rubberTreeVarietyPtr string

func RubberTreeVarietyPtr(v string) RubberTreeVarietyPtrInput {
	return (*rubberTreeVarietyPtr)(&v)
}

func (*rubberTreeVarietyPtr) ElementType() reflect.Type {
	return rubberTreeVarietyPtrType
}

func (in *rubberTreeVarietyPtr) ToRubberTreeVarietyPtrOutput() RubberTreeVarietyPtrOutput {
	return khulnasoft.ToOutput(in).(RubberTreeVarietyPtrOutput)
}

func (in *rubberTreeVarietyPtr) ToRubberTreeVarietyPtrOutputWithContext(ctx context.Context) RubberTreeVarietyPtrOutput {
	return khulnasoft.ToOutputWithContext(ctx, in).(RubberTreeVarietyPtrOutput)
}

func init() {
	khulnasoft.RegisterInputType(reflect.TypeOf((*RubberTreeVarietyInput)(nil)).Elem(), RubberTreeVariety("Burgundy"))
	khulnasoft.RegisterInputType(reflect.TypeOf((*RubberTreeVarietyPtrInput)(nil)).Elem(), RubberTreeVariety("Burgundy"))
	khulnasoft.RegisterOutputType(OutputOnlyEnumTypeOutput{})
	khulnasoft.RegisterOutputType(OutputOnlyEnumTypePtrOutput{})
	khulnasoft.RegisterOutputType(OutputOnlyEnumTypeMapOutput{})
	khulnasoft.RegisterOutputType(RubberTreeVarietyOutput{})
	khulnasoft.RegisterOutputType(RubberTreeVarietyPtrOutput{})
}