// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"context"
	"reflect"

	"errors"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"plain-object-defaults/example/internal"
)

// test new feature with resoruces
type Foo struct {
	khulnasoft.CustomResourceState

	// A test for plain types
	DefaultKubeClientSettings KubeClientSettingsPtrOutput `khulnasoft:"defaultKubeClientSettings"`
}

// NewFoo registers a new resource with the given unique name, arguments, and options.
func NewFoo(ctx *khulnasoft.Context,
	name string, args *FooArgs, opts ...khulnasoft.ResourceOption) (*Foo, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackupKubeClientSettings == nil {
		return nil, errors.New("invalid value for required argument 'BackupKubeClientSettings'")
	}
	args.BackupKubeClientSettings = args.BackupKubeClientSettings.ToKubeClientSettingsOutput().ApplyT(func(v KubeClientSettings) KubeClientSettings { return *v.Defaults() }).(KubeClientSettingsOutput)
	if args.KubeClientSettings != nil {
		args.KubeClientSettings = args.KubeClientSettings.ToKubeClientSettingsPtrOutput().ApplyT(func(v *KubeClientSettings) *KubeClientSettings { return v.Defaults() }).(KubeClientSettingsPtrOutput)
	}
	if args.Settings != nil {
		args.Settings = args.Settings.ToLayeredTypePtrOutput().ApplyT(func(v *LayeredType) *LayeredType { return v.Defaults() }).(LayeredTypePtrOutput)
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Foo
	err := ctx.RegisterResource("example:index:Foo", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFoo gets an existing Foo resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFoo(ctx *khulnasoft.Context,
	name string, id khulnasoft.IDInput, state *FooState, opts ...khulnasoft.ResourceOption) (*Foo, error) {
	var resource Foo
	err := ctx.ReadResource("example:index:Foo", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Foo resources.
type fooState struct {
}

type FooState struct {
}

func (FooState) ElementType() reflect.Type {
	return reflect.TypeOf((*fooState)(nil)).Elem()
}

type fooArgs struct {
	Argument *string `khulnasoft:"argument"`
	// Options for tuning the Kubernetes client used by a Provider.
	BackupKubeClientSettings KubeClientSettings `khulnasoft:"backupKubeClientSettings"`
	// Options for tuning the Kubernetes client used by a Provider.
	KubeClientSettings *KubeClientSettings `khulnasoft:"kubeClientSettings"`
	// describing things
	Settings *LayeredType `khulnasoft:"settings"`
}

// The set of arguments for constructing a Foo resource.
type FooArgs struct {
	Argument *string
	// Options for tuning the Kubernetes client used by a Provider.
	BackupKubeClientSettings KubeClientSettingsInput
	// Options for tuning the Kubernetes client used by a Provider.
	KubeClientSettings KubeClientSettingsPtrInput
	// describing things
	Settings LayeredTypePtrInput
}

func (FooArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fooArgs)(nil)).Elem()
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

// A test for plain types
func (o FooOutput) DefaultKubeClientSettings() KubeClientSettingsPtrOutput {
	return o.ApplyT(func(v *Foo) KubeClientSettingsPtrOutput { return v.DefaultKubeClientSettings }).(KubeClientSettingsPtrOutput)
}

func init() {
	khulnasoft.RegisterInputType(reflect.TypeOf((*FooInput)(nil)).Elem(), &Foo{})
	khulnasoft.RegisterOutputType(FooOutput{})
}