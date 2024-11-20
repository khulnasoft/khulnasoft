// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"context"
	"reflect"

	"cyclic-types/example/internal"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

var _ = internal.GetEnvOrDefault

type AcyclicReferent struct {
	Bar  IndirectCycleS `khulnasoft:"bar"`
	Baz  IndirectCycleT `khulnasoft:"baz"`
	Foo4 DirectCycle    `khulnasoft:"foo4"`
}

// AcyclicReferentInput is an input type that accepts AcyclicReferentArgs and AcyclicReferentOutput values.
// You can construct a concrete instance of `AcyclicReferentInput` via:
//
//	AcyclicReferentArgs{...}
type AcyclicReferentInput interface {
	khulnasoft.Input

	ToAcyclicReferentOutput() AcyclicReferentOutput
	ToAcyclicReferentOutputWithContext(context.Context) AcyclicReferentOutput
}

type AcyclicReferentArgs struct {
	Bar  IndirectCycleSInput `khulnasoft:"bar"`
	Baz  IndirectCycleTInput `khulnasoft:"baz"`
	Foo4 DirectCycleInput    `khulnasoft:"foo4"`
}

func (AcyclicReferentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AcyclicReferent)(nil)).Elem()
}

func (i AcyclicReferentArgs) ToAcyclicReferentOutput() AcyclicReferentOutput {
	return i.ToAcyclicReferentOutputWithContext(context.Background())
}

func (i AcyclicReferentArgs) ToAcyclicReferentOutputWithContext(ctx context.Context) AcyclicReferentOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(AcyclicReferentOutput)
}

type AcyclicReferentOutput struct{ *khulnasoft.OutputState }

func (AcyclicReferentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AcyclicReferent)(nil)).Elem()
}

func (o AcyclicReferentOutput) ToAcyclicReferentOutput() AcyclicReferentOutput {
	return o
}

func (o AcyclicReferentOutput) ToAcyclicReferentOutputWithContext(ctx context.Context) AcyclicReferentOutput {
	return o
}

func (o AcyclicReferentOutput) Bar() IndirectCycleSOutput {
	return o.ApplyT(func(v AcyclicReferent) IndirectCycleS { return v.Bar }).(IndirectCycleSOutput)
}

func (o AcyclicReferentOutput) Baz() IndirectCycleTOutput {
	return o.ApplyT(func(v AcyclicReferent) IndirectCycleT { return v.Baz }).(IndirectCycleTOutput)
}

func (o AcyclicReferentOutput) Foo4() DirectCycleOutput {
	return o.ApplyT(func(v AcyclicReferent) DirectCycle { return v.Foo4 }).(DirectCycleOutput)
}

type AcyclicS struct {
	Foo5 string `khulnasoft:"foo5"`
}

// AcyclicSInput is an input type that accepts AcyclicSArgs and AcyclicSOutput values.
// You can construct a concrete instance of `AcyclicSInput` via:
//
//	AcyclicSArgs{...}
type AcyclicSInput interface {
	khulnasoft.Input

	ToAcyclicSOutput() AcyclicSOutput
	ToAcyclicSOutputWithContext(context.Context) AcyclicSOutput
}

type AcyclicSArgs struct {
	Foo5 khulnasoft.StringInput `khulnasoft:"foo5"`
}

func (AcyclicSArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AcyclicS)(nil)).Elem()
}

func (i AcyclicSArgs) ToAcyclicSOutput() AcyclicSOutput {
	return i.ToAcyclicSOutputWithContext(context.Background())
}

func (i AcyclicSArgs) ToAcyclicSOutputWithContext(ctx context.Context) AcyclicSOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(AcyclicSOutput)
}

type AcyclicSOutput struct{ *khulnasoft.OutputState }

func (AcyclicSOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AcyclicS)(nil)).Elem()
}

func (o AcyclicSOutput) ToAcyclicSOutput() AcyclicSOutput {
	return o
}

func (o AcyclicSOutput) ToAcyclicSOutputWithContext(ctx context.Context) AcyclicSOutput {
	return o
}

func (o AcyclicSOutput) Foo5() khulnasoft.StringOutput {
	return o.ApplyT(func(v AcyclicS) string { return v.Foo5 }).(khulnasoft.StringOutput)
}

type AcyclicT struct {
	Foo6 AcyclicS `khulnasoft:"foo6"`
}

// AcyclicTInput is an input type that accepts AcyclicTArgs and AcyclicTOutput values.
// You can construct a concrete instance of `AcyclicTInput` via:
//
//	AcyclicTArgs{...}
type AcyclicTInput interface {
	khulnasoft.Input

	ToAcyclicTOutput() AcyclicTOutput
	ToAcyclicTOutputWithContext(context.Context) AcyclicTOutput
}

type AcyclicTArgs struct {
	Foo6 AcyclicSInput `khulnasoft:"foo6"`
}

func (AcyclicTArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AcyclicT)(nil)).Elem()
}

func (i AcyclicTArgs) ToAcyclicTOutput() AcyclicTOutput {
	return i.ToAcyclicTOutputWithContext(context.Background())
}

func (i AcyclicTArgs) ToAcyclicTOutputWithContext(ctx context.Context) AcyclicTOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(AcyclicTOutput)
}

type AcyclicTOutput struct{ *khulnasoft.OutputState }

func (AcyclicTOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AcyclicT)(nil)).Elem()
}

func (o AcyclicTOutput) ToAcyclicTOutput() AcyclicTOutput {
	return o
}

func (o AcyclicTOutput) ToAcyclicTOutputWithContext(ctx context.Context) AcyclicTOutput {
	return o
}

func (o AcyclicTOutput) Foo6() AcyclicSOutput {
	return o.ApplyT(func(v AcyclicT) AcyclicS { return v.Foo6 }).(AcyclicSOutput)
}

type DirectCycle struct {
	Foo *DirectCycle `khulnasoft:"foo"`
}

// DirectCycleInput is an input type that accepts DirectCycleArgs and DirectCycleOutput values.
// You can construct a concrete instance of `DirectCycleInput` via:
//
//	DirectCycleArgs{...}
type DirectCycleInput interface {
	khulnasoft.Input

	ToDirectCycleOutput() DirectCycleOutput
	ToDirectCycleOutputWithContext(context.Context) DirectCycleOutput
}

type DirectCycleArgs struct {
	Foo DirectCyclePtrInput `khulnasoft:"foo"`
}

func (DirectCycleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DirectCycle)(nil)).Elem()
}

func (i DirectCycleArgs) ToDirectCycleOutput() DirectCycleOutput {
	return i.ToDirectCycleOutputWithContext(context.Background())
}

func (i DirectCycleArgs) ToDirectCycleOutputWithContext(ctx context.Context) DirectCycleOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(DirectCycleOutput)
}

func (i DirectCycleArgs) ToDirectCyclePtrOutput() DirectCyclePtrOutput {
	return i.ToDirectCyclePtrOutputWithContext(context.Background())
}

func (i DirectCycleArgs) ToDirectCyclePtrOutputWithContext(ctx context.Context) DirectCyclePtrOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(DirectCycleOutput).ToDirectCyclePtrOutputWithContext(ctx)
}

// DirectCyclePtrInput is an input type that accepts DirectCycleArgs, DirectCyclePtr and DirectCyclePtrOutput values.
// You can construct a concrete instance of `DirectCyclePtrInput` via:
//
//	        DirectCycleArgs{...}
//
//	or:
//
//	        nil
type DirectCyclePtrInput interface {
	khulnasoft.Input

	ToDirectCyclePtrOutput() DirectCyclePtrOutput
	ToDirectCyclePtrOutputWithContext(context.Context) DirectCyclePtrOutput
}

type directCyclePtrType DirectCycleArgs

func DirectCyclePtr(v *DirectCycleArgs) DirectCyclePtrInput {
	return (*directCyclePtrType)(v)
}

func (*directCyclePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DirectCycle)(nil)).Elem()
}

func (i *directCyclePtrType) ToDirectCyclePtrOutput() DirectCyclePtrOutput {
	return i.ToDirectCyclePtrOutputWithContext(context.Background())
}

func (i *directCyclePtrType) ToDirectCyclePtrOutputWithContext(ctx context.Context) DirectCyclePtrOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(DirectCyclePtrOutput)
}

type DirectCycleOutput struct{ *khulnasoft.OutputState }

func (DirectCycleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DirectCycle)(nil)).Elem()
}

func (o DirectCycleOutput) ToDirectCycleOutput() DirectCycleOutput {
	return o
}

func (o DirectCycleOutput) ToDirectCycleOutputWithContext(ctx context.Context) DirectCycleOutput {
	return o
}

func (o DirectCycleOutput) ToDirectCyclePtrOutput() DirectCyclePtrOutput {
	return o.ToDirectCyclePtrOutputWithContext(context.Background())
}

func (o DirectCycleOutput) ToDirectCyclePtrOutputWithContext(ctx context.Context) DirectCyclePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DirectCycle) *DirectCycle {
		return &v
	}).(DirectCyclePtrOutput)
}

func (o DirectCycleOutput) Foo() DirectCyclePtrOutput {
	return o.ApplyT(func(v DirectCycle) *DirectCycle { return v.Foo }).(DirectCyclePtrOutput)
}

type DirectCyclePtrOutput struct{ *khulnasoft.OutputState }

func (DirectCyclePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DirectCycle)(nil)).Elem()
}

func (o DirectCyclePtrOutput) ToDirectCyclePtrOutput() DirectCyclePtrOutput {
	return o
}

func (o DirectCyclePtrOutput) ToDirectCyclePtrOutputWithContext(ctx context.Context) DirectCyclePtrOutput {
	return o
}

func (o DirectCyclePtrOutput) Elem() DirectCycleOutput {
	return o.ApplyT(func(v *DirectCycle) DirectCycle {
		if v != nil {
			return *v
		}
		var ret DirectCycle
		return ret
	}).(DirectCycleOutput)
}

func (o DirectCyclePtrOutput) Foo() DirectCyclePtrOutput {
	return o.ApplyT(func(v *DirectCycle) *DirectCycle {
		if v == nil {
			return nil
		}
		return v.Foo
	}).(DirectCyclePtrOutput)
}

type IndirectCycleS struct {
	Foo2 *IndirectCycleT `khulnasoft:"foo2"`
}

// IndirectCycleSInput is an input type that accepts IndirectCycleSArgs and IndirectCycleSOutput values.
// You can construct a concrete instance of `IndirectCycleSInput` via:
//
//	IndirectCycleSArgs{...}
type IndirectCycleSInput interface {
	khulnasoft.Input

	ToIndirectCycleSOutput() IndirectCycleSOutput
	ToIndirectCycleSOutputWithContext(context.Context) IndirectCycleSOutput
}

type IndirectCycleSArgs struct {
	Foo2 IndirectCycleTPtrInput `khulnasoft:"foo2"`
}

func (IndirectCycleSArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IndirectCycleS)(nil)).Elem()
}

func (i IndirectCycleSArgs) ToIndirectCycleSOutput() IndirectCycleSOutput {
	return i.ToIndirectCycleSOutputWithContext(context.Background())
}

func (i IndirectCycleSArgs) ToIndirectCycleSOutputWithContext(ctx context.Context) IndirectCycleSOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(IndirectCycleSOutput)
}

func (i IndirectCycleSArgs) ToIndirectCycleSPtrOutput() IndirectCycleSPtrOutput {
	return i.ToIndirectCycleSPtrOutputWithContext(context.Background())
}

func (i IndirectCycleSArgs) ToIndirectCycleSPtrOutputWithContext(ctx context.Context) IndirectCycleSPtrOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(IndirectCycleSOutput).ToIndirectCycleSPtrOutputWithContext(ctx)
}

// IndirectCycleSPtrInput is an input type that accepts IndirectCycleSArgs, IndirectCycleSPtr and IndirectCycleSPtrOutput values.
// You can construct a concrete instance of `IndirectCycleSPtrInput` via:
//
//	        IndirectCycleSArgs{...}
//
//	or:
//
//	        nil
type IndirectCycleSPtrInput interface {
	khulnasoft.Input

	ToIndirectCycleSPtrOutput() IndirectCycleSPtrOutput
	ToIndirectCycleSPtrOutputWithContext(context.Context) IndirectCycleSPtrOutput
}

type indirectCycleSPtrType IndirectCycleSArgs

func IndirectCycleSPtr(v *IndirectCycleSArgs) IndirectCycleSPtrInput {
	return (*indirectCycleSPtrType)(v)
}

func (*indirectCycleSPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IndirectCycleS)(nil)).Elem()
}

func (i *indirectCycleSPtrType) ToIndirectCycleSPtrOutput() IndirectCycleSPtrOutput {
	return i.ToIndirectCycleSPtrOutputWithContext(context.Background())
}

func (i *indirectCycleSPtrType) ToIndirectCycleSPtrOutputWithContext(ctx context.Context) IndirectCycleSPtrOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(IndirectCycleSPtrOutput)
}

type IndirectCycleSOutput struct{ *khulnasoft.OutputState }

func (IndirectCycleSOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IndirectCycleS)(nil)).Elem()
}

func (o IndirectCycleSOutput) ToIndirectCycleSOutput() IndirectCycleSOutput {
	return o
}

func (o IndirectCycleSOutput) ToIndirectCycleSOutputWithContext(ctx context.Context) IndirectCycleSOutput {
	return o
}

func (o IndirectCycleSOutput) ToIndirectCycleSPtrOutput() IndirectCycleSPtrOutput {
	return o.ToIndirectCycleSPtrOutputWithContext(context.Background())
}

func (o IndirectCycleSOutput) ToIndirectCycleSPtrOutputWithContext(ctx context.Context) IndirectCycleSPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IndirectCycleS) *IndirectCycleS {
		return &v
	}).(IndirectCycleSPtrOutput)
}

func (o IndirectCycleSOutput) Foo2() IndirectCycleTPtrOutput {
	return o.ApplyT(func(v IndirectCycleS) *IndirectCycleT { return v.Foo2 }).(IndirectCycleTPtrOutput)
}

type IndirectCycleSPtrOutput struct{ *khulnasoft.OutputState }

func (IndirectCycleSPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IndirectCycleS)(nil)).Elem()
}

func (o IndirectCycleSPtrOutput) ToIndirectCycleSPtrOutput() IndirectCycleSPtrOutput {
	return o
}

func (o IndirectCycleSPtrOutput) ToIndirectCycleSPtrOutputWithContext(ctx context.Context) IndirectCycleSPtrOutput {
	return o
}

func (o IndirectCycleSPtrOutput) Elem() IndirectCycleSOutput {
	return o.ApplyT(func(v *IndirectCycleS) IndirectCycleS {
		if v != nil {
			return *v
		}
		var ret IndirectCycleS
		return ret
	}).(IndirectCycleSOutput)
}

func (o IndirectCycleSPtrOutput) Foo2() IndirectCycleTPtrOutput {
	return o.ApplyT(func(v *IndirectCycleS) *IndirectCycleT {
		if v == nil {
			return nil
		}
		return v.Foo2
	}).(IndirectCycleTPtrOutput)
}

type IndirectCycleT struct {
	Foo3 *IndirectCycleS `khulnasoft:"foo3"`
}

// IndirectCycleTInput is an input type that accepts IndirectCycleTArgs and IndirectCycleTOutput values.
// You can construct a concrete instance of `IndirectCycleTInput` via:
//
//	IndirectCycleTArgs{...}
type IndirectCycleTInput interface {
	khulnasoft.Input

	ToIndirectCycleTOutput() IndirectCycleTOutput
	ToIndirectCycleTOutputWithContext(context.Context) IndirectCycleTOutput
}

type IndirectCycleTArgs struct {
	Foo3 IndirectCycleSPtrInput `khulnasoft:"foo3"`
}

func (IndirectCycleTArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IndirectCycleT)(nil)).Elem()
}

func (i IndirectCycleTArgs) ToIndirectCycleTOutput() IndirectCycleTOutput {
	return i.ToIndirectCycleTOutputWithContext(context.Background())
}

func (i IndirectCycleTArgs) ToIndirectCycleTOutputWithContext(ctx context.Context) IndirectCycleTOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(IndirectCycleTOutput)
}

func (i IndirectCycleTArgs) ToIndirectCycleTPtrOutput() IndirectCycleTPtrOutput {
	return i.ToIndirectCycleTPtrOutputWithContext(context.Background())
}

func (i IndirectCycleTArgs) ToIndirectCycleTPtrOutputWithContext(ctx context.Context) IndirectCycleTPtrOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(IndirectCycleTOutput).ToIndirectCycleTPtrOutputWithContext(ctx)
}

// IndirectCycleTPtrInput is an input type that accepts IndirectCycleTArgs, IndirectCycleTPtr and IndirectCycleTPtrOutput values.
// You can construct a concrete instance of `IndirectCycleTPtrInput` via:
//
//	        IndirectCycleTArgs{...}
//
//	or:
//
//	        nil
type IndirectCycleTPtrInput interface {
	khulnasoft.Input

	ToIndirectCycleTPtrOutput() IndirectCycleTPtrOutput
	ToIndirectCycleTPtrOutputWithContext(context.Context) IndirectCycleTPtrOutput
}

type indirectCycleTPtrType IndirectCycleTArgs

func IndirectCycleTPtr(v *IndirectCycleTArgs) IndirectCycleTPtrInput {
	return (*indirectCycleTPtrType)(v)
}

func (*indirectCycleTPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IndirectCycleT)(nil)).Elem()
}

func (i *indirectCycleTPtrType) ToIndirectCycleTPtrOutput() IndirectCycleTPtrOutput {
	return i.ToIndirectCycleTPtrOutputWithContext(context.Background())
}

func (i *indirectCycleTPtrType) ToIndirectCycleTPtrOutputWithContext(ctx context.Context) IndirectCycleTPtrOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(IndirectCycleTPtrOutput)
}

type IndirectCycleTOutput struct{ *khulnasoft.OutputState }

func (IndirectCycleTOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IndirectCycleT)(nil)).Elem()
}

func (o IndirectCycleTOutput) ToIndirectCycleTOutput() IndirectCycleTOutput {
	return o
}

func (o IndirectCycleTOutput) ToIndirectCycleTOutputWithContext(ctx context.Context) IndirectCycleTOutput {
	return o
}

func (o IndirectCycleTOutput) ToIndirectCycleTPtrOutput() IndirectCycleTPtrOutput {
	return o.ToIndirectCycleTPtrOutputWithContext(context.Background())
}

func (o IndirectCycleTOutput) ToIndirectCycleTPtrOutputWithContext(ctx context.Context) IndirectCycleTPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IndirectCycleT) *IndirectCycleT {
		return &v
	}).(IndirectCycleTPtrOutput)
}

func (o IndirectCycleTOutput) Foo3() IndirectCycleSPtrOutput {
	return o.ApplyT(func(v IndirectCycleT) *IndirectCycleS { return v.Foo3 }).(IndirectCycleSPtrOutput)
}

type IndirectCycleTPtrOutput struct{ *khulnasoft.OutputState }

func (IndirectCycleTPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IndirectCycleT)(nil)).Elem()
}

func (o IndirectCycleTPtrOutput) ToIndirectCycleTPtrOutput() IndirectCycleTPtrOutput {
	return o
}

func (o IndirectCycleTPtrOutput) ToIndirectCycleTPtrOutputWithContext(ctx context.Context) IndirectCycleTPtrOutput {
	return o
}

func (o IndirectCycleTPtrOutput) Elem() IndirectCycleTOutput {
	return o.ApplyT(func(v *IndirectCycleT) IndirectCycleT {
		if v != nil {
			return *v
		}
		var ret IndirectCycleT
		return ret
	}).(IndirectCycleTOutput)
}

func (o IndirectCycleTPtrOutput) Foo3() IndirectCycleSPtrOutput {
	return o.ApplyT(func(v *IndirectCycleT) *IndirectCycleS {
		if v == nil {
			return nil
		}
		return v.Foo3
	}).(IndirectCycleSPtrOutput)
}

func init() {
	khulnasoft.RegisterInputType(reflect.TypeOf((*AcyclicReferentInput)(nil)).Elem(), AcyclicReferentArgs{})
	khulnasoft.RegisterInputType(reflect.TypeOf((*AcyclicSInput)(nil)).Elem(), AcyclicSArgs{})
	khulnasoft.RegisterInputType(reflect.TypeOf((*AcyclicTInput)(nil)).Elem(), AcyclicTArgs{})
	khulnasoft.RegisterInputType(reflect.TypeOf((*DirectCycleInput)(nil)).Elem(), DirectCycleArgs{})
	khulnasoft.RegisterInputType(reflect.TypeOf((*DirectCyclePtrInput)(nil)).Elem(), DirectCycleArgs{})
	khulnasoft.RegisterInputType(reflect.TypeOf((*IndirectCycleSInput)(nil)).Elem(), IndirectCycleSArgs{})
	khulnasoft.RegisterInputType(reflect.TypeOf((*IndirectCycleSPtrInput)(nil)).Elem(), IndirectCycleSArgs{})
	khulnasoft.RegisterInputType(reflect.TypeOf((*IndirectCycleTInput)(nil)).Elem(), IndirectCycleTArgs{})
	khulnasoft.RegisterInputType(reflect.TypeOf((*IndirectCycleTPtrInput)(nil)).Elem(), IndirectCycleTArgs{})
	khulnasoft.RegisterOutputType(AcyclicReferentOutput{})
	khulnasoft.RegisterOutputType(AcyclicSOutput{})
	khulnasoft.RegisterOutputType(AcyclicTOutput{})
	khulnasoft.RegisterOutputType(DirectCycleOutput{})
	khulnasoft.RegisterOutputType(DirectCyclePtrOutput{})
	khulnasoft.RegisterOutputType(IndirectCycleSOutput{})
	khulnasoft.RegisterOutputType(IndirectCycleSPtrOutput{})
	khulnasoft.RegisterOutputType(IndirectCycleTOutput{})
	khulnasoft.RegisterOutputType(IndirectCycleTPtrOutput{})
}
