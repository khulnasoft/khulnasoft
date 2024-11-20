// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package credentials

import (
	"context"
	"reflect"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"using-shared-types-in-config/credentials/internal"
)

var _ = internal.GetEnvOrDefault

type Shared struct {
	Foo *string `khulnasoft:"foo"`
}

// SharedInput is an input type that accepts SharedArgs and SharedOutput values.
// You can construct a concrete instance of `SharedInput` via:
//
//	SharedArgs{...}
type SharedInput interface {
	khulnasoft.Input

	ToSharedOutput() SharedOutput
	ToSharedOutputWithContext(context.Context) SharedOutput
}

type SharedArgs struct {
	Foo khulnasoft.StringPtrInput `khulnasoft:"foo"`
}

func (SharedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Shared)(nil)).Elem()
}

func (i SharedArgs) ToSharedOutput() SharedOutput {
	return i.ToSharedOutputWithContext(context.Background())
}

func (i SharedArgs) ToSharedOutputWithContext(ctx context.Context) SharedOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(SharedOutput)
}

func (i SharedArgs) ToSharedPtrOutput() SharedPtrOutput {
	return i.ToSharedPtrOutputWithContext(context.Background())
}

func (i SharedArgs) ToSharedPtrOutputWithContext(ctx context.Context) SharedPtrOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(SharedOutput).ToSharedPtrOutputWithContext(ctx)
}

// SharedPtrInput is an input type that accepts SharedArgs, SharedPtr and SharedPtrOutput values.
// You can construct a concrete instance of `SharedPtrInput` via:
//
//	        SharedArgs{...}
//
//	or:
//
//	        nil
type SharedPtrInput interface {
	khulnasoft.Input

	ToSharedPtrOutput() SharedPtrOutput
	ToSharedPtrOutputWithContext(context.Context) SharedPtrOutput
}

type sharedPtrType SharedArgs

func SharedPtr(v *SharedArgs) SharedPtrInput {
	return (*sharedPtrType)(v)
}

func (*sharedPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Shared)(nil)).Elem()
}

func (i *sharedPtrType) ToSharedPtrOutput() SharedPtrOutput {
	return i.ToSharedPtrOutputWithContext(context.Background())
}

func (i *sharedPtrType) ToSharedPtrOutputWithContext(ctx context.Context) SharedPtrOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(SharedPtrOutput)
}

type SharedOutput struct{ *khulnasoft.OutputState }

func (SharedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Shared)(nil)).Elem()
}

func (o SharedOutput) ToSharedOutput() SharedOutput {
	return o
}

func (o SharedOutput) ToSharedOutputWithContext(ctx context.Context) SharedOutput {
	return o
}

func (o SharedOutput) ToSharedPtrOutput() SharedPtrOutput {
	return o.ToSharedPtrOutputWithContext(context.Background())
}

func (o SharedOutput) ToSharedPtrOutputWithContext(ctx context.Context) SharedPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Shared) *Shared {
		return &v
	}).(SharedPtrOutput)
}

func (o SharedOutput) Foo() khulnasoft.StringPtrOutput {
	return o.ApplyT(func(v Shared) *string { return v.Foo }).(khulnasoft.StringPtrOutput)
}

type SharedPtrOutput struct{ *khulnasoft.OutputState }

func (SharedPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Shared)(nil)).Elem()
}

func (o SharedPtrOutput) ToSharedPtrOutput() SharedPtrOutput {
	return o
}

func (o SharedPtrOutput) ToSharedPtrOutputWithContext(ctx context.Context) SharedPtrOutput {
	return o
}

func (o SharedPtrOutput) Elem() SharedOutput {
	return o.ApplyT(func(v *Shared) Shared {
		if v != nil {
			return *v
		}
		var ret Shared
		return ret
	}).(SharedOutput)
}

func (o SharedPtrOutput) Foo() khulnasoft.StringPtrOutput {
	return o.ApplyT(func(v *Shared) *string {
		if v == nil {
			return nil
		}
		return v.Foo
	}).(khulnasoft.StringPtrOutput)
}

func init() {
	khulnasoft.RegisterInputType(reflect.TypeOf((*SharedInput)(nil)).Elem(), SharedArgs{})
	khulnasoft.RegisterInputType(reflect.TypeOf((*SharedPtrInput)(nil)).Elem(), SharedArgs{})
	khulnasoft.RegisterOutputType(SharedOutput{})
	khulnasoft.RegisterOutputType(SharedPtrOutput{})
}
