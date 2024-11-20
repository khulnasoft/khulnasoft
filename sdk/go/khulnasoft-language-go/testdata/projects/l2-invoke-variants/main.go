package main

import (
	"example.com/khulnasoft-simple-invoke/sdk/go/v10/simpleinvoke"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		res, err := simpleinvoke.NewStringResource(ctx, "res", nil)
		if err != nil {
			return err
		}
		ctx.Export("outputInput", simpleinvoke.MyInvokeOutput(ctx, simpleinvoke.MyInvokeOutputArgs{
			Value: res.Text,
		}, nil).ApplyT(func(invoke simpleinvoke.MyInvokeResult) (string, error) {
			return invoke.Result, nil
		}).(khulnasoft.StringOutput))
		ctx.Export("unit", simpleinvoke.UnitOutput(ctx, simpleinvoke.UnitOutputArgs{}, nil).ApplyT(func(invoke simpleinvoke.UnitResult) (string, error) {
			return invoke.Result, nil
		}).(khulnasoft.StringOutput))
		return nil
	})
}
