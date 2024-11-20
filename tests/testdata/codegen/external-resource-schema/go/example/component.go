// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"context"
	"reflect"

	"errors"
	"external-resource-schema/example/internal"
	"github.com/khulnasoft/khulnasoft-aws/sdk/v4/go/aws/ec2"
	"github.com/khulnasoft/khulnasoft-kubernetes/sdk/v3/go/kubernetes"
	metav1 "github.com/khulnasoft/khulnasoft-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	storagev1 "github.com/khulnasoft/khulnasoft-kubernetes/sdk/v3/go/kubernetes/storage/v1"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

type Component struct {
	khulnasoft.CustomResourceState

	Provider       kubernetes.ProviderOutput       `khulnasoft:"provider"`
	SecurityGroup  ec2.SecurityGroupOutput         `khulnasoft:"securityGroup"`
	StorageClasses storagev1.StorageClassMapOutput `khulnasoft:"storageClasses"`
}

// NewComponent registers a new resource with the given unique name, arguments, and options.
func NewComponent(ctx *khulnasoft.Context,
	name string, args *ComponentArgs, opts ...khulnasoft.ResourceOption) (*Component, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RequiredMetadata == nil {
		return nil, errors.New("invalid value for required argument 'RequiredMetadata'")
	}
	if args.RequiredMetadataArray == nil {
		return nil, errors.New("invalid value for required argument 'RequiredMetadataArray'")
	}
	if args.RequiredMetadataMap == nil {
		return nil, errors.New("invalid value for required argument 'RequiredMetadataMap'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Component
	err := ctx.RegisterResource("example::Component", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComponent gets an existing Component resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComponent(ctx *khulnasoft.Context,
	name string, id khulnasoft.IDInput, state *ComponentState, opts ...khulnasoft.ResourceOption) (*Component, error) {
	var resource Component
	err := ctx.ReadResource("example::Component", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Component resources.
type componentState struct {
}

type ComponentState struct {
}

func (ComponentState) ElementType() reflect.Type {
	return reflect.TypeOf((*componentState)(nil)).Elem()
}

type componentArgs struct {
	Metadata              *metav1.ObjectMeta           `khulnasoft:"metadata"`
	MetadataArray         []metav1.ObjectMeta          `khulnasoft:"metadataArray"`
	MetadataMap           map[string]metav1.ObjectMeta `khulnasoft:"metadataMap"`
	RequiredMetadata      metav1.ObjectMeta            `khulnasoft:"requiredMetadata"`
	RequiredMetadataArray []metav1.ObjectMeta          `khulnasoft:"requiredMetadataArray"`
	RequiredMetadataMap   map[string]metav1.ObjectMeta `khulnasoft:"requiredMetadataMap"`
}

// The set of arguments for constructing a Component resource.
type ComponentArgs struct {
	Metadata              metav1.ObjectMetaPtrInput
	MetadataArray         metav1.ObjectMetaArrayInput
	MetadataMap           metav1.ObjectMetaMapInput
	RequiredMetadata      metav1.ObjectMetaInput
	RequiredMetadataArray metav1.ObjectMetaArrayInput
	RequiredMetadataMap   metav1.ObjectMetaMapInput
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

// ComponentArrayInput is an input type that accepts ComponentArray and ComponentArrayOutput values.
// You can construct a concrete instance of `ComponentArrayInput` via:
//
//	ComponentArray{ ComponentArgs{...} }
type ComponentArrayInput interface {
	khulnasoft.Input

	ToComponentArrayOutput() ComponentArrayOutput
	ToComponentArrayOutputWithContext(context.Context) ComponentArrayOutput
}

type ComponentArray []ComponentInput

func (ComponentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Component)(nil)).Elem()
}

func (i ComponentArray) ToComponentArrayOutput() ComponentArrayOutput {
	return i.ToComponentArrayOutputWithContext(context.Background())
}

func (i ComponentArray) ToComponentArrayOutputWithContext(ctx context.Context) ComponentArrayOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(ComponentArrayOutput)
}

// ComponentMapInput is an input type that accepts ComponentMap and ComponentMapOutput values.
// You can construct a concrete instance of `ComponentMapInput` via:
//
//	ComponentMap{ "key": ComponentArgs{...} }
type ComponentMapInput interface {
	khulnasoft.Input

	ToComponentMapOutput() ComponentMapOutput
	ToComponentMapOutputWithContext(context.Context) ComponentMapOutput
}

type ComponentMap map[string]ComponentInput

func (ComponentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Component)(nil)).Elem()
}

func (i ComponentMap) ToComponentMapOutput() ComponentMapOutput {
	return i.ToComponentMapOutputWithContext(context.Background())
}

func (i ComponentMap) ToComponentMapOutputWithContext(ctx context.Context) ComponentMapOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(ComponentMapOutput)
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

func (o ComponentOutput) Provider() kubernetes.ProviderOutput {
	return o.ApplyT(func(v *Component) kubernetes.ProviderOutput { return v.Provider }).(kubernetes.ProviderOutput)
}

func (o ComponentOutput) SecurityGroup() ec2.SecurityGroupOutput {
	return o.ApplyT(func(v *Component) ec2.SecurityGroupOutput { return v.SecurityGroup }).(ec2.SecurityGroupOutput)
}

func (o ComponentOutput) StorageClasses() storagev1.StorageClassMapOutput {
	return o.ApplyT(func(v *Component) storagev1.StorageClassMapOutput { return v.StorageClasses }).(storagev1.StorageClassMapOutput)
}

type ComponentArrayOutput struct{ *khulnasoft.OutputState }

func (ComponentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Component)(nil)).Elem()
}

func (o ComponentArrayOutput) ToComponentArrayOutput() ComponentArrayOutput {
	return o
}

func (o ComponentArrayOutput) ToComponentArrayOutputWithContext(ctx context.Context) ComponentArrayOutput {
	return o
}

func (o ComponentArrayOutput) Index(i khulnasoft.IntInput) ComponentOutput {
	return khulnasoft.All(o, i).ApplyT(func(vs []interface{}) *Component {
		return vs[0].([]*Component)[vs[1].(int)]
	}).(ComponentOutput)
}

type ComponentMapOutput struct{ *khulnasoft.OutputState }

func (ComponentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Component)(nil)).Elem()
}

func (o ComponentMapOutput) ToComponentMapOutput() ComponentMapOutput {
	return o
}

func (o ComponentMapOutput) ToComponentMapOutputWithContext(ctx context.Context) ComponentMapOutput {
	return o
}

func (o ComponentMapOutput) MapIndex(k khulnasoft.StringInput) ComponentOutput {
	return khulnasoft.All(o, k).ApplyT(func(vs []interface{}) *Component {
		return vs[0].(map[string]*Component)[vs[1].(string)]
	}).(ComponentOutput)
}

func init() {
	khulnasoft.RegisterInputType(reflect.TypeOf((*ComponentInput)(nil)).Elem(), &Component{})
	khulnasoft.RegisterInputType(reflect.TypeOf((*ComponentArrayInput)(nil)).Elem(), ComponentArray{})
	khulnasoft.RegisterInputType(reflect.TypeOf((*ComponentMapInput)(nil)).Elem(), ComponentMap{})
	khulnasoft.RegisterOutputType(ComponentOutput{})
	khulnasoft.RegisterOutputType(ComponentArrayOutput{})
	khulnasoft.RegisterOutputType(ComponentMapOutput{})
}