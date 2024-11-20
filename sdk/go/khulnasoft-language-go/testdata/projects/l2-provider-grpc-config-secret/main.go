package main

import (
	"example.com/khulnasoft-config-grpc/sdk/go/configgrpc"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		// This provider covers scenarios where user passes secret values to the provider.
		config_grpc_provider, err := configgrpc.NewProvider(ctx, "config_grpc_provider", &configgrpc.ProviderArgs{
			String1: configgrpc.ToSecretOutput(ctx, configgrpc.ToSecretOutputArgs{
				String1: khulnasoft.String("SECRET"),
			}, nil).ApplyT(func(invoke configgrpc.ToSecretResult) (string, error) {
				return invoke.String1, nil
			}).(khulnasoft.StringOutput),
			Int1: configgrpc.ToSecretOutput(ctx, configgrpc.ToSecretOutputArgs{
				Int1: khulnasoft.Int(1234567890),
			}, nil).ApplyT(func(invoke configgrpc.ToSecretResult) (int, error) {
				return invoke.Int1, nil
			}).(khulnasoft.IntOutput),
			Num1: configgrpc.ToSecretOutput(ctx, configgrpc.ToSecretOutputArgs{
				Num1: khulnasoft.Float64(123456.789),
			}, nil).ApplyT(func(invoke configgrpc.ToSecretResult) (float64, error) {
				return invoke.Num1, nil
			}).(khulnasoft.Float64Output),
			Bool1: configgrpc.ToSecretOutput(ctx, configgrpc.ToSecretOutputArgs{
				Bool1: khulnasoft.Bool(true),
			}, nil).ApplyT(func(invoke configgrpc.ToSecretResult) (bool, error) {
				return invoke.Bool1, nil
			}).(khulnasoft.BoolOutput),
			ListString1: configgrpc.ToSecretOutput(ctx, configgrpc.ToSecretOutputArgs{
				ListString1: khulnasoft.StringArray{
					khulnasoft.String("SECRET"),
					khulnasoft.String("SECRET2"),
				},
			}, nil).ApplyT(func(invoke configgrpc.ToSecretResult) ([]string, error) {
				return invoke.ListString1, nil
			}).(khulnasoft.StringArrayOutput),
			ListString2: khulnasoft.StringArray{
				khulnasoft.String("VALUE"),
				configgrpc.ToSecretOutput(ctx, configgrpc.ToSecretOutputArgs{
					String1: khulnasoft.String("SECRET"),
				}, nil).ApplyT(func(invoke configgrpc.ToSecretResult) (string, error) {
					return invoke.String1, nil
				}).(khulnasoft.StringOutput),
			},
			MapString2: khulnasoft.StringMap{
				"key1": khulnasoft.String("value1"),
				"key2": configgrpc.ToSecretOutput(ctx, configgrpc.ToSecretOutputArgs{
					String1: khulnasoft.String("SECRET"),
				}, nil).ApplyT(func(invoke configgrpc.ToSecretResult) (string, error) {
					return invoke.String1, nil
				}).(khulnasoft.StringOutput),
			},
			ObjString2: &configgrpc.Tstring2Args{
				X: configgrpc.ToSecretOutput(ctx, configgrpc.ToSecretOutputArgs{
					String1: khulnasoft.String("SECRET"),
				}, nil).ApplyT(func(invoke configgrpc.ToSecretResult) (string, error) {
					return invoke.String1, nil
				}).(khulnasoft.StringOutput),
			},
		})
		if err != nil {
			return err
		}
		_, err = configgrpc.NewConfigFetcher(ctx, "config", nil, khulnasoft.Provider(config_grpc_provider))
		if err != nil {
			return err
		}
		return nil
	})
}
