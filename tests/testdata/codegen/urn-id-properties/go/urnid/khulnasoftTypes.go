// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package urnid

import (
	"context"
	"reflect"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"urn-id-properties/urnid/internal"
)

var _ = internal.GetEnvOrDefault

// It's fine to use urn and id in nested objects
type InnerType struct {
	Id  *string `khulnasoft:"id"`
	Urn *string `khulnasoft:"urn"`
}

// It's fine to use urn and id in nested objects
type InnerTypeOutput struct{ *khulnasoft.OutputState }

func (InnerTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InnerType)(nil)).Elem()
}

func (o InnerTypeOutput) ToInnerTypeOutput() InnerTypeOutput {
	return o
}

func (o InnerTypeOutput) ToInnerTypeOutputWithContext(ctx context.Context) InnerTypeOutput {
	return o
}

func (o InnerTypeOutput) Id() khulnasoft.StringPtrOutput {
	return o.ApplyT(func(v InnerType) *string { return v.Id }).(khulnasoft.StringPtrOutput)
}

func (o InnerTypeOutput) Urn() khulnasoft.StringPtrOutput {
	return o.ApplyT(func(v InnerType) *string { return v.Urn }).(khulnasoft.StringPtrOutput)
}

type InnerTypePtrOutput struct{ *khulnasoft.OutputState }

func (InnerTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InnerType)(nil)).Elem()
}

func (o InnerTypePtrOutput) ToInnerTypePtrOutput() InnerTypePtrOutput {
	return o
}

func (o InnerTypePtrOutput) ToInnerTypePtrOutputWithContext(ctx context.Context) InnerTypePtrOutput {
	return o
}

func (o InnerTypePtrOutput) Elem() InnerTypeOutput {
	return o.ApplyT(func(v *InnerType) InnerType {
		if v != nil {
			return *v
		}
		var ret InnerType
		return ret
	}).(InnerTypeOutput)
}

func (o InnerTypePtrOutput) Id() khulnasoft.StringPtrOutput {
	return o.ApplyT(func(v *InnerType) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(khulnasoft.StringPtrOutput)
}

func (o InnerTypePtrOutput) Urn() khulnasoft.StringPtrOutput {
	return o.ApplyT(func(v *InnerType) *string {
		if v == nil {
			return nil
		}
		return v.Urn
	}).(khulnasoft.StringPtrOutput)
}

func init() {
	khulnasoft.RegisterOutputType(InnerTypeOutput{})
	khulnasoft.RegisterOutputType(InnerTypePtrOutput{})
}