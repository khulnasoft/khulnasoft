// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

type Diameter float64

const (
	DiameterSixinch    = Diameter(6)
	DiameterTwelveinch = Diameter(12)
)

func (Diameter) ElementType() reflect.Type {
	return reflect.TypeOf((*Diameter)(nil)).Elem()
}

func (e Diameter) ToDiameterOutput() DiameterOutput {
	return khulnasoft.ToOutput(e).(DiameterOutput)
}

func (e Diameter) ToDiameterOutputWithContext(ctx context.Context) DiameterOutput {
	return khulnasoft.ToOutputWithContext(ctx, e).(DiameterOutput)
}

func (e Diameter) ToDiameterPtrOutput() DiameterPtrOutput {
	return e.ToDiameterPtrOutputWithContext(context.Background())
}

func (e Diameter) ToDiameterPtrOutputWithContext(ctx context.Context) DiameterPtrOutput {
	return Diameter(e).ToDiameterOutputWithContext(ctx).ToDiameterPtrOutputWithContext(ctx)
}

func (e Diameter) ToFloat64Output() khulnasoft.Float64Output {
	return khulnasoft.ToOutput(khulnasoft.Float64(e)).(khulnasoft.Float64Output)
}

func (e Diameter) ToFloat64OutputWithContext(ctx context.Context) khulnasoft.Float64Output {
	return khulnasoft.ToOutputWithContext(ctx, khulnasoft.Float64(e)).(khulnasoft.Float64Output)
}

func (e Diameter) ToFloat64PtrOutput() khulnasoft.Float64PtrOutput {
	return khulnasoft.Float64(e).ToFloat64PtrOutputWithContext(context.Background())
}

func (e Diameter) ToFloat64PtrOutputWithContext(ctx context.Context) khulnasoft.Float64PtrOutput {
	return khulnasoft.Float64(e).ToFloat64OutputWithContext(ctx).ToFloat64PtrOutputWithContext(ctx)
}

type DiameterOutput struct{ *khulnasoft.OutputState }

func (DiameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Diameter)(nil)).Elem()
}

func (o DiameterOutput) ToDiameterOutput() DiameterOutput {
	return o
}

func (o DiameterOutput) ToDiameterOutputWithContext(ctx context.Context) DiameterOutput {
	return o
}

func (o DiameterOutput) ToDiameterPtrOutput() DiameterPtrOutput {
	return o.ToDiameterPtrOutputWithContext(context.Background())
}

func (o DiameterOutput) ToDiameterPtrOutputWithContext(ctx context.Context) DiameterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Diameter) *Diameter {
		return &v
	}).(DiameterPtrOutput)
}

func (o DiameterOutput) ToFloat64Output() khulnasoft.Float64Output {
	return o.ToFloat64OutputWithContext(context.Background())
}

func (o DiameterOutput) ToFloat64OutputWithContext(ctx context.Context) khulnasoft.Float64Output {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Diameter) float64 {
		return float64(e)
	}).(khulnasoft.Float64Output)
}

func (o DiameterOutput) ToFloat64PtrOutput() khulnasoft.Float64PtrOutput {
	return o.ToFloat64PtrOutputWithContext(context.Background())
}

func (o DiameterOutput) ToFloat64PtrOutputWithContext(ctx context.Context) khulnasoft.Float64PtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Diameter) *float64 {
		v := float64(e)
		return &v
	}).(khulnasoft.Float64PtrOutput)
}

type DiameterPtrOutput struct{ *khulnasoft.OutputState }

func (DiameterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Diameter)(nil)).Elem()
}

func (o DiameterPtrOutput) ToDiameterPtrOutput() DiameterPtrOutput {
	return o
}

func (o DiameterPtrOutput) ToDiameterPtrOutputWithContext(ctx context.Context) DiameterPtrOutput {
	return o
}

func (o DiameterPtrOutput) Elem() DiameterOutput {
	return o.ApplyT(func(v *Diameter) Diameter {
		if v != nil {
			return *v
		}
		var ret Diameter
		return ret
	}).(DiameterOutput)
}

func (o DiameterPtrOutput) ToFloat64PtrOutput() khulnasoft.Float64PtrOutput {
	return o.ToFloat64PtrOutputWithContext(context.Background())
}

func (o DiameterPtrOutput) ToFloat64PtrOutputWithContext(ctx context.Context) khulnasoft.Float64PtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Diameter) *float64 {
		if e == nil {
			return nil
		}
		v := float64(*e)
		return &v
	}).(khulnasoft.Float64PtrOutput)
}

// DiameterInput is an input type that accepts values of the Diameter enum
// A concrete instance of `DiameterInput` can be one of the following:
//
//	DiameterSixinch
//	DiameterTwelveinch
type DiameterInput interface {
	khulnasoft.Input

	ToDiameterOutput() DiameterOutput
	ToDiameterOutputWithContext(context.Context) DiameterOutput
}

var diameterPtrType = reflect.TypeOf((**Diameter)(nil)).Elem()

type DiameterPtrInput interface {
	khulnasoft.Input

	ToDiameterPtrOutput() DiameterPtrOutput
	ToDiameterPtrOutputWithContext(context.Context) DiameterPtrOutput
}

type diameterPtr float64

func DiameterPtr(v float64) DiameterPtrInput {
	return (*diameterPtr)(&v)
}

func (*diameterPtr) ElementType() reflect.Type {
	return diameterPtrType
}

func (in *diameterPtr) ToDiameterPtrOutput() DiameterPtrOutput {
	return khulnasoft.ToOutput(in).(DiameterPtrOutput)
}

func (in *diameterPtr) ToDiameterPtrOutputWithContext(ctx context.Context) DiameterPtrOutput {
	return khulnasoft.ToOutputWithContext(ctx, in).(DiameterPtrOutput)
}

type Farm string

const (
	Farm_Pulumi_Planters_Inc_ = Farm("Pulumi Planters Inc.")
	Farm_Plants_R_Us          = Farm("Plants'R'Us")
)

func (Farm) ElementType() reflect.Type {
	return reflect.TypeOf((*Farm)(nil)).Elem()
}

func (e Farm) ToFarmOutput() FarmOutput {
	return khulnasoft.ToOutput(e).(FarmOutput)
}

func (e Farm) ToFarmOutputWithContext(ctx context.Context) FarmOutput {
	return khulnasoft.ToOutputWithContext(ctx, e).(FarmOutput)
}

func (e Farm) ToFarmPtrOutput() FarmPtrOutput {
	return e.ToFarmPtrOutputWithContext(context.Background())
}

func (e Farm) ToFarmPtrOutputWithContext(ctx context.Context) FarmPtrOutput {
	return Farm(e).ToFarmOutputWithContext(ctx).ToFarmPtrOutputWithContext(ctx)
}

func (e Farm) ToStringOutput() khulnasoft.StringOutput {
	return khulnasoft.ToOutput(khulnasoft.String(e)).(khulnasoft.StringOutput)
}

func (e Farm) ToStringOutputWithContext(ctx context.Context) khulnasoft.StringOutput {
	return khulnasoft.ToOutputWithContext(ctx, khulnasoft.String(e)).(khulnasoft.StringOutput)
}

func (e Farm) ToStringPtrOutput() khulnasoft.StringPtrOutput {
	return khulnasoft.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Farm) ToStringPtrOutputWithContext(ctx context.Context) khulnasoft.StringPtrOutput {
	return khulnasoft.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type FarmOutput struct{ *khulnasoft.OutputState }

func (FarmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Farm)(nil)).Elem()
}

func (o FarmOutput) ToFarmOutput() FarmOutput {
	return o
}

func (o FarmOutput) ToFarmOutputWithContext(ctx context.Context) FarmOutput {
	return o
}

func (o FarmOutput) ToFarmPtrOutput() FarmPtrOutput {
	return o.ToFarmPtrOutputWithContext(context.Background())
}

func (o FarmOutput) ToFarmPtrOutputWithContext(ctx context.Context) FarmPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Farm) *Farm {
		return &v
	}).(FarmPtrOutput)
}

func (o FarmOutput) ToStringOutput() khulnasoft.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FarmOutput) ToStringOutputWithContext(ctx context.Context) khulnasoft.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Farm) string {
		return string(e)
	}).(khulnasoft.StringOutput)
}

func (o FarmOutput) ToStringPtrOutput() khulnasoft.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FarmOutput) ToStringPtrOutputWithContext(ctx context.Context) khulnasoft.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Farm) *string {
		v := string(e)
		return &v
	}).(khulnasoft.StringPtrOutput)
}

type FarmPtrOutput struct{ *khulnasoft.OutputState }

func (FarmPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Farm)(nil)).Elem()
}

func (o FarmPtrOutput) ToFarmPtrOutput() FarmPtrOutput {
	return o
}

func (o FarmPtrOutput) ToFarmPtrOutputWithContext(ctx context.Context) FarmPtrOutput {
	return o
}

func (o FarmPtrOutput) Elem() FarmOutput {
	return o.ApplyT(func(v *Farm) Farm {
		if v != nil {
			return *v
		}
		var ret Farm
		return ret
	}).(FarmOutput)
}

func (o FarmPtrOutput) ToStringPtrOutput() khulnasoft.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FarmPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) khulnasoft.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Farm) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(khulnasoft.StringPtrOutput)
}

// FarmInput is an input type that accepts values of the Farm enum
// A concrete instance of `FarmInput` can be one of the following:
//
//	Farm_Pulumi_Planters_Inc_
//	Farm_Plants_R_Us
type FarmInput interface {
	khulnasoft.Input

	ToFarmOutput() FarmOutput
	ToFarmOutputWithContext(context.Context) FarmOutput
}

var farmPtrType = reflect.TypeOf((**Farm)(nil)).Elem()

type FarmPtrInput interface {
	khulnasoft.Input

	ToFarmPtrOutput() FarmPtrOutput
	ToFarmPtrOutputWithContext(context.Context) FarmPtrOutput
}

type farmPtr string

func FarmPtr(v string) FarmPtrInput {
	return (*farmPtr)(&v)
}

func (*farmPtr) ElementType() reflect.Type {
	return farmPtrType
}

func (in *farmPtr) ToFarmPtrOutput() FarmPtrOutput {
	return khulnasoft.ToOutput(in).(FarmPtrOutput)
}

func (in *farmPtr) ToFarmPtrOutputWithContext(ctx context.Context) FarmPtrOutput {
	return khulnasoft.ToOutputWithContext(ctx, in).(FarmPtrOutput)
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

// RubberTreeVarietyArrayInput is an input type that accepts RubberTreeVarietyArray and RubberTreeVarietyArrayOutput values.
// You can construct a concrete instance of `RubberTreeVarietyArrayInput` via:
//
//	RubberTreeVarietyArray{ RubberTreeVarietyArgs{...} }
type RubberTreeVarietyArrayInput interface {
	khulnasoft.Input

	ToRubberTreeVarietyArrayOutput() RubberTreeVarietyArrayOutput
	ToRubberTreeVarietyArrayOutputWithContext(context.Context) RubberTreeVarietyArrayOutput
}

type RubberTreeVarietyArray []RubberTreeVariety

func (RubberTreeVarietyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RubberTreeVariety)(nil)).Elem()
}

func (i RubberTreeVarietyArray) ToRubberTreeVarietyArrayOutput() RubberTreeVarietyArrayOutput {
	return i.ToRubberTreeVarietyArrayOutputWithContext(context.Background())
}

func (i RubberTreeVarietyArray) ToRubberTreeVarietyArrayOutputWithContext(ctx context.Context) RubberTreeVarietyArrayOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(RubberTreeVarietyArrayOutput)
}

type RubberTreeVarietyArrayOutput struct{ *khulnasoft.OutputState }

func (RubberTreeVarietyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RubberTreeVariety)(nil)).Elem()
}

func (o RubberTreeVarietyArrayOutput) ToRubberTreeVarietyArrayOutput() RubberTreeVarietyArrayOutput {
	return o
}

func (o RubberTreeVarietyArrayOutput) ToRubberTreeVarietyArrayOutputWithContext(ctx context.Context) RubberTreeVarietyArrayOutput {
	return o
}

func (o RubberTreeVarietyArrayOutput) Index(i khulnasoft.IntInput) RubberTreeVarietyOutput {
	return khulnasoft.All(o, i).ApplyT(func(vs []interface{}) RubberTreeVariety {
		return vs[0].([]RubberTreeVariety)[vs[1].(int)]
	}).(RubberTreeVarietyOutput)
}

type TreeSize string

const (
	TreeSizeSmall  = TreeSize("small")
	TreeSizeMedium = TreeSize("medium")
	TreeSizeLarge  = TreeSize("large")
)

func (TreeSize) ElementType() reflect.Type {
	return reflect.TypeOf((*TreeSize)(nil)).Elem()
}

func (e TreeSize) ToTreeSizeOutput() TreeSizeOutput {
	return khulnasoft.ToOutput(e).(TreeSizeOutput)
}

func (e TreeSize) ToTreeSizeOutputWithContext(ctx context.Context) TreeSizeOutput {
	return khulnasoft.ToOutputWithContext(ctx, e).(TreeSizeOutput)
}

func (e TreeSize) ToTreeSizePtrOutput() TreeSizePtrOutput {
	return e.ToTreeSizePtrOutputWithContext(context.Background())
}

func (e TreeSize) ToTreeSizePtrOutputWithContext(ctx context.Context) TreeSizePtrOutput {
	return TreeSize(e).ToTreeSizeOutputWithContext(ctx).ToTreeSizePtrOutputWithContext(ctx)
}

func (e TreeSize) ToStringOutput() khulnasoft.StringOutput {
	return khulnasoft.ToOutput(khulnasoft.String(e)).(khulnasoft.StringOutput)
}

func (e TreeSize) ToStringOutputWithContext(ctx context.Context) khulnasoft.StringOutput {
	return khulnasoft.ToOutputWithContext(ctx, khulnasoft.String(e)).(khulnasoft.StringOutput)
}

func (e TreeSize) ToStringPtrOutput() khulnasoft.StringPtrOutput {
	return khulnasoft.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TreeSize) ToStringPtrOutputWithContext(ctx context.Context) khulnasoft.StringPtrOutput {
	return khulnasoft.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TreeSizeOutput struct{ *khulnasoft.OutputState }

func (TreeSizeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TreeSize)(nil)).Elem()
}

func (o TreeSizeOutput) ToTreeSizeOutput() TreeSizeOutput {
	return o
}

func (o TreeSizeOutput) ToTreeSizeOutputWithContext(ctx context.Context) TreeSizeOutput {
	return o
}

func (o TreeSizeOutput) ToTreeSizePtrOutput() TreeSizePtrOutput {
	return o.ToTreeSizePtrOutputWithContext(context.Background())
}

func (o TreeSizeOutput) ToTreeSizePtrOutputWithContext(ctx context.Context) TreeSizePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TreeSize) *TreeSize {
		return &v
	}).(TreeSizePtrOutput)
}

func (o TreeSizeOutput) ToStringOutput() khulnasoft.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TreeSizeOutput) ToStringOutputWithContext(ctx context.Context) khulnasoft.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TreeSize) string {
		return string(e)
	}).(khulnasoft.StringOutput)
}

func (o TreeSizeOutput) ToStringPtrOutput() khulnasoft.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TreeSizeOutput) ToStringPtrOutputWithContext(ctx context.Context) khulnasoft.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TreeSize) *string {
		v := string(e)
		return &v
	}).(khulnasoft.StringPtrOutput)
}

type TreeSizePtrOutput struct{ *khulnasoft.OutputState }

func (TreeSizePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TreeSize)(nil)).Elem()
}

func (o TreeSizePtrOutput) ToTreeSizePtrOutput() TreeSizePtrOutput {
	return o
}

func (o TreeSizePtrOutput) ToTreeSizePtrOutputWithContext(ctx context.Context) TreeSizePtrOutput {
	return o
}

func (o TreeSizePtrOutput) Elem() TreeSizeOutput {
	return o.ApplyT(func(v *TreeSize) TreeSize {
		if v != nil {
			return *v
		}
		var ret TreeSize
		return ret
	}).(TreeSizeOutput)
}

func (o TreeSizePtrOutput) ToStringPtrOutput() khulnasoft.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TreeSizePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) khulnasoft.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TreeSize) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(khulnasoft.StringPtrOutput)
}

// TreeSizeInput is an input type that accepts values of the TreeSize enum
// A concrete instance of `TreeSizeInput` can be one of the following:
//
//	TreeSizeSmall
//	TreeSizeMedium
//	TreeSizeLarge
type TreeSizeInput interface {
	khulnasoft.Input

	ToTreeSizeOutput() TreeSizeOutput
	ToTreeSizeOutputWithContext(context.Context) TreeSizeOutput
}

var treeSizePtrType = reflect.TypeOf((**TreeSize)(nil)).Elem()

type TreeSizePtrInput interface {
	khulnasoft.Input

	ToTreeSizePtrOutput() TreeSizePtrOutput
	ToTreeSizePtrOutputWithContext(context.Context) TreeSizePtrOutput
}

type treeSizePtr string

func TreeSizePtr(v string) TreeSizePtrInput {
	return (*treeSizePtr)(&v)
}

func (*treeSizePtr) ElementType() reflect.Type {
	return treeSizePtrType
}

func (in *treeSizePtr) ToTreeSizePtrOutput() TreeSizePtrOutput {
	return khulnasoft.ToOutput(in).(TreeSizePtrOutput)
}

func (in *treeSizePtr) ToTreeSizePtrOutputWithContext(ctx context.Context) TreeSizePtrOutput {
	return khulnasoft.ToOutputWithContext(ctx, in).(TreeSizePtrOutput)
}

// TreeSizeMapInput is an input type that accepts TreeSizeMap and TreeSizeMapOutput values.
// You can construct a concrete instance of `TreeSizeMapInput` via:
//
//	TreeSizeMap{ "key": TreeSizeArgs{...} }
type TreeSizeMapInput interface {
	khulnasoft.Input

	ToTreeSizeMapOutput() TreeSizeMapOutput
	ToTreeSizeMapOutputWithContext(context.Context) TreeSizeMapOutput
}

type TreeSizeMap map[string]TreeSize

func (TreeSizeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TreeSize)(nil)).Elem()
}

func (i TreeSizeMap) ToTreeSizeMapOutput() TreeSizeMapOutput {
	return i.ToTreeSizeMapOutputWithContext(context.Background())
}

func (i TreeSizeMap) ToTreeSizeMapOutputWithContext(ctx context.Context) TreeSizeMapOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(TreeSizeMapOutput)
}

type TreeSizeMapOutput struct{ *khulnasoft.OutputState }

func (TreeSizeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TreeSize)(nil)).Elem()
}

func (o TreeSizeMapOutput) ToTreeSizeMapOutput() TreeSizeMapOutput {
	return o
}

func (o TreeSizeMapOutput) ToTreeSizeMapOutputWithContext(ctx context.Context) TreeSizeMapOutput {
	return o
}

func (o TreeSizeMapOutput) MapIndex(k khulnasoft.StringInput) TreeSizeOutput {
	return khulnasoft.All(o, k).ApplyT(func(vs []interface{}) TreeSize {
		return vs[0].(map[string]TreeSize)[vs[1].(string)]
	}).(TreeSizeOutput)
}

func init() {
	khulnasoft.RegisterInputType(reflect.TypeOf((*DiameterInput)(nil)).Elem(), Diameter(6))
	khulnasoft.RegisterInputType(reflect.TypeOf((*DiameterPtrInput)(nil)).Elem(), Diameter(6))
	khulnasoft.RegisterInputType(reflect.TypeOf((*FarmInput)(nil)).Elem(), Farm("Pulumi Planters Inc."))
	khulnasoft.RegisterInputType(reflect.TypeOf((*FarmPtrInput)(nil)).Elem(), Farm("Pulumi Planters Inc."))
	khulnasoft.RegisterInputType(reflect.TypeOf((*RubberTreeVarietyInput)(nil)).Elem(), RubberTreeVariety("Burgundy"))
	khulnasoft.RegisterInputType(reflect.TypeOf((*RubberTreeVarietyPtrInput)(nil)).Elem(), RubberTreeVariety("Burgundy"))
	khulnasoft.RegisterInputType(reflect.TypeOf((*RubberTreeVarietyArrayInput)(nil)).Elem(), RubberTreeVarietyArray{})
	khulnasoft.RegisterInputType(reflect.TypeOf((*TreeSizeInput)(nil)).Elem(), TreeSize("small"))
	khulnasoft.RegisterInputType(reflect.TypeOf((*TreeSizePtrInput)(nil)).Elem(), TreeSize("small"))
	khulnasoft.RegisterInputType(reflect.TypeOf((*TreeSizeMapInput)(nil)).Elem(), TreeSizeMap{})
	khulnasoft.RegisterOutputType(DiameterOutput{})
	khulnasoft.RegisterOutputType(DiameterPtrOutput{})
	khulnasoft.RegisterOutputType(FarmOutput{})
	khulnasoft.RegisterOutputType(FarmPtrOutput{})
	khulnasoft.RegisterOutputType(RubberTreeVarietyOutput{})
	khulnasoft.RegisterOutputType(RubberTreeVarietyPtrOutput{})
	khulnasoft.RegisterOutputType(RubberTreeVarietyArrayOutput{})
	khulnasoft.RegisterOutputType(TreeSizeOutput{})
	khulnasoft.RegisterOutputType(TreeSizePtrOutput{})
	khulnasoft.RegisterOutputType(TreeSizeMapOutput{})
}