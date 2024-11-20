// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package legacy_names

import (
	"context"
	"reflect"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"legacy-names/legacy_names/http_module"
	"legacy-names/legacy_names/internal"
)

type Example_resource struct {
	khulnasoft.CustomResourceState

	URL       khulnasoft.StringPtrOutput   `khulnasoft:"URL"`
	Good_URLs khulnasoft.StringArrayOutput `khulnasoft:"good_URLs"`
	Map_enum  Enum_XYZMapArrayOutput   `khulnasoft:"map_enum"`
}

// NewExample_resource registers a new resource with the given unique name, arguments, and options.
func NewExample_resource(ctx *khulnasoft.Context,
	name string, args *Example_resourceArgs, opts ...khulnasoft.ResourceOption) (*Example_resource, error) {
	if args == nil {
		args = &Example_resourceArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Example_resource
	err := ctx.RegisterResource("legacy_names:index:example_resource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExample_resource gets an existing Example_resource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExample_resource(ctx *khulnasoft.Context,
	name string, id khulnasoft.IDInput, state *Example_resourceState, opts ...khulnasoft.ResourceOption) (*Example_resource, error) {
	var resource Example_resource
	err := ctx.ReadResource("legacy_names:index:example_resource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Example_resource resources.
type example_resourceState struct {
}

type Example_resourceState struct {
}

func (Example_resourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*example_resourceState)(nil)).Elem()
}

type example_resourceArgs struct {
	Map_enum     []map[string]Enum_XYZ `khulnasoft:"map_enum"`
	Request_HTTP *http_module.Request  `khulnasoft:"request_HTTP"`
}

// The set of arguments for constructing a Example_resource resource.
type Example_resourceArgs struct {
	Map_enum     Enum_XYZMapArrayInput
	Request_HTTP http_module.RequestPtrInput
}

func (Example_resourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*example_resourceArgs)(nil)).Elem()
}

type Example_resourceInput interface {
	khulnasoft.Input

	ToExample_resourceOutput() Example_resourceOutput
	ToExample_resourceOutputWithContext(ctx context.Context) Example_resourceOutput
}

func (*Example_resource) ElementType() reflect.Type {
	return reflect.TypeOf((**Example_resource)(nil)).Elem()
}

func (i *Example_resource) ToExample_resourceOutput() Example_resourceOutput {
	return i.ToExample_resourceOutputWithContext(context.Background())
}

func (i *Example_resource) ToExample_resourceOutputWithContext(ctx context.Context) Example_resourceOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(Example_resourceOutput)
}

type Example_resourceOutput struct{ *khulnasoft.OutputState }

func (Example_resourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Example_resource)(nil)).Elem()
}

func (o Example_resourceOutput) ToExample_resourceOutput() Example_resourceOutput {
	return o
}

func (o Example_resourceOutput) ToExample_resourceOutputWithContext(ctx context.Context) Example_resourceOutput {
	return o
}

func (o Example_resourceOutput) URL() khulnasoft.StringPtrOutput {
	return o.ApplyT(func(v *Example_resource) khulnasoft.StringPtrOutput { return v.URL }).(khulnasoft.StringPtrOutput)
}

func (o Example_resourceOutput) Good_URLs() khulnasoft.StringArrayOutput {
	return o.ApplyT(func(v *Example_resource) khulnasoft.StringArrayOutput { return v.Good_URLs }).(khulnasoft.StringArrayOutput)
}

func (o Example_resourceOutput) Map_enum() Enum_XYZMapArrayOutput {
	return o.ApplyT(func(v *Example_resource) Enum_XYZMapArrayOutput { return v.Map_enum }).(Enum_XYZMapArrayOutput)
}

func init() {
	khulnasoft.RegisterInputType(reflect.TypeOf((*Example_resourceInput)(nil)).Elem(), &Example_resource{})
	khulnasoft.RegisterOutputType(Example_resourceOutput{})
}