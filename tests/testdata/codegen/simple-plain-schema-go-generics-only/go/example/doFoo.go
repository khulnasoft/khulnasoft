// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"simple-plain-schema-go-generics-only/example/internal"
)

func DoFoo(ctx *khulnasoft.Context, args *DoFooArgs, opts ...khulnasoft.InvokeOption) error {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv struct{}
	err := ctx.Invoke("example::doFoo", args, &rv, opts...)
	return err
}

type DoFooArgs struct {
	Foo Foo `khulnasoft:"foo"`
}