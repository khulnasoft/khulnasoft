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

// Check codegen of functions with a List parameter.
func FuncWithListParam(ctx *khulnasoft.Context, args *FuncWithListParamArgs, opts ...khulnasoft.InvokeOption) (*FuncWithListParamResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv FuncWithListParamResult
	err := ctx.Invoke("mypkg::funcWithListParam", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type FuncWithListParamArgs struct {
	A []string `khulnasoft:"a"`
	B *string  `khulnasoft:"b"`
}

type FuncWithListParamResult struct {
	R string `khulnasoft:"r"`
}

func FuncWithListParamOutput(ctx *khulnasoft.Context, args FuncWithListParamOutputArgs, opts ...khulnasoft.InvokeOption) FuncWithListParamResultOutput {
	outputResult := khulnasoftx.ApplyErr[*FuncWithListParamArgs](args.ToOutput(), func(plainArgs *FuncWithListParamArgs) (*FuncWithListParamResult, error) {
		return FuncWithListParam(ctx, plainArgs, opts...)
	})

	return khulnasoftx.Cast[FuncWithListParamResultOutput, *FuncWithListParamResult](outputResult)
}

type FuncWithListParamOutputArgs struct {
	A khulnasoftx.Input[[]string] `khulnasoft:"a"`
	B khulnasoftx.Input[*string]  `khulnasoft:"b"`
}

func (args FuncWithListParamOutputArgs) ToOutput() khulnasoftx.Output[*FuncWithListParamArgs] {
	allArgs := khulnasoftx.All(
		args.A.ToOutput(context.Background()).AsAny(),
		args.B.ToOutput(context.Background()).AsAny())
	return khulnasoftx.Apply[[]any](allArgs, func(resolvedArgs []interface{}) *FuncWithListParamArgs {
		return &FuncWithListParamArgs{
			A: resolvedArgs[0].([]string),
			B: resolvedArgs[1].(*string),
		}
	})
}

type FuncWithListParamResultOutput struct{ *khulnasoft.OutputState }

func (FuncWithListParamResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FuncWithListParamResult)(nil)).Elem()
}

func (o FuncWithListParamResultOutput) ToOutput(context.Context) khulnasoftx.Output[*FuncWithListParamResult] {
	return khulnasoftx.Output[*FuncWithListParamResult]{
		OutputState: o.OutputState,
	}
}

func (o FuncWithListParamResultOutput) R() khulnasoftx.Output[string] {
	return khulnasoftx.Apply[*FuncWithListParamResult](o, func(v *FuncWithListParamResult) string { return v.R })
}
