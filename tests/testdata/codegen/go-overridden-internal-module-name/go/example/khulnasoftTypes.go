// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"context"
	"reflect"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"go-overridden-internal-module-name/example/utilities"
)

var _ = utilities.GetEnvOrDefault

type ConfigMap struct {
	Config *string `khulnasoft:"config"`
}

// ConfigMapInput is an input type that accepts ConfigMap and ConfigMapOutput values.
// You can construct a concrete instance of `ConfigMapInput` via:
//
//	ConfigMap{ "key": ConfigArgs{...} }
type ConfigMapInput interface {
	khulnasoft.Input

	ToConfigMapOutput() ConfigMapOutput
	ToConfigMapOutputWithContext(context.Context) ConfigMapOutput
}

type ConfigMapArgs struct {
	Config khulnasoft.StringPtrInput `khulnasoft:"config"`
}

func (ConfigMapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigMap)(nil)).Elem()
}

func (i ConfigMapArgs) ToConfigMapOutput() ConfigMapOutput {
	return i.ToConfigMapOutputWithContext(context.Background())
}

func (i ConfigMapArgs) ToConfigMapOutputWithContext(ctx context.Context) ConfigMapOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(ConfigMapOutput)
}

// ConfigMapArrayInput is an input type that accepts ConfigMapArray and ConfigMapArrayOutput values.
// You can construct a concrete instance of `ConfigMapArrayInput` via:
//
//	ConfigMapArray{ ConfigMapArgs{...} }
type ConfigMapArrayInput interface {
	khulnasoft.Input

	ToConfigMapArrayOutput() ConfigMapArrayOutput
	ToConfigMapArrayOutputWithContext(context.Context) ConfigMapArrayOutput
}

type ConfigMapArray []ConfigMapInput

func (ConfigMapArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigMap)(nil)).Elem()
}

func (i ConfigMapArray) ToConfigMapArrayOutput() ConfigMapArrayOutput {
	return i.ToConfigMapArrayOutputWithContext(context.Background())
}

func (i ConfigMapArray) ToConfigMapArrayOutputWithContext(ctx context.Context) ConfigMapArrayOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(ConfigMapArrayOutput)
}

type ConfigMapOutput struct{ *khulnasoft.OutputState }

func (ConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigMap)(nil)).Elem()
}

func (o ConfigMapOutput) ToConfigMapOutput() ConfigMapOutput {
	return o
}

func (o ConfigMapOutput) ToConfigMapOutputWithContext(ctx context.Context) ConfigMapOutput {
	return o
}

func (o ConfigMapOutput) Config() khulnasoft.StringPtrOutput {
	return o.ApplyT(func(v ConfigMap) *string { return v.Config }).(khulnasoft.StringPtrOutput)
}

type ConfigMapArrayOutput struct{ *khulnasoft.OutputState }

func (ConfigMapArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigMap)(nil)).Elem()
}

func (o ConfigMapArrayOutput) ToConfigMapArrayOutput() ConfigMapArrayOutput {
	return o
}

func (o ConfigMapArrayOutput) ToConfigMapArrayOutputWithContext(ctx context.Context) ConfigMapArrayOutput {
	return o
}

func (o ConfigMapArrayOutput) Index(i khulnasoft.IntInput) ConfigMapOutput {
	return khulnasoft.All(o, i).ApplyT(func(vs []interface{}) ConfigMap {
		return vs[0].([]ConfigMap)[vs[1].(int)]
	}).(ConfigMapOutput)
}

type Object struct {
	Bar     *string     `khulnasoft:"bar"`
	Configs []ConfigMap `khulnasoft:"configs"`
	Foo     *Resource   `khulnasoft:"foo"`
	// List of lists of other objects
	Others [][]SomeOtherObject `khulnasoft:"others"`
	// Mapping from string to list of some other object
	StillOthers map[string][]SomeOtherObject `khulnasoft:"stillOthers"`
}

// ObjectInput is an input type that accepts ObjectArgs and ObjectOutput values.
// You can construct a concrete instance of `ObjectInput` via:
//
//	ObjectArgs{...}
type ObjectInput interface {
	khulnasoft.Input

	ToObjectOutput() ObjectOutput
	ToObjectOutputWithContext(context.Context) ObjectOutput
}

type ObjectArgs struct {
	Bar     khulnasoft.StringPtrInput `khulnasoft:"bar"`
	Configs ConfigMapArrayInput   `khulnasoft:"configs"`
	Foo     ResourceInput         `khulnasoft:"foo"`
	// List of lists of other objects
	Others SomeOtherObjectArrayArrayInput `khulnasoft:"others"`
	// Mapping from string to list of some other object
	StillOthers SomeOtherObjectArrayMapInput `khulnasoft:"stillOthers"`
}

func (ObjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Object)(nil)).Elem()
}

func (i ObjectArgs) ToObjectOutput() ObjectOutput {
	return i.ToObjectOutputWithContext(context.Background())
}

func (i ObjectArgs) ToObjectOutputWithContext(ctx context.Context) ObjectOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(ObjectOutput)
}

func (i ObjectArgs) ToObjectPtrOutput() ObjectPtrOutput {
	return i.ToObjectPtrOutputWithContext(context.Background())
}

func (i ObjectArgs) ToObjectPtrOutputWithContext(ctx context.Context) ObjectPtrOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(ObjectOutput).ToObjectPtrOutputWithContext(ctx)
}

// ObjectPtrInput is an input type that accepts ObjectArgs, ObjectPtr and ObjectPtrOutput values.
// You can construct a concrete instance of `ObjectPtrInput` via:
//
//	        ObjectArgs{...}
//
//	or:
//
//	        nil
type ObjectPtrInput interface {
	khulnasoft.Input

	ToObjectPtrOutput() ObjectPtrOutput
	ToObjectPtrOutputWithContext(context.Context) ObjectPtrOutput
}

type objectPtrType ObjectArgs

func ObjectPtr(v *ObjectArgs) ObjectPtrInput {
	return (*objectPtrType)(v)
}

func (*objectPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Object)(nil)).Elem()
}

func (i *objectPtrType) ToObjectPtrOutput() ObjectPtrOutput {
	return i.ToObjectPtrOutputWithContext(context.Background())
}

func (i *objectPtrType) ToObjectPtrOutputWithContext(ctx context.Context) ObjectPtrOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(ObjectPtrOutput)
}

type ObjectOutput struct{ *khulnasoft.OutputState }

func (ObjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Object)(nil)).Elem()
}

func (o ObjectOutput) ToObjectOutput() ObjectOutput {
	return o
}

func (o ObjectOutput) ToObjectOutputWithContext(ctx context.Context) ObjectOutput {
	return o
}

func (o ObjectOutput) ToObjectPtrOutput() ObjectPtrOutput {
	return o.ToObjectPtrOutputWithContext(context.Background())
}

func (o ObjectOutput) ToObjectPtrOutputWithContext(ctx context.Context) ObjectPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Object) *Object {
		return &v
	}).(ObjectPtrOutput)
}

func (o ObjectOutput) Bar() khulnasoft.StringPtrOutput {
	return o.ApplyT(func(v Object) *string { return v.Bar }).(khulnasoft.StringPtrOutput)
}

func (o ObjectOutput) Configs() ConfigMapArrayOutput {
	return o.ApplyT(func(v Object) []ConfigMap { return v.Configs }).(ConfigMapArrayOutput)
}

func (o ObjectOutput) Foo() ResourceOutput {
	return o.ApplyT(func(v Object) *Resource { return v.Foo }).(ResourceOutput)
}

// List of lists of other objects
func (o ObjectOutput) Others() SomeOtherObjectArrayArrayOutput {
	return o.ApplyT(func(v Object) [][]SomeOtherObject { return v.Others }).(SomeOtherObjectArrayArrayOutput)
}

// Mapping from string to list of some other object
func (o ObjectOutput) StillOthers() SomeOtherObjectArrayMapOutput {
	return o.ApplyT(func(v Object) map[string][]SomeOtherObject { return v.StillOthers }).(SomeOtherObjectArrayMapOutput)
}

type ObjectPtrOutput struct{ *khulnasoft.OutputState }

func (ObjectPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Object)(nil)).Elem()
}

func (o ObjectPtrOutput) ToObjectPtrOutput() ObjectPtrOutput {
	return o
}

func (o ObjectPtrOutput) ToObjectPtrOutputWithContext(ctx context.Context) ObjectPtrOutput {
	return o
}

func (o ObjectPtrOutput) Elem() ObjectOutput {
	return o.ApplyT(func(v *Object) Object {
		if v != nil {
			return *v
		}
		var ret Object
		return ret
	}).(ObjectOutput)
}

func (o ObjectPtrOutput) Bar() khulnasoft.StringPtrOutput {
	return o.ApplyT(func(v *Object) *string {
		if v == nil {
			return nil
		}
		return v.Bar
	}).(khulnasoft.StringPtrOutput)
}

func (o ObjectPtrOutput) Configs() ConfigMapArrayOutput {
	return o.ApplyT(func(v *Object) []ConfigMap {
		if v == nil {
			return nil
		}
		return v.Configs
	}).(ConfigMapArrayOutput)
}

func (o ObjectPtrOutput) Foo() ResourceOutput {
	return o.ApplyT(func(v *Object) *Resource {
		if v == nil {
			return nil
		}
		return v.Foo
	}).(ResourceOutput)
}

// List of lists of other objects
func (o ObjectPtrOutput) Others() SomeOtherObjectArrayArrayOutput {
	return o.ApplyT(func(v *Object) [][]SomeOtherObject {
		if v == nil {
			return nil
		}
		return v.Others
	}).(SomeOtherObjectArrayArrayOutput)
}

// Mapping from string to list of some other object
func (o ObjectPtrOutput) StillOthers() SomeOtherObjectArrayMapOutput {
	return o.ApplyT(func(v *Object) map[string][]SomeOtherObject {
		if v == nil {
			return nil
		}
		return v.StillOthers
	}).(SomeOtherObjectArrayMapOutput)
}

type ObjectWithNodeOptionalInputs struct {
	Bar *int   `khulnasoft:"bar"`
	Foo string `khulnasoft:"foo"`
}

// ObjectWithNodeOptionalInputsInput is an input type that accepts ObjectWithNodeOptionalInputsArgs and ObjectWithNodeOptionalInputsOutput values.
// You can construct a concrete instance of `ObjectWithNodeOptionalInputsInput` via:
//
//	ObjectWithNodeOptionalInputsArgs{...}
type ObjectWithNodeOptionalInputsInput interface {
	khulnasoft.Input

	ToObjectWithNodeOptionalInputsOutput() ObjectWithNodeOptionalInputsOutput
	ToObjectWithNodeOptionalInputsOutputWithContext(context.Context) ObjectWithNodeOptionalInputsOutput
}

type ObjectWithNodeOptionalInputsArgs struct {
	Bar khulnasoft.IntPtrInput `khulnasoft:"bar"`
	Foo khulnasoft.StringInput `khulnasoft:"foo"`
}

func (ObjectWithNodeOptionalInputsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ObjectWithNodeOptionalInputs)(nil)).Elem()
}

func (i ObjectWithNodeOptionalInputsArgs) ToObjectWithNodeOptionalInputsOutput() ObjectWithNodeOptionalInputsOutput {
	return i.ToObjectWithNodeOptionalInputsOutputWithContext(context.Background())
}

func (i ObjectWithNodeOptionalInputsArgs) ToObjectWithNodeOptionalInputsOutputWithContext(ctx context.Context) ObjectWithNodeOptionalInputsOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(ObjectWithNodeOptionalInputsOutput)
}

func (i ObjectWithNodeOptionalInputsArgs) ToObjectWithNodeOptionalInputsPtrOutput() ObjectWithNodeOptionalInputsPtrOutput {
	return i.ToObjectWithNodeOptionalInputsPtrOutputWithContext(context.Background())
}

func (i ObjectWithNodeOptionalInputsArgs) ToObjectWithNodeOptionalInputsPtrOutputWithContext(ctx context.Context) ObjectWithNodeOptionalInputsPtrOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(ObjectWithNodeOptionalInputsOutput).ToObjectWithNodeOptionalInputsPtrOutputWithContext(ctx)
}

// ObjectWithNodeOptionalInputsPtrInput is an input type that accepts ObjectWithNodeOptionalInputsArgs, ObjectWithNodeOptionalInputsPtr and ObjectWithNodeOptionalInputsPtrOutput values.
// You can construct a concrete instance of `ObjectWithNodeOptionalInputsPtrInput` via:
//
//	        ObjectWithNodeOptionalInputsArgs{...}
//
//	or:
//
//	        nil
type ObjectWithNodeOptionalInputsPtrInput interface {
	khulnasoft.Input

	ToObjectWithNodeOptionalInputsPtrOutput() ObjectWithNodeOptionalInputsPtrOutput
	ToObjectWithNodeOptionalInputsPtrOutputWithContext(context.Context) ObjectWithNodeOptionalInputsPtrOutput
}

type objectWithNodeOptionalInputsPtrType ObjectWithNodeOptionalInputsArgs

func ObjectWithNodeOptionalInputsPtr(v *ObjectWithNodeOptionalInputsArgs) ObjectWithNodeOptionalInputsPtrInput {
	return (*objectWithNodeOptionalInputsPtrType)(v)
}

func (*objectWithNodeOptionalInputsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectWithNodeOptionalInputs)(nil)).Elem()
}

func (i *objectWithNodeOptionalInputsPtrType) ToObjectWithNodeOptionalInputsPtrOutput() ObjectWithNodeOptionalInputsPtrOutput {
	return i.ToObjectWithNodeOptionalInputsPtrOutputWithContext(context.Background())
}

func (i *objectWithNodeOptionalInputsPtrType) ToObjectWithNodeOptionalInputsPtrOutputWithContext(ctx context.Context) ObjectWithNodeOptionalInputsPtrOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(ObjectWithNodeOptionalInputsPtrOutput)
}

type ObjectWithNodeOptionalInputsOutput struct{ *khulnasoft.OutputState }

func (ObjectWithNodeOptionalInputsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ObjectWithNodeOptionalInputs)(nil)).Elem()
}

func (o ObjectWithNodeOptionalInputsOutput) ToObjectWithNodeOptionalInputsOutput() ObjectWithNodeOptionalInputsOutput {
	return o
}

func (o ObjectWithNodeOptionalInputsOutput) ToObjectWithNodeOptionalInputsOutputWithContext(ctx context.Context) ObjectWithNodeOptionalInputsOutput {
	return o
}

func (o ObjectWithNodeOptionalInputsOutput) ToObjectWithNodeOptionalInputsPtrOutput() ObjectWithNodeOptionalInputsPtrOutput {
	return o.ToObjectWithNodeOptionalInputsPtrOutputWithContext(context.Background())
}

func (o ObjectWithNodeOptionalInputsOutput) ToObjectWithNodeOptionalInputsPtrOutputWithContext(ctx context.Context) ObjectWithNodeOptionalInputsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ObjectWithNodeOptionalInputs) *ObjectWithNodeOptionalInputs {
		return &v
	}).(ObjectWithNodeOptionalInputsPtrOutput)
}

func (o ObjectWithNodeOptionalInputsOutput) Bar() khulnasoft.IntPtrOutput {
	return o.ApplyT(func(v ObjectWithNodeOptionalInputs) *int { return v.Bar }).(khulnasoft.IntPtrOutput)
}

func (o ObjectWithNodeOptionalInputsOutput) Foo() khulnasoft.StringOutput {
	return o.ApplyT(func(v ObjectWithNodeOptionalInputs) string { return v.Foo }).(khulnasoft.StringOutput)
}

type ObjectWithNodeOptionalInputsPtrOutput struct{ *khulnasoft.OutputState }

func (ObjectWithNodeOptionalInputsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectWithNodeOptionalInputs)(nil)).Elem()
}

func (o ObjectWithNodeOptionalInputsPtrOutput) ToObjectWithNodeOptionalInputsPtrOutput() ObjectWithNodeOptionalInputsPtrOutput {
	return o
}

func (o ObjectWithNodeOptionalInputsPtrOutput) ToObjectWithNodeOptionalInputsPtrOutputWithContext(ctx context.Context) ObjectWithNodeOptionalInputsPtrOutput {
	return o
}

func (o ObjectWithNodeOptionalInputsPtrOutput) Elem() ObjectWithNodeOptionalInputsOutput {
	return o.ApplyT(func(v *ObjectWithNodeOptionalInputs) ObjectWithNodeOptionalInputs {
		if v != nil {
			return *v
		}
		var ret ObjectWithNodeOptionalInputs
		return ret
	}).(ObjectWithNodeOptionalInputsOutput)
}

func (o ObjectWithNodeOptionalInputsPtrOutput) Bar() khulnasoft.IntPtrOutput {
	return o.ApplyT(func(v *ObjectWithNodeOptionalInputs) *int {
		if v == nil {
			return nil
		}
		return v.Bar
	}).(khulnasoft.IntPtrOutput)
}

func (o ObjectWithNodeOptionalInputsPtrOutput) Foo() khulnasoft.StringPtrOutput {
	return o.ApplyT(func(v *ObjectWithNodeOptionalInputs) *string {
		if v == nil {
			return nil
		}
		return &v.Foo
	}).(khulnasoft.StringPtrOutput)
}

type OtherResourceOutputType struct {
	Foo *string `khulnasoft:"foo"`
}

// OtherResourceOutputTypeInput is an input type that accepts OtherResourceOutputTypeArgs and OtherResourceOutputTypeOutput values.
// You can construct a concrete instance of `OtherResourceOutputTypeInput` via:
//
//	OtherResourceOutputTypeArgs{...}
type OtherResourceOutputTypeInput interface {
	khulnasoft.Input

	ToOtherResourceOutputTypeOutput() OtherResourceOutputTypeOutput
	ToOtherResourceOutputTypeOutputWithContext(context.Context) OtherResourceOutputTypeOutput
}

type OtherResourceOutputTypeArgs struct {
	Foo khulnasoft.StringPtrInput `khulnasoft:"foo"`
}

func (OtherResourceOutputTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OtherResourceOutputType)(nil)).Elem()
}

func (i OtherResourceOutputTypeArgs) ToOtherResourceOutputTypeOutput() OtherResourceOutputTypeOutput {
	return i.ToOtherResourceOutputTypeOutputWithContext(context.Background())
}

func (i OtherResourceOutputTypeArgs) ToOtherResourceOutputTypeOutputWithContext(ctx context.Context) OtherResourceOutputTypeOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(OtherResourceOutputTypeOutput)
}

type OtherResourceOutputTypeOutput struct{ *khulnasoft.OutputState }

func (OtherResourceOutputTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OtherResourceOutputType)(nil)).Elem()
}

func (o OtherResourceOutputTypeOutput) ToOtherResourceOutputTypeOutput() OtherResourceOutputTypeOutput {
	return o
}

func (o OtherResourceOutputTypeOutput) ToOtherResourceOutputTypeOutputWithContext(ctx context.Context) OtherResourceOutputTypeOutput {
	return o
}

func (o OtherResourceOutputTypeOutput) Foo() khulnasoft.StringPtrOutput {
	return o.ApplyT(func(v OtherResourceOutputType) *string { return v.Foo }).(khulnasoft.StringPtrOutput)
}

type SomeOtherObject struct {
	Baz *string `khulnasoft:"baz"`
}

// SomeOtherObjectInput is an input type that accepts SomeOtherObjectArgs and SomeOtherObjectOutput values.
// You can construct a concrete instance of `SomeOtherObjectInput` via:
//
//	SomeOtherObjectArgs{...}
type SomeOtherObjectInput interface {
	khulnasoft.Input

	ToSomeOtherObjectOutput() SomeOtherObjectOutput
	ToSomeOtherObjectOutputWithContext(context.Context) SomeOtherObjectOutput
}

type SomeOtherObjectArgs struct {
	Baz khulnasoft.StringPtrInput `khulnasoft:"baz"`
}

func (SomeOtherObjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SomeOtherObject)(nil)).Elem()
}

func (i SomeOtherObjectArgs) ToSomeOtherObjectOutput() SomeOtherObjectOutput {
	return i.ToSomeOtherObjectOutputWithContext(context.Background())
}

func (i SomeOtherObjectArgs) ToSomeOtherObjectOutputWithContext(ctx context.Context) SomeOtherObjectOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(SomeOtherObjectOutput)
}

func (i SomeOtherObjectArgs) ToSomeOtherObjectPtrOutput() SomeOtherObjectPtrOutput {
	return i.ToSomeOtherObjectPtrOutputWithContext(context.Background())
}

func (i SomeOtherObjectArgs) ToSomeOtherObjectPtrOutputWithContext(ctx context.Context) SomeOtherObjectPtrOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(SomeOtherObjectOutput).ToSomeOtherObjectPtrOutputWithContext(ctx)
}

// SomeOtherObjectPtrInput is an input type that accepts SomeOtherObjectArgs, SomeOtherObjectPtr and SomeOtherObjectPtrOutput values.
// You can construct a concrete instance of `SomeOtherObjectPtrInput` via:
//
//	        SomeOtherObjectArgs{...}
//
//	or:
//
//	        nil
type SomeOtherObjectPtrInput interface {
	khulnasoft.Input

	ToSomeOtherObjectPtrOutput() SomeOtherObjectPtrOutput
	ToSomeOtherObjectPtrOutputWithContext(context.Context) SomeOtherObjectPtrOutput
}

type someOtherObjectPtrType SomeOtherObjectArgs

func SomeOtherObjectPtr(v *SomeOtherObjectArgs) SomeOtherObjectPtrInput {
	return (*someOtherObjectPtrType)(v)
}

func (*someOtherObjectPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SomeOtherObject)(nil)).Elem()
}

func (i *someOtherObjectPtrType) ToSomeOtherObjectPtrOutput() SomeOtherObjectPtrOutput {
	return i.ToSomeOtherObjectPtrOutputWithContext(context.Background())
}

func (i *someOtherObjectPtrType) ToSomeOtherObjectPtrOutputWithContext(ctx context.Context) SomeOtherObjectPtrOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(SomeOtherObjectPtrOutput)
}

// SomeOtherObjectArrayInput is an input type that accepts SomeOtherObjectArray and SomeOtherObjectArrayOutput values.
// You can construct a concrete instance of `SomeOtherObjectArrayInput` via:
//
//	SomeOtherObjectArray{ SomeOtherObjectArgs{...} }
type SomeOtherObjectArrayInput interface {
	khulnasoft.Input

	ToSomeOtherObjectArrayOutput() SomeOtherObjectArrayOutput
	ToSomeOtherObjectArrayOutputWithContext(context.Context) SomeOtherObjectArrayOutput
}

type SomeOtherObjectArray []SomeOtherObjectInput

func (SomeOtherObjectArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SomeOtherObject)(nil)).Elem()
}

func (i SomeOtherObjectArray) ToSomeOtherObjectArrayOutput() SomeOtherObjectArrayOutput {
	return i.ToSomeOtherObjectArrayOutputWithContext(context.Background())
}

func (i SomeOtherObjectArray) ToSomeOtherObjectArrayOutputWithContext(ctx context.Context) SomeOtherObjectArrayOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(SomeOtherObjectArrayOutput)
}

type SomeOtherObjectOutput struct{ *khulnasoft.OutputState }

func (SomeOtherObjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SomeOtherObject)(nil)).Elem()
}

func (o SomeOtherObjectOutput) ToSomeOtherObjectOutput() SomeOtherObjectOutput {
	return o
}

func (o SomeOtherObjectOutput) ToSomeOtherObjectOutputWithContext(ctx context.Context) SomeOtherObjectOutput {
	return o
}

func (o SomeOtherObjectOutput) ToSomeOtherObjectPtrOutput() SomeOtherObjectPtrOutput {
	return o.ToSomeOtherObjectPtrOutputWithContext(context.Background())
}

func (o SomeOtherObjectOutput) ToSomeOtherObjectPtrOutputWithContext(ctx context.Context) SomeOtherObjectPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SomeOtherObject) *SomeOtherObject {
		return &v
	}).(SomeOtherObjectPtrOutput)
}

func (o SomeOtherObjectOutput) Baz() khulnasoft.StringPtrOutput {
	return o.ApplyT(func(v SomeOtherObject) *string { return v.Baz }).(khulnasoft.StringPtrOutput)
}

type SomeOtherObjectPtrOutput struct{ *khulnasoft.OutputState }

func (SomeOtherObjectPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SomeOtherObject)(nil)).Elem()
}

func (o SomeOtherObjectPtrOutput) ToSomeOtherObjectPtrOutput() SomeOtherObjectPtrOutput {
	return o
}

func (o SomeOtherObjectPtrOutput) ToSomeOtherObjectPtrOutputWithContext(ctx context.Context) SomeOtherObjectPtrOutput {
	return o
}

func (o SomeOtherObjectPtrOutput) Elem() SomeOtherObjectOutput {
	return o.ApplyT(func(v *SomeOtherObject) SomeOtherObject {
		if v != nil {
			return *v
		}
		var ret SomeOtherObject
		return ret
	}).(SomeOtherObjectOutput)
}

func (o SomeOtherObjectPtrOutput) Baz() khulnasoft.StringPtrOutput {
	return o.ApplyT(func(v *SomeOtherObject) *string {
		if v == nil {
			return nil
		}
		return v.Baz
	}).(khulnasoft.StringPtrOutput)
}

type SomeOtherObjectArrayOutput struct{ *khulnasoft.OutputState }

func (SomeOtherObjectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SomeOtherObject)(nil)).Elem()
}

func (o SomeOtherObjectArrayOutput) ToSomeOtherObjectArrayOutput() SomeOtherObjectArrayOutput {
	return o
}

func (o SomeOtherObjectArrayOutput) ToSomeOtherObjectArrayOutputWithContext(ctx context.Context) SomeOtherObjectArrayOutput {
	return o
}

func (o SomeOtherObjectArrayOutput) Index(i khulnasoft.IntInput) SomeOtherObjectOutput {
	return khulnasoft.All(o, i).ApplyT(func(vs []interface{}) SomeOtherObject {
		return vs[0].([]SomeOtherObject)[vs[1].(int)]
	}).(SomeOtherObjectOutput)
}

type SomeOtherObjectArrayArray []SomeOtherObjectArrayInput

func (SomeOtherObjectArrayArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[][]SomeOtherObject)(nil)).Elem()
}

func (i SomeOtherObjectArrayArray) ToSomeOtherObjectArrayArrayOutput() SomeOtherObjectArrayArrayOutput {
	return i.ToSomeOtherObjectArrayArrayOutputWithContext(context.Background())
}

func (i SomeOtherObjectArrayArray) ToSomeOtherObjectArrayArrayOutputWithContext(ctx context.Context) SomeOtherObjectArrayArrayOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(SomeOtherObjectArrayArrayOutput)
}

// SomeOtherObjectArrayArrayInput is an input type that accepts SomeOtherObjectArrayArray and SomeOtherObjectArrayArrayOutput values.
// You can construct a concrete instance of `SomeOtherObjectArrayArrayInput` via:
//
//	SomeOtherObjectArrayArray{ SomeOtherObjectArray{ SomeOtherObjectArgs{...} } }
type SomeOtherObjectArrayArrayInput interface {
	khulnasoft.Input

	ToSomeOtherObjectArrayArrayOutput() SomeOtherObjectArrayArrayOutput
	ToSomeOtherObjectArrayArrayOutputWithContext(context.Context) SomeOtherObjectArrayArrayOutput
}

type SomeOtherObjectArrayArrayOutput struct{ *khulnasoft.OutputState }

func (SomeOtherObjectArrayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[][]SomeOtherObject)(nil)).Elem()
}

func (o SomeOtherObjectArrayArrayOutput) ToSomeOtherObjectArrayArrayOutput() SomeOtherObjectArrayArrayOutput {
	return o
}

func (o SomeOtherObjectArrayArrayOutput) ToSomeOtherObjectArrayArrayOutputWithContext(ctx context.Context) SomeOtherObjectArrayArrayOutput {
	return o
}

func (o SomeOtherObjectArrayArrayOutput) Index(i khulnasoft.IntInput) SomeOtherObjectArrayOutput {
	return khulnasoft.All(o, i).ApplyT(func(vs []interface{}) []SomeOtherObject {
		return vs[0].([][]SomeOtherObject)[vs[1].(int)]
	}).(SomeOtherObjectArrayOutput)
}

type SomeOtherObjectArrayMap map[string]SomeOtherObjectArrayInput

func (SomeOtherObjectArrayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string][]SomeOtherObject)(nil)).Elem()
}

func (i SomeOtherObjectArrayMap) ToSomeOtherObjectArrayMapOutput() SomeOtherObjectArrayMapOutput {
	return i.ToSomeOtherObjectArrayMapOutputWithContext(context.Background())
}

func (i SomeOtherObjectArrayMap) ToSomeOtherObjectArrayMapOutputWithContext(ctx context.Context) SomeOtherObjectArrayMapOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(SomeOtherObjectArrayMapOutput)
}

// SomeOtherObjectArrayMapInput is an input type that accepts SomeOtherObjectArrayMap and SomeOtherObjectArrayMapOutput values.
// You can construct a concrete instance of `SomeOtherObjectArrayMapInput` via:
//
//	SomeOtherObjectArrayMap{ "key": SomeOtherObjectArray{ SomeOtherObjectArgs{...} } }
type SomeOtherObjectArrayMapInput interface {
	khulnasoft.Input

	ToSomeOtherObjectArrayMapOutput() SomeOtherObjectArrayMapOutput
	ToSomeOtherObjectArrayMapOutputWithContext(context.Context) SomeOtherObjectArrayMapOutput
}

type SomeOtherObjectArrayMapOutput struct{ *khulnasoft.OutputState }

func (SomeOtherObjectArrayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string][]SomeOtherObject)(nil)).Elem()
}

func (o SomeOtherObjectArrayMapOutput) ToSomeOtherObjectArrayMapOutput() SomeOtherObjectArrayMapOutput {
	return o
}

func (o SomeOtherObjectArrayMapOutput) ToSomeOtherObjectArrayMapOutputWithContext(ctx context.Context) SomeOtherObjectArrayMapOutput {
	return o
}

func (o SomeOtherObjectArrayMapOutput) MapIndex(k khulnasoft.StringInput) SomeOtherObjectArrayOutput {
	return khulnasoft.All(o, k).ApplyT(func(vs []interface{}) []SomeOtherObject {
		return vs[0].(map[string][]SomeOtherObject)[vs[1].(string)]
	}).(SomeOtherObjectArrayOutput)
}

func init() {
	khulnasoft.RegisterInputType(reflect.TypeOf((*ConfigMapInput)(nil)).Elem(), ConfigMapArgs{})
	khulnasoft.RegisterInputType(reflect.TypeOf((*ConfigMapArrayInput)(nil)).Elem(), ConfigMapArray{})
	khulnasoft.RegisterInputType(reflect.TypeOf((*ObjectInput)(nil)).Elem(), ObjectArgs{})
	khulnasoft.RegisterInputType(reflect.TypeOf((*ObjectPtrInput)(nil)).Elem(), ObjectArgs{})
	khulnasoft.RegisterInputType(reflect.TypeOf((*ObjectWithNodeOptionalInputsInput)(nil)).Elem(), ObjectWithNodeOptionalInputsArgs{})
	khulnasoft.RegisterInputType(reflect.TypeOf((*ObjectWithNodeOptionalInputsPtrInput)(nil)).Elem(), ObjectWithNodeOptionalInputsArgs{})
	khulnasoft.RegisterInputType(reflect.TypeOf((*OtherResourceOutputTypeInput)(nil)).Elem(), OtherResourceOutputTypeArgs{})
	khulnasoft.RegisterInputType(reflect.TypeOf((*SomeOtherObjectInput)(nil)).Elem(), SomeOtherObjectArgs{})
	khulnasoft.RegisterInputType(reflect.TypeOf((*SomeOtherObjectPtrInput)(nil)).Elem(), SomeOtherObjectArgs{})
	khulnasoft.RegisterInputType(reflect.TypeOf((*SomeOtherObjectArrayInput)(nil)).Elem(), SomeOtherObjectArray{})
	khulnasoft.RegisterInputType(reflect.TypeOf((*SomeOtherObjectArrayArrayInput)(nil)).Elem(), SomeOtherObjectArrayArray{})
	khulnasoft.RegisterInputType(reflect.TypeOf((*SomeOtherObjectArrayMapInput)(nil)).Elem(), SomeOtherObjectArrayMap{})
	khulnasoft.RegisterOutputType(ConfigMapOutput{})
	khulnasoft.RegisterOutputType(ConfigMapArrayOutput{})
	khulnasoft.RegisterOutputType(ObjectOutput{})
	khulnasoft.RegisterOutputType(ObjectPtrOutput{})
	khulnasoft.RegisterOutputType(ObjectWithNodeOptionalInputsOutput{})
	khulnasoft.RegisterOutputType(ObjectWithNodeOptionalInputsPtrOutput{})
	khulnasoft.RegisterOutputType(OtherResourceOutputTypeOutput{})
	khulnasoft.RegisterOutputType(SomeOtherObjectOutput{})
	khulnasoft.RegisterOutputType(SomeOtherObjectPtrOutput{})
	khulnasoft.RegisterOutputType(SomeOtherObjectArrayOutput{})
	khulnasoft.RegisterOutputType(SomeOtherObjectArrayArrayOutput{})
	khulnasoft.RegisterOutputType(SomeOtherObjectArrayMapOutput{})
}