package main

import (
	storage "github.com/khulnasoft/khulnasoft-azure-native/sdk/go/azure/storage"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		someString := "foobar"
		typeVar := "Block"
		staticwebsite, err := storage.NewStorageAccountStaticWebsite(ctx, "staticwebsite", &storage.StorageAccountStaticWebsiteArgs{
			ResourceGroupName: khulnasoft.String(someString),
			AccountName:       khulnasoft.String(someString),
		})
		if err != nil {
			return err
		}
		_, err = storage.NewBlob(ctx, "faviconpng", &storage.BlobArgs{
			ResourceGroupName: khulnasoft.String(someString),
			AccountName:       khulnasoft.String(someString),
			ContainerName:     khulnasoft.String(someString),
			Type:              storage.BlobTypeBlock,
		})
		if err != nil {
			return err
		}
		_, err = storage.NewBlob(ctx, "_404html", &storage.BlobArgs{
			ResourceGroupName: khulnasoft.String(someString),
			AccountName:       khulnasoft.String(someString),
			ContainerName:     khulnasoft.String(someString),
			Type:              staticwebsite.IndexDocument.ApplyT(func(x *string) storage.BlobType { return storage.BlobType(*x) }).(storage.BlobTypeOutput),
		})
		if err != nil {
			return err
		}
		_, err = storage.NewBlob(ctx, "another", &storage.BlobArgs{
			ResourceGroupName: khulnasoft.String(someString),
			AccountName:       khulnasoft.String(someString),
			ContainerName:     khulnasoft.String(someString),
			Type:              storage.BlobType(typeVar),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
