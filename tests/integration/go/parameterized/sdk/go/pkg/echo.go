// Code generated by khulnasoft-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pkg

import (
	"context"
	"reflect"

	"example.com/khulnasoft-pkg/sdk/go/pkg/internal"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

// A test resource that echoes its input.
type Echo struct {
	khulnasoft.CustomResourceState

	// Input to echo.
	Echo khulnasoft.AnyOutput `khulnasoft:"echo"`
}

// NewEcho registers a new resource with the given unique name, arguments, and options.
func NewEcho(ctx *khulnasoft.Context,
	name string, args *EchoArgs, opts ...khulnasoft.ResourceOption) (*Echo, error) {
	if args == nil {
		args = &EchoArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource Echo
	err = ctx.RegisterPackageResource("pkg:index:Echo", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEcho gets an existing Echo resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEcho(ctx *khulnasoft.Context,
	name string, id khulnasoft.IDInput, state *EchoState, opts ...khulnasoft.ResourceOption) (*Echo, error) {
	var resource Echo
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("pkg:index:Echo", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Echo resources.
type echoState struct {
}

type EchoState struct {
}

func (EchoState) ElementType() reflect.Type {
	return reflect.TypeOf((*echoState)(nil)).Elem()
}

type echoArgs struct {
	// An echoed input.
	Echo interface{} `khulnasoft:"echo"`
}

// The set of arguments for constructing a Echo resource.
type EchoArgs struct {
	// An echoed input.
	Echo khulnasoft.Input
}

func (EchoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*echoArgs)(nil)).Elem()
}

// A test call that echoes its input.
func (r *Echo) DoEchoMethod(ctx *khulnasoft.Context, args *EchoDoEchoMethodArgs) (EchoDoEchoMethodResultOutput, error) {
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return EchoDoEchoMethodResultOutput{}, err
	}
	out, err := ctx.CallPackage("pkg:index:Echo/doEchoMethod", args, EchoDoEchoMethodResultOutput{}, r, ref)
	if err != nil {
		return EchoDoEchoMethodResultOutput{}, err
	}
	return out.(EchoDoEchoMethodResultOutput), nil
}

type echoDoEchoMethodArgs struct {
	Echo *string `khulnasoft:"echo"`
}

// The set of arguments for the DoEchoMethod method of the Echo resource.
type EchoDoEchoMethodArgs struct {
	Echo khulnasoft.StringPtrInput
}

func (EchoDoEchoMethodArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*echoDoEchoMethodArgs)(nil)).Elem()
}

type EchoDoEchoMethodResult struct {
	Echo *string `khulnasoft:"echo"`
}

type EchoDoEchoMethodResultOutput struct{ *khulnasoft.OutputState }

func (EchoDoEchoMethodResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EchoDoEchoMethodResult)(nil)).Elem()
}

func (o EchoDoEchoMethodResultOutput) Echo() khulnasoft.StringPtrOutput {
	return o.ApplyT(func(v EchoDoEchoMethodResult) *string { return v.Echo }).(khulnasoft.StringPtrOutput)
}

type EchoInput interface {
	khulnasoft.Input

	ToEchoOutput() EchoOutput
	ToEchoOutputWithContext(ctx context.Context) EchoOutput
}

func (*Echo) ElementType() reflect.Type {
	return reflect.TypeOf((**Echo)(nil)).Elem()
}

func (i *Echo) ToEchoOutput() EchoOutput {
	return i.ToEchoOutputWithContext(context.Background())
}

func (i *Echo) ToEchoOutputWithContext(ctx context.Context) EchoOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(EchoOutput)
}

type EchoOutput struct{ *khulnasoft.OutputState }

func (EchoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Echo)(nil)).Elem()
}

func (o EchoOutput) ToEchoOutput() EchoOutput {
	return o
}

func (o EchoOutput) ToEchoOutputWithContext(ctx context.Context) EchoOutput {
	return o
}

// Input to echo.
func (o EchoOutput) Echo() khulnasoft.AnyOutput {
	return o.ApplyT(func(v *Echo) khulnasoft.AnyOutput { return v.Echo }).(khulnasoft.AnyOutput)
}

func init() {
	khulnasoft.RegisterInputType(reflect.TypeOf((*EchoInput)(nil)).Elem(), &Echo{})
	khulnasoft.RegisterOutputType(EchoOutput{})
	khulnasoft.RegisterOutputType(EchoDoEchoMethodResultOutput{})
}
