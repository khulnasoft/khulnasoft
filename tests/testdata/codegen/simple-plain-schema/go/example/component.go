// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"context"
	"reflect"

	"errors"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoftx"
	"simple-plain-schema/example/internal"
)

type Component struct {
	khulnasoft.ResourceState

	A   khulnasoft.BoolOutput      `khulnasoft:"a"`
	B   khulnasoft.BoolPtrOutput   `khulnasoft:"b"`
	Bar FooPtrOutput           `khulnasoft:"bar"`
	Baz FooArrayOutput         `khulnasoft:"baz"`
	C   khulnasoft.IntOutput       `khulnasoft:"c"`
	D   khulnasoft.IntPtrOutput    `khulnasoft:"d"`
	E   khulnasoft.StringOutput    `khulnasoft:"e"`
	F   khulnasoft.StringPtrOutput `khulnasoft:"f"`
	Foo FooPtrOutput           `khulnasoft:"foo"`
}

// NewComponent registers a new resource with the given unique name, arguments, and options.
func NewComponent(ctx *khulnasoft.Context,
	name string, args *ComponentArgs, opts ...khulnasoft.ResourceOption) (*Component, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Component
	err := ctx.RegisterRemoteComponentResource("example::Component", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type componentArgs struct {
	A      bool           `khulnasoft:"a"`
	B      *bool          `khulnasoft:"b"`
	Bar    *Foo           `khulnasoft:"bar"`
	Baz    []Foo          `khulnasoft:"baz"`
	BazMap map[string]Foo `khulnasoft:"bazMap"`
	C      int            `khulnasoft:"c"`
	D      *int           `khulnasoft:"d"`
	E      string         `khulnasoft:"e"`
	F      *string        `khulnasoft:"f"`
	Foo    *Foo           `khulnasoft:"foo"`
}

// The set of arguments for constructing a Component resource.
type ComponentArgs struct {
	A      bool
	B      *bool
	Bar    *FooArgs
	Baz    []FooInput
	BazMap map[string]FooInput
	C      int
	D      *int
	E      string
	F      *string
	Foo    FooPtrInput
}

func (ComponentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentArgs)(nil)).Elem()
}

type ComponentInput interface {
	khulnasoft.Input

	ToComponentOutput() ComponentOutput
	ToComponentOutputWithContext(ctx context.Context) ComponentOutput
}

func (*Component) ElementType() reflect.Type {
	return reflect.TypeOf((**Component)(nil)).Elem()
}

func (i *Component) ToComponentOutput() ComponentOutput {
	return i.ToComponentOutputWithContext(context.Background())
}

func (i *Component) ToComponentOutputWithContext(ctx context.Context) ComponentOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(ComponentOutput)
}

func (i *Component) ToOutput(ctx context.Context) khulnasoftx.Output[*Component] {
	return khulnasoftx.Output[*Component]{
		OutputState: i.ToComponentOutputWithContext(ctx).OutputState,
	}
}

type ComponentOutput struct{ *khulnasoft.OutputState }

func (ComponentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Component)(nil)).Elem()
}

func (o ComponentOutput) ToComponentOutput() ComponentOutput {
	return o
}

func (o ComponentOutput) ToComponentOutputWithContext(ctx context.Context) ComponentOutput {
	return o
}

func (o ComponentOutput) ToOutput(ctx context.Context) khulnasoftx.Output[*Component] {
	return khulnasoftx.Output[*Component]{
		OutputState: o.OutputState,
	}
}

func (o ComponentOutput) A() khulnasoft.BoolOutput {
	return o.ApplyT(func(v *Component) khulnasoft.BoolOutput { return v.A }).(khulnasoft.BoolOutput)
}

func (o ComponentOutput) B() khulnasoft.BoolPtrOutput {
	return o.ApplyT(func(v *Component) khulnasoft.BoolPtrOutput { return v.B }).(khulnasoft.BoolPtrOutput)
}

func (o ComponentOutput) Bar() FooPtrOutput {
	return o.ApplyT(func(v *Component) FooPtrOutput { return v.Bar }).(FooPtrOutput)
}

func (o ComponentOutput) Baz() FooArrayOutput {
	return o.ApplyT(func(v *Component) FooArrayOutput { return v.Baz }).(FooArrayOutput)
}

func (o ComponentOutput) C() khulnasoft.IntOutput {
	return o.ApplyT(func(v *Component) khulnasoft.IntOutput { return v.C }).(khulnasoft.IntOutput)
}

func (o ComponentOutput) D() khulnasoft.IntPtrOutput {
	return o.ApplyT(func(v *Component) khulnasoft.IntPtrOutput { return v.D }).(khulnasoft.IntPtrOutput)
}

func (o ComponentOutput) E() khulnasoft.StringOutput {
	return o.ApplyT(func(v *Component) khulnasoft.StringOutput { return v.E }).(khulnasoft.StringOutput)
}

func (o ComponentOutput) F() khulnasoft.StringPtrOutput {
	return o.ApplyT(func(v *Component) khulnasoft.StringPtrOutput { return v.F }).(khulnasoft.StringPtrOutput)
}

func (o ComponentOutput) Foo() FooPtrOutput {
	return o.ApplyT(func(v *Component) FooPtrOutput { return v.Foo }).(FooPtrOutput)
}

func init() {
	khulnasoft.RegisterInputType(reflect.TypeOf((*ComponentInput)(nil)).Elem(), &Component{})
	khulnasoft.RegisterOutputType(ComponentOutput{})
}