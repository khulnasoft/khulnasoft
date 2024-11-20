// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mypkg

import (
	"context"
	"reflect"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"output-funcs-tfbridge20/mypkg/internal"
)

// Taken from khulnasoft-AWS to regress an issue
//
// Deprecated: aws.getAmiIds has been deprecated in favor of aws.ec2.getAmiIds
func GetAmiIds(ctx *khulnasoft.Context, args *GetAmiIdsArgs, opts ...khulnasoft.InvokeOption) (*GetAmiIdsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAmiIdsResult
	err := ctx.Invoke("mypkg::getAmiIds", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAmiIds.
type GetAmiIdsArgs struct {
	// Limit search to users with *explicit* launch
	// permission on  the image. Valid items are the numeric account ID or `self`.
	ExecutableUsers []string `khulnasoft:"executableUsers"`
	// One or more name/value pairs to filter off of. There
	// are several valid keys, for a full reference, check out
	// [describe-images in the AWS CLI reference][1].
	Filters []GetAmiIdsFilter `khulnasoft:"filters"`
	// A regex string to apply to the AMI list returned
	// by AWS. This allows more advanced filtering not supported from the AWS API.
	// This filtering is done locally on what AWS returns, and could have a performance
	// impact if the result is large. It is recommended to combine this with other
	// options to narrow down the list AWS returns.
	NameRegex *string `khulnasoft:"nameRegex"`
	// List of AMI owners to limit search. At least 1 value must be specified. Valid values: an AWS account ID, `self` (the current account), or an AWS owner alias (e.g. `amazon`, `aws-marketplace`, `microsoft`).
	Owners []string `khulnasoft:"owners"`
	// Used to sort AMIs by creation time.
	SortAscending *bool `khulnasoft:"sortAscending"`
}

// A collection of values returned by getAmiIds.
type GetAmiIdsResult struct {
	ExecutableUsers []string          `khulnasoft:"executableUsers"`
	Filters         []GetAmiIdsFilter `khulnasoft:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id            string   `khulnasoft:"id"`
	Ids           []string `khulnasoft:"ids"`
	NameRegex     *string  `khulnasoft:"nameRegex"`
	Owners        []string `khulnasoft:"owners"`
	SortAscending *bool    `khulnasoft:"sortAscending"`
}

func GetAmiIdsOutput(ctx *khulnasoft.Context, args GetAmiIdsOutputArgs, opts ...khulnasoft.InvokeOption) GetAmiIdsResultOutput {
	return khulnasoft.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAmiIdsResultOutput, error) {
			args := v.(GetAmiIdsArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv GetAmiIdsResult
			secret, err := ctx.InvokePackageRaw("mypkg::getAmiIds", args, &rv, "", opts...)
			if err != nil {
				return GetAmiIdsResultOutput{}, err
			}

			output := khulnasoft.ToOutput(rv).(GetAmiIdsResultOutput)
			if secret {
				return khulnasoft.ToSecret(output).(GetAmiIdsResultOutput), nil
			}
			return output, nil
		}).(GetAmiIdsResultOutput)
}

// A collection of arguments for invoking getAmiIds.
type GetAmiIdsOutputArgs struct {
	// Limit search to users with *explicit* launch
	// permission on  the image. Valid items are the numeric account ID or `self`.
	ExecutableUsers khulnasoft.StringArrayInput `khulnasoft:"executableUsers"`
	// One or more name/value pairs to filter off of. There
	// are several valid keys, for a full reference, check out
	// [describe-images in the AWS CLI reference][1].
	Filters GetAmiIdsFilterArrayInput `khulnasoft:"filters"`
	// A regex string to apply to the AMI list returned
	// by AWS. This allows more advanced filtering not supported from the AWS API.
	// This filtering is done locally on what AWS returns, and could have a performance
	// impact if the result is large. It is recommended to combine this with other
	// options to narrow down the list AWS returns.
	NameRegex khulnasoft.StringPtrInput `khulnasoft:"nameRegex"`
	// List of AMI owners to limit search. At least 1 value must be specified. Valid values: an AWS account ID, `self` (the current account), or an AWS owner alias (e.g. `amazon`, `aws-marketplace`, `microsoft`).
	Owners khulnasoft.StringArrayInput `khulnasoft:"owners"`
	// Used to sort AMIs by creation time.
	SortAscending khulnasoft.BoolPtrInput `khulnasoft:"sortAscending"`
}

func (GetAmiIdsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAmiIdsArgs)(nil)).Elem()
}

// A collection of values returned by getAmiIds.
type GetAmiIdsResultOutput struct{ *khulnasoft.OutputState }

func (GetAmiIdsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAmiIdsResult)(nil)).Elem()
}

func (o GetAmiIdsResultOutput) ToGetAmiIdsResultOutput() GetAmiIdsResultOutput {
	return o
}

func (o GetAmiIdsResultOutput) ToGetAmiIdsResultOutputWithContext(ctx context.Context) GetAmiIdsResultOutput {
	return o
}

func (o GetAmiIdsResultOutput) ExecutableUsers() khulnasoft.StringArrayOutput {
	return o.ApplyT(func(v GetAmiIdsResult) []string { return v.ExecutableUsers }).(khulnasoft.StringArrayOutput)
}

func (o GetAmiIdsResultOutput) Filters() GetAmiIdsFilterArrayOutput {
	return o.ApplyT(func(v GetAmiIdsResult) []GetAmiIdsFilter { return v.Filters }).(GetAmiIdsFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAmiIdsResultOutput) Id() khulnasoft.StringOutput {
	return o.ApplyT(func(v GetAmiIdsResult) string { return v.Id }).(khulnasoft.StringOutput)
}

func (o GetAmiIdsResultOutput) Ids() khulnasoft.StringArrayOutput {
	return o.ApplyT(func(v GetAmiIdsResult) []string { return v.Ids }).(khulnasoft.StringArrayOutput)
}

func (o GetAmiIdsResultOutput) NameRegex() khulnasoft.StringPtrOutput {
	return o.ApplyT(func(v GetAmiIdsResult) *string { return v.NameRegex }).(khulnasoft.StringPtrOutput)
}

func (o GetAmiIdsResultOutput) Owners() khulnasoft.StringArrayOutput {
	return o.ApplyT(func(v GetAmiIdsResult) []string { return v.Owners }).(khulnasoft.StringArrayOutput)
}

func (o GetAmiIdsResultOutput) SortAscending() khulnasoft.BoolPtrOutput {
	return o.ApplyT(func(v GetAmiIdsResult) *bool { return v.SortAscending }).(khulnasoft.BoolPtrOutput)
}

func init() {
	khulnasoft.RegisterOutputType(GetAmiIdsResultOutput{})
}