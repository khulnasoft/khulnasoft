package main

import (
	"fmt"

	"github.com/khulnasoft/khulnasoft-aws/sdk/v5/go/aws/s3"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		mybucket, err := s3.NewBucket(ctx, "mybucket", &s3.BucketArgs{
			Website: &s3.BucketWebsiteArgs{
				IndexDocument: khulnasoft.String("index.html"),
			},
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucketObject(ctx, "indexhtml", &s3.BucketObjectArgs{
			Bucket:      mybucket.ID(),
			Source:      khulnasoft.NewStringAsset("<h1>Hello, world!</h1>"),
			Acl:         khulnasoft.String("public-read"),
			ContentType: khulnasoft.String("text/html"),
		})
		if err != nil {
			return err
		}
		ctx.Export("bucketEndpoint", mybucket.WebsiteEndpoint.ApplyT(func(websiteEndpoint string) (string, error) {
			return fmt.Sprintf("http://%v", websiteEndpoint), nil
		}).(khulnasoft.StringOutput))
		return nil
	})
}
