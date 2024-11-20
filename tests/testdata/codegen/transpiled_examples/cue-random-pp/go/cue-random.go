package main

import (
	"github.com/khulnasoft/khulnasoft-random/sdk/v4/go/random"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		randomPassword, err := random.NewRandomPassword(ctx, "randomPassword", &random.RandomPasswordArgs{
			Length:          khulnasoft.Int(16),
			Special:         khulnasoft.Bool(true),
			OverrideSpecial: khulnasoft.String("_%@"),
		})
		if err != nil {
			return err
		}
		ctx.Export("password", randomPassword.Result)
		return nil
	})
}
