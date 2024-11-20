package main

import (
	"github.com/khulnasoft/khulnasoft-azure/sdk/v4/go/azure/core"
	"github.com/khulnasoft/khulnasoft-azure/sdk/v4/go/azure/storage"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft/config"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		cfg := config.New(ctx, "")
		// The name of the storage account
		storageAccountNameParam := cfg.Require("storageAccountNameParam")
		// The name of the resource group
		resourceGroupNameParam := cfg.Require("resourceGroupNameParam")
		resourceGroupVar, err := core.LookupResourceGroup(ctx, &core.LookupResourceGroupArgs{
			Name: resourceGroupNameParam,
		}, nil)
		if err != nil {
			return err
		}
		locationParam := resourceGroupVar.Location
		if param := cfg.Get("locationParam"); param != "" {
			locationParam = param
		}
		storageAccountTierParam := "Standard"
		if param := cfg.Get("storageAccountTierParam"); param != "" {
			storageAccountTierParam = param
		}
		storageAccountTypeReplicationParam := "LRS"
		if param := cfg.Get("storageAccountTypeReplicationParam"); param != "" {
			storageAccountTypeReplicationParam = param
		}
		storageAccountResource, err := storage.NewAccount(ctx, "storageAccountResource", &storage.AccountArgs{
			Name:                   khulnasoft.String(storageAccountNameParam),
			AccountKind:            khulnasoft.String("StorageV2"),
			Location:               khulnasoft.String(locationParam),
			ResourceGroupName:      khulnasoft.String(resourceGroupNameParam),
			AccountTier:            khulnasoft.String(storageAccountTierParam),
			AccountReplicationType: khulnasoft.String(storageAccountTypeReplicationParam),
		})
		if err != nil {
			return err
		}
		ctx.Export("storageAccountNameOut", storageAccountResource.Name)
		return nil
	})
}
