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
	outputResult := khulnasoftx.ApplyErr[*FuncWithDictParamArgs](args.ToOutput(), func(plainArgs *FuncWithDictParamArgs) (*FuncWithDictParamResult, error) {
		return FuncWithDictParam(ctx, plainArgs, opts...)
	})

	return khulnasoftx.Cast[FuncWithDictParamResultOutput, *FuncWithDictParamResult](outputResult)
}

type FuncWithDictParamOutputArgs struct {
	A khulnasoftx.Input[map[string]string] `khulnasoft:"a"`
	B khulnasoftx.Input[*string]           `khulnasoft:"b"`
}

func (args FuncWithDictParamOutputArgs) ToOutput() khulnasoftx.Output[*FuncWithDictParamArgs] {
	allArgs := khulnasoftx.All(
		args.A.ToOutput(context.Background()).AsAny(),
		args.B.ToOutput(context.Background()).AsAny())
	return khulnasoftx.Apply[[]any](allArgs, func(resolvedArgs []interface{}) *FuncWithDictParamArgs {
		return &FuncWithDictParamArgs{
			A: resolvedArgs[0].(map[string]string),
			B: resolvedArgs[1].(*string),
		}
	})
}

type FuncWithDictParamResultOutput struct{ *khulnasoft.OutputState }

func (FuncWithDictParamResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FuncWithDictParamResult)(nil)).Elem()
}

func (o FuncWithDictParamResultOutput) ToOutput(context.Context) khulnasoftx.Output[*FuncWithDictParamResult] {
	return khulnasoftx.Output[*FuncWithDictParamResult]{
		OutputState: o.OutputState,
	}
}

func (o FuncWithDictParamResultOutput) R() khulnasoftx.Output[string] {
	return khulnasoftx.Apply[*FuncWithDictParamResult](o, func(v *FuncWithDictParamResult) string { return v.R })
}
