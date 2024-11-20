// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mypkg

import (
	"context"
	"reflect"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"output-funcs-tfbridge20/mypkg/internal"
)

var _ = internal.GetEnvOrDefault

// An access key for the storage account.
type StorageAccountKeyResponse struct {
	// Creation time of the key, in round trip date format.
	CreationTime string `khulnasoft:"creationTime"`
	// Name of the key.
	KeyName string `khulnasoft:"keyName"`
	// Permissions for the key -- read-only or full permissions.
	Permissions string `khulnasoft:"permissions"`
	// Base 64-encoded value of the key.
	Value string `khulnasoft:"value"`
}

// StorageAccountKeyResponseInput is an input type that accepts StorageAccountKeyResponseArgs and StorageAccountKeyResponseOutput values.
// You can construct a concrete instance of `StorageAccountKeyResponseInput` via:
//
//	StorageAccountKeyResponseArgs{...}
type StorageAccountKeyResponseInput interface {
	khulnasoft.Input

	ToStorageAccountKeyResponseOutput() StorageAccountKeyResponseOutput
	ToStorageAccountKeyResponseOutputWithContext(context.Context) StorageAccountKeyResponseOutput
}

// An access key for the storage account.
type StorageAccountKeyResponseArgs struct {
	// Creation time of the key, in round trip date format.
	CreationTime khulnasoft.StringInput `khulnasoft:"creationTime"`
	// Name of the key.
	KeyName khulnasoft.StringInput `khulnasoft:"keyName"`
	// Permissions for the key -- read-only or full permissions.
	Permissions khulnasoft.StringInput `khulnasoft:"permissions"`
	// Base 64-encoded value of the key.
	Value khulnasoft.StringInput `khulnasoft:"value"`
}

func (StorageAccountKeyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountKeyResponse)(nil)).Elem()
}

func (i StorageAccountKeyResponseArgs) ToStorageAccountKeyResponseOutput() StorageAccountKeyResponseOutput {
	return i.ToStorageAccountKeyResponseOutputWithContext(context.Background())
}

func (i StorageAccountKeyResponseArgs) ToStorageAccountKeyResponseOutputWithContext(ctx context.Context) StorageAccountKeyResponseOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(StorageAccountKeyResponseOutput)
}

// StorageAccountKeyResponseArrayInput is an input type that accepts StorageAccountKeyResponseArray and StorageAccountKeyResponseArrayOutput values.
// You can construct a concrete instance of `StorageAccountKeyResponseArrayInput` via:
//
//	StorageAccountKeyResponseArray{ StorageAccountKeyResponseArgs{...} }
type StorageAccountKeyResponseArrayInput interface {
	khulnasoft.Input

	ToStorageAccountKeyResponseArrayOutput() StorageAccountKeyResponseArrayOutput
	ToStorageAccountKeyResponseArrayOutputWithContext(context.Context) StorageAccountKeyResponseArrayOutput
}

type StorageAccountKeyResponseArray []StorageAccountKeyResponseInput

func (StorageAccountKeyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageAccountKeyResponse)(nil)).Elem()
}

func (i StorageAccountKeyResponseArray) ToStorageAccountKeyResponseArrayOutput() StorageAccountKeyResponseArrayOutput {
	return i.ToStorageAccountKeyResponseArrayOutputWithContext(context.Background())
}

func (i StorageAccountKeyResponseArray) ToStorageAccountKeyResponseArrayOutputWithContext(ctx context.Context) StorageAccountKeyResponseArrayOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(StorageAccountKeyResponseArrayOutput)
}

// An access key for the storage account.
type StorageAccountKeyResponseOutput struct{ *khulnasoft.OutputState }

func (StorageAccountKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountKeyResponse)(nil)).Elem()
}

func (o StorageAccountKeyResponseOutput) ToStorageAccountKeyResponseOutput() StorageAccountKeyResponseOutput {
	return o
}

func (o StorageAccountKeyResponseOutput) ToStorageAccountKeyResponseOutputWithContext(ctx context.Context) StorageAccountKeyResponseOutput {
	return o
}

// Creation time of the key, in round trip date format.
func (o StorageAccountKeyResponseOutput) CreationTime() khulnasoft.StringOutput {
	return o.ApplyT(func(v StorageAccountKeyResponse) string { return v.CreationTime }).(khulnasoft.StringOutput)
}

// Name of the key.
func (o StorageAccountKeyResponseOutput) KeyName() khulnasoft.StringOutput {
	return o.ApplyT(func(v StorageAccountKeyResponse) string { return v.KeyName }).(khulnasoft.StringOutput)
}

// Permissions for the key -- read-only or full permissions.
func (o StorageAccountKeyResponseOutput) Permissions() khulnasoft.StringOutput {
	return o.ApplyT(func(v StorageAccountKeyResponse) string { return v.Permissions }).(khulnasoft.StringOutput)
}

// Base 64-encoded value of the key.
func (o StorageAccountKeyResponseOutput) Value() khulnasoft.StringOutput {
	return o.ApplyT(func(v StorageAccountKeyResponse) string { return v.Value }).(khulnasoft.StringOutput)
}

type StorageAccountKeyResponseArrayOutput struct{ *khulnasoft.OutputState }

func (StorageAccountKeyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageAccountKeyResponse)(nil)).Elem()
}

func (o StorageAccountKeyResponseArrayOutput) ToStorageAccountKeyResponseArrayOutput() StorageAccountKeyResponseArrayOutput {
	return o
}

func (o StorageAccountKeyResponseArrayOutput) ToStorageAccountKeyResponseArrayOutputWithContext(ctx context.Context) StorageAccountKeyResponseArrayOutput {
	return o
}

func (o StorageAccountKeyResponseArrayOutput) Index(i khulnasoft.IntInput) StorageAccountKeyResponseOutput {
	return khulnasoft.All(o, i).ApplyT(func(vs []interface{}) StorageAccountKeyResponse {
		return vs[0].([]StorageAccountKeyResponse)[vs[1].(int)]
	}).(StorageAccountKeyResponseOutput)
}

type GetAmiIdsFilter struct {
	Name   string   `khulnasoft:"name"`
	Values []string `khulnasoft:"values"`
}

// GetAmiIdsFilterInput is an input type that accepts GetAmiIdsFilterArgs and GetAmiIdsFilterOutput values.
// You can construct a concrete instance of `GetAmiIdsFilterInput` via:
//
//	GetAmiIdsFilterArgs{...}
type GetAmiIdsFilterInput interface {
	khulnasoft.Input

	ToGetAmiIdsFilterOutput() GetAmiIdsFilterOutput
	ToGetAmiIdsFilterOutputWithContext(context.Context) GetAmiIdsFilterOutput
}

type GetAmiIdsFilterArgs struct {
	Name   khulnasoft.StringInput      `khulnasoft:"name"`
	Values khulnasoft.StringArrayInput `khulnasoft:"values"`
}

func (GetAmiIdsFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAmiIdsFilter)(nil)).Elem()
}

func (i GetAmiIdsFilterArgs) ToGetAmiIdsFilterOutput() GetAmiIdsFilterOutput {
	return i.ToGetAmiIdsFilterOutputWithContext(context.Background())
}

func (i GetAmiIdsFilterArgs) ToGetAmiIdsFilterOutputWithContext(ctx context.Context) GetAmiIdsFilterOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(GetAmiIdsFilterOutput)
}

// GetAmiIdsFilterArrayInput is an input type that accepts GetAmiIdsFilterArray and GetAmiIdsFilterArrayOutput values.
// You can construct a concrete instance of `GetAmiIdsFilterArrayInput` via:
//
//	GetAmiIdsFilterArray{ GetAmiIdsFilterArgs{...} }
type GetAmiIdsFilterArrayInput interface {
	khulnasoft.Input

	ToGetAmiIdsFilterArrayOutput() GetAmiIdsFilterArrayOutput
	ToGetAmiIdsFilterArrayOutputWithContext(context.Context) GetAmiIdsFilterArrayOutput
}

type GetAmiIdsFilterArray []GetAmiIdsFilterInput

func (GetAmiIdsFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetAmiIdsFilter)(nil)).Elem()
}

func (i GetAmiIdsFilterArray) ToGetAmiIdsFilterArrayOutput() GetAmiIdsFilterArrayOutput {
	return i.ToGetAmiIdsFilterArrayOutputWithContext(context.Background())
}

func (i GetAmiIdsFilterArray) ToGetAmiIdsFilterArrayOutputWithContext(ctx context.Context) GetAmiIdsFilterArrayOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(GetAmiIdsFilterArrayOutput)
}

type GetAmiIdsFilterOutput struct{ *khulnasoft.OutputState }

func (GetAmiIdsFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAmiIdsFilter)(nil)).Elem()
}

func (o GetAmiIdsFilterOutput) ToGetAmiIdsFilterOutput() GetAmiIdsFilterOutput {
	return o
}

func (o GetAmiIdsFilterOutput) ToGetAmiIdsFilterOutputWithContext(ctx context.Context) GetAmiIdsFilterOutput {
	return o
}

func (o GetAmiIdsFilterOutput) Name() khulnasoft.StringOutput {
	return o.ApplyT(func(v GetAmiIdsFilter) string { return v.Name }).(khulnasoft.StringOutput)
}

func (o GetAmiIdsFilterOutput) Values() khulnasoft.StringArrayOutput {
	return o.ApplyT(func(v GetAmiIdsFilter) []string { return v.Values }).(khulnasoft.StringArrayOutput)
}

type GetAmiIdsFilterArrayOutput struct{ *khulnasoft.OutputState }

func (GetAmiIdsFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetAmiIdsFilter)(nil)).Elem()
}

func (o GetAmiIdsFilterArrayOutput) ToGetAmiIdsFilterArrayOutput() GetAmiIdsFilterArrayOutput {
	return o
}

func (o GetAmiIdsFilterArrayOutput) ToGetAmiIdsFilterArrayOutputWithContext(ctx context.Context) GetAmiIdsFilterArrayOutput {
	return o
}

func (o GetAmiIdsFilterArrayOutput) Index(i khulnasoft.IntInput) GetAmiIdsFilterOutput {
	return khulnasoft.All(o, i).ApplyT(func(vs []interface{}) GetAmiIdsFilter {
		return vs[0].([]GetAmiIdsFilter)[vs[1].(int)]
	}).(GetAmiIdsFilterOutput)
}

func init() {
	khulnasoft.RegisterInputType(reflect.TypeOf((*StorageAccountKeyResponseInput)(nil)).Elem(), StorageAccountKeyResponseArgs{})
	khulnasoft.RegisterInputType(reflect.TypeOf((*StorageAccountKeyResponseArrayInput)(nil)).Elem(), StorageAccountKeyResponseArray{})
	khulnasoft.RegisterInputType(reflect.TypeOf((*GetAmiIdsFilterInput)(nil)).Elem(), GetAmiIdsFilterArgs{})
	khulnasoft.RegisterInputType(reflect.TypeOf((*GetAmiIdsFilterArrayInput)(nil)).Elem(), GetAmiIdsFilterArray{})
	khulnasoft.RegisterOutputType(StorageAccountKeyResponseOutput{})
	khulnasoft.RegisterOutputType(StorageAccountKeyResponseArrayOutput{})
	khulnasoft.RegisterOutputType(GetAmiIdsFilterOutput{})
	khulnasoft.RegisterOutputType(GetAmiIdsFilterArrayOutput{})
}