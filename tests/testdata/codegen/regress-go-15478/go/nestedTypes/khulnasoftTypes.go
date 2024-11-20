// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nestedTypes

import (
	"context"
	"reflect"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"regress-go-15478/nestedTypes/internal"
)

var _ = internal.GetEnvOrDefault

type NestedType struct {
	Data       [][][]string                            `khulnasoft:"data"`
	NestedMaps map[string]map[string]map[string]string `khulnasoft:"nestedMaps"`
}

type NestedTypeOutput struct{ *khulnasoft.OutputState }

func (NestedTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NestedType)(nil)).Elem()
}

func (o NestedTypeOutput) ToNestedTypeOutput() NestedTypeOutput {
	return o
}

func (o NestedTypeOutput) ToNestedTypeOutputWithContext(ctx context.Context) NestedTypeOutput {
	return o
}

func (o NestedTypeOutput) Data() khulnasoft.StringArrayArrayArrayOutput {
	return o.ApplyT(func(v NestedType) [][][]string { return v.Data }).(khulnasoft.StringArrayArrayArrayOutput)
}

func (o NestedTypeOutput) NestedMaps() khulnasoft.StringMapMapMapOutput {
	return o.ApplyT(func(v NestedType) map[string]map[string]map[string]string { return v.NestedMaps }).(khulnasoft.StringMapMapMapOutput)
}

type NestedTypePtrOutput struct{ *khulnasoft.OutputState }

func (NestedTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NestedType)(nil)).Elem()
}

func (o NestedTypePtrOutput) ToNestedTypePtrOutput() NestedTypePtrOutput {
	return o
}

func (o NestedTypePtrOutput) ToNestedTypePtrOutputWithContext(ctx context.Context) NestedTypePtrOutput {
	return o
}

func (o NestedTypePtrOutput) Elem() NestedTypeOutput {
	return o.ApplyT(func(v *NestedType) NestedType {
		if v != nil {
			return *v
		}
		var ret NestedType
		return ret
	}).(NestedTypeOutput)
}

func (o NestedTypePtrOutput) Data() khulnasoft.StringArrayArrayArrayOutput {
	return o.ApplyT(func(v *NestedType) [][][]string {
		if v == nil {
			return nil
		}
		return v.Data
	}).(khulnasoft.StringArrayArrayArrayOutput)
}

func (o NestedTypePtrOutput) NestedMaps() khulnasoft.StringMapMapMapOutput {
	return o.ApplyT(func(v *NestedType) map[string]map[string]map[string]string {
		if v == nil {
			return nil
		}
		return v.NestedMaps
	}).(khulnasoft.StringMapMapMapOutput)
}

func init() {
	khulnasoft.RegisterOutputType(NestedTypeOutput{})
	khulnasoft.RegisterOutputType(NestedTypePtrOutput{})
}