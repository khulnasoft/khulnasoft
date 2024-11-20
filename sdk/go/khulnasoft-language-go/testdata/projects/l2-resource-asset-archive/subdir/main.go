package main

import (
	"example.com/khulnasoft-asset-archive/sdk/go/v5/assetarchive"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		_, err := assetarchive.NewAssetResource(ctx, "ass", &assetarchive.AssetResourceArgs{
			Value: khulnasoft.NewFileAsset("../test.txt"),
		})
		if err != nil {
			return err
		}
		_, err = assetarchive.NewArchiveResource(ctx, "arc", &assetarchive.ArchiveResourceArgs{
			Value: khulnasoft.NewFileArchive("../archive.tar"),
		})
		if err != nil {
			return err
		}
		_, err = assetarchive.NewArchiveResource(ctx, "dir", &assetarchive.ArchiveResourceArgs{
			Value: khulnasoft.NewFileArchive("../folder"),
		})
		if err != nil {
			return err
		}
		_, err = assetarchive.NewArchiveResource(ctx, "assarc", &assetarchive.ArchiveResourceArgs{
			Value: khulnasoft.NewAssetArchive(map[string]interface{}{
				"string":  khulnasoft.NewStringAsset("file contents"),
				"file":    khulnasoft.NewFileAsset("../test.txt"),
				"folder":  khulnasoft.NewFileArchive("../folder"),
				"archive": khulnasoft.NewFileArchive("../archive.tar"),
			}),
		})
		if err != nil {
			return err
		}
		_, err = assetarchive.NewAssetResource(ctx, "remoteass", &assetarchive.AssetResourceArgs{
			Value: khulnasoft.NewRemoteAsset("https://raw.githubusercontent.com/khulnasoft/khulnasoft/master/cmd/khulnasoft-test-language/testdata/l2-resource-asset-archive/test.txt"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
