// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mypkg

import (
	"context"
	"reflect"

	"errors"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoftx"
	"secrets-go-generics-only/mypkg/internal"
)

type Resource struct {
	khulnasoft.CustomResourceState

	Config      khulnasoftx.GPtrOutput[Config, ConfigOutput]   `khulnasoft:"config"`
	ConfigArray khulnasoftx.GArrayOutput[Config, ConfigOutput] `khulnasoft:"configArray"`
	ConfigMap   khulnasoftx.GMapOutput[Config, ConfigOutput]   `khulnasoft:"configMap"`
	Foo         khulnasoftx.Output[string]                     `khulnasoft:"foo"`
	FooArray    khulnasoftx.ArrayOutput[string]                `khulnasoft:"fooArray"`
	FooMap      khulnasoftx.MapOutput[string]                  `khulnasoft:"fooMap"`
}

// NewResource registers a new resource with the given unique name, arguments, and options.
func NewResource(ctx *khulnasoft.Context,
	name string, args *ResourceArgs, opts ...khulnasoft.ResourceOption) (*Resource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Config == nil {
		return nil, errors.New("invalid value for required argument 'Config'")
	}
	if args.ConfigArray == nil {
		return nil, errors.New("invalid value for required argument 'ConfigArray'")
	}
	if args.ConfigMap == nil {
		return nil, errors.New("invalid value for required argument 'ConfigMap'")
	}
	if args.Foo == nil {
		return nil, errors.New("invalid value for required argument 'Foo'")
	}
	if args.FooArray == nil {
		return nil, errors.New("invalid value for required argument 'FooArray'")
	}
	if args.FooMap == nil {
		return nil, errors.New("invalid value for required argument 'FooMap'")
	}
	if args.Config != nil {
		untypedSecretValue := khulnasoft.ToSecret(args.Config.ToOutput(ctx.Context()).Untyped())
		args.Config = khulnasoftx.MustConvertTyped[*ConfigArgs](untypedSecretValue)
	}
	if args.ConfigArray != nil {
		untypedSecretValue := khulnasoft.ToSecret(args.ConfigArray.ToOutput(ctx.Context()).Untyped())
		args.ConfigArray = khulnasoftx.MustConvertTyped[[]*ConfigArgs](untypedSecretValue)
	}
	if args.ConfigMap != nil {
		untypedSecretValue := khulnasoft.ToSecret(args.ConfigMap.ToOutput(ctx.Context()).Untyped())
		args.ConfigMap = khulnasoftx.MustConvertTyped[map[string]*ConfigArgs](untypedSecretValue)
	}
	if args.Foo != nil {
		untypedSecretValue := khulnasoft.ToSecret(args.Foo.ToOutput(ctx.Context()).Untyped())
		args.Foo = khulnasoftx.MustConvertTyped[string](untypedSecretValue)
	}
	if args.FooArray != nil {
		untypedSecretValue := khulnasoft.ToSecret(args.FooArray.ToOutput(ctx.Context()).Untyped())
		args.FooArray = khulnasoftx.MustConvertTyped[[]string](untypedSecretValue)
	}
	if args.FooMap != nil {
		untypedSecretValue := khulnasoft.ToSecret(args.FooMap.ToOutput(ctx.Context()).Untyped())
		args.FooMap = khulnasoftx.MustConvertTyped[map[string]string](untypedSecretValue)
	}
	secrets := khulnasoft.AdditionalSecretOutputs([]string{
		"config",
		"configArray",
		"configMap",
		"foo",
		"fooArray",
		"fooMap",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Resource
	err := ctx.RegisterResource("mypkg::Resource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResource gets an existing Resource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResource(ctx *khulnasoft.Context,
	name string, id khulnasoft.IDInput, state *ResourceState, opts ...khulnasoft.ResourceOption) (*Resource, error) {
	var resource Resource
	err := ctx.ReadResource("mypkg::Resource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Resource resources.
type resourceState struct {
}

type ResourceState struct {
}

func (ResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceState)(nil)).Elem()
}

type resourceArgs struct {
	Config      Config            `khulnasoft:"config"`
	ConfigArray []Config          `khulnasoft:"configArray"`
	ConfigMap   map[string]Config `khulnasoft:"configMap"`
	Foo         string            `khulnasoft:"foo"`
	FooArray    []string          `khulnasoft:"fooArray"`
	FooMap      map[string]string `khulnasoft:"fooMap"`
}

// The set of arguments for constructing a Resource resource.
type ResourceArgs struct {
	Config      khulnasoftx.Input[*ConfigArgs]
	ConfigArray khulnasoftx.Input[[]*ConfigArgs]
	ConfigMap   khulnasoftx.Input[map[string]*ConfigArgs]
	Foo         khulnasoftx.Input[string]
	FooArray    khulnasoftx.Input[[]string]
	FooMap      khulnasoftx.Input[map[string]string]
}

func (ResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceArgs)(nil)).Elem()
}

type ResourceOutput struct{ *khulnasoft.OutputState }

func (ResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Resource)(nil)).Elem()
}

func (o ResourceOutput) ToResourceOutput() ResourceOutput {
	return o
}

func (o ResourceOutput) ToResourceOutputWithContext(ctx context.Context) ResourceOutput {
	return o
}

func (o ResourceOutput) ToOutput(ctx context.Context) khulnasoftx.Output[Resource] {
	return khulnasoftx.Output[Resource]{
		OutputState: o.OutputState,
	}
}

func (o ResourceOutput) Config() khulnasoftx.GPtrOutput[Config, ConfigOutput] {
	value := khulnasoftx.Apply[Resource](o, func(v Resource) khulnasoftx.GPtrOutput[Config, ConfigOutput] { return v.Config })
	unwrapped := khulnasoftx.Flatten[*Config, khulnasoftx.GPtrOutput[Config, ConfigOutput]](value)
	return khulnasoftx.GPtrOutput[Config, ConfigOutput]{OutputState: unwrapped.OutputState}
}

func (o ResourceOutput) ConfigArray() khulnasoftx.GArrayOutput[Config, ConfigOutput] {
	value := khulnasoftx.Apply[Resource](o, func(v Resource) khulnasoftx.GArrayOutput[Config, ConfigOutput] { return v.ConfigArray })
	unwrapped := khulnasoftx.Flatten[[]Config, khulnasoftx.GArrayOutput[Config, ConfigOutput]](value)
	return khulnasoftx.GArrayOutput[Config, ConfigOutput]{OutputState: unwrapped.OutputState}
}

func (o ResourceOutput) ConfigMap() khulnasoftx.GMapOutput[Config, ConfigOutput] {
	value := khulnasoftx.Apply[Resource](o, func(v Resource) khulnasoftx.GMapOutput[Config, ConfigOutput] { return v.ConfigMap })
	unwrapped := khulnasoftx.Flatten[map[string]Config, khulnasoftx.GMapOutput[Config, ConfigOutput]](value)
	return khulnasoftx.GMapOutput[Config, ConfigOutput]{OutputState: unwrapped.OutputState}
}

func (o ResourceOutput) Foo() khulnasoftx.Output[string] {
	value := khulnasoftx.Apply[Resource](o, func(v Resource) khulnasoftx.Output[string] { return v.Foo })
	return khulnasoftx.Flatten[string, khulnasoftx.Output[string]](value)
}

func (o ResourceOutput) FooArray() khulnasoftx.ArrayOutput[string] {
	value := khulnasoftx.Apply[Resource](o, func(v Resource) khulnasoftx.ArrayOutput[string] { return v.FooArray })
	unwrapped := khulnasoftx.Flatten[[]string, khulnasoftx.ArrayOutput[string]](value)
	return khulnasoftx.ArrayOutput[string]{OutputState: unwrapped.OutputState}
}

func (o ResourceOutput) FooMap() khulnasoftx.MapOutput[string] {
	value := khulnasoftx.Apply[Resource](o, func(v Resource) khulnasoftx.MapOutput[string] { return v.FooMap })
	unwrapped := khulnasoftx.Flatten[map[string]string, khulnasoftx.MapOutput[string]](value)
	return khulnasoftx.MapOutput[string]{OutputState: unwrapped.OutputState}
}

func init() {
	khulnasoft.RegisterOutputType(ResourceOutput{})
}
