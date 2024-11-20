// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mypkg

import (
	"context"
	"reflect"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoftx"
	"output-funcs/mypkg/internal"
)

// Another failing example. A list of SSIS object metadata.
// API Version: 2018-06-01.
func GetIntegrationRuntimeObjectMetadatum(ctx *khulnasoft.Context, args *GetIntegrationRuntimeObjectMetadatumArgs, opts ...khulnasoft.InvokeOption) (*GetIntegrationRuntimeObjectMetadatumResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetIntegrationRuntimeObjectMetadatumResult
	err := ctx.Invoke("mypkg::getIntegrationRuntimeObjectMetadatum", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetIntegrationRuntimeObjectMetadatumArgs struct {
	// The factory name.
	FactoryName string `khulnasoft:"factoryName"`
	// The integration runtime name.
	IntegrationRuntimeName string `khulnasoft:"integrationRuntimeName"`
	// Metadata path.
	MetadataPath *string `khulnasoft:"metadataPath"`
	// The resource group name.
	ResourceGroupName string `khulnasoft:"resourceGroupName"`
}

// A list of SSIS object metadata.
type GetIntegrationRuntimeObjectMetadatumResult struct {
	// The link to the next page of results, if any remaining results exist.
	NextLink *string `khulnasoft:"nextLink"`
	// List of SSIS object metadata.
	Value []interface{} `khulnasoft:"value"`
}

func GetIntegrationRuntimeObjectMetadatumOutput(ctx *khulnasoft.Context, args GetIntegrationRuntimeObjectMetadatumOutputArgs, opts ...khulnasoft.InvokeOption) GetIntegrationRuntimeObjectMetadatumResultOutput {
	outputResult := khulnasoftx.ApplyErr[*GetIntegrationRuntimeObjectMetadatumArgs](args.ToOutput(), func(plainArgs *GetIntegrationRuntimeObjectMetadatumArgs) (*GetIntegrationRuntimeObjectMetadatumResult, error) {
		return GetIntegrationRuntimeObjectMetadatum(ctx, plainArgs, opts...)
	})

	return khulnasoftx.Cast[GetIntegrationRuntimeObjectMetadatumResultOutput, *GetIntegrationRuntimeObjectMetadatumResult](outputResult)
}

type GetIntegrationRuntimeObjectMetadatumOutputArgs struct {
	// The factory name.
	FactoryName khulnasoftx.Input[string] `khulnasoft:"factoryName"`
	// The integration runtime name.
	IntegrationRuntimeName khulnasoftx.Input[string] `khulnasoft:"integrationRuntimeName"`
	// Metadata path.
	MetadataPath khulnasoftx.Input[*string] `khulnasoft:"metadataPath"`
	// The resource group name.
	ResourceGroupName khulnasoftx.Input[string] `khulnasoft:"resourceGroupName"`
}

func (args GetIntegrationRuntimeObjectMetadatumOutputArgs) ToOutput() khulnasoftx.Output[*GetIntegrationRuntimeObjectMetadatumArgs] {
	allArgs := khulnasoftx.All(
		args.FactoryName.ToOutput(context.Background()).AsAny(),
		args.IntegrationRuntimeName.ToOutput(context.Background()).AsAny(),
		args.MetadataPath.ToOutput(context.Background()).AsAny(),
		args.ResourceGroupName.ToOutput(context.Background()).AsAny())
	return khulnasoftx.Apply[[]any](allArgs, func(resolvedArgs []interface{}) *GetIntegrationRuntimeObjectMetadatumArgs {
		return &GetIntegrationRuntimeObjectMetadatumArgs{
			FactoryName:            resolvedArgs[0].(string),
			IntegrationRuntimeName: resolvedArgs[1].(string),
			MetadataPath:           resolvedArgs[2].(*string),
			ResourceGroupName:      resolvedArgs[3].(string),
		}
	})
}

type GetIntegrationRuntimeObjectMetadatumResultOutput struct{ *khulnasoft.OutputState }

func (GetIntegrationRuntimeObjectMetadatumResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIntegrationRuntimeObjectMetadatumResult)(nil)).Elem()
}

func (o GetIntegrationRuntimeObjectMetadatumResultOutput) ToOutput(context.Context) khulnasoftx.Output[*GetIntegrationRuntimeObjectMetadatumResult] {
	return khulnasoftx.Output[*GetIntegrationRuntimeObjectMetadatumResult]{
		OutputState: o.OutputState,
	}
}

func (o GetIntegrationRuntimeObjectMetadatumResultOutput) NextLink() khulnasoftx.Output[*string] {
	return khulnasoftx.Apply[*GetIntegrationRuntimeObjectMetadatumResult](o, func(v *GetIntegrationRuntimeObjectMetadatumResult) *string { return v.NextLink })
}

func (o GetIntegrationRuntimeObjectMetadatumResultOutput) Value() khulnasoftx.ArrayOutput[any] {
	value := khulnasoftx.Apply[*GetIntegrationRuntimeObjectMetadatumResult](o, func(v *GetIntegrationRuntimeObjectMetadatumResult) []interface{} { return v.Value })
	return khulnasoftx.ArrayOutput[any]{
		OutputState: value.OutputState,
	}
}