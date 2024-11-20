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
	return khulnasoft.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListStorageAccountKeysResultOutput, error) {
			args := v.(ListStorageAccountKeysArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv ListStorageAccountKeysResult
			secret, err := ctx.InvokePackageRaw("mypkg::listStorageAccountKeys", args, &rv, "", opts...)
			if err != nil {
				return ListStorageAccountKeysResultOutput{}, err
			}

			output := khulnasoft.ToOutput(rv).(ListStorageAccountKeysResultOutput)
			if secret {
				return khulnasoft.ToSecret(output).(ListStorageAccountKeysResultOutput), nil
			}
			return output, nil
		}).(ListStorageAccountKeysResultOutput)
}

type ListStorageAccountKeysOutputArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName khulnasoft.StringInput `khulnasoft:"accountName"`
	// Specifies type of the key to be listed. Possible value is kerb.
	Expand khulnasoft.StringPtrInput `khulnasoft:"expand"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName khulnasoft.StringInput `khulnasoft:"resourceGroupName"`
}

func (ListStorageAccountKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStorageAccountKeysArgs)(nil)).Elem()
}

// The response from the ListKeys operation.
type ListStorageAccountKeysResultOutput struct{ *khulnasoft.OutputState }

func (ListStorageAccountKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStorageAccountKeysResult)(nil)).Elem()
}

func (o ListStorageAccountKeysResultOutput) ToListStorageAccountKeysResultOutput() ListStorageAccountKeysResultOutput {
	return o
}

func (o ListStorageAccountKeysResultOutput) ToListStorageAccountKeysResultOutputWithContext(ctx context.Context) ListStorageAccountKeysResultOutput {
	return o
}

func (o ListStorageAccountKeysResultOutput) ToOutput(ctx context.Context) khulnasoftx.Output[ListStorageAccountKeysResult] {
	return khulnasoftx.Output[ListStorageAccountKeysResult]{
		OutputState: o.OutputState,
	}
}

// Gets the list of storage account keys and their properties for the specified storage account.
func (o ListStorageAccountKeysResultOutput) Keys() StorageAccountKeyResponseArrayOutput {
	return o.ApplyT(func(v ListStorageAccountKeysResult) []StorageAccountKeyResponse { return v.Keys }).(StorageAccountKeyResponseArrayOutput)
}

func init() {
	khulnasoft.RegisterOutputType(ListStorageAccountKeysResultOutput{})
}