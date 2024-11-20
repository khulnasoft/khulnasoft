package main

import (
	"github.com/khulnasoft/khulnasoft-random/sdk/v4/go/random"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		_, err := random.NewRandomPet(ctx, "random-pet", &random.RandomPetArgs{
			Prefix: khulnasoft.String("doggo"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
