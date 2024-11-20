// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mongodbatlas

import (
	"context"
	"reflect"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"regress-8403/mongodbatlas/internal"
)

func LookupCustomDbRoles(ctx *khulnasoft.Context, args *LookupCustomDbRolesArgs, opts ...khulnasoft.InvokeOption) (*LookupCustomDbRolesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupCustomDbRolesResult
	err := ctx.Invoke("mongodbatlas::getCustomDbRoles", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCustomDbRolesArgs struct {
}

type LookupCustomDbRolesResult struct {
	Result *GetCustomDbRolesResult `khulnasoft:"result"`
}

func LookupCustomDbRolesOutput(ctx *khulnasoft.Context, args LookupCustomDbRolesOutputArgs, opts ...khulnasoft.InvokeOption) LookupCustomDbRolesResultOutput {
	return khulnasoft.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCustomDbRolesResultOutput, error) {
			args := v.(LookupCustomDbRolesArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupCustomDbRolesResult
			secret, err := ctx.InvokePackageRaw("mongodbatlas::getCustomDbRoles", args, &rv, "", opts...)
			if err != nil {
				return LookupCustomDbRolesResultOutput{}, err
			}

			output := khulnasoft.ToOutput(rv).(LookupCustomDbRolesResultOutput)
			if secret {
				return khulnasoft.ToSecret(output).(LookupCustomDbRolesResultOutput), nil
			}
			return output, nil
		}).(LookupCustomDbRolesResultOutput)
}

type LookupCustomDbRolesOutputArgs struct {
}

func (LookupCustomDbRolesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomDbRolesArgs)(nil)).Elem()
}

type LookupCustomDbRolesResultOutput struct{ *khulnasoft.OutputState }

func (LookupCustomDbRolesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomDbRolesResult)(nil)).Elem()
}

func (o LookupCustomDbRolesResultOutput) ToLookupCustomDbRolesResultOutput() LookupCustomDbRolesResultOutput {
	return o
}

func (o LookupCustomDbRolesResultOutput) ToLookupCustomDbRolesResultOutputWithContext(ctx context.Context) LookupCustomDbRolesResultOutput {
	return o
}

func (o LookupCustomDbRolesResultOutput) Result() GetCustomDbRolesResultPtrOutput {
	return o.ApplyT(func(v LookupCustomDbRolesResult) *GetCustomDbRolesResult { return v.Result }).(GetCustomDbRolesResultPtrOutput)
}

func init() {
	khulnasoft.RegisterOutputType(LookupCustomDbRolesResultOutput{})
}
