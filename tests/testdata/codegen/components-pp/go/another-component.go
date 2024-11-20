package main

import (
	"fmt"

	"github.com/khulnasoft/khulnasoft-random/sdk/v4/go/random"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

type AnotherComponentArgs struct {
}

type AnotherComponent struct {
	khulnasoft.ResourceState
}

func NewAnotherComponent(
	ctx *khulnasoft.Context,
	name string,
	args *AnotherComponentArgs,
	opts ...khulnasoft.ResourceOption,
) (*AnotherComponent, error) {
	var componentResource AnotherComponent
	err := ctx.RegisterComponentResource("components:index:AnotherComponent", name, &componentResource, opts...)
	if err != nil {
		return nil, err
	}
	_, err = random.NewRandomPassword(ctx, fmt.Sprintf("%s-firstPassword", name), &random.RandomPasswordArgs{
		Length:  khulnasoft.Int(16),
		Special: khulnasoft.Bool(true),
	}, khulnasoft.Parent(&componentResource))
	if err != nil {
		return nil, err
	}
	err = ctx.RegisterResourceOutputs(&componentResource, khulnasoft.Map{})
	if err != nil {
		return nil, err
	}
	return &componentResource, nil
}
