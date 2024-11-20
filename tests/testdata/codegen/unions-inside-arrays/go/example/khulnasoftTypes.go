// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"context"
	"reflect"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoftx"
	"unions-inside-arrays/example/internal"
)

var _ = internal.GetEnvOrDefault

type ServerPropertiesForReplica struct {
	CreateMode string  `khulnasoft:"createMode"`
	Version    *string `khulnasoft:"version"`
}

// ServerPropertiesForReplicaInput is an input type that accepts ServerPropertiesForReplicaArgs and ServerPropertiesForReplicaOutput values.
// You can construct a concrete instance of `ServerPropertiesForReplicaInput` via:
//
//	ServerPropertiesForReplicaArgs{...}
type ServerPropertiesForReplicaInput interface {
	khulnasoft.Input

	ToServerPropertiesForReplicaOutput() ServerPropertiesForReplicaOutput
	ToServerPropertiesForReplicaOutputWithContext(context.Context) ServerPropertiesForReplicaOutput
}

type ServerPropertiesForReplicaArgs struct {
	CreateMode khulnasoft.StringInput    `khulnasoft:"createMode"`
	Version    khulnasoft.StringPtrInput `khulnasoft:"version"`
}

func (ServerPropertiesForReplicaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerPropertiesForReplica)(nil)).Elem()
}

func (i ServerPropertiesForReplicaArgs) ToServerPropertiesForReplicaOutput() ServerPropertiesForReplicaOutput {
	return i.ToServerPropertiesForReplicaOutputWithContext(context.Background())
}

func (i ServerPropertiesForReplicaArgs) ToServerPropertiesForReplicaOutputWithContext(ctx context.Context) ServerPropertiesForReplicaOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(ServerPropertiesForReplicaOutput)
}

func (i ServerPropertiesForReplicaArgs) ToOutput(ctx context.Context) khulnasoftx.Output[ServerPropertiesForReplica] {
	return khulnasoftx.Output[ServerPropertiesForReplica]{
		OutputState: i.ToServerPropertiesForReplicaOutputWithContext(ctx).OutputState,
	}
}

type ServerPropertiesForReplicaOutput struct{ *khulnasoft.OutputState }

func (ServerPropertiesForReplicaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerPropertiesForReplica)(nil)).Elem()
}

func (o ServerPropertiesForReplicaOutput) ToServerPropertiesForReplicaOutput() ServerPropertiesForReplicaOutput {
	return o
}

func (o ServerPropertiesForReplicaOutput) ToServerPropertiesForReplicaOutputWithContext(ctx context.Context) ServerPropertiesForReplicaOutput {
	return o
}

func (o ServerPropertiesForReplicaOutput) ToOutput(ctx context.Context) khulnasoftx.Output[ServerPropertiesForReplica] {
	return khulnasoftx.Output[ServerPropertiesForReplica]{
		OutputState: o.OutputState,
	}
}

func (o ServerPropertiesForReplicaOutput) CreateMode() khulnasoft.StringOutput {
	return o.ApplyT(func(v ServerPropertiesForReplica) string { return v.CreateMode }).(khulnasoft.StringOutput)
}

func (o ServerPropertiesForReplicaOutput) Version() khulnasoft.StringPtrOutput {
	return o.ApplyT(func(v ServerPropertiesForReplica) *string { return v.Version }).(khulnasoft.StringPtrOutput)
}

type ServerPropertiesForRestore struct {
	CreateMode         string `khulnasoft:"createMode"`
	RestorePointInTime string `khulnasoft:"restorePointInTime"`
}

// ServerPropertiesForRestoreInput is an input type that accepts ServerPropertiesForRestoreArgs and ServerPropertiesForRestoreOutput values.
// You can construct a concrete instance of `ServerPropertiesForRestoreInput` via:
//
//	ServerPropertiesForRestoreArgs{...}
type ServerPropertiesForRestoreInput interface {
	khulnasoft.Input

	ToServerPropertiesForRestoreOutput() ServerPropertiesForRestoreOutput
	ToServerPropertiesForRestoreOutputWithContext(context.Context) ServerPropertiesForRestoreOutput
}

type ServerPropertiesForRestoreArgs struct {
	CreateMode         khulnasoft.StringInput `khulnasoft:"createMode"`
	RestorePointInTime khulnasoft.StringInput `khulnasoft:"restorePointInTime"`
}

func (ServerPropertiesForRestoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerPropertiesForRestore)(nil)).Elem()
}

func (i ServerPropertiesForRestoreArgs) ToServerPropertiesForRestoreOutput() ServerPropertiesForRestoreOutput {
	return i.ToServerPropertiesForRestoreOutputWithContext(context.Background())
}

func (i ServerPropertiesForRestoreArgs) ToServerPropertiesForRestoreOutputWithContext(ctx context.Context) ServerPropertiesForRestoreOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(ServerPropertiesForRestoreOutput)
}

func (i ServerPropertiesForRestoreArgs) ToOutput(ctx context.Context) khulnasoftx.Output[ServerPropertiesForRestore] {
	return khulnasoftx.Output[ServerPropertiesForRestore]{
		OutputState: i.ToServerPropertiesForRestoreOutputWithContext(ctx).OutputState,
	}
}

type ServerPropertiesForRestoreOutput struct{ *khulnasoft.OutputState }

func (ServerPropertiesForRestoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerPropertiesForRestore)(nil)).Elem()
}

func (o ServerPropertiesForRestoreOutput) ToServerPropertiesForRestoreOutput() ServerPropertiesForRestoreOutput {
	return o
}

func (o ServerPropertiesForRestoreOutput) ToServerPropertiesForRestoreOutputWithContext(ctx context.Context) ServerPropertiesForRestoreOutput {
	return o
}

func (o ServerPropertiesForRestoreOutput) ToOutput(ctx context.Context) khulnasoftx.Output[ServerPropertiesForRestore] {
	return khulnasoftx.Output[ServerPropertiesForRestore]{
		OutputState: o.OutputState,
	}
}

func (o ServerPropertiesForRestoreOutput) CreateMode() khulnasoft.StringOutput {
	return o.ApplyT(func(v ServerPropertiesForRestore) string { return v.CreateMode }).(khulnasoft.StringOutput)
}

func (o ServerPropertiesForRestoreOutput) RestorePointInTime() khulnasoft.StringOutput {
	return o.ApplyT(func(v ServerPropertiesForRestore) string { return v.RestorePointInTime }).(khulnasoft.StringOutput)
}

func init() {
	khulnasoft.RegisterInputType(reflect.TypeOf((*ServerPropertiesForReplicaInput)(nil)).Elem(), ServerPropertiesForReplicaArgs{})
	khulnasoft.RegisterInputType(reflect.TypeOf((*ServerPropertiesForRestoreInput)(nil)).Elem(), ServerPropertiesForRestoreArgs{})
	khulnasoft.RegisterOutputType(ServerPropertiesForReplicaOutput{})
	khulnasoft.RegisterOutputType(ServerPropertiesForRestoreOutput{})
}