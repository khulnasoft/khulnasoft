// Copyright 2016-2022, Pulumi Corporation.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/blang/semver"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft/internals"
)

type Child struct {
	khulnasoft.ResourceState

	Message khulnasoft.StringOutput `khulnasoft:"message"`
}

func NewChild(ctx *khulnasoft.Context, name string, message khulnasoft.StringInput,
	opts ...khulnasoft.ResourceOption,
) (*Child, error) {
	component := &Child{}
	if err := ctx.RegisterComponentResource("test:index:Child", name, component, opts...); err != nil {
		return nil, err
	}
	if err := ctx.RegisterResourceOutputs(component, khulnasoft.Map{
		"message": message,
	}); err != nil {
		return nil, err
	}
	// Wait to make sure RegisterResourceOutputs has actually finished registering the resource outputs.
	//
	// See also the comment in NewContainer below for a more thorough explanation.
	//
	// TODO: make RegisterResourceOutputs not racy [khulnasoft/khulnasoft#16896]
	componentURNResult, err := internals.UnsafeAwaitOutput(ctx.Context(), component.URN())
	if err != nil {
		return nil, err
	}
	componentURN := componentURNResult.Value.(khulnasoft.URN)
	sleep := 20 * time.Millisecond
	for i := 0; ; i++ {
		roundTrippedComponent := &Child{}

		if err := ctx.RegisterComponentResource("test:index:Child", "mychild", roundTrippedComponent,
			khulnasoft.URN_(string(componentURN))); err != nil {
			return nil, err
		}
		message, err := internals.UnsafeAwaitOutput(ctx.Context(), roundTrippedComponent.Message)
		if err != nil {
			return nil, err
		}
		if message.Value.(string) != "" {
			break
		} else if i > 10 {
			return nil, errors.New("outputs were not registered successfully")
		}
		time.Sleep(sleep)
		sleep *= 2
	}
	return component, nil
}

type ChildInput interface {
	khulnasoft.Input
}

func (*Child) ElementType() reflect.Type {
	return reflect.TypeOf((**Child)(nil)).Elem()
}

type ChildOutput struct{ *khulnasoft.OutputState }

func (ChildOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Child)(nil)).Elem()
}

type Container struct {
	khulnasoft.ResourceState

	Child ChildOutput `khulnasoft:"child"`
}

type CustomContainer struct {
	khulnasoft.CustomResource

	Child ChildOutput `khulnasoft:"child"`
}

func NewContainer(ctx *khulnasoft.Context, name string, child ChildInput,
	opts ...khulnasoft.ResourceOption,
) (*Container, error) {
	component := &Container{}
	if err := ctx.RegisterComponentResource("test:index:Container", name, component, opts...); err != nil {
		return nil, err
	}
	if err := ctx.RegisterResourceOutputs(component, khulnasoft.Map{
		"child": child,
	}); err != nil {
		return nil, err
	}
	// Wait to make sure RegisterResourceOutputs has actually finished registering the resource outputs.
	//
	// RegisterResourceOutputs does most of its work in a a goroutine, as does RegisterComponentResource.  This
	// means RegisterResourceOutputs is inheritly racy with the resource being read later.  This test explicitly
	// tests roundtripping a container component resource, which means we need to read the outputs registered
	// through RegisterResourceOutputs later, making the test racy.  We can work around this by making sure the
	// outputs are registered before we return the container.  Ideally we should find a way to make this non-racy
	// (see the issue linked below)
	//
	// TODO: make RegisterResourceOutputs not racy [khulnasoft/khulnasoft#16896]
	containerURNResult, err := internals.UnsafeAwaitOutput(ctx.Context(), component.URN())
	if err != nil {
		return nil, err
	}
	containerURN := containerURNResult.Value.(khulnasoft.URN)
	sleep := 20 * time.Millisecond
	for i := 0; ; i++ {
		roundTrippedContainer := &Container{}

		if err := ctx.RegisterComponentResource("test:index:Container", "mycontainer", roundTrippedContainer,
			khulnasoft.URN_(string(containerURN))); err != nil {
			return nil, err
		}
		child, err := internals.UnsafeAwaitOutput(ctx.Context(), roundTrippedContainer.Child)
		if err != nil {
			return nil, err
		}
		if child.Value.(*Child) != nil {
			break
		} else if i > 10 {
			return nil, errors.New("outputs were not registered successfully")
		}
		time.Sleep(sleep)
		sleep *= 2
	}
	return component, nil
}

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *khulnasoft.Context, name, typ, urn string) (r khulnasoft.Resource, err error) {
	switch typ {
	case "test:index:Child":
		r = &Child{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}
	err = ctx.RegisterResource(typ, name, nil, r, khulnasoft.URN_(urn))
	return
}

func main() {
	khulnasoft.RegisterOutputType(ChildOutput{})
	khulnasoft.RegisterResourceModule(
		"test",
		"index",
		&module{semver.MustParse("1.0.0")},
	)

	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		child, err := NewChild(ctx, "mychild", khulnasoft.String("hello world!"))
		if err != nil {
			return err
		}

		container, err := NewContainer(ctx, "mycontainer", child)
		if err != nil {
			return err
		}

		containerURNResult, err := internals.UnsafeAwaitOutput(ctx.Context(), container.URN())
		if err != nil {
			return err
		}
		containerURN := containerURNResult.Value.(khulnasoft.URN)

		roundTrippedContainer := &Container{}
		if err := ctx.RegisterComponentResource("test:index:Container", "mycontainer", roundTrippedContainer,
			khulnasoft.URN_(string(containerURN))); err != nil {
			return err
		}

		roundTrippedContainerChildResult, err := internals.UnsafeAwaitOutput(ctx.Context(), roundTrippedContainer.Child)
		if err != nil {
			return err
		}
		roundTrippedContainerChild := roundTrippedContainerChildResult.Value.(*Child)

		khulnasoft.All(child.URN(), roundTrippedContainerChild.URN(), roundTrippedContainerChild.Message).ApplyT(
			func(args []interface{}) (*string, error) {
				const expectedMessage = "hello world!"
				expectedURN := args[0].(khulnasoft.URN)
				actualURN := args[1].(khulnasoft.URN)
				actualMessage := args[2].(string)
				if expectedURN != actualURN {
					panic(fmt.Errorf("expected urn %q not equal to actual urn %q", expectedURN, actualMessage))
				}
				if expectedMessage != actualMessage {
					panic(fmt.Errorf(
						"expected message %q not equal to actual message %q", expectedMessage, actualMessage))
				}
				return nil, nil
			})

		return nil
	})
}
