package main

import (
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		_, err := khulnasoft.NewStackReference(ctx, "stackRef", &khulnasoft.StackReferenceArgs{
			Name: khulnasoft.String("foo/bar/dev"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
