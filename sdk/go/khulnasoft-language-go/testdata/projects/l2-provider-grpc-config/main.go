package main

import (
	"example.com/khulnasoft-config-grpc/sdk/go/configgrpc"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		// Cover interesting schema shapes.
		config_grpc_provider, err := configgrpc.NewProvider(ctx, "config_grpc_provider", &configgrpc.ProviderArgs{
			String1:     khulnasoft.String(""),
			String2:     khulnasoft.String("x"),
			String3:     khulnasoft.String("{}"),
			Int1:        khulnasoft.Int(0),
			Int2:        khulnasoft.Int(42),
			Num1:        khulnasoft.Float64(0),
			Num2:        khulnasoft.Float64(42.42),
			Bool1:       khulnasoft.Bool(true),
			Bool2:       khulnasoft.Bool(false),
			ListString1: khulnasoft.StringArray{},
			ListString2: khulnasoft.StringArray{
				khulnasoft.String(""),
				khulnasoft.String("foo"),
			},
			ListInt1: khulnasoft.IntArray{
				khulnasoft.Int(1),
				khulnasoft.Int(2),
			},
			MapString1: khulnasoft.StringMap{},
			MapString2: khulnasoft.StringMap{
				"key1": khulnasoft.String("value1"),
				"key2": khulnasoft.String("value2"),
			},
			MapInt1: khulnasoft.IntMap{
				"key1": khulnasoft.Int(0),
				"key2": khulnasoft.Int(42),
			},
			ObjString1: &configgrpc.Tstring1Args{},
			ObjString2: &configgrpc.Tstring2Args{
				X: khulnasoft.String("x-value"),
			},
			ObjInt1: &configgrpc.Tint1Args{
				X: khulnasoft.Int(42),
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
