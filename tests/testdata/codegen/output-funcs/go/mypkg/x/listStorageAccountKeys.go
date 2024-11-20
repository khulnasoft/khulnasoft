// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mypkg

import (
	"context"
	"reflect"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoftx"
	"output-funcs/mypkg/internal"
)

// The response from the ListKeys operation.
// API Version: 2021-02-01.
func ListStorageAccountKeys(ctx *khulnasoft.Context, args *ListStorageAccountKeysArgs, opts ...khulnasoft.InvokeOption) (*ListStorageAccountKeysResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListStorageAccountKeysResult
	err := ctx.Invoke("mypkg::listStorageAccountKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListStorageAccountKeysArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName string `khulnasoft:"accountName"`
	// Specifies type of the key to be listed. Possible value is kerb.
	Expand *string `khulnasoft:"expand"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `khulnasoft:"resourceGroupName"`
}

// The response from the ListKeys operation.
type ListStorageAccountKeysResult struct {
	// Gets the list of storage account keys and their properties for the specified storage account.
	Keys []StorageAccountKeyResponse `khulnasoft:"keys"`
}

func ListStorageAccountKeysOutput(ctx *khulnasoft.Context, args ListStorageAccountKeysOutputArgs, opts ...khulnasoft.InvokeOption) ListStorageAccountKeysResultOutput {
	outputResult := khulnasoftx.ApplyErr[*ListStorageAccountKeysArgs](args.ToOutput(), func(plainArgs *ListStorageAccountKeysArgs) (*ListStorageAccountKeysResult, error) {
		return ListStorageAccountKeys(ctx, plainArgs, opts...)
	})

	return khulnasoftx.Cast[ListStorageAccountKeysResultOutput, *ListStorageAccountKeysResult](outputResult)
}

type ListStorageAccountKeysOutputArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName khulnasoftx.Input[string] `khulnasoft:"accountName"`
	// Specifies type of the key to be listed. Possible value is kerb.
	Expand khulnasoftx.Input[*string] `khulnasoft:"expand"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName khulnasoftx.Input[string] `khulnasoft:"resourceGroupName"`
}

func (args ListStorageAccountKeysOutputArgs) ToOutput() khulnasoftx.Output[*ListStorageAccountKeysArgs] {
	allArgs := khulnasoftx.All(
		args.AccountName.ToOutput(context.Background()).AsAny(),
		args.Expand.ToOutput(context.Background()).AsAny(),
		args.ResourceGroupName.ToOutput(context.Background()).AsAny())
	return khulnasoftx.Apply[[]any](allArgs, func(resolvedArgs []interface{}) *ListStorageAccountKeysArgs {
		return &ListStorageAccountKeysArgs{
			AccountName:       resolvedArgs[0].(string),
			Expand:            resolvedArgs[1].(*string),
			ResourceGroupName: resolvedArgs[2].(string),
		}
	})
}

type ListStorageAccountKeysResultOutput struct{ *khulnasoft.OutputState }

func (ListStorageAccountKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStorageAccountKeysResult)(nil)).Elem()
}

func (o ListStorageAccountKeysResultOutput) ToOutput(context.Context) khulnasoftx.Output[*ListStorageAccountKeysResult] {
	return khulnasoftx.Output[*ListStorageAccountKeysResult]{
		OutputState: o.OutputState,
	}
}

func (o ListStorageAccountKeysResultOutput) Keys() khulnasoftx.GArrayOutput[StorageAccountKeyResponse, StorageAccountKeyResponseOutput] {
	value := khulnasoftx.Apply[*ListStorageAccountKeysResult](o, func(v *ListStorageAccountKeysResult) []StorageAccountKeyResponse { return v.Keys })
	return khulnasoftx.GArrayOutput[StorageAccountKeyResponse, StorageAccountKeyResponseOutput]{
		OutputState: value.OutputState,
	}
}