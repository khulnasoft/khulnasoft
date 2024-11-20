package main

import (
	"example.com/khulnasoft-fail_on_create/sdk/go/v4/fail_on_create"
	"example.com/khulnasoft-simple/sdk/go/v2/simple"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		failing, err := fail_on_create.NewResource(ctx, "failing", &fail_on_create.ResourceArgs{
			Value: khulnasoft.Bool(false),
		})
		if err != nil {
			return err
		}
		_, err = simple.NewResource(ctx, "dependent", &simple.ResourceArgs{
			Value: khulnasoft.Bool(true),
		}, khulnasoft.DependsOn([]khulnasoft.Resource{
			failing,
		}))
		if err != nil {
			return err
		}
		dependent_on_output, err := simple.NewResource(ctx, "dependent_on_output", &simple.ResourceArgs{
			Value: failing.Value,
		})
		if err != nil {
			return err
		}
		independent, err := simple.NewResource(ctx, "independent", &simple.ResourceArgs{
			Value: khulnasoft.Bool(true),
		})
		if err != nil {
			return err
		}
		_, err = simple.NewResource(ctx, "double_dependency", &simple.ResourceArgs{
			Value: khulnasoft.Bool(true),
		}, khulnasoft.DependsOn([]khulnasoft.Resource{
			independent,
			dependent_on_output,
		}))
		if err != nil {
			return err
		}
		return nil
	})
}
