// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"fmt"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		r, err := NewRandom(ctx, "resource", &RandomArgs{Length: khulnasoft.Int(10)})
		if err != nil {
			return err
		}
		_, err = NewComponent(ctx, "component", &ComponentArgs{
			Message: r.ID().ApplyT(func(id khulnasoft.ID) string {
				return fmt.Sprintf("message %v", id)
			}).(khulnasoft.StringOutput),
			Nested: &ComponentNestedArgs{
				Value: r.ID().ApplyT(func(id khulnasoft.ID) string {
					return fmt.Sprintf("nested.value %v", id)
				}).(khulnasoft.StringOutput),
			},
		})
		return err
	})
}
