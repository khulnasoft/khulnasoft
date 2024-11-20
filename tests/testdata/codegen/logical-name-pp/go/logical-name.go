package main

import (
	"github.com/khulnasoft/khulnasoft-random/sdk/v4/go/random"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft/config"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		cfg := config.New(ctx, "")
		configLexicalName := cfg.Require("cC-Charlie_charlie.😃⁉️")
		resourceLexicalName, err := random.NewRandomPet(ctx, "aA-Alpha_alpha.🤯⁉️", &random.RandomPetArgs{
			Prefix: khulnasoft.String(configLexicalName),
		})
		if err != nil {
			return err
		}
		ctx.Export("bB-Beta_beta.💜⁉", resourceLexicalName.ID())
		ctx.Export("dD-Delta_delta.🔥⁉", resourceLexicalName.ID())
		return nil
	})
}
