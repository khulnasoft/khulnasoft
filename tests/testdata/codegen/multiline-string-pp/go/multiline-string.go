package main

import (
	"github.com/khulnasoft/khulnasoft-random/sdk/v4/go/random"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		_, err := random.NewRandomShuffle(ctx, "foo", &random.RandomShuffleArgs{
			Inputs: khulnasoft.StringArray{
				khulnasoft.String("just one\nnewline"),
				khulnasoft.String("foo\nbar\nbaz\nqux\nquux\nqux"),
				khulnasoft.String(`{
    "a": 1,
    "b": 2,
    "c": [
      "foo",
      "bar",
      "baz",
      "qux",
      "quux"
    ]
}
`),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
