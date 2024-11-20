package main

import (
	"example.com/khulnasoft-simple-invoke/sdk/go/v10/simpleinvoke"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		ctx.Export("hello", simpleinvoke.MyInvokeOutput(ctx, simpleinvoke.MyInvokeOutputArgs{
			Value: khulnasoft.String("hello"),
		}, nil).ApplyT(func(invoke simpleinvoke.MyInvokeResult) (string, error) {
			return invoke.Result, nil
		}).(khulnasoft.StringOutput))
		ctx.Export("goodbye", simpleinvoke.MyInvokeOutput(ctx, simpleinvoke.MyInvokeOutputArgs{
			Value: khulnasoft.String("goodbye"),
		}, nil).ApplyT(func(invoke simpleinvoke.MyInvokeResult) (string, error) {
			return invoke.Result, nil
		}).(khulnasoft.StringOutput))
		return nil
	})
}
