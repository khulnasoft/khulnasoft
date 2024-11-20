package main

import (
	"github.com/khulnasoft/khulnasoft-azure-native/sdk/go/azure/storage"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		_, err := storage.NewStorageAccount(ctx, "storageAccounts", &storage.StorageAccountArgs{
			AccountName:       khulnasoft.String("sto4445"),
			Kind:              khulnasoft.String(storage.KindBlockBlobStorage),
			Location:          khulnasoft.String("eastus"),
			ResourceGroupName: khulnasoft.String("res9101"),
			Sku: &storage.SkuArgs{
				Name: khulnasoft.String(storage.SkuName_Premium_LRS),
			},
			NetworkRuleSet: &storage.NetworkRuleSetArgs{
				DefaultAction: storage.DefaultActionAllow,
				IpRules:       storage.IPRuleArray{},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
