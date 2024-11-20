package main

import (
	"fmt"

	"github.com/khulnasoft/khulnasoft-random/sdk/v4/go/random"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		var numbers []*random.RandomInteger
		for index := 0; index < 2; index++ {
			key0 := index
			val0 := index
			__res, err := random.NewRandomInteger(ctx, fmt.Sprintf("numbers-%v", key0), &random.RandomIntegerArgs{
				Min:  khulnasoft.Int(1),
				Max:  khulnasoft.Int(val0),
				Seed: khulnasoft.Sprintf("seed%v", val0),
			})
			if err != nil {
				return err
			}
			numbers = append(numbers, __res)
		}
		ctx.Export("first", numbers[0].ID())
		ctx.Export("second", numbers[1].ID())
		return nil
	})
}
