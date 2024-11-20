package main

import (
	"example.com/khulnasoft-simple-invoke/sdk/go/v10/simpleinvoke"
	"example.com/khulnasoft-simple/sdk/go/v2/simple"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		res, err := simple.NewResource(ctx, "res", &simple.ResourceArgs{
			Value: khulnasoft.Bool(true),
		})
		if err != nil {
			return err
		}
		ctx.Export("nonSecret", simpleinvoke.SecretInvokeOutput(ctx, simpleinvoke.SecretInvokeOutputArgs{
			Value:          khulnasoft.String("hello"),
			SecretResponse: khulnasoft.Bool(false),
		}, nil).ApplyT(func(invoke simpleinvoke.SecretInvokeResult) (string, error) {
			return invoke.Response, nil
		}).(khulnasoft.StringOutput))
		ctx.Export("firstSecret", simpleinvoke.SecretInvokeOutput(ctx, simpleinvoke.SecretInvokeOutputArgs{
			Value:          khulnasoft.String("hello"),
			SecretResponse: res.Value,
		}, nil).ApplyT(func(invoke simpleinvoke.SecretInvokeResult) (string, error) {
			return invoke.Response, nil
		}).(khulnasoft.StringOutput))
		ctx.Export("secondSecret", simpleinvoke.SecretInvokeOutput(ctx, simpleinvoke.SecretInvokeOutputArgs{
			Value:          khulnasoft.ToSecret("goodbye").(khulnasoft.StringOutput),
			SecretResponse: khulnasoft.Bool(false),
		}, nil).ApplyT(func(invoke simpleinvoke.SecretInvokeResult) (string, error) {
			return invoke.Response, nil
		}).(khulnasoft.StringOutput))
		return nil
	})
}
