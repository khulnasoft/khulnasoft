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

// Check codegen of functions with all optional inputs.
func FuncWithAllOptionalInputs(ctx *khulnasoft.Context, args *FuncWithAllOptionalInputsArgs, opts ...khulnasoft.InvokeOption) (*FuncWithAllOptionalInputsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv FuncWithAllOptionalInputsResult
	err := ctx.Invoke("mypkg::funcWithAllOptionalInputs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type FuncWithAllOptionalInputsArgs struct {
	// Property A
	A *string `khulnasoft:"a"`
	// Property B
	B *string `khulnasoft:"b"`
}

type FuncWithAllOptionalInputsResult struct {
	R string `khulnasoft:"r"`
}

func FuncWithAllOptionalInputsOutput(ctx *khulnasoft.Context, args FuncWithAllOptionalInputsOutputArgs, opts ...khulnasoft.InvokeOption) FuncWithAllOptionalInputsResultOutput {
	return khulnasoft.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (FuncWithAllOptionalInputsResultOutput, error) {
			args := v.(FuncWithAllOptionalInputsArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv FuncWithAllOptionalInputsResult
			secret, err := ctx.InvokePackageRaw("mypkg::funcWithAllOptionalInputs", args, &rv, "", opts...)
			if err != nil {
				return FuncWithAllOptionalInputsResultOutput{}, err
			}

			output := khulnasoft.ToOutput(rv).(FuncWithAllOptionalInputsResultOutput)
			if secret {
				return khulnasoft.ToSecret(output).(FuncWithAllOptionalInputsResultOutput), nil
			}
			return output, nil
		}).(FuncWithAllOptionalInputsResultOutput)
}

type FuncWithAllOptionalInputsOutputArgs struct {
	// Property A
	A khulnasoft.StringPtrInput `khulnasoft:"a"`
	// Property B
	B khulnasoft.StringPtrInput `khulnasoft:"b"`
}

func (FuncWithAllOptionalInputsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FuncWithAllOptionalInputsArgs)(nil)).Elem()
}

type FuncWithAllOptionalInputsResultOutput struct{ *khulnasoft.OutputState }

func (FuncWithAllOptionalInputsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FuncWithAllOptionalInputsResult)(nil)).Elem()
}

func (o FuncWithAllOptionalInputsResultOutput) ToFuncWithAllOptionalInputsResultOutput() FuncWithAllOptionalInputsResultOutput {
	return o
}

func (o FuncWithAllOptionalInputsResultOutput) ToFuncWithAllOptionalInputsResultOutputWithContext(ctx context.Context) FuncWithAllOptionalInputsResultOutput {
	return o
}

func (o FuncWithAllOptionalInputsResultOutput) ToOutput(ctx context.Context) khulnasoftx.Output[FuncWithAllOptionalInputsResult] {
	return khulnasoftx.Output[FuncWithAllOptionalInputsResult]{
		OutputState: o.OutputState,
	}
}

func (o FuncWithAllOptionalInputsResultOutput) R() khulnasoft.StringOutput {
	return o.ApplyT(func(v FuncWithAllOptionalInputsResult) string { return v.R }).(khulnasoft.StringOutput)
}

func init() {
	khulnasoft.RegisterOutputType(FuncWithAllOptionalInputsResultOutput{})
}