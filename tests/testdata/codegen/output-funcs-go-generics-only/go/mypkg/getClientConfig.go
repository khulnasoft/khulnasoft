// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mypkg

import (
	"context"
	"reflect"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoftx"
	"output-funcs-go-generics-only/mypkg/internal"
)

// Failing example taken from azure-native. Original doc: Use this function to access the current configuration of the native Azure provider.
func GetClientConfig(ctx *khulnasoft.Context, opts ...khulnasoft.InvokeOption) (*GetClientConfigResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetClientConfigResult
	err := ctx.Invoke("mypkg::getClientConfig", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// Configuration values returned by getClientConfig.
type GetClientConfigResult struct {
	// Azure Client ID (Application Object ID).
	ClientId string `khulnasoft:"clientId"`
	// Azure Object ID of the current user or service principal.
	ObjectId string `khulnasoft:"objectId"`
	// Azure Subscription ID
	SubscriptionId string `khulnasoft:"subscriptionId"`
	// Azure Tenant ID
	TenantId string `khulnasoft:"tenantId"`
}

func GetClientConfigOutput(ctx *khulnasoft.Context, opts ...khulnasoft.InvokeOption) GetClientConfigResultOutput {
	outputResult := khulnasoftx.ApplyErr[int](khulnasoftx.Val(0), func(_ int) (*GetClientConfigResult, error) {
		return GetClientConfig(ctx, opts...)
	})

	return khulnasoftx.Cast[GetClientConfigResultOutput, *GetClientConfigResult](outputResult)
}

type GetClientConfigResultOutput struct{ *khulnasoft.OutputState }

func (GetClientConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClientConfigResult)(nil)).Elem()
}

func (o GetClientConfigResultOutput) ToOutput(context.Context) khulnasoftx.Output[*GetClientConfigResult] {
	return khulnasoftx.Output[*GetClientConfigResult]{
		OutputState: o.OutputState,
	}
}

func (o GetClientConfigResultOutput) ClientId() khulnasoftx.Output[string] {
	return khulnasoftx.Apply[*GetClientConfigResult](o, func(v *GetClientConfigResult) string { return v.ClientId })
}

func (o GetClientConfigResultOutput) ObjectId() khulnasoftx.Output[string] {
	return khulnasoftx.Apply[*GetClientConfigResult](o, func(v *GetClientConfigResult) string { return v.ObjectId })
}

func (o GetClientConfigResultOutput) SubscriptionId() khulnasoftx.Output[string] {
	return khulnasoftx.Apply[*GetClientConfigResult](o, func(v *GetClientConfigResult) string { return v.SubscriptionId })
}

func (o GetClientConfigResultOutput) TenantId() khulnasoftx.Output[string] {
	return khulnasoftx.Apply[*GetClientConfigResult](o, func(v *GetClientConfigResult) string { return v.TenantId })
}