// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		rand, err := NewRandom(ctx, "random", &RandomArgs{Length: khulnasoft.Int(10)})
		if err != nil {
			return err
		}

		_, err = NewFailsOnDelete(ctx, "failsondelete", khulnasoft.DeletedWith(rand))
		if err != nil {
			return err
		}

		return nil
	})
}
