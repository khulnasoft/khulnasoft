// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mymodule

import (
	"context"
	"reflect"

	"enum-reference/example/internal"
	iam "github.com/khulnasoft/khulnasoft-google-native/sdk/go/google/iam/v1"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

type IamResource struct {
	khulnasoft.ResourceState
}

// NewIamResource registers a new resource with the given unique name, arguments, and options.
func NewIamResource(ctx *khulnasoft.Context,
	name string, args *IamResourceArgs, opts ...khulnasoft.ResourceOption) (*IamResource, error) {
	if args == nil {
		args = &IamResourceArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IamResource
	err := ctx.RegisterRemoteComponentResource("example:myModule:IamResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type iamResourceArgs struct {
	Config *iam.AuditConfig `khulnasoft:"config"`
}

// The set of arguments for constructing a IamResource resource.
type IamResourceArgs struct {
	Config *iam.AuditConfigInput
}

func (IamResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iamResourceArgs)(nil)).Elem()
}

type IamResourceInput interface {
	khulnasoft.Input

	ToIamResourceOutput() IamResourceOutput
	ToIamResourceOutputWithContext(ctx context.Context) IamResourceOutput
}

func (*IamResource) ElementType() reflect.Type {
	return reflect.TypeOf((**IamResource)(nil)).Elem()
}

func (i *IamResource) ToIamResourceOutput() IamResourceOutput {
	return i.ToIamResourceOutputWithContext(context.Background())
}

func (i *IamResource) ToIamResourceOutputWithContext(ctx context.Context) IamResourceOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(IamResourceOutput)
}

type IamResourceOutput struct{ *khulnasoft.OutputState }

func (IamResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IamResource)(nil)).Elem()
}

func (o IamResourceOutput) ToIamResourceOutput() IamResourceOutput {
	return o
}

func (o IamResourceOutput) ToIamResourceOutputWithContext(ctx context.Context) IamResourceOutput {
	return o
}

func init() {
	khulnasoft.RegisterInputType(reflect.TypeOf((*IamResourceInput)(nil)).Elem(), &IamResource{})
	khulnasoft.RegisterOutputType(IamResourceOutput{})
}