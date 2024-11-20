// Copyright 2016-2021, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//go:build !all
// +build !all

package main

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

type Random struct {
	khulnasoft.CustomResourceState

	Length khulnasoft.IntOutput    `khulnasoft:"length"`
	Result khulnasoft.StringOutput `khulnasoft:"result"`
}

func NewRandom(ctx *khulnasoft.Context,
	name string, args *RandomArgs, opts ...khulnasoft.ResourceOption,
) (*Random, error) {
	if args == nil || args.Length == nil {
		return nil, errors.New("missing required argument 'Length'")
	}
	var resource Random
	err := ctx.RegisterResource("testprovider:index:Random", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type randomArgs struct {
	Length int    `khulnasoft:"length"`
	Prefix string `khulnasoft:"prefix"`
}

type RandomArgs struct {
	Length khulnasoft.IntInput
	Prefix khulnasoft.StringInput
}

func (RandomArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*randomArgs)(nil)).Elem()
}

type Component struct {
	khulnasoft.ResourceState

	ChildID khulnasoft.IDOutput `khulnasoft:"childId"`
}

type ComponentArgs struct {
	Length int `khulnasoft:"length"`
}

func NewComponent(ctx *khulnasoft.Context, name string, args *ComponentArgs,
	opts ...khulnasoft.ResourceOption,
) (*Component, error) {
	if args == nil {
		return nil, errors.New("args is required")
	}

	component := &Component{}
	err := ctx.RegisterComponentResource("testprovider:index:Component", name, component, opts...)
	if err != nil {
		return nil, err
	}

	res, err := NewRandom(ctx, fmt.Sprintf("child-%s", name), &RandomArgs{
		Length: khulnasoft.Int(args.Length),
	}, khulnasoft.Parent(component))
	if err != nil {
		return nil, err
	}

	component.ChildID = res.ID()

	if err := ctx.RegisterResourceOutputs(component, khulnasoft.Map{
		"childId": component.ChildID,
	}); err != nil {
		return nil, err
	}

	return component, nil
}
