//go:build !all
// +build !all

package main

import (
	"reflect"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		newComponent := func(name string, opts ...khulnasoft.ResourceOption) (*Component, error) {
			return NewComponent(ctx, name, &ComponentArgs{Echo: khulnasoft.String(name)}, opts...)
		}

		dep1, err := newComponent("Dep1")
		if err != nil {
			return err
		}
		dep2, err := newComponent("Dep2")
		if err != nil {
			return err
		}
		_, err = newComponent("DependsOn", khulnasoft.DependsOn([]khulnasoft.Resource{dep1, dep2}))
		if err != nil {
			return err
		}

		_, err = newComponent("Protect", khulnasoft.Protect(true))
		if err != nil {
			return err
		}

		_, err = newComponent("AdditionalSecretOutputs", khulnasoft.AdditionalSecretOutputs([]string{"foo"}))
		if err != nil {
			return err
		}

		_, err = newComponent("CustomTimeouts", khulnasoft.Timeouts(&khulnasoft.CustomTimeouts{
			Create: ("1m"),
			Update: ("2m"),
			Delete: ("3m"),
		}))
		if err != nil {
			return err
		}

		getDeletedWithMe, err := newComponent("getDeletedWithMe")
		if err != nil {
			return err
		}
		_, err = newComponent("DeletedWith", khulnasoft.DeletedWith(getDeletedWithMe))
		if err != nil {
			return err
		}

		_, err = newComponent("RetainOnDelete", khulnasoft.RetainOnDelete(true))
		if err != nil {
			return err
		}

		return nil
	})
}

type Component struct {
	khulnasoft.ResourceState

	Echo khulnasoft.StringOutput `khulnasoft:"echo"`
	Foo  khulnasoft.StringOutput `khulnasoft:"foo"`
	Bar  khulnasoft.StringOutput `khulnasoft:"bar"`
}

func NewComponent(ctx *khulnasoft.Context, name string, args *ComponentArgs, opts ...khulnasoft.ResourceOption) (*Component, error) {
	var resource Component
	err := ctx.RegisterRemoteComponentResource("testcomponent:index:Component", name, args, &resource, opts...)
	return &resource, err
}

type ComponentArgs struct {
	Echo khulnasoft.StringInput
}

func (ComponentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentArgs)(nil)).Elem()
}

type componentArgs struct {
	Echo string `khulnasoft:"echo"`
}
