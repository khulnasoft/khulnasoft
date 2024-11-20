// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"fmt"
	"os"

	"github.com/khulnasoft/khulnasoft/pkg/v3/resource/provider"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/common/util/cmdutil"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	khulnasoftprovider "github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft/provider"
)

const (
	providerName = "testcomponent"
	version      = "0.0.1"
)

func main() {
	var schema string
	if _, ok := os.LookupEnv("INCLUDE_SCHEMA"); ok {
		schema = `{"hello": "world"}`
	}
	err := provider.MainWithOptions(provider.Options{
		Name:    providerName,
		Version: version,
		Schema:  []byte(schema),
		Construct: func(ctx *khulnasoft.Context, typ, name string,
			inputs khulnasoftprovider.ConstructInputs, options khulnasoft.ResourceOption,
		) (*khulnasoftprovider.ConstructResult, error) {
			return nil, fmt.Errorf("unknown resource type %s", typ)
		},
	})
	if err != nil {
		cmdutil.ExitError(err.Error())
	}
}
