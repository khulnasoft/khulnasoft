package main

import (
	other "git.example.org/thirdparty/sdk/go/pkg"
	"git.example.org/thirdparty/sdk/go/pkg/module"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		_, err := other.NewThing(ctx, "Other", &other.ThingArgs{
			Idea: khulnasoft.String("Support Third Party"),
		})
		if err != nil {
			return err
		}
		_, err = module.NewObject(ctx, "Question", &module.ObjectArgs{
			Answer: khulnasoft.Float64(42),
		})
		if err != nil {
			return err
		}
		_, err = module.NewObject(ctx, "Question2", &module.ObjectArgs{
			Answer: khulnasoft.Float64(24),
		})
		if err != nil {
			return err
		}
		_, err = other.NewProvider(ctx, "Provider", &other.ProviderArgs{
			ObjectProp: khulnasoft.StringMap{
				"prop1": khulnasoft.String("foo"),
				"prop2": khulnasoft.String("bar"),
				"prop3": khulnasoft.String("fizz"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
