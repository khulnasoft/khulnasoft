package main

import (
	"example.com/khulnasoft-config-grpc/sdk/go/configgrpc"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		// This provider covers scenarios where configuration properties are marked as secret in the schema.
		config_grpc_provider, err := configgrpc.NewProvider(ctx, "config_grpc_provider", &configgrpc.ProviderArgs{
			SecretString1: khulnasoft.String("SECRET"),
			SecretInt1:    khulnasoft.Int(16),
			SecretNum1:    khulnasoft.Float64(123456.789),
			SecretBool1:   khulnasoft.Bool(true),
			ListSecretString1: khulnasoft.StringArray{
				khulnasoft.String("SECRET"),
				khulnasoft.String("SECRET2"),
			},
			MapSecretString1: khulnasoft.StringMap{
				"key1": khulnasoft.String("SECRET"),
				"key2": khulnasoft.String("SECRET2"),
			},
			ObjSecretString1: &configgrpc.TsecretString1Args{
				SecretX: khulnasoft.String("SECRET"),
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
