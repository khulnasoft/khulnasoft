// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package world

import (
	"context"
	"reflect"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"regress-go-10527/world/internal"
)

type WorldMap struct {
	khulnasoft.CustomResourceState

	Name khulnasoft.StringPtrOutput `khulnasoft:"name"`
}

// NewWorldMap registers a new resource with the given unique name, arguments, and options.
func NewWorldMap(ctx *khulnasoft.Context,
	name string, args *WorldMapArgs, opts ...khulnasoft.ResourceOption) (*WorldMap, error) {
	if args == nil {
		args = &WorldMapArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource WorldMap
	err := ctx.RegisterResource("world::WorldMap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorldMap gets an existing WorldMap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorldMap(ctx *khulnasoft.Context,
	name string, id khulnasoft.IDInput, state *WorldMapState, opts ...khulnasoft.ResourceOption) (*WorldMap, error) {
	var resource WorldMap
	err := ctx.ReadResource("world::WorldMap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorldMap resources.
type worldMapState struct {
}

type WorldMapState struct {
}

func (WorldMapState) ElementType() reflect.Type {
	return reflect.TypeOf((*worldMapState)(nil)).Elem()
}

type worldMapArgs struct {
}

// The set of arguments for constructing a WorldMap resource.
type WorldMapArgs struct {
}

func (WorldMapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*worldMapArgs)(nil)).Elem()
}

type WorldMapInput interface {
	khulnasoft.Input

	ToWorldMapOutput() WorldMapOutput
	ToWorldMapOutputWithContext(ctx context.Context) WorldMapOutput
}

func (*WorldMap) ElementType() reflect.Type {
	return reflect.TypeOf((**WorldMap)(nil)).Elem()
}

func (i *WorldMap) ToWorldMapOutput() WorldMapOutput {
	return i.ToWorldMapOutputWithContext(context.Background())
}

func (i *WorldMap) ToWorldMapOutputWithContext(ctx context.Context) WorldMapOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(WorldMapOutput)
}

type WorldMapOutput struct{ *khulnasoft.OutputState }

func (WorldMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorldMap)(nil)).Elem()
}

func (o WorldMapOutput) ToWorldMapOutput() WorldMapOutput {
	return o
}

func (o WorldMapOutput) ToWorldMapOutputWithContext(ctx context.Context) WorldMapOutput {
	return o
}

func (o WorldMapOutput) Name() khulnasoft.StringPtrOutput {
	return o.ApplyT(func(v *WorldMap) khulnasoft.StringPtrOutput { return v.Name }).(khulnasoft.StringPtrOutput)
}

func init() {
	khulnasoft.RegisterInputType(reflect.TypeOf((*WorldMapInput)(nil)).Elem(), &WorldMap{})
	khulnasoft.RegisterOutputType(WorldMapOutput{})
}
