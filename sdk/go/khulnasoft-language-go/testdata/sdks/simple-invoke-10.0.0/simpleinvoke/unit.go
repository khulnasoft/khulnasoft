// Code generated by khulnasoft-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package simpleinvoke

import (
	"context"
	"reflect"

	"example.com/khulnasoft-simple-invoke/sdk/go/v10/simpleinvoke/internal"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func Unit(ctx *khulnasoft.Context, args *UnitArgs, opts ...khulnasoft.InvokeOption) (*UnitResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv UnitResult
	err := ctx.Invoke("simple-invoke:index:unit", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type UnitArgs struct {
}

type UnitResult struct {
	Result string `khulnasoft:"result"`
}

func UnitOutput(ctx *khulnasoft.Context, args UnitOutputArgs, opts ...khulnasoft.InvokeOption) UnitResultOutput {
	return khulnasoft.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (UnitResultOutput, error) {
			args := v.(UnitArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv UnitResult
			secret, err := ctx.InvokePackageRaw("simple-invoke:index:unit", args, &rv, "", opts...)
			if err != nil {
				return UnitResultOutput{}, err
			}

			output := khulnasoft.ToOutput(rv).(UnitResultOutput)
			if secret {
				return khulnasoft.ToSecret(output).(UnitResultOutput), nil
			}
			return output, nil
		}).(UnitResultOutput)
}

type UnitOutputArgs struct {
}

func (UnitOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UnitArgs)(nil)).Elem()
}

type UnitResultOutput struct{ *khulnasoft.OutputState }

func (UnitResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UnitResult)(nil)).Elem()
}

func (o UnitResultOutput) ToUnitResultOutput() UnitResultOutput {
	return o
}

func (o UnitResultOutput) ToUnitResultOutputWithContext(ctx context.Context) UnitResultOutput {
	return o
}

func (o UnitResultOutput) Result() khulnasoft.StringOutput {
	return o.ApplyT(func(v UnitResult) string { return v.Result }).(khulnasoft.StringOutput)
}

func init() {
	khulnasoft.RegisterOutputType(UnitResultOutput{})
}
