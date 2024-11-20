package main

import (
	usingdashes "example.com/khulnasoft-using-dashes/sdk/go/using-dashes"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		_, err := usingdashes.NewDash(ctx, "main", &usingdashes.DashArgs{
			Stack: khulnasoft.String("dev"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
