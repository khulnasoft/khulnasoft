// Stress-test the engine handling many resources with many aliases.
//go:build !all
// +build !all

package main

import (
	"fmt"

	random "github.com/khulnasoft/khulnasoft-random/sdk/v4/go/random"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft/config"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		conf := config.New(ctx, "")
		mode := conf.Require("mode")
		n := conf.RequireInt("n")

		parent, err := makeResource(ctx, nil, nil, 0, "parent")
		if err != nil {
			return err
		}

		var prev []*random.RandomInteger
		for i := 0; i < n; i++ {
			r, err := makeResource(ctx, parent, prev, i, mode)
			if err != nil {
				return err
			}
			// uncomment for quadratic deps
			// prev = append(prev, r)
			prev = []*random.RandomInteger{r}
		}
		return nil
	})
}

func nameResource(i int, mode string) string {
	return fmt.Sprintf("resource-%s-%d", mode, i)
}

func makeResource(
	ctx *khulnasoft.Context,
	parent *random.RandomInteger,
	prev []*random.RandomInteger,
	i int,
	mode string,
) (*random.RandomInteger, error) {
	name := nameResource(i, mode)
	opts := []khulnasoft.ResourceOption{}
	if len(prev) != 0 {
		deps := []khulnasoft.Resource{}
		for _, p := range prev {
			deps = append(deps, p)
		}
		opts = append(opts, khulnasoft.DependsOn(deps))
	} else if parent != nil {
		opts = append(opts, khulnasoft.DependsOn([]khulnasoft.Resource{parent}))
	}
	if mode == "alias" {
		alias := khulnasoft.Alias{
			Name:     khulnasoft.String(nameResource(i, "new")),
			NoParent: khulnasoft.Bool(true),
		}
		opts = append(opts, khulnasoft.Aliases([]khulnasoft.Alias{alias}))
		opts = append(opts, khulnasoft.Parent(parent))
	}

	if len(prev) != 0 {
		ints := []interface{}{}
		for _, p := range prev {
			ints = append(ints, p.Result)
		}

		var derived khulnasoft.IntOutput = khulnasoft.All(ints...).ApplyT(func(data []interface{}) int {
			s := 10
			return s
		}).(khulnasoft.IntOutput)

		return random.NewRandomInteger(ctx,
			name,
			&random.RandomIntegerArgs{
				Min: derived,
				Max: derived,
			},
			opts...)

	} else {
		return random.NewRandomInteger(ctx,
			name,
			&random.RandomIntegerArgs{
				Min: khulnasoft.Int(0),
				Max: khulnasoft.Int(100),
			},
			opts...)
	}
}
