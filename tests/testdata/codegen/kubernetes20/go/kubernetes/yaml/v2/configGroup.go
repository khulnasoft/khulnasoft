// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/khulnasoft/khulnasoft-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

// A non-overlay component resource.
type ConfigGroup struct {
	khulnasoft.ResourceState

	// Resources created by the ConfigGroup.
	Resources khulnasoft.ArrayOutput `khulnasoft:"resources"`
}

// NewConfigGroup registers a new resource with the given unique name, arguments, and options.
func NewConfigGroup(ctx *khulnasoft.Context,
	name string, args *ConfigGroupArgs, opts ...khulnasoft.ResourceOption) (*ConfigGroup, error) {
	if args == nil {
		args = &ConfigGroupArgs{}
	}

	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ConfigGroup
	err := ctx.RegisterRemoteComponentResource("kubernetes:yaml/v2:ConfigGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type configGroupArgs struct {
	// Set of paths or a URLs that uniquely identify files.
	Files interface{} `khulnasoft:"files"`
	// Objects representing Kubernetes resources.
	Objs interface{} `khulnasoft:"objs"`
	// An optional prefix for the auto-generated resource names. Example: A resource created with resourcePrefix="foo" would produce a resource named "foo-resourceName".
	ResourcePrefix *string `khulnasoft:"resourcePrefix"`
	// YAML text containing Kubernetes resource definitions.
	Yaml interface{} `khulnasoft:"yaml"`
}

// The set of arguments for constructing a ConfigGroup resource.
type ConfigGroupArgs struct {
	// Set of paths or a URLs that uniquely identify files.
	Files khulnasoft.Input
	// Objects representing Kubernetes resources.
	Objs khulnasoft.Input
	// An optional prefix for the auto-generated resource names. Example: A resource created with resourcePrefix="foo" would produce a resource named "foo-resourceName".
	ResourcePrefix khulnasoft.StringPtrInput
	// YAML text containing Kubernetes resource definitions.
	Yaml khulnasoft.Input
}

func (ConfigGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configGroupArgs)(nil)).Elem()
}

type ConfigGroupInput interface {
	khulnasoft.Input

	ToConfigGroupOutput() ConfigGroupOutput
	ToConfigGroupOutputWithContext(ctx context.Context) ConfigGroupOutput
}

func (*ConfigGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigGroup)(nil)).Elem()
}

func (i *ConfigGroup) ToConfigGroupOutput() ConfigGroupOutput {
	return i.ToConfigGroupOutputWithContext(context.Background())
}

func (i *ConfigGroup) ToConfigGroupOutputWithContext(ctx context.Context) ConfigGroupOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(ConfigGroupOutput)
}

// ConfigGroupArrayInput is an input type that accepts ConfigGroupArray and ConfigGroupArrayOutput values.
// You can construct a concrete instance of `ConfigGroupArrayInput` via:
//
//	ConfigGroupArray{ ConfigGroupArgs{...} }
type ConfigGroupArrayInput interface {
	khulnasoft.Input

	ToConfigGroupArrayOutput() ConfigGroupArrayOutput
	ToConfigGroupArrayOutputWithContext(context.Context) ConfigGroupArrayOutput
}

type ConfigGroupArray []ConfigGroupInput

func (ConfigGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConfigGroup)(nil)).Elem()
}

func (i ConfigGroupArray) ToConfigGroupArrayOutput() ConfigGroupArrayOutput {
	return i.ToConfigGroupArrayOutputWithContext(context.Background())
}

func (i ConfigGroupArray) ToConfigGroupArrayOutputWithContext(ctx context.Context) ConfigGroupArrayOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(ConfigGroupArrayOutput)
}

// ConfigGroupMapInput is an input type that accepts ConfigGroupMap and ConfigGroupMapOutput values.
// You can construct a concrete instance of `ConfigGroupMapInput` via:
//
//	ConfigGroupMap{ "key": ConfigGroupArgs{...} }
type ConfigGroupMapInput interface {
	khulnasoft.Input

	ToConfigGroupMapOutput() ConfigGroupMapOutput
	ToConfigGroupMapOutputWithContext(context.Context) ConfigGroupMapOutput
}

type ConfigGroupMap map[string]ConfigGroupInput

func (ConfigGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConfigGroup)(nil)).Elem()
}

func (i ConfigGroupMap) ToConfigGroupMapOutput() ConfigGroupMapOutput {
	return i.ToConfigGroupMapOutputWithContext(context.Background())
}

func (i ConfigGroupMap) ToConfigGroupMapOutputWithContext(ctx context.Context) ConfigGroupMapOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(ConfigGroupMapOutput)
}

type ConfigGroupOutput struct{ *khulnasoft.OutputState }

func (ConfigGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigGroup)(nil)).Elem()
}

func (o ConfigGroupOutput) ToConfigGroupOutput() ConfigGroupOutput {
	return o
}

func (o ConfigGroupOutput) ToConfigGroupOutputWithContext(ctx context.Context) ConfigGroupOutput {
	return o
}

// Resources created by the ConfigGroup.
func (o ConfigGroupOutput) Resources() khulnasoft.ArrayOutput {
	return o.ApplyT(func(v *ConfigGroup) khulnasoft.ArrayOutput { return v.Resources }).(khulnasoft.ArrayOutput)
}

type ConfigGroupArrayOutput struct{ *khulnasoft.OutputState }

func (ConfigGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConfigGroup)(nil)).Elem()
}

func (o ConfigGroupArrayOutput) ToConfigGroupArrayOutput() ConfigGroupArrayOutput {
	return o
}

func (o ConfigGroupArrayOutput) ToConfigGroupArrayOutputWithContext(ctx context.Context) ConfigGroupArrayOutput {
	return o
}

func (o ConfigGroupArrayOutput) Index(i khulnasoft.IntInput) ConfigGroupOutput {
	return khulnasoft.All(o, i).ApplyT(func(vs []interface{}) *ConfigGroup {
		return vs[0].([]*ConfigGroup)[vs[1].(int)]
	}).(ConfigGroupOutput)
}

type ConfigGroupMapOutput struct{ *khulnasoft.OutputState }

func (ConfigGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConfigGroup)(nil)).Elem()
}

func (o ConfigGroupMapOutput) ToConfigGroupMapOutput() ConfigGroupMapOutput {
	return o
}

func (o ConfigGroupMapOutput) ToConfigGroupMapOutputWithContext(ctx context.Context) ConfigGroupMapOutput {
	return o
}

func (o ConfigGroupMapOutput) MapIndex(k khulnasoft.StringInput) ConfigGroupOutput {
	return khulnasoft.All(o, k).ApplyT(func(vs []interface{}) *ConfigGroup {
		return vs[0].(map[string]*ConfigGroup)[vs[1].(string)]
	}).(ConfigGroupOutput)
}

func init() {
	khulnasoft.RegisterInputType(reflect.TypeOf((*ConfigGroupInput)(nil)).Elem(), &ConfigGroup{})
	khulnasoft.RegisterInputType(reflect.TypeOf((*ConfigGroupArrayInput)(nil)).Elem(), ConfigGroupArray{})
	khulnasoft.RegisterInputType(reflect.TypeOf((*ConfigGroupMapInput)(nil)).Elem(), ConfigGroupMap{})
	khulnasoft.RegisterOutputType(ConfigGroupOutput{})
	khulnasoft.RegisterOutputType(ConfigGroupArrayOutput{})
	khulnasoft.RegisterOutputType(ConfigGroupMapOutput{})
}
