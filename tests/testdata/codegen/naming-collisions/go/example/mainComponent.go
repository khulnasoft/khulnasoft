// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"context"
	"reflect"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"naming-collisions/example/internal"
)

type MainComponent struct {
	khulnasoft.CustomResourceState
}

// NewMainComponent registers a new resource with the given unique name, arguments, and options.
func NewMainComponent(ctx *khulnasoft.Context,
	name string, args *MainComponentArgs, opts ...khulnasoft.ResourceOption) (*MainComponent, error) {
	if args == nil {
		args = &MainComponentArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MainComponent
	err := ctx.RegisterResource("example::MainComponent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMainComponent gets an existing MainComponent resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMainComponent(ctx *khulnasoft.Context,
	name string, id khulnasoft.IDInput, state *MainComponentState, opts ...khulnasoft.ResourceOption) (*MainComponent, error) {
	var resource MainComponent
	err := ctx.ReadResource("example::MainComponent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MainComponent resources.
type mainComponentState struct {
}

type MainComponentState struct {
}

func (MainComponentState) ElementType() reflect.Type {
	return reflect.TypeOf((*mainComponentState)(nil)).Elem()
}

type mainComponentArgs struct {
}

// The set of arguments for constructing a MainComponent resource.
type MainComponentArgs struct {
}

func (MainComponentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mainComponentArgs)(nil)).Elem()
}

type MainComponentInput interface {
	khulnasoft.Input

	ToMainComponentOutput() MainComponentOutput
	ToMainComponentOutputWithContext(ctx context.Context) MainComponentOutput
}

func (*MainComponent) ElementType() reflect.Type {
	return reflect.TypeOf((**MainComponent)(nil)).Elem()
}

func (i *MainComponent) ToMainComponentOutput() MainComponentOutput {
	return i.ToMainComponentOutputWithContext(context.Background())
}

func (i *MainComponent) ToMainComponentOutputWithContext(ctx context.Context) MainComponentOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(MainComponentOutput)
}

type MainComponentOutput struct{ *khulnasoft.OutputState }

func (MainComponentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MainComponent)(nil)).Elem()
}

func (o MainComponentOutput) ToMainComponentOutput() MainComponentOutput {
	return o
}

func (o MainComponentOutput) ToMainComponentOutputWithContext(ctx context.Context) MainComponentOutput {
	return o
}

func init() {
	khulnasoft.RegisterInputType(reflect.TypeOf((*MainComponentInput)(nil)).Elem(), &MainComponent{})
	khulnasoft.RegisterOutputType(MainComponentOutput{})
}