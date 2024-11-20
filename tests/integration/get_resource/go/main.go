//go:build !all
// +build !all

package main

import (
	"reflect"

	"github.com/khulnasoft/khulnasoft-random/sdk/v4/go/random"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft/config"
)

type MyResource struct {
	khulnasoft.ResourceState

	Length khulnasoft.IntOutput       `khulnasoft:"length"`
	Prefix khulnasoft.StringPtrOutput `khulnasoft:"prefix"`
}

type (
	myResourceArgs struct{}
	MyResourceArgs struct{}
)

func (MyResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*myResourceArgs)(nil)).Elem()
}

func GetResource(ctx *khulnasoft.Context, urn khulnasoft.URN) (*MyResource, error) {
	var resource MyResource
	err := ctx.RegisterResource("unused:unused:unused", "unused", &MyResourceArgs{}, &resource,
		khulnasoft.URN_(string(urn)))
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		c := config.New(ctx, "")
		bar := c.RequireSecret("bar")
		pet, err := random.NewRandomPet(ctx, "cat", &random.RandomPetArgs{
			Length: khulnasoft.Int(2),
			Prefix: bar,
		})
		if err != nil {
			return err
		}

		getPetLength := pet.URN().ApplyT(func(urn khulnasoft.URN) (khulnasoft.IntInput, error) {
			r, err := GetResource(ctx, urn)
			if err != nil {
				return nil, err
			}
			return r.Length, nil
		})
		getPetSecret := pet.URN().ApplyT(func(urn khulnasoft.URN) (khulnasoft.StringPtrInput, error) {
			r, err := GetResource(ctx, urn)
			if err != nil {
				return nil, err
			}
			return r.Prefix, nil
		})
		ctx.Export("getPetLength", getPetLength)
		ctx.Export("secret", getPetSecret)

		return nil
	})
}
