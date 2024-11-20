package main

import (
	"fmt"

	"github.com/khulnasoft/khulnasoft-random/sdk/v4/go/random"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

type SimpleComponentArgs struct {
}

type SimpleComponent struct {
	khulnasoft.ResourceState
}

func NewSimpleComponent(
	ctx *khulnasoft.Context,
	name string,
	args *SimpleComponentArgs,
	opts ...khulnasoft.ResourceOption,
) (*SimpleComponent, error) {
	var componentResource SimpleComponent
	err := ctx.RegisterComponentResource("components:index:SimpleComponent", name, &componentResource, opts...)
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
	_, err = random.NewRandomPassword(ctx, fmt.Sprintf("%s-secondPassword", name), &random.RandomPasswordArgs{
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
