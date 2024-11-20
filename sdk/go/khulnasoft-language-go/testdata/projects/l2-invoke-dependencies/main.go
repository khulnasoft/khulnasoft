package main

import (
	"example.com/khulnasoft-simple-invoke/sdk/go/v10/simpleinvoke"
	"example.com/khulnasoft-simple/sdk/go/v2/simple"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		first, err := simple.NewResource(ctx, "first", &simple.ResourceArgs{
			Value: khulnasoft.Bool(false),
		})
		if err != nil {
			return err
		}
		// assert that resource second depends on resource first
		// because it uses .secret from the invoke which depends on first
		_, err = simple.NewResource(ctx, "second", &simple.ResourceArgs{
			Value: simpleinvoke.SecretInvokeOutput(ctx, simpleinvoke.SecretInvokeOutputArgs{
				Value:          khulnasoft.String("hello"),
				SecretResponse: first.Value,
			}, nil).ApplyT(func(invoke simpleinvoke.SecretInvokeResult) (bool, error) {
				return invoke.Secret, nil
			}).(khulnasoft.BoolOutput),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
