// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package plantprovider

import (
	"context"
	"reflect"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

// The log_name to populate in the Cloud Audit Record. This is added to regress khulnasoft/khulnasoft issue #7913
type CloudAuditOptionsLogName string

const (
	// Default. Should not be used.
	CloudAuditOptionsLogNameUnspecifiedLogName = CloudAuditOptionsLogName("UNSPECIFIED_LOG_NAME")
	// Corresponds to "cloudaudit.googleapis.com/activity"
	CloudAuditOptionsLogNameAdminActivity = CloudAuditOptionsLogName("ADMIN_ACTIVITY")
	// Corresponds to "cloudaudit.googleapis.com/data_access"
	CloudAuditOptionsLogNameDataAccess = CloudAuditOptionsLogName("DATA_ACCESS")
	// What if triple quotes """ are used in the description
	CloudAuditOptionsLogNameSynthetic = CloudAuditOptionsLogName("SYNTHETIC")
)

func (CloudAuditOptionsLogName) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudAuditOptionsLogName)(nil)).Elem()
}

func (e CloudAuditOptionsLogName) ToCloudAuditOptionsLogNameOutput() CloudAuditOptionsLogNameOutput {
	return khulnasoft.ToOutput(e).(CloudAuditOptionsLogNameOutput)
}

func (e CloudAuditOptionsLogName) ToCloudAuditOptionsLogNameOutputWithContext(ctx context.Context) CloudAuditOptionsLogNameOutput {
	return khulnasoft.ToOutputWithContext(ctx, e).(CloudAuditOptionsLogNameOutput)
}

func (e CloudAuditOptionsLogName) ToCloudAuditOptionsLogNamePtrOutput() CloudAuditOptionsLogNamePtrOutput {
	return e.ToCloudAuditOptionsLogNamePtrOutputWithContext(context.Background())
}

func (e CloudAuditOptionsLogName) ToCloudAuditOptionsLogNamePtrOutputWithContext(ctx context.Context) CloudAuditOptionsLogNamePtrOutput {
	return CloudAuditOptionsLogName(e).ToCloudAuditOptionsLogNameOutputWithContext(ctx).ToCloudAuditOptionsLogNamePtrOutputWithContext(ctx)
}

func (e CloudAuditOptionsLogName) ToStringOutput() khulnasoft.StringOutput {
	return khulnasoft.ToOutput(khulnasoft.String(e)).(khulnasoft.StringOutput)
}

func (e CloudAuditOptionsLogName) ToStringOutputWithContext(ctx context.Context) khulnasoft.StringOutput {
	return khulnasoft.ToOutputWithContext(ctx, khulnasoft.String(e)).(khulnasoft.StringOutput)
}

func (e CloudAuditOptionsLogName) ToStringPtrOutput() khulnasoft.StringPtrOutput {
	return khulnasoft.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CloudAuditOptionsLogName) ToStringPtrOutputWithContext(ctx context.Context) khulnasoft.StringPtrOutput {
	return khulnasoft.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CloudAuditOptionsLogNameOutput struct{ *khulnasoft.OutputState }

func (CloudAuditOptionsLogNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudAuditOptionsLogName)(nil)).Elem()
}

func (o CloudAuditOptionsLogNameOutput) ToCloudAuditOptionsLogNameOutput() CloudAuditOptionsLogNameOutput {
	return o
}

func (o CloudAuditOptionsLogNameOutput) ToCloudAuditOptionsLogNameOutputWithContext(ctx context.Context) CloudAuditOptionsLogNameOutput {
	return o
}

func (o CloudAuditOptionsLogNameOutput) ToCloudAuditOptionsLogNamePtrOutput() CloudAuditOptionsLogNamePtrOutput {
	return o.ToCloudAuditOptionsLogNamePtrOutputWithContext(context.Background())
}

func (o CloudAuditOptionsLogNameOutput) ToCloudAuditOptionsLogNamePtrOutputWithContext(ctx context.Context) CloudAuditOptionsLogNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CloudAuditOptionsLogName) *CloudAuditOptionsLogName {
		return &v
	}).(CloudAuditOptionsLogNamePtrOutput)
}

func (o CloudAuditOptionsLogNameOutput) ToStringOutput() khulnasoft.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CloudAuditOptionsLogNameOutput) ToStringOutputWithContext(ctx context.Context) khulnasoft.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CloudAuditOptionsLogName) string {
		return string(e)
	}).(khulnasoft.StringOutput)
}

func (o CloudAuditOptionsLogNameOutput) ToStringPtrOutput() khulnasoft.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CloudAuditOptionsLogNameOutput) ToStringPtrOutputWithContext(ctx context.Context) khulnasoft.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CloudAuditOptionsLogName) *string {
		v := string(e)
		return &v
	}).(khulnasoft.StringPtrOutput)
}

type CloudAuditOptionsLogNamePtrOutput struct{ *khulnasoft.OutputState }

func (CloudAuditOptionsLogNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudAuditOptionsLogName)(nil)).Elem()
}

func (o CloudAuditOptionsLogNamePtrOutput) ToCloudAuditOptionsLogNamePtrOutput() CloudAuditOptionsLogNamePtrOutput {
	return o
}

func (o CloudAuditOptionsLogNamePtrOutput) ToCloudAuditOptionsLogNamePtrOutputWithContext(ctx context.Context) CloudAuditOptionsLogNamePtrOutput {
	return o
}

func (o CloudAuditOptionsLogNamePtrOutput) Elem() CloudAuditOptionsLogNameOutput {
	return o.ApplyT(func(v *CloudAuditOptionsLogName) CloudAuditOptionsLogName {
		if v != nil {
			return *v
		}
		var ret CloudAuditOptionsLogName
		return ret
	}).(CloudAuditOptionsLogNameOutput)
}

func (o CloudAuditOptionsLogNamePtrOutput) ToStringPtrOutput() khulnasoft.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CloudAuditOptionsLogNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) khulnasoft.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CloudAuditOptionsLogName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(khulnasoft.StringPtrOutput)
}

// CloudAuditOptionsLogNameInput is an input type that accepts values of the CloudAuditOptionsLogName enum
// A concrete instance of `CloudAuditOptionsLogNameInput` can be one of the following:
//
//	CloudAuditOptionsLogNameUnspecifiedLogName
//	CloudAuditOptionsLogNameAdminActivity
//	CloudAuditOptionsLogNameDataAccess
//	CloudAuditOptionsLogNameSynthetic
type CloudAuditOptionsLogNameInput interface {
	khulnasoft.Input

	ToCloudAuditOptionsLogNameOutput() CloudAuditOptionsLogNameOutput
	ToCloudAuditOptionsLogNameOutputWithContext(context.Context) CloudAuditOptionsLogNameOutput
}

var cloudAuditOptionsLogNamePtrType = reflect.TypeOf((**CloudAuditOptionsLogName)(nil)).Elem()

type CloudAuditOptionsLogNamePtrInput interface {
	khulnasoft.Input

	ToCloudAuditOptionsLogNamePtrOutput() CloudAuditOptionsLogNamePtrOutput
	ToCloudAuditOptionsLogNamePtrOutputWithContext(context.Context) CloudAuditOptionsLogNamePtrOutput
}

type cloudAuditOptionsLogNamePtr string

func CloudAuditOptionsLogNamePtr(v string) CloudAuditOptionsLogNamePtrInput {
	return (*cloudAuditOptionsLogNamePtr)(&v)
}

func (*cloudAuditOptionsLogNamePtr) ElementType() reflect.Type {
	return cloudAuditOptionsLogNamePtrType
}

func (in *cloudAuditOptionsLogNamePtr) ToCloudAuditOptionsLogNamePtrOutput() CloudAuditOptionsLogNamePtrOutput {
	return khulnasoft.ToOutput(in).(CloudAuditOptionsLogNamePtrOutput)
}

func (in *cloudAuditOptionsLogNamePtr) ToCloudAuditOptionsLogNamePtrOutputWithContext(ctx context.Context) CloudAuditOptionsLogNamePtrOutput {
	return khulnasoft.ToOutputWithContext(ctx, in).(CloudAuditOptionsLogNamePtrOutput)
}

type ContainerBrightness float64

const (
	ContainerBrightnessZeroPointOne = ContainerBrightness(0.1)
	ContainerBrightnessOne          = ContainerBrightness(1)
)

func (ContainerBrightness) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerBrightness)(nil)).Elem()
}

func (e ContainerBrightness) ToContainerBrightnessOutput() ContainerBrightnessOutput {
	return khulnasoft.ToOutput(e).(ContainerBrightnessOutput)
}

func (e ContainerBrightness) ToContainerBrightnessOutputWithContext(ctx context.Context) ContainerBrightnessOutput {
	return khulnasoft.ToOutputWithContext(ctx, e).(ContainerBrightnessOutput)
}

func (e ContainerBrightness) ToContainerBrightnessPtrOutput() ContainerBrightnessPtrOutput {
	return e.ToContainerBrightnessPtrOutputWithContext(context.Background())
}

func (e ContainerBrightness) ToContainerBrightnessPtrOutputWithContext(ctx context.Context) ContainerBrightnessPtrOutput {
	return ContainerBrightness(e).ToContainerBrightnessOutputWithContext(ctx).ToContainerBrightnessPtrOutputWithContext(ctx)
}

func (e ContainerBrightness) ToFloat64Output() khulnasoft.Float64Output {
	return khulnasoft.ToOutput(khulnasoft.Float64(e)).(khulnasoft.Float64Output)
}

func (e ContainerBrightness) ToFloat64OutputWithContext(ctx context.Context) khulnasoft.Float64Output {
	return khulnasoft.ToOutputWithContext(ctx, khulnasoft.Float64(e)).(khulnasoft.Float64Output)
}

func (e ContainerBrightness) ToFloat64PtrOutput() khulnasoft.Float64PtrOutput {
	return khulnasoft.Float64(e).ToFloat64PtrOutputWithContext(context.Background())
}

func (e ContainerBrightness) ToFloat64PtrOutputWithContext(ctx context.Context) khulnasoft.Float64PtrOutput {
	return khulnasoft.Float64(e).ToFloat64OutputWithContext(ctx).ToFloat64PtrOutputWithContext(ctx)
}

type ContainerBrightnessOutput struct{ *khulnasoft.OutputState }

func (ContainerBrightnessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerBrightness)(nil)).Elem()
}

func (o ContainerBrightnessOutput) ToContainerBrightnessOutput() ContainerBrightnessOutput {
	return o
}

func (o ContainerBrightnessOutput) ToContainerBrightnessOutputWithContext(ctx context.Context) ContainerBrightnessOutput {
	return o
}

func (o ContainerBrightnessOutput) ToContainerBrightnessPtrOutput() ContainerBrightnessPtrOutput {
	return o.ToContainerBrightnessPtrOutputWithContext(context.Background())
}

func (o ContainerBrightnessOutput) ToContainerBrightnessPtrOutputWithContext(ctx context.Context) ContainerBrightnessPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerBrightness) *ContainerBrightness {
		return &v
	}).(ContainerBrightnessPtrOutput)
}

func (o ContainerBrightnessOutput) ToFloat64Output() khulnasoft.Float64Output {
	return o.ToFloat64OutputWithContext(context.Background())
}

func (o ContainerBrightnessOutput) ToFloat64OutputWithContext(ctx context.Context) khulnasoft.Float64Output {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContainerBrightness) float64 {
		return float64(e)
	}).(khulnasoft.Float64Output)
}

func (o ContainerBrightnessOutput) ToFloat64PtrOutput() khulnasoft.Float64PtrOutput {
	return o.ToFloat64PtrOutputWithContext(context.Background())
}

func (o ContainerBrightnessOutput) ToFloat64PtrOutputWithContext(ctx context.Context) khulnasoft.Float64PtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContainerBrightness) *float64 {
		v := float64(e)
		return &v
	}).(khulnasoft.Float64PtrOutput)
}

type ContainerBrightnessPtrOutput struct{ *khulnasoft.OutputState }

func (ContainerBrightnessPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerBrightness)(nil)).Elem()
}

func (o ContainerBrightnessPtrOutput) ToContainerBrightnessPtrOutput() ContainerBrightnessPtrOutput {
	return o
}

func (o ContainerBrightnessPtrOutput) ToContainerBrightnessPtrOutputWithContext(ctx context.Context) ContainerBrightnessPtrOutput {
	return o
}

func (o ContainerBrightnessPtrOutput) Elem() ContainerBrightnessOutput {
	return o.ApplyT(func(v *ContainerBrightness) ContainerBrightness {
		if v != nil {
			return *v
		}
		var ret ContainerBrightness
		return ret
	}).(ContainerBrightnessOutput)
}

func (o ContainerBrightnessPtrOutput) ToFloat64PtrOutput() khulnasoft.Float64PtrOutput {
	return o.ToFloat64PtrOutputWithContext(context.Background())
}

func (o ContainerBrightnessPtrOutput) ToFloat64PtrOutputWithContext(ctx context.Context) khulnasoft.Float64PtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ContainerBrightness) *float64 {
		if e == nil {
			return nil
		}
		v := float64(*e)
		return &v
	}).(khulnasoft.Float64PtrOutput)
}

// ContainerBrightnessInput is an input type that accepts values of the ContainerBrightness enum
// A concrete instance of `ContainerBrightnessInput` can be one of the following:
//
//	ContainerBrightnessZeroPointOne
//	ContainerBrightnessOne
type ContainerBrightnessInput interface {
	khulnasoft.Input

	ToContainerBrightnessOutput() ContainerBrightnessOutput
	ToContainerBrightnessOutputWithContext(context.Context) ContainerBrightnessOutput
}

var containerBrightnessPtrType = reflect.TypeOf((**ContainerBrightness)(nil)).Elem()

type ContainerBrightnessPtrInput interface {
	khulnasoft.Input

	ToContainerBrightnessPtrOutput() ContainerBrightnessPtrOutput
	ToContainerBrightnessPtrOutputWithContext(context.Context) ContainerBrightnessPtrOutput
}

type containerBrightnessPtr float64

func ContainerBrightnessPtr(v float64) ContainerBrightnessPtrInput {
	return (*containerBrightnessPtr)(&v)
}

func (*containerBrightnessPtr) ElementType() reflect.Type {
	return containerBrightnessPtrType
}

func (in *containerBrightnessPtr) ToContainerBrightnessPtrOutput() ContainerBrightnessPtrOutput {
	return khulnasoft.ToOutput(in).(ContainerBrightnessPtrOutput)
}

func (in *containerBrightnessPtr) ToContainerBrightnessPtrOutputWithContext(ctx context.Context) ContainerBrightnessPtrOutput {
	return khulnasoft.ToOutputWithContext(ctx, in).(ContainerBrightnessPtrOutput)
}

// plant container colors
type ContainerColor string

const (
	ContainerColorRed    = ContainerColor("red")
	ContainerColorBlue   = ContainerColor("blue")
	ContainerColorYellow = ContainerColor("yellow")
)

func (ContainerColor) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerColor)(nil)).Elem()
}

func (e ContainerColor) ToContainerColorOutput() ContainerColorOutput {
	return khulnasoft.ToOutput(e).(ContainerColorOutput)
}

func (e ContainerColor) ToContainerColorOutputWithContext(ctx context.Context) ContainerColorOutput {
	return khulnasoft.ToOutputWithContext(ctx, e).(ContainerColorOutput)
}

func (e ContainerColor) ToContainerColorPtrOutput() ContainerColorPtrOutput {
	return e.ToContainerColorPtrOutputWithContext(context.Background())
}

func (e ContainerColor) ToContainerColorPtrOutputWithContext(ctx context.Context) ContainerColorPtrOutput {
	return ContainerColor(e).ToContainerColorOutputWithContext(ctx).ToContainerColorPtrOutputWithContext(ctx)
}

func (e ContainerColor) ToStringOutput() khulnasoft.StringOutput {
	return khulnasoft.ToOutput(khulnasoft.String(e)).(khulnasoft.StringOutput)
}

func (e ContainerColor) ToStringOutputWithContext(ctx context.Context) khulnasoft.StringOutput {
	return khulnasoft.ToOutputWithContext(ctx, khulnasoft.String(e)).(khulnasoft.StringOutput)
}

func (e ContainerColor) ToStringPtrOutput() khulnasoft.StringPtrOutput {
	return khulnasoft.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ContainerColor) ToStringPtrOutputWithContext(ctx context.Context) khulnasoft.StringPtrOutput {
	return khulnasoft.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ContainerColorOutput struct{ *khulnasoft.OutputState }

func (ContainerColorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerColor)(nil)).Elem()
}

func (o ContainerColorOutput) ToContainerColorOutput() ContainerColorOutput {
	return o
}

func (o ContainerColorOutput) ToContainerColorOutputWithContext(ctx context.Context) ContainerColorOutput {
	return o
}

func (o ContainerColorOutput) ToContainerColorPtrOutput() ContainerColorPtrOutput {
	return o.ToContainerColorPtrOutputWithContext(context.Background())
}

func (o ContainerColorOutput) ToContainerColorPtrOutputWithContext(ctx context.Context) ContainerColorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerColor) *ContainerColor {
		return &v
	}).(ContainerColorPtrOutput)
}

func (o ContainerColorOutput) ToStringOutput() khulnasoft.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ContainerColorOutput) ToStringOutputWithContext(ctx context.Context) khulnasoft.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContainerColor) string {
		return string(e)
	}).(khulnasoft.StringOutput)
}

func (o ContainerColorOutput) ToStringPtrOutput() khulnasoft.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContainerColorOutput) ToStringPtrOutputWithContext(ctx context.Context) khulnasoft.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContainerColor) *string {
		v := string(e)
		return &v
	}).(khulnasoft.StringPtrOutput)
}

type ContainerColorPtrOutput struct{ *khulnasoft.OutputState }

func (ContainerColorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerColor)(nil)).Elem()
}

func (o ContainerColorPtrOutput) ToContainerColorPtrOutput() ContainerColorPtrOutput {
	return o
}

func (o ContainerColorPtrOutput) ToContainerColorPtrOutputWithContext(ctx context.Context) ContainerColorPtrOutput {
	return o
}

func (o ContainerColorPtrOutput) Elem() ContainerColorOutput {
	return o.ApplyT(func(v *ContainerColor) ContainerColor {
		if v != nil {
			return *v
		}
		var ret ContainerColor
		return ret
	}).(ContainerColorOutput)
}

func (o ContainerColorPtrOutput) ToStringPtrOutput() khulnasoft.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContainerColorPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) khulnasoft.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ContainerColor) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(khulnasoft.StringPtrOutput)
}

// ContainerColorInput is an input type that accepts values of the ContainerColor enum
// A concrete instance of `ContainerColorInput` can be one of the following:
//
//	ContainerColorRed
//	ContainerColorBlue
//	ContainerColorYellow
type ContainerColorInput interface {
	khulnasoft.Input

	ToContainerColorOutput() ContainerColorOutput
	ToContainerColorOutputWithContext(context.Context) ContainerColorOutput
}

var containerColorPtrType = reflect.TypeOf((**ContainerColor)(nil)).Elem()

type ContainerColorPtrInput interface {
	khulnasoft.Input

	ToContainerColorPtrOutput() ContainerColorPtrOutput
	ToContainerColorPtrOutputWithContext(context.Context) ContainerColorPtrOutput
}

type containerColorPtr string

func ContainerColorPtr(v string) ContainerColorPtrInput {
	return (*containerColorPtr)(&v)
}

func (*containerColorPtr) ElementType() reflect.Type {
	return containerColorPtrType
}

func (in *containerColorPtr) ToContainerColorPtrOutput() ContainerColorPtrOutput {
	return khulnasoft.ToOutput(in).(ContainerColorPtrOutput)
}

func (in *containerColorPtr) ToContainerColorPtrOutputWithContext(ctx context.Context) ContainerColorPtrOutput {
	return khulnasoft.ToOutputWithContext(ctx, in).(ContainerColorPtrOutput)
}

// plant container sizes
type ContainerSize int

const (
	ContainerSizeFourInch = ContainerSize(4)
	ContainerSizeSixInch  = ContainerSize(6)
	// Deprecated: Eight inch pots are no longer supported.
	ContainerSizeEightInch = ContainerSize(8)
)

func (ContainerSize) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerSize)(nil)).Elem()
}

func (e ContainerSize) ToContainerSizeOutput() ContainerSizeOutput {
	return khulnasoft.ToOutput(e).(ContainerSizeOutput)
}

func (e ContainerSize) ToContainerSizeOutputWithContext(ctx context.Context) ContainerSizeOutput {
	return khulnasoft.ToOutputWithContext(ctx, e).(ContainerSizeOutput)
}

func (e ContainerSize) ToContainerSizePtrOutput() ContainerSizePtrOutput {
	return e.ToContainerSizePtrOutputWithContext(context.Background())
}

func (e ContainerSize) ToContainerSizePtrOutputWithContext(ctx context.Context) ContainerSizePtrOutput {
	return ContainerSize(e).ToContainerSizeOutputWithContext(ctx).ToContainerSizePtrOutputWithContext(ctx)
}

func (e ContainerSize) ToIntOutput() khulnasoft.IntOutput {
	return khulnasoft.ToOutput(khulnasoft.Int(e)).(khulnasoft.IntOutput)
}

func (e ContainerSize) ToIntOutputWithContext(ctx context.Context) khulnasoft.IntOutput {
	return khulnasoft.ToOutputWithContext(ctx, khulnasoft.Int(e)).(khulnasoft.IntOutput)
}

func (e ContainerSize) ToIntPtrOutput() khulnasoft.IntPtrOutput {
	return khulnasoft.Int(e).ToIntPtrOutputWithContext(context.Background())
}

func (e ContainerSize) ToIntPtrOutputWithContext(ctx context.Context) khulnasoft.IntPtrOutput {
	return khulnasoft.Int(e).ToIntOutputWithContext(ctx).ToIntPtrOutputWithContext(ctx)
}

type ContainerSizeOutput struct{ *khulnasoft.OutputState }

func (ContainerSizeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerSize)(nil)).Elem()
}

func (o ContainerSizeOutput) ToContainerSizeOutput() ContainerSizeOutput {
	return o
}

func (o ContainerSizeOutput) ToContainerSizeOutputWithContext(ctx context.Context) ContainerSizeOutput {
	return o
}

func (o ContainerSizeOutput) ToContainerSizePtrOutput() ContainerSizePtrOutput {
	return o.ToContainerSizePtrOutputWithContext(context.Background())
}

func (o ContainerSizeOutput) ToContainerSizePtrOutputWithContext(ctx context.Context) ContainerSizePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerSize) *ContainerSize {
		return &v
	}).(ContainerSizePtrOutput)
}

func (o ContainerSizeOutput) ToIntOutput() khulnasoft.IntOutput {
	return o.ToIntOutputWithContext(context.Background())
}

func (o ContainerSizeOutput) ToIntOutputWithContext(ctx context.Context) khulnasoft.IntOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContainerSize) int {
		return int(e)
	}).(khulnasoft.IntOutput)
}

func (o ContainerSizeOutput) ToIntPtrOutput() khulnasoft.IntPtrOutput {
	return o.ToIntPtrOutputWithContext(context.Background())
}

func (o ContainerSizeOutput) ToIntPtrOutputWithContext(ctx context.Context) khulnasoft.IntPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContainerSize) *int {
		v := int(e)
		return &v
	}).(khulnasoft.IntPtrOutput)
}

type ContainerSizePtrOutput struct{ *khulnasoft.OutputState }

func (ContainerSizePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerSize)(nil)).Elem()
}

func (o ContainerSizePtrOutput) ToContainerSizePtrOutput() ContainerSizePtrOutput {
	return o
}

func (o ContainerSizePtrOutput) ToContainerSizePtrOutputWithContext(ctx context.Context) ContainerSizePtrOutput {
	return o
}

func (o ContainerSizePtrOutput) Elem() ContainerSizeOutput {
	return o.ApplyT(func(v *ContainerSize) ContainerSize {
		if v != nil {
			return *v
		}
		var ret ContainerSize
		return ret
	}).(ContainerSizeOutput)
}

func (o ContainerSizePtrOutput) ToIntPtrOutput() khulnasoft.IntPtrOutput {
	return o.ToIntPtrOutputWithContext(context.Background())
}

func (o ContainerSizePtrOutput) ToIntPtrOutputWithContext(ctx context.Context) khulnasoft.IntPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ContainerSize) *int {
		if e == nil {
			return nil
		}
		v := int(*e)
		return &v
	}).(khulnasoft.IntPtrOutput)
}

// ContainerSizeInput is an input type that accepts values of the ContainerSize enum
// A concrete instance of `ContainerSizeInput` can be one of the following:
//
//	ContainerSizeFourInch
//	ContainerSizeSixInch
type ContainerSizeInput interface {
	khulnasoft.Input

	ToContainerSizeOutput() ContainerSizeOutput
	ToContainerSizeOutputWithContext(context.Context) ContainerSizeOutput
}

var containerSizePtrType = reflect.TypeOf((**ContainerSize)(nil)).Elem()

type ContainerSizePtrInput interface {
	khulnasoft.Input

	ToContainerSizePtrOutput() ContainerSizePtrOutput
	ToContainerSizePtrOutputWithContext(context.Context) ContainerSizePtrOutput
}

type containerSizePtr int

func ContainerSizePtr(v int) ContainerSizePtrInput {
	return (*containerSizePtr)(&v)
}

func (*containerSizePtr) ElementType() reflect.Type {
	return containerSizePtrType
}

func (in *containerSizePtr) ToContainerSizePtrOutput() ContainerSizePtrOutput {
	return khulnasoft.ToOutput(in).(ContainerSizePtrOutput)
}

func (in *containerSizePtr) ToContainerSizePtrOutputWithContext(ctx context.Context) ContainerSizePtrOutput {
	return khulnasoft.ToOutputWithContext(ctx, in).(ContainerSizePtrOutput)
}

func init() {
	khulnasoft.RegisterInputType(reflect.TypeOf((*CloudAuditOptionsLogNameInput)(nil)).Elem(), CloudAuditOptionsLogName("UNSPECIFIED_LOG_NAME"))
	khulnasoft.RegisterInputType(reflect.TypeOf((*CloudAuditOptionsLogNamePtrInput)(nil)).Elem(), CloudAuditOptionsLogName("UNSPECIFIED_LOG_NAME"))
	khulnasoft.RegisterInputType(reflect.TypeOf((*ContainerBrightnessInput)(nil)).Elem(), ContainerBrightness(0.1))
	khulnasoft.RegisterInputType(reflect.TypeOf((*ContainerBrightnessPtrInput)(nil)).Elem(), ContainerBrightness(0.1))
	khulnasoft.RegisterInputType(reflect.TypeOf((*ContainerColorInput)(nil)).Elem(), ContainerColor("red"))
	khulnasoft.RegisterInputType(reflect.TypeOf((*ContainerColorPtrInput)(nil)).Elem(), ContainerColor("red"))
	khulnasoft.RegisterInputType(reflect.TypeOf((*ContainerSizeInput)(nil)).Elem(), ContainerSize(4))
	khulnasoft.RegisterInputType(reflect.TypeOf((*ContainerSizePtrInput)(nil)).Elem(), ContainerSize(4))
	khulnasoft.RegisterOutputType(CloudAuditOptionsLogNameOutput{})
	khulnasoft.RegisterOutputType(CloudAuditOptionsLogNamePtrOutput{})
	khulnasoft.RegisterOutputType(ContainerBrightnessOutput{})
	khulnasoft.RegisterOutputType(ContainerBrightnessPtrOutput{})
	khulnasoft.RegisterOutputType(ContainerColorOutput{})
	khulnasoft.RegisterOutputType(ContainerColorPtrOutput{})
	khulnasoft.RegisterOutputType(ContainerSizeOutput{})
	khulnasoft.RegisterOutputType(ContainerSizePtrOutput{})
}
