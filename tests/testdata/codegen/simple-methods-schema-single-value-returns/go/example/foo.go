// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"context"
	"reflect"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"simple-methods-schema-single-value-returns/example/internal"
)

type Foo struct {
	khulnasoft.ResourceState
}

// NewFoo registers a new resource with the given unique name, arguments, and options.
func NewFoo(ctx *khulnasoft.Context,
	name string, args *FooArgs, opts ...khulnasoft.ResourceOption) (*Foo, error) {
	if args == nil {
		args = &FooArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Foo
	err := ctx.RegisterRemoteComponentResource("example::Foo", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type fooArgs struct {
}

// The set of arguments for constructing a Foo resource.
type FooArgs struct {
}

func (FooArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fooArgs)(nil)).Elem()
}

func (r *Foo) GetKubeconfig(ctx *khulnasoft.Context, args *FooGetKubeconfigArgs) (khulnasoft.StringOutput, error) {
	out, err := ctx.Call("example::Foo/getKubeconfig", args, fooGetKubeconfigResultOutput{}, r)
	if err != nil {
		return khulnasoft.StringOutput{}, err
	}
	return out.(fooGetKubeconfigResultOutput).Kubeconfig(), nil
}

type fooGetKubeconfigArgs struct {
	ProfileName *string `khulnasoft:"profileName"`
	RoleArn     *string `khulnasoft:"roleArn"`
}

// The set of arguments for the GetKubeconfig method of the Foo resource.
type FooGetKubeconfigArgs struct {
	ProfileName khulnasoft.StringPtrInput
	RoleArn     khulnasoft.StringPtrInput
}

func (FooGetKubeconfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fooGetKubeconfigArgs)(nil)).Elem()
}

type fooGetKubeconfigResult struct {
	Kubeconfig string `khulnasoft:"kubeconfig"`
}

type fooGetKubeconfigResultOutput struct{ *khulnasoft.OutputState }

func (fooGetKubeconfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*fooGetKubeconfigResult)(nil)).Elem()
}

func (o fooGetKubeconfigResultOutput) Kubeconfig() khulnasoft.StringOutput {
	return o.ApplyT(func(v fooGetKubeconfigResult) string { return v.Kubeconfig }).(khulnasoft.StringOutput)
}

type FooInput interface {
	khulnasoft.Input

	ToFooOutput() FooOutput
	ToFooOutputWithContext(ctx context.Context) FooOutput
}

func (*Foo) ElementType() reflect.Type {
	return reflect.TypeOf((**Foo)(nil)).Elem()
}

func (i *Foo) ToFooOutput() FooOutput {
	return i.ToFooOutputWithContext(context.Background())
}

func (i *Foo) ToFooOutputWithContext(ctx context.Context) FooOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(FooOutput)
}

type FooOutput struct{ *khulnasoft.OutputState }

func (FooOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Foo)(nil)).Elem()
}

func (o FooOutput) ToFooOutput() FooOutput {
	return o
}

func (o FooOutput) ToFooOutputWithContext(ctx context.Context) FooOutput {
	return o
}

func init() {
	khulnasoft.RegisterInputType(reflect.TypeOf((*FooInput)(nil)).Elem(), &Foo{})
	khulnasoft.RegisterOutputType(FooOutput{})
	khulnasoft.RegisterOutputType(fooGetKubeconfigResultOutput{})
}