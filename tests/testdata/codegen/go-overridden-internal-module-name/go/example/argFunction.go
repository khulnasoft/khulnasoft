// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"context"
	"reflect"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"go-overridden-internal-module-name/example/utilities"
)

func ArgFunction(ctx *khulnasoft.Context, args *ArgFunctionArgs, opts ...khulnasoft.InvokeOption) (*ArgFunctionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ArgFunctionResult
	err := ctx.Invoke("example::argFunction", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ArgFunctionArgs struct {
	Arg1 *Resource `khulnasoft:"arg1"`
}

type ArgFunctionResult struct {
	Result *Resource `khulnasoft:"result"`
}

func ArgFunctionOutput(ctx *khulnasoft.Context, args ArgFunctionOutputArgs, opts ...khulnasoft.InvokeOption) ArgFunctionResultOutput {
	return khulnasoft.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ArgFunctionResultOutput, error) {
			args := v.(ArgFunctionArgs)
			opts = utilities.PkgInvokeDefaultOpts(opts)
			var rv ArgFunctionResult
			secret, err := ctx.InvokePackageRaw("example::argFunction", args, &rv, "", opts...)
			if err != nil {
				return ArgFunctionResultOutput{}, err
			}

			output := khulnasoft.ToOutput(rv).(ArgFunctionResultOutput)
			if secret {
				return khulnasoft.ToSecret(output).(ArgFunctionResultOutput), nil
			}
			return output, nil
		}).(ArgFunctionResultOutput)
}

type ArgFunctionOutputArgs struct {
	Arg1 ResourceInput `khulnasoft:"arg1"`
}

func (ArgFunctionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ArgFunctionArgs)(nil)).Elem()
}

type ArgFunctionResultOutput struct{ *khulnasoft.OutputState }

func (ArgFunctionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArgFunctionResult)(nil)).Elem()
}

func (o ArgFunctionResultOutput) ToArgFunctionResultOutput() ArgFunctionResultOutput {
	return o
}

func (o ArgFunctionResultOutput) ToArgFunctionResultOutputWithContext(ctx context.Context) ArgFunctionResultOutput {
	return o
}

func (o ArgFunctionResultOutput) Result() ResourceOutput {
	return o.ApplyT(func(v ArgFunctionResult) *Resource { return v.Result }).(ResourceOutput)
}

func init() {
	khulnasoft.RegisterOutputType(ArgFunctionResultOutput{})
}
