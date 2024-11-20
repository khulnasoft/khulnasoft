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

// Check codegen of functions with a Dict<str,str> parameter.
func FuncWithDictParam(ctx *khulnasoft.Context, args *FuncWithDictParamArgs, opts ...khulnasoft.InvokeOption) (*FuncWithDictParamResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv FuncWithDictParamResult
	err := ctx.Invoke("mypkg::funcWithDictParam", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type FuncWithDictParamArgs struct {
	A map[string]string `khulnasoft:"a"`
	B *string           `khulnasoft:"b"`
}

type FuncWithDictParamResult struct {
	R string `khulnasoft:"r"`
}

func FuncWithDictParamOutput(ctx *khulnasoft.Context, args FuncWithDictParamOutputArgs, opts ...khulnasoft.InvokeOption) FuncWithDictParamResultOutput {
	return khulnasoft.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (FuncWithDictParamResultOutput, error) {
			args := v.(FuncWithDictParamArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv FuncWithDictParamResult
			secret, err := ctx.InvokePackageRaw("mypkg::funcWithDictParam", args, &rv, "", opts...)
			if err != nil {
				return FuncWithDictParamResultOutput{}, err
			}

			output := khulnasoft.ToOutput(rv).(FuncWithDictParamResultOutput)
			if secret {
				return khulnasoft.ToSecret(output).(FuncWithDictParamResultOutput), nil
			}
			return output, nil
		}).(FuncWithDictParamResultOutput)
}

type FuncWithDictParamOutputArgs struct {
	A khulnasoft.StringMapInput `khulnasoft:"a"`
	B khulnasoft.StringPtrInput `khulnasoft:"b"`
}

func (FuncWithDictParamOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FuncWithDictParamArgs)(nil)).Elem()
}

type FuncWithDictParamResultOutput struct{ *khulnasoft.OutputState }

func (FuncWithDictParamResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FuncWithDictParamResult)(nil)).Elem()
}

func (o FuncWithDictParamResultOutput) ToFuncWithDictParamResultOutput() FuncWithDictParamResultOutput {
	return o
}

func (o FuncWithDictParamResultOutput) ToFuncWithDictParamResultOutputWithContext(ctx context.Context) FuncWithDictParamResultOutput {
	return o
}

func (o FuncWithDictParamResultOutput) ToOutput(ctx context.Context) khulnasoftx.Output[FuncWithDictParamResult] {
	return khulnasoftx.Output[FuncWithDictParamResult]{
		OutputState: o.OutputState,
	}
}

func (o FuncWithDictParamResultOutput) R() khulnasoft.StringOutput {
	return o.ApplyT(func(v FuncWithDictParamResult) string { return v.R }).(khulnasoft.StringOutput)
}

func init() {
	khulnasoft.RegisterOutputType(FuncWithDictParamResultOutput{})
}