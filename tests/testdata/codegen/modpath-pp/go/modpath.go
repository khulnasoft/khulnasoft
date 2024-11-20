package main

import (
	other "git.example.org/thirdparty/sdk/go/pkg"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		_, err := other.NewThing(ctx, "thing", &other.ThingArgs{
			Idea: khulnasoft.String("myIdea"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
