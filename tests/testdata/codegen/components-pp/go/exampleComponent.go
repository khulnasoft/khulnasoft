package main

import (
	"fmt"

	"github.com/khulnasoft/khulnasoft-random/sdk/v4/go/random"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft/config"
)

type DeploymentZonesArgs struct {
	Zone khulnasoft.StringInput
}

type GithubAppArgs struct {
	Id            khulnasoft.StringInput
	KeyBase64     khulnasoft.StringInput
	WebhookSecret khulnasoft.StringInput
}

type ServersArgs struct {
	Name khulnasoft.StringInput
}

type ExampleComponentArgs struct {
	Input           khulnasoft.StringInput
	CidrBlocks      map[string]khulnasoft.StringInput
	GithubApp       *GithubAppArgs
	Servers         []*ServersArgs
	DeploymentZones map[string]*DeploymentZonesArgs
	IpAddress       []khulnasoft.IntInput
}

type ExampleComponent struct {
	khulnasoft.ResourceState
	Result khulnasoft.AnyOutput
}

func NewExampleComponent(
	ctx *khulnasoft.Context,
	name string,
	args *ExampleComponentArgs,
	opts ...khulnasoft.ResourceOption,
) (*ExampleComponent, error) {
	var componentResource ExampleComponent
	err := ctx.RegisterComponentResource("components:index:ExampleComponent", name, &componentResource, opts...)
	if err != nil {
		return nil, err
	}
	password, err := random.NewRandomPassword(ctx, fmt.Sprintf("%s-password", name), &random.RandomPasswordArgs{
		Length:          khulnasoft.Int(16),
		Special:         khulnasoft.Bool(true),
		OverrideSpecial: args.Input,
	}, khulnasoft.Parent(&componentResource))
	if err != nil {
		return nil, err
	}
	_, err = random.NewRandomPassword(ctx, fmt.Sprintf("%s-githubPassword", name), &random.RandomPasswordArgs{
		Length:          khulnasoft.Int(16),
		Special:         khulnasoft.Bool(true),
		OverrideSpecial: args.GithubApp.WebhookSecret,
	}, khulnasoft.Parent(&componentResource))
	if err != nil {
		return nil, err
	}
	// Example of iterating a list of objects
	var serverPasswords []*random.RandomPassword
	for index := 0; index < len(args.Servers); index++ {
		key0 := index
		val0 := index
		__res, err := random.NewRandomPassword(ctx, fmt.Sprintf("%s-serverPasswords-%v", name, key0), &random.RandomPasswordArgs{
			Length:          khulnasoft.Int(16),
			Special:         khulnasoft.Bool(true),
			OverrideSpecial: args.Servers[val0].Name,
		}, khulnasoft.Parent(&componentResource))
		if err != nil {
			return nil, err
		}
		serverPasswords = append(serverPasswords, __res)
	}
	// Example of iterating a map of objects
	var zonePasswords []*random.RandomPassword
	for key0, val0 := range args.DeploymentZones {
		__res, err := random.NewRandomPassword(ctx, fmt.Sprintf("%s-zonePasswords-%v", name, key0), &random.RandomPasswordArgs{
			Length:          khulnasoft.Int(16),
			Special:         khulnasoft.Bool(true),
			OverrideSpecial: khulnasoft.String(val0),
		}, khulnasoft.Parent(&componentResource))
		if err != nil {
			return nil, err
		}
		zonePasswords = append(zonePasswords, __res)
	}
	_, err = NewSimpleComponent(ctx, fmt.Sprintf("%s-simpleComponent", name), nil, khulnasoft.Parent(&componentResource))
	if err != nil {
		return nil, err
	}
	err = ctx.RegisterResourceOutputs(&componentResource, khulnasoft.Map{
		"result": password.Result,
	})
	if err != nil {
		return nil, err
	}
	componentResource.Result = password.Result
	return &componentResource, nil
}
