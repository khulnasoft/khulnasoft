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

var _ = internal.GetEnvOrDefault

// Bastion Shareable Link.
type BastionShareableLink struct {
	// Reference of the virtual machine resource.
	Vm string `khulnasoft:"vm"`
}

// Bastion Shareable Link.
type BastionShareableLinkArgs struct {
	// Reference of the virtual machine resource.
	Vm khulnasoftx.Input[string] `khulnasoft:"vm"`
}

func (BastionShareableLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BastionShareableLink)(nil)).Elem()
}

func (i BastionShareableLinkArgs) ToBastionShareableLinkOutput() BastionShareableLinkOutput {
	return i.ToBastionShareableLinkOutputWithContext(context.Background())
}

func (i BastionShareableLinkArgs) ToBastionShareableLinkOutputWithContext(ctx context.Context) BastionShareableLinkOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(BastionShareableLinkOutput)
}

func (i *BastionShareableLinkArgs) ToOutput(ctx context.Context) khulnasoftx.Output[*BastionShareableLinkArgs] {
	return khulnasoftx.Val(i)
}

// Bastion Shareable Link.
type BastionShareableLinkOutput struct{ *khulnasoft.OutputState }

func (BastionShareableLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BastionShareableLink)(nil)).Elem()
}

func (o BastionShareableLinkOutput) ToBastionShareableLinkOutput() BastionShareableLinkOutput {
	return o
}

func (o BastionShareableLinkOutput) ToBastionShareableLinkOutputWithContext(ctx context.Context) BastionShareableLinkOutput {
	return o
}

func (o BastionShareableLinkOutput) ToOutput(ctx context.Context) khulnasoftx.Output[BastionShareableLink] {
	return khulnasoftx.Output[BastionShareableLink]{
		OutputState: o.OutputState,
	}
}

// Reference of the virtual machine resource.
func (o BastionShareableLinkOutput) Vm() khulnasoftx.Output[string] {
	return khulnasoftx.Apply[BastionShareableLink](o, func(v BastionShareableLink) string { return v.Vm })
}

// Ssis environment reference.
type SsisEnvironmentReferenceResponse struct {
	// Environment folder name.
	EnvironmentFolderName *string `khulnasoft:"environmentFolderName"`
	// Environment name.
	EnvironmentName *string `khulnasoft:"environmentName"`
	// Environment reference id.
	Id *float64 `khulnasoft:"id"`
	// Reference type
	ReferenceType *string `khulnasoft:"referenceType"`
}

// Ssis environment reference.
type SsisEnvironmentReferenceResponseArgs struct {
	// Environment folder name.
	EnvironmentFolderName khulnasoftx.Input[*string] `khulnasoft:"environmentFolderName"`
	// Environment name.
	EnvironmentName khulnasoftx.Input[*string] `khulnasoft:"environmentName"`
	// Environment reference id.
	Id khulnasoftx.Input[*float64] `khulnasoft:"id"`
	// Reference type
	ReferenceType khulnasoftx.Input[*string] `khulnasoft:"referenceType"`
}

func (SsisEnvironmentReferenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SsisEnvironmentReferenceResponse)(nil)).Elem()
}

func (i SsisEnvironmentReferenceResponseArgs) ToSsisEnvironmentReferenceResponseOutput() SsisEnvironmentReferenceResponseOutput {
	return i.ToSsisEnvironmentReferenceResponseOutputWithContext(context.Background())
}

func (i SsisEnvironmentReferenceResponseArgs) ToSsisEnvironmentReferenceResponseOutputWithContext(ctx context.Context) SsisEnvironmentReferenceResponseOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(SsisEnvironmentReferenceResponseOutput)
}

func (i *SsisEnvironmentReferenceResponseArgs) ToOutput(ctx context.Context) khulnasoftx.Output[*SsisEnvironmentReferenceResponseArgs] {
	return khulnasoftx.Val(i)
}

// Ssis environment reference.
type SsisEnvironmentReferenceResponseOutput struct{ *khulnasoft.OutputState }

func (SsisEnvironmentReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SsisEnvironmentReferenceResponse)(nil)).Elem()
}

func (o SsisEnvironmentReferenceResponseOutput) ToSsisEnvironmentReferenceResponseOutput() SsisEnvironmentReferenceResponseOutput {
	return o
}

func (o SsisEnvironmentReferenceResponseOutput) ToSsisEnvironmentReferenceResponseOutputWithContext(ctx context.Context) SsisEnvironmentReferenceResponseOutput {
	return o
}

func (o SsisEnvironmentReferenceResponseOutput) ToOutput(ctx context.Context) khulnasoftx.Output[SsisEnvironmentReferenceResponse] {
	return khulnasoftx.Output[SsisEnvironmentReferenceResponse]{
		OutputState: o.OutputState,
	}
}

// Environment folder name.
func (o SsisEnvironmentReferenceResponseOutput) EnvironmentFolderName() khulnasoftx.Output[*string] {
	return khulnasoftx.Apply[SsisEnvironmentReferenceResponse](o, func(v SsisEnvironmentReferenceResponse) *string { return v.EnvironmentFolderName })
}

// Environment name.
func (o SsisEnvironmentReferenceResponseOutput) EnvironmentName() khulnasoftx.Output[*string] {
	return khulnasoftx.Apply[SsisEnvironmentReferenceResponse](o, func(v SsisEnvironmentReferenceResponse) *string { return v.EnvironmentName })
}

// Environment reference id.
func (o SsisEnvironmentReferenceResponseOutput) Id() khulnasoftx.Output[*float64] {
	return khulnasoftx.Apply[SsisEnvironmentReferenceResponse](o, func(v SsisEnvironmentReferenceResponse) *float64 { return v.Id })
}

// Reference type
func (o SsisEnvironmentReferenceResponseOutput) ReferenceType() khulnasoftx.Output[*string] {
	return khulnasoftx.Apply[SsisEnvironmentReferenceResponse](o, func(v SsisEnvironmentReferenceResponse) *string { return v.ReferenceType })
}

// Ssis environment.
type SsisEnvironmentResponse struct {
	// Metadata description.
	Description *string `khulnasoft:"description"`
	// Folder id which contains environment.
	FolderId *float64 `khulnasoft:"folderId"`
	// Metadata id.
	Id *float64 `khulnasoft:"id"`
	// Metadata name.
	Name *string `khulnasoft:"name"`
	// The type of SSIS object metadata.
	// Expected value is 'Environment'.
	Type string `khulnasoft:"type"`
	// Variable in environment
	Variables []*SsisVariableResponse `khulnasoft:"variables"`
}

// Ssis environment.
type SsisEnvironmentResponseArgs struct {
	// Metadata description.
	Description khulnasoftx.Input[*string] `khulnasoft:"description"`
	// Folder id which contains environment.
	FolderId khulnasoftx.Input[*float64] `khulnasoft:"folderId"`
	// Metadata id.
	Id khulnasoftx.Input[*float64] `khulnasoft:"id"`
	// Metadata name.
	Name khulnasoftx.Input[*string] `khulnasoft:"name"`
	// The type of SSIS object metadata.
	// Expected value is 'Environment'.
	Type khulnasoftx.Input[string] `khulnasoft:"type"`
	// Variable in environment
	Variables khulnasoftx.Input[[]*SsisVariableResponseArgs] `khulnasoft:"variables"`
}

func (SsisEnvironmentResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SsisEnvironmentResponse)(nil)).Elem()
}

func (i SsisEnvironmentResponseArgs) ToSsisEnvironmentResponseOutput() SsisEnvironmentResponseOutput {
	return i.ToSsisEnvironmentResponseOutputWithContext(context.Background())
}

func (i SsisEnvironmentResponseArgs) ToSsisEnvironmentResponseOutputWithContext(ctx context.Context) SsisEnvironmentResponseOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(SsisEnvironmentResponseOutput)
}

func (i *SsisEnvironmentResponseArgs) ToOutput(ctx context.Context) khulnasoftx.Output[*SsisEnvironmentResponseArgs] {
	return khulnasoftx.Val(i)
}

// Ssis environment.
type SsisEnvironmentResponseOutput struct{ *khulnasoft.OutputState }

func (SsisEnvironmentResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SsisEnvironmentResponse)(nil)).Elem()
}

func (o SsisEnvironmentResponseOutput) ToSsisEnvironmentResponseOutput() SsisEnvironmentResponseOutput {
	return o
}

func (o SsisEnvironmentResponseOutput) ToSsisEnvironmentResponseOutputWithContext(ctx context.Context) SsisEnvironmentResponseOutput {
	return o
}

func (o SsisEnvironmentResponseOutput) ToOutput(ctx context.Context) khulnasoftx.Output[SsisEnvironmentResponse] {
	return khulnasoftx.Output[SsisEnvironmentResponse]{
		OutputState: o.OutputState,
	}
}

// Metadata description.
func (o SsisEnvironmentResponseOutput) Description() khulnasoftx.Output[*string] {
	return khulnasoftx.Apply[SsisEnvironmentResponse](o, func(v SsisEnvironmentResponse) *string { return v.Description })
}

// Folder id which contains environment.
func (o SsisEnvironmentResponseOutput) FolderId() khulnasoftx.Output[*float64] {
	return khulnasoftx.Apply[SsisEnvironmentResponse](o, func(v SsisEnvironmentResponse) *float64 { return v.FolderId })
}

// Metadata id.
func (o SsisEnvironmentResponseOutput) Id() khulnasoftx.Output[*float64] {
	return khulnasoftx.Apply[SsisEnvironmentResponse](o, func(v SsisEnvironmentResponse) *float64 { return v.Id })
}

// Metadata name.
func (o SsisEnvironmentResponseOutput) Name() khulnasoftx.Output[*string] {
	return khulnasoftx.Apply[SsisEnvironmentResponse](o, func(v SsisEnvironmentResponse) *string { return v.Name })
}

// The type of SSIS object metadata.
// Expected value is 'Environment'.
func (o SsisEnvironmentResponseOutput) Type() khulnasoftx.Output[string] {
	return khulnasoftx.Apply[SsisEnvironmentResponse](o, func(v SsisEnvironmentResponse) string { return v.Type })
}

// Variable in environment
func (o SsisEnvironmentResponseOutput) Variables() khulnasoftx.GArrayOutput[SsisVariableResponse, SsisVariableResponseOutput] {
	value := khulnasoftx.Apply[SsisEnvironmentResponse](o, func(v SsisEnvironmentResponse) []*SsisVariableResponse { return v.Variables })
	return khulnasoftx.GArrayOutput[SsisVariableResponse, SsisVariableResponseOutput]{OutputState: value.OutputState}
}

// Ssis folder.
type SsisFolderResponse struct {
	// Metadata description.
	Description *string `khulnasoft:"description"`
	// Metadata id.
	Id *float64 `khulnasoft:"id"`
	// Metadata name.
	Name *string `khulnasoft:"name"`
	// The type of SSIS object metadata.
	// Expected value is 'Folder'.
	Type string `khulnasoft:"type"`
}

// Ssis folder.
type SsisFolderResponseArgs struct {
	// Metadata description.
	Description khulnasoftx.Input[*string] `khulnasoft:"description"`
	// Metadata id.
	Id khulnasoftx.Input[*float64] `khulnasoft:"id"`
	// Metadata name.
	Name khulnasoftx.Input[*string] `khulnasoft:"name"`
	// The type of SSIS object metadata.
	// Expected value is 'Folder'.
	Type khulnasoftx.Input[string] `khulnasoft:"type"`
}

func (SsisFolderResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SsisFolderResponse)(nil)).Elem()
}

func (i SsisFolderResponseArgs) ToSsisFolderResponseOutput() SsisFolderResponseOutput {
	return i.ToSsisFolderResponseOutputWithContext(context.Background())
}

func (i SsisFolderResponseArgs) ToSsisFolderResponseOutputWithContext(ctx context.Context) SsisFolderResponseOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(SsisFolderResponseOutput)
}

func (i *SsisFolderResponseArgs) ToOutput(ctx context.Context) khulnasoftx.Output[*SsisFolderResponseArgs] {
	return khulnasoftx.Val(i)
}

// Ssis folder.
type SsisFolderResponseOutput struct{ *khulnasoft.OutputState }

func (SsisFolderResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SsisFolderResponse)(nil)).Elem()
}

func (o SsisFolderResponseOutput) ToSsisFolderResponseOutput() SsisFolderResponseOutput {
	return o
}

func (o SsisFolderResponseOutput) ToSsisFolderResponseOutputWithContext(ctx context.Context) SsisFolderResponseOutput {
	return o
}

func (o SsisFolderResponseOutput) ToOutput(ctx context.Context) khulnasoftx.Output[SsisFolderResponse] {
	return khulnasoftx.Output[SsisFolderResponse]{
		OutputState: o.OutputState,
	}
}

// Metadata description.
func (o SsisFolderResponseOutput) Description() khulnasoftx.Output[*string] {
	return khulnasoftx.Apply[SsisFolderResponse](o, func(v SsisFolderResponse) *string { return v.Description })
}

// Metadata id.
func (o SsisFolderResponseOutput) Id() khulnasoftx.Output[*float64] {
	return khulnasoftx.Apply[SsisFolderResponse](o, func(v SsisFolderResponse) *float64 { return v.Id })
}

// Metadata name.
func (o SsisFolderResponseOutput) Name() khulnasoftx.Output[*string] {
	return khulnasoftx.Apply[SsisFolderResponse](o, func(v SsisFolderResponse) *string { return v.Name })
}

// The type of SSIS object metadata.
// Expected value is 'Folder'.
func (o SsisFolderResponseOutput) Type() khulnasoftx.Output[string] {
	return khulnasoftx.Apply[SsisFolderResponse](o, func(v SsisFolderResponse) string { return v.Type })
}

// Ssis Package.
type SsisPackageResponse struct {
	// Metadata description.
	Description *string `khulnasoft:"description"`
	// Folder id which contains package.
	FolderId *float64 `khulnasoft:"folderId"`
	// Metadata id.
	Id *float64 `khulnasoft:"id"`
	// Metadata name.
	Name *string `khulnasoft:"name"`
	// Parameters in package
	Parameters []*SsisParameterResponse `khulnasoft:"parameters"`
	// Project id which contains package.
	ProjectId *float64 `khulnasoft:"projectId"`
	// Project version which contains package.
	ProjectVersion *float64 `khulnasoft:"projectVersion"`
	// The type of SSIS object metadata.
	// Expected value is 'Package'.
	Type string `khulnasoft:"type"`
}

// Ssis Package.
type SsisPackageResponseArgs struct {
	// Metadata description.
	Description khulnasoftx.Input[*string] `khulnasoft:"description"`
	// Folder id which contains package.
	FolderId khulnasoftx.Input[*float64] `khulnasoft:"folderId"`
	// Metadata id.
	Id khulnasoftx.Input[*float64] `khulnasoft:"id"`
	// Metadata name.
	Name khulnasoftx.Input[*string] `khulnasoft:"name"`
	// Parameters in package
	Parameters khulnasoftx.Input[[]*SsisParameterResponseArgs] `khulnasoft:"parameters"`
	// Project id which contains package.
	ProjectId khulnasoftx.Input[*float64] `khulnasoft:"projectId"`
	// Project version which contains package.
	ProjectVersion khulnasoftx.Input[*float64] `khulnasoft:"projectVersion"`
	// The type of SSIS object metadata.
	// Expected value is 'Package'.
	Type khulnasoftx.Input[string] `khulnasoft:"type"`
}

func (SsisPackageResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SsisPackageResponse)(nil)).Elem()
}

func (i SsisPackageResponseArgs) ToSsisPackageResponseOutput() SsisPackageResponseOutput {
	return i.ToSsisPackageResponseOutputWithContext(context.Background())
}

func (i SsisPackageResponseArgs) ToSsisPackageResponseOutputWithContext(ctx context.Context) SsisPackageResponseOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(SsisPackageResponseOutput)
}

func (i *SsisPackageResponseArgs) ToOutput(ctx context.Context) khulnasoftx.Output[*SsisPackageResponseArgs] {
	return khulnasoftx.Val(i)
}

// Ssis Package.
type SsisPackageResponseOutput struct{ *khulnasoft.OutputState }

func (SsisPackageResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SsisPackageResponse)(nil)).Elem()
}

func (o SsisPackageResponseOutput) ToSsisPackageResponseOutput() SsisPackageResponseOutput {
	return o
}

func (o SsisPackageResponseOutput) ToSsisPackageResponseOutputWithContext(ctx context.Context) SsisPackageResponseOutput {
	return o
}

func (o SsisPackageResponseOutput) ToOutput(ctx context.Context) khulnasoftx.Output[SsisPackageResponse] {
	return khulnasoftx.Output[SsisPackageResponse]{
		OutputState: o.OutputState,
	}
}

// Metadata description.
func (o SsisPackageResponseOutput) Description() khulnasoftx.Output[*string] {
	return khulnasoftx.Apply[SsisPackageResponse](o, func(v SsisPackageResponse) *string { return v.Description })
}

// Folder id which contains package.
func (o SsisPackageResponseOutput) FolderId() khulnasoftx.Output[*float64] {
	return khulnasoftx.Apply[SsisPackageResponse](o, func(v SsisPackageResponse) *float64 { return v.FolderId })
}

// Metadata id.
func (o SsisPackageResponseOutput) Id() khulnasoftx.Output[*float64] {
	return khulnasoftx.Apply[SsisPackageResponse](o, func(v SsisPackageResponse) *float64 { return v.Id })
}

// Metadata name.
func (o SsisPackageResponseOutput) Name() khulnasoftx.Output[*string] {
	return khulnasoftx.Apply[SsisPackageResponse](o, func(v SsisPackageResponse) *string { return v.Name })
}

// Parameters in package
func (o SsisPackageResponseOutput) Parameters() khulnasoftx.GArrayOutput[SsisParameterResponse, SsisParameterResponseOutput] {
	value := khulnasoftx.Apply[SsisPackageResponse](o, func(v SsisPackageResponse) []*SsisParameterResponse { return v.Parameters })
	return khulnasoftx.GArrayOutput[SsisParameterResponse, SsisParameterResponseOutput]{OutputState: value.OutputState}
}

// Project id which contains package.
func (o SsisPackageResponseOutput) ProjectId() khulnasoftx.Output[*float64] {
	return khulnasoftx.Apply[SsisPackageResponse](o, func(v SsisPackageResponse) *float64 { return v.ProjectId })
}

// Project version which contains package.
func (o SsisPackageResponseOutput) ProjectVersion() khulnasoftx.Output[*float64] {
	return khulnasoftx.Apply[SsisPackageResponse](o, func(v SsisPackageResponse) *float64 { return v.ProjectVersion })
}

// The type of SSIS object metadata.
// Expected value is 'Package'.
func (o SsisPackageResponseOutput) Type() khulnasoftx.Output[string] {
	return khulnasoftx.Apply[SsisPackageResponse](o, func(v SsisPackageResponse) string { return v.Type })
}

// Ssis parameter.
type SsisParameterResponse struct {
	// Parameter type.
	DataType *string `khulnasoft:"dataType"`
	// Default value of parameter.
	DefaultValue *string `khulnasoft:"defaultValue"`
	// Parameter description.
	Description *string `khulnasoft:"description"`
	// Design default value of parameter.
	DesignDefaultValue *string `khulnasoft:"designDefaultValue"`
	// Parameter id.
	Id *float64 `khulnasoft:"id"`
	// Parameter name.
	Name *string `khulnasoft:"name"`
	// Whether parameter is required.
	Required *bool `khulnasoft:"required"`
	// Whether parameter is sensitive.
	Sensitive *bool `khulnasoft:"sensitive"`
	// Default sensitive value of parameter.
	SensitiveDefaultValue *string `khulnasoft:"sensitiveDefaultValue"`
	// Parameter value set.
	ValueSet *bool `khulnasoft:"valueSet"`
	// Parameter value type.
	ValueType *string `khulnasoft:"valueType"`
	// Parameter reference variable.
	Variable *string `khulnasoft:"variable"`
}

// Ssis parameter.
type SsisParameterResponseArgs struct {
	// Parameter type.
	DataType khulnasoftx.Input[*string] `khulnasoft:"dataType"`
	// Default value of parameter.
	DefaultValue khulnasoftx.Input[*string] `khulnasoft:"defaultValue"`
	// Parameter description.
	Description khulnasoftx.Input[*string] `khulnasoft:"description"`
	// Design default value of parameter.
	DesignDefaultValue khulnasoftx.Input[*string] `khulnasoft:"designDefaultValue"`
	// Parameter id.
	Id khulnasoftx.Input[*float64] `khulnasoft:"id"`
	// Parameter name.
	Name khulnasoftx.Input[*string] `khulnasoft:"name"`
	// Whether parameter is required.
	Required khulnasoftx.Input[*bool] `khulnasoft:"required"`
	// Whether parameter is sensitive.
	Sensitive khulnasoftx.Input[*bool] `khulnasoft:"sensitive"`
	// Default sensitive value of parameter.
	SensitiveDefaultValue khulnasoftx.Input[*string] `khulnasoft:"sensitiveDefaultValue"`
	// Parameter value set.
	ValueSet khulnasoftx.Input[*bool] `khulnasoft:"valueSet"`
	// Parameter value type.
	ValueType khulnasoftx.Input[*string] `khulnasoft:"valueType"`
	// Parameter reference variable.
	Variable khulnasoftx.Input[*string] `khulnasoft:"variable"`
}

func (SsisParameterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SsisParameterResponse)(nil)).Elem()
}

func (i SsisParameterResponseArgs) ToSsisParameterResponseOutput() SsisParameterResponseOutput {
	return i.ToSsisParameterResponseOutputWithContext(context.Background())
}

func (i SsisParameterResponseArgs) ToSsisParameterResponseOutputWithContext(ctx context.Context) SsisParameterResponseOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(SsisParameterResponseOutput)
}

func (i *SsisParameterResponseArgs) ToOutput(ctx context.Context) khulnasoftx.Output[*SsisParameterResponseArgs] {
	return khulnasoftx.Val(i)
}

// Ssis parameter.
type SsisParameterResponseOutput struct{ *khulnasoft.OutputState }

func (SsisParameterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SsisParameterResponse)(nil)).Elem()
}

func (o SsisParameterResponseOutput) ToSsisParameterResponseOutput() SsisParameterResponseOutput {
	return o
}

func (o SsisParameterResponseOutput) ToSsisParameterResponseOutputWithContext(ctx context.Context) SsisParameterResponseOutput {
	return o
}

func (o SsisParameterResponseOutput) ToOutput(ctx context.Context) khulnasoftx.Output[SsisParameterResponse] {
	return khulnasoftx.Output[SsisParameterResponse]{
		OutputState: o.OutputState,
	}
}

// Parameter type.
func (o SsisParameterResponseOutput) DataType() khulnasoftx.Output[*string] {
	return khulnasoftx.Apply[SsisParameterResponse](o, func(v SsisParameterResponse) *string { return v.DataType })
}

// Default value of parameter.
func (o SsisParameterResponseOutput) DefaultValue() khulnasoftx.Output[*string] {
	return khulnasoftx.Apply[SsisParameterResponse](o, func(v SsisParameterResponse) *string { return v.DefaultValue })
}

// Parameter description.
func (o SsisParameterResponseOutput) Description() khulnasoftx.Output[*string] {
	return khulnasoftx.Apply[SsisParameterResponse](o, func(v SsisParameterResponse) *string { return v.Description })
}

// Design default value of parameter.
func (o SsisParameterResponseOutput) DesignDefaultValue() khulnasoftx.Output[*string] {
	return khulnasoftx.Apply[SsisParameterResponse](o, func(v SsisParameterResponse) *string { return v.DesignDefaultValue })
}

// Parameter id.
func (o SsisParameterResponseOutput) Id() khulnasoftx.Output[*float64] {
	return khulnasoftx.Apply[SsisParameterResponse](o, func(v SsisParameterResponse) *float64 { return v.Id })
}

// Parameter name.
func (o SsisParameterResponseOutput) Name() khulnasoftx.Output[*string] {
	return khulnasoftx.Apply[SsisParameterResponse](o, func(v SsisParameterResponse) *string { return v.Name })
}

// Whether parameter is required.
func (o SsisParameterResponseOutput) Required() khulnasoftx.Output[*bool] {
	return khulnasoftx.Apply[SsisParameterResponse](o, func(v SsisParameterResponse) *bool { return v.Required })
}

// Whether parameter is sensitive.
func (o SsisParameterResponseOutput) Sensitive() khulnasoftx.Output[*bool] {
	return khulnasoftx.Apply[SsisParameterResponse](o, func(v SsisParameterResponse) *bool { return v.Sensitive })
}

// Default sensitive value of parameter.
func (o SsisParameterResponseOutput) SensitiveDefaultValue() khulnasoftx.Output[*string] {
	return khulnasoftx.Apply[SsisParameterResponse](o, func(v SsisParameterResponse) *string { return v.SensitiveDefaultValue })
}

// Parameter value set.
func (o SsisParameterResponseOutput) ValueSet() khulnasoftx.Output[*bool] {
	return khulnasoftx.Apply[SsisParameterResponse](o, func(v SsisParameterResponse) *bool { return v.ValueSet })
}

// Parameter value type.
func (o SsisParameterResponseOutput) ValueType() khulnasoftx.Output[*string] {
	return khulnasoftx.Apply[SsisParameterResponse](o, func(v SsisParameterResponse) *string { return v.ValueType })
}

// Parameter reference variable.
func (o SsisParameterResponseOutput) Variable() khulnasoftx.Output[*string] {
	return khulnasoftx.Apply[SsisParameterResponse](o, func(v SsisParameterResponse) *string { return v.Variable })
}

// Ssis project.
type SsisProjectResponse struct {
	// Metadata description.
	Description *string `khulnasoft:"description"`
	// Environment reference in project
	EnvironmentRefs []*SsisEnvironmentReferenceResponse `khulnasoft:"environmentRefs"`
	// Folder id which contains project.
	FolderId *float64 `khulnasoft:"folderId"`
	// Metadata id.
	Id *float64 `khulnasoft:"id"`
	// Metadata name.
	Name *string `khulnasoft:"name"`
	// Parameters in project
	Parameters []*SsisParameterResponse `khulnasoft:"parameters"`
	// The type of SSIS object metadata.
	// Expected value is 'Project'.
	Type string `khulnasoft:"type"`
	// Project version.
	Version *float64 `khulnasoft:"version"`
}

// Ssis project.
type SsisProjectResponseArgs struct {
	// Metadata description.
	Description khulnasoftx.Input[*string] `khulnasoft:"description"`
	// Environment reference in project
	EnvironmentRefs khulnasoftx.Input[[]*SsisEnvironmentReferenceResponseArgs] `khulnasoft:"environmentRefs"`
	// Folder id which contains project.
	FolderId khulnasoftx.Input[*float64] `khulnasoft:"folderId"`
	// Metadata id.
	Id khulnasoftx.Input[*float64] `khulnasoft:"id"`
	// Metadata name.
	Name khulnasoftx.Input[*string] `khulnasoft:"name"`
	// Parameters in project
	Parameters khulnasoftx.Input[[]*SsisParameterResponseArgs] `khulnasoft:"parameters"`
	// The type of SSIS object metadata.
	// Expected value is 'Project'.
	Type khulnasoftx.Input[string] `khulnasoft:"type"`
	// Project version.
	Version khulnasoftx.Input[*float64] `khulnasoft:"version"`
}

func (SsisProjectResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SsisProjectResponse)(nil)).Elem()
}

func (i SsisProjectResponseArgs) ToSsisProjectResponseOutput() SsisProjectResponseOutput {
	return i.ToSsisProjectResponseOutputWithContext(context.Background())
}

func (i SsisProjectResponseArgs) ToSsisProjectResponseOutputWithContext(ctx context.Context) SsisProjectResponseOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(SsisProjectResponseOutput)
}

func (i *SsisProjectResponseArgs) ToOutput(ctx context.Context) khulnasoftx.Output[*SsisProjectResponseArgs] {
	return khulnasoftx.Val(i)
}

// Ssis project.
type SsisProjectResponseOutput struct{ *khulnasoft.OutputState }

func (SsisProjectResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SsisProjectResponse)(nil)).Elem()
}

func (o SsisProjectResponseOutput) ToSsisProjectResponseOutput() SsisProjectResponseOutput {
	return o
}

func (o SsisProjectResponseOutput) ToSsisProjectResponseOutputWithContext(ctx context.Context) SsisProjectResponseOutput {
	return o
}

func (o SsisProjectResponseOutput) ToOutput(ctx context.Context) khulnasoftx.Output[SsisProjectResponse] {
	return khulnasoftx.Output[SsisProjectResponse]{
		OutputState: o.OutputState,
	}
}

// Metadata description.
func (o SsisProjectResponseOutput) Description() khulnasoftx.Output[*string] {
	return khulnasoftx.Apply[SsisProjectResponse](o, func(v SsisProjectResponse) *string { return v.Description })
}

// Environment reference in project
func (o SsisProjectResponseOutput) EnvironmentRefs() khulnasoftx.GArrayOutput[SsisEnvironmentReferenceResponse, SsisEnvironmentReferenceResponseOutput] {
	value := khulnasoftx.Apply[SsisProjectResponse](o, func(v SsisProjectResponse) []*SsisEnvironmentReferenceResponse { return v.EnvironmentRefs })
	return khulnasoftx.GArrayOutput[SsisEnvironmentReferenceResponse, SsisEnvironmentReferenceResponseOutput]{OutputState: value.OutputState}
}

// Folder id which contains project.
func (o SsisProjectResponseOutput) FolderId() khulnasoftx.Output[*float64] {
	return khulnasoftx.Apply[SsisProjectResponse](o, func(v SsisProjectResponse) *float64 { return v.FolderId })
}

// Metadata id.
func (o SsisProjectResponseOutput) Id() khulnasoftx.Output[*float64] {
	return khulnasoftx.Apply[SsisProjectResponse](o, func(v SsisProjectResponse) *float64 { return v.Id })
}

// Metadata name.
func (o SsisProjectResponseOutput) Name() khulnasoftx.Output[*string] {
	return khulnasoftx.Apply[SsisProjectResponse](o, func(v SsisProjectResponse) *string { return v.Name })
}

// Parameters in project
func (o SsisProjectResponseOutput) Parameters() khulnasoftx.GArrayOutput[SsisParameterResponse, SsisParameterResponseOutput] {
	value := khulnasoftx.Apply[SsisProjectResponse](o, func(v SsisProjectResponse) []*SsisParameterResponse { return v.Parameters })
	return khulnasoftx.GArrayOutput[SsisParameterResponse, SsisParameterResponseOutput]{OutputState: value.OutputState}
}

// The type of SSIS object metadata.
// Expected value is 'Project'.
func (o SsisProjectResponseOutput) Type() khulnasoftx.Output[string] {
	return khulnasoftx.Apply[SsisProjectResponse](o, func(v SsisProjectResponse) string { return v.Type })
}

// Project version.
func (o SsisProjectResponseOutput) Version() khulnasoftx.Output[*float64] {
	return khulnasoftx.Apply[SsisProjectResponse](o, func(v SsisProjectResponse) *float64 { return v.Version })
}

// Ssis variable.
type SsisVariableResponse struct {
	// Variable type.
	DataType *string `khulnasoft:"dataType"`
	// Variable description.
	Description *string `khulnasoft:"description"`
	// Variable id.
	Id *float64 `khulnasoft:"id"`
	// Variable name.
	Name *string `khulnasoft:"name"`
	// Whether variable is sensitive.
	Sensitive *bool `khulnasoft:"sensitive"`
	// Variable sensitive value.
	SensitiveValue *string `khulnasoft:"sensitiveValue"`
	// Variable value.
	Value *string `khulnasoft:"value"`
}

// Ssis variable.
type SsisVariableResponseArgs struct {
	// Variable type.
	DataType khulnasoftx.Input[*string] `khulnasoft:"dataType"`
	// Variable description.
	Description khulnasoftx.Input[*string] `khulnasoft:"description"`
	// Variable id.
	Id khulnasoftx.Input[*float64] `khulnasoft:"id"`
	// Variable name.
	Name khulnasoftx.Input[*string] `khulnasoft:"name"`
	// Whether variable is sensitive.
	Sensitive khulnasoftx.Input[*bool] `khulnasoft:"sensitive"`
	// Variable sensitive value.
	SensitiveValue khulnasoftx.Input[*string] `khulnasoft:"sensitiveValue"`
	// Variable value.
	Value khulnasoftx.Input[*string] `khulnasoft:"value"`
}

func (SsisVariableResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SsisVariableResponse)(nil)).Elem()
}

func (i SsisVariableResponseArgs) ToSsisVariableResponseOutput() SsisVariableResponseOutput {
	return i.ToSsisVariableResponseOutputWithContext(context.Background())
}

func (i SsisVariableResponseArgs) ToSsisVariableResponseOutputWithContext(ctx context.Context) SsisVariableResponseOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(SsisVariableResponseOutput)
}

func (i *SsisVariableResponseArgs) ToOutput(ctx context.Context) khulnasoftx.Output[*SsisVariableResponseArgs] {
	return khulnasoftx.Val(i)
}

// Ssis variable.
type SsisVariableResponseOutput struct{ *khulnasoft.OutputState }

func (SsisVariableResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SsisVariableResponse)(nil)).Elem()
}

func (o SsisVariableResponseOutput) ToSsisVariableResponseOutput() SsisVariableResponseOutput {
	return o
}

func (o SsisVariableResponseOutput) ToSsisVariableResponseOutputWithContext(ctx context.Context) SsisVariableResponseOutput {
	return o
}

func (o SsisVariableResponseOutput) ToOutput(ctx context.Context) khulnasoftx.Output[SsisVariableResponse] {
	return khulnasoftx.Output[SsisVariableResponse]{
		OutputState: o.OutputState,
	}
}

// Variable type.
func (o SsisVariableResponseOutput) DataType() khulnasoftx.Output[*string] {
	return khulnasoftx.Apply[SsisVariableResponse](o, func(v SsisVariableResponse) *string { return v.DataType })
}

// Variable description.
func (o SsisVariableResponseOutput) Description() khulnasoftx.Output[*string] {
	return khulnasoftx.Apply[SsisVariableResponse](o, func(v SsisVariableResponse) *string { return v.Description })
}

// Variable id.
func (o SsisVariableResponseOutput) Id() khulnasoftx.Output[*float64] {
	return khulnasoftx.Apply[SsisVariableResponse](o, func(v SsisVariableResponse) *float64 { return v.Id })
}

// Variable name.
func (o SsisVariableResponseOutput) Name() khulnasoftx.Output[*string] {
	return khulnasoftx.Apply[SsisVariableResponse](o, func(v SsisVariableResponse) *string { return v.Name })
}

// Whether variable is sensitive.
func (o SsisVariableResponseOutput) Sensitive() khulnasoftx.Output[*bool] {
	return khulnasoftx.Apply[SsisVariableResponse](o, func(v SsisVariableResponse) *bool { return v.Sensitive })
}

// Variable sensitive value.
func (o SsisVariableResponseOutput) SensitiveValue() khulnasoftx.Output[*string] {
	return khulnasoftx.Apply[SsisVariableResponse](o, func(v SsisVariableResponse) *string { return v.SensitiveValue })
}

// Variable value.
func (o SsisVariableResponseOutput) Value() khulnasoftx.Output[*string] {
	return khulnasoftx.Apply[SsisVariableResponse](o, func(v SsisVariableResponse) *string { return v.Value })
}

// An access key for the storage account.
type StorageAccountKeyResponse struct {
	// Creation time of the key, in round trip date format.
	CreationTime string `khulnasoft:"creationTime"`
	// Name of the key.
	KeyName string `khulnasoft:"keyName"`
	// Permissions for the key -- read-only or full permissions.
	Permissions string `khulnasoft:"permissions"`
	// Base 64-encoded value of the key.
	Value string `khulnasoft:"value"`
}

// An access key for the storage account.
type StorageAccountKeyResponseArgs struct {
	// Creation time of the key, in round trip date format.
	CreationTime khulnasoftx.Input[string] `khulnasoft:"creationTime"`
	// Name of the key.
	KeyName khulnasoftx.Input[string] `khulnasoft:"keyName"`
	// Permissions for the key -- read-only or full permissions.
	Permissions khulnasoftx.Input[string] `khulnasoft:"permissions"`
	// Base 64-encoded value of the key.
	Value khulnasoftx.Input[string] `khulnasoft:"value"`
}

func (StorageAccountKeyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountKeyResponse)(nil)).Elem()
}

func (i StorageAccountKeyResponseArgs) ToStorageAccountKeyResponseOutput() StorageAccountKeyResponseOutput {
	return i.ToStorageAccountKeyResponseOutputWithContext(context.Background())
}

func (i StorageAccountKeyResponseArgs) ToStorageAccountKeyResponseOutputWithContext(ctx context.Context) StorageAccountKeyResponseOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(StorageAccountKeyResponseOutput)
}

func (i *StorageAccountKeyResponseArgs) ToOutput(ctx context.Context) khulnasoftx.Output[*StorageAccountKeyResponseArgs] {
	return khulnasoftx.Val(i)
}

// An access key for the storage account.
type StorageAccountKeyResponseOutput struct{ *khulnasoft.OutputState }

func (StorageAccountKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountKeyResponse)(nil)).Elem()
}

func (o StorageAccountKeyResponseOutput) ToStorageAccountKeyResponseOutput() StorageAccountKeyResponseOutput {
	return o
}

func (o StorageAccountKeyResponseOutput) ToStorageAccountKeyResponseOutputWithContext(ctx context.Context) StorageAccountKeyResponseOutput {
	return o
}

func (o StorageAccountKeyResponseOutput) ToOutput(ctx context.Context) khulnasoftx.Output[StorageAccountKeyResponse] {
	return khulnasoftx.Output[StorageAccountKeyResponse]{
		OutputState: o.OutputState,
	}
}

// Creation time of the key, in round trip date format.
func (o StorageAccountKeyResponseOutput) CreationTime() khulnasoftx.Output[string] {
	return khulnasoftx.Apply[StorageAccountKeyResponse](o, func(v StorageAccountKeyResponse) string { return v.CreationTime })
}

// Name of the key.
func (o StorageAccountKeyResponseOutput) KeyName() khulnasoftx.Output[string] {
	return khulnasoftx.Apply[StorageAccountKeyResponse](o, func(v StorageAccountKeyResponse) string { return v.KeyName })
}

// Permissions for the key -- read-only or full permissions.
func (o StorageAccountKeyResponseOutput) Permissions() khulnasoftx.Output[string] {
	return khulnasoftx.Apply[StorageAccountKeyResponse](o, func(v StorageAccountKeyResponse) string { return v.Permissions })
}

// Base 64-encoded value of the key.
func (o StorageAccountKeyResponseOutput) Value() khulnasoftx.Output[string] {
	return khulnasoftx.Apply[StorageAccountKeyResponse](o, func(v StorageAccountKeyResponse) string { return v.Value })
}

func init() {
	khulnasoft.RegisterOutputType(BastionShareableLinkOutput{})
	khulnasoft.RegisterOutputType(SsisEnvironmentReferenceResponseOutput{})
	khulnasoft.RegisterOutputType(SsisEnvironmentResponseOutput{})
	khulnasoft.RegisterOutputType(SsisFolderResponseOutput{})
	khulnasoft.RegisterOutputType(SsisPackageResponseOutput{})
	khulnasoft.RegisterOutputType(SsisParameterResponseOutput{})
	khulnasoft.RegisterOutputType(SsisProjectResponseOutput{})
	khulnasoft.RegisterOutputType(SsisVariableResponseOutput{})
	khulnasoft.RegisterOutputType(StorageAccountKeyResponseOutput{})
}
