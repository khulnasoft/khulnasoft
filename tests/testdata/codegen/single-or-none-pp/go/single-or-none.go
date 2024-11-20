package main

import (
	"fmt"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func singleOrNone[T any](elements []T) T {
	if len(elements) != 1 {
		panic(fmt.Errorf("singleOrNone expected input slice to have a single element"))
	}
	return elements[0]
}

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		ctx.Export("result", khulnasoft.Float64(singleOrNone([]float64{
			1,
		})))
		return nil
	})
}
