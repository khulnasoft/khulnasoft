// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"reflect"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

type componentArgs struct {
	Message string              `khulnasoft:"message"`
	Nested  componentNestedArgs `khulnasoft:"nested"`
}

type ComponentArgs struct {
	Message khulnasoft.StringInput
	Nested  ComponentNestedInput
}

func (ComponentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentArgs)(nil)).Elem()
}

type componentNestedArgs struct {
	Value string `khulnasoft:"Value"`
}

type ComponentNestedArgs struct {
	Value khulnasoft.StringInput
}

type ComponentNestedInput interface {
	khulnasoft.Input
}

func (ComponentNestedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentNestedArgs)(nil)).Elem()
}

type Component struct {
	khulnasoft.ResourceState
}

func NewComponent(
	ctx *khulnasoft.Context, name string, args *ComponentArgs, opts ...khulnasoft.ResourceOption,
) (*Component, error) {
	var resource Component
	err := ctx.RegisterRemoteComponentResource("testcomponent:index:Component", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}

	return &resource, nil
}
