package main

import (
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		parent, err := NewRandom(ctx, "parent", &RandomArgs{
			Length: khulnasoft.Int(8),
		})
		if err != nil {
			return err
		}

		child, err := NewRandom(ctx, "child", &RandomArgs{
			Length: khulnasoft.Int(4),
		}, khulnasoft.Parent(parent))
		if err != nil {
			return err
		}

		ctx.Export("parent", parent.Result)
		ctx.Export("child", child.Result)
		return nil
	})
}
