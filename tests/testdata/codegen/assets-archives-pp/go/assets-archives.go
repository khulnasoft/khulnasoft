package main

import (
	"github.com/khulnasoft/khulnasoft-aws/sdk/v5/go/aws/lambda"
	"github.com/khulnasoft/khulnasoft-aws/sdk/v5/go/aws/s3"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		siteBucket, err := s3.NewBucket(ctx, "siteBucket", nil)
		if err != nil {
			return err
		}
		_, err = s3.NewBucketObject(ctx, "testFileAsset", &s3.BucketObjectArgs{
			Bucket: siteBucket.ID(),
			Source: khulnasoft.NewFileAsset("file.txt"),
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucketObject(ctx, "testStringAsset", &s3.BucketObjectArgs{
			Bucket: siteBucket.ID(),
			Source: khulnasoft.NewStringAsset("<h1>File contents</h1>"),
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucketObject(ctx, "testRemoteAsset", &s3.BucketObjectArgs{
			Bucket: siteBucket.ID(),
			Source: khulnasoft.NewRemoteAsset("https://khulnasoft.test"),
		})
		if err != nil {
			return err
		}
		_, err = lambda.NewFunction(ctx, "testFileArchive", &lambda.FunctionArgs{
			Role: siteBucket.Arn,
			Code: khulnasoft.NewFileArchive("file.tar.gz"),
		})
		if err != nil {
			return err
		}
		_, err = lambda.NewFunction(ctx, "testRemoteArchive", &lambda.FunctionArgs{
			Role: siteBucket.Arn,
			Code: khulnasoft.NewRemoteArchive("https://khulnasoft.test/foo.tar.gz"),
		})
		if err != nil {
			return err
		}
		_, err = lambda.NewFunction(ctx, "testAssetArchive", &lambda.FunctionArgs{
			Role: siteBucket.Arn,
			Code: khulnasoft.NewAssetArchive(map[string]interface{}{
				"file.txt":   khulnasoft.NewFileAsset("file.txt"),
				"string.txt": khulnasoft.NewStringAsset("<h1>File contents</h1>"),
				"remote.txt": khulnasoft.NewRemoteAsset("https://khulnasoft.test"),
				"file.tar":   khulnasoft.NewFileArchive("file.tar.gz"),
				"remote.tar": khulnasoft.NewRemoteArchive("https://khulnasoft.test/foo.tar.gz"),
				".nestedDir": khulnasoft.NewAssetArchive(map[string]interface{}{
					"file.txt":   khulnasoft.NewFileAsset("file.txt"),
					"string.txt": khulnasoft.NewStringAsset("<h1>File contents</h1>"),
					"remote.txt": khulnasoft.NewRemoteAsset("https://khulnasoft.test"),
					"file.tar":   khulnasoft.NewFileArchive("file.tar.gz"),
					"remote.tar": khulnasoft.NewRemoteArchive("https://khulnasoft.test/foo.tar.gz"),
				}),
			}),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
