//go:build !all
// +build !all

package main

import (
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		_, err := NewComponent(ctx, "foo")
		return err
	})
}

type Component struct {
	khulnasoft.ResourceState
}

func NewComponent(ctx *khulnasoft.Context, name string, opts ...khulnasoft.ResourceOption) (*Component, error) {
	var component Component
	err := ctx.RegisterRemoteComponentResource("testcomponent:index:Component", name, nil, &component, opts...)
	return &component, err
}
