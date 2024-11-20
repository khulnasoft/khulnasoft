// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"dashed-import-schema/plant-provider"
	"dashed-import-schema/plant-provider/internal"
	"errors"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

type RubberTree struct {
	khulnasoft.CustomResourceState

	Container plantprovider.ContainerPtrOutput `khulnasoft:"container"`
	Diameter  DiameterOutput                   `khulnasoft:"diameter"`
	Farm      khulnasoft.StringPtrOutput           `khulnasoft:"farm"`
	Size      TreeSizePtrOutput                `khulnasoft:"size"`
	Type      RubberTreeVarietyOutput          `khulnasoft:"type"`
}

// NewRubberTree registers a new resource with the given unique name, arguments, and options.
func NewRubberTree(ctx *khulnasoft.Context,
	name string, args *RubberTreeArgs, opts ...khulnasoft.ResourceOption) (*RubberTree, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Container != nil {
		args.Container = args.Container.ToContainerPtrOutput().ApplyT(func(v *plantprovider.Container) *plantprovider.Container { return v.Defaults() }).(plantprovider.ContainerPtrOutput)
	}
	if args.Diameter == nil {
		args.Diameter = Diameter(6.0)
	}
	if args.Farm == nil {
		args.Farm = khulnasoft.StringPtr("(unknown)")
	}
	if args.Size == nil {
		args.Size = TreeSize("medium")
	}
	if args.Type == nil {
		args.Type = RubberTreeVariety("Burgundy")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RubberTree
	err := ctx.RegisterResource("plant:tree/v1:RubberTree", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRubberTree gets an existing RubberTree resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRubberTree(ctx *khulnasoft.Context,
	name string, id khulnasoft.IDInput, state *RubberTreeState, opts ...khulnasoft.ResourceOption) (*RubberTree, error) {
	var resource RubberTree
	err := ctx.ReadResource("plant:tree/v1:RubberTree", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RubberTree resources.
type rubberTreeState struct {
	Farm *string `khulnasoft:"farm"`
}

type RubberTreeState struct {
	Farm khulnasoft.StringPtrInput
}

func (RubberTreeState) ElementType() reflect.Type {
	return reflect.TypeOf((*rubberTreeState)(nil)).Elem()
}

type rubberTreeArgs struct {
	Container *plantprovider.Container `khulnasoft:"container"`
	Diameter  Diameter                 `khulnasoft:"diameter"`
	Farm      *string                  `khulnasoft:"farm"`
	Size      *TreeSize                `khulnasoft:"size"`
	Type      RubberTreeVariety        `khulnasoft:"type"`
}

// The set of arguments for constructing a RubberTree resource.
type RubberTreeArgs struct {
	Container plantprovider.ContainerPtrInput
	Diameter  DiameterInput
	Farm      khulnasoft.StringPtrInput
	Size      TreeSizePtrInput
	Type      RubberTreeVarietyInput
}

func (RubberTreeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rubberTreeArgs)(nil)).Elem()
}

type RubberTreeInput interface {
	khulnasoft.Input

	ToRubberTreeOutput() RubberTreeOutput
	ToRubberTreeOutputWithContext(ctx context.Context) RubberTreeOutput
}

func (*RubberTree) ElementType() reflect.Type {
	return reflect.TypeOf((**RubberTree)(nil)).Elem()
}

func (i *RubberTree) ToRubberTreeOutput() RubberTreeOutput {
	return i.ToRubberTreeOutputWithContext(context.Background())
}

func (i *RubberTree) ToRubberTreeOutputWithContext(ctx context.Context) RubberTreeOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(RubberTreeOutput)
}

type RubberTreeOutput struct{ *khulnasoft.OutputState }

func (RubberTreeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RubberTree)(nil)).Elem()
}

func (o RubberTreeOutput) ToRubberTreeOutput() RubberTreeOutput {
	return o
}

func (o RubberTreeOutput) ToRubberTreeOutputWithContext(ctx context.Context) RubberTreeOutput {
	return o
}

func (o RubberTreeOutput) Container() plantprovider.ContainerPtrOutput {
	return o.ApplyT(func(v *RubberTree) plantprovider.ContainerPtrOutput { return v.Container }).(plantprovider.ContainerPtrOutput)
}

func (o RubberTreeOutput) Diameter() DiameterOutput {
	return o.ApplyT(func(v *RubberTree) DiameterOutput { return v.Diameter }).(DiameterOutput)
}

func (o RubberTreeOutput) Farm() khulnasoft.StringPtrOutput {
	return o.ApplyT(func(v *RubberTree) khulnasoft.StringPtrOutput { return v.Farm }).(khulnasoft.StringPtrOutput)
}

func (o RubberTreeOutput) Size() TreeSizePtrOutput {
	return o.ApplyT(func(v *RubberTree) TreeSizePtrOutput { return v.Size }).(TreeSizePtrOutput)
}

func (o RubberTreeOutput) Type() RubberTreeVarietyOutput {
	return o.ApplyT(func(v *RubberTree) RubberTreeVarietyOutput { return v.Type }).(RubberTreeVarietyOutput)
}

func init() {
	khulnasoft.RegisterInputType(reflect.TypeOf((*RubberTreeInput)(nil)).Elem(), &RubberTree{})
	khulnasoft.RegisterOutputType(RubberTreeOutput{})
}