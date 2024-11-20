package main

import (
	"fmt"

	"github.com/khulnasoft/khulnasoft-aws/sdk/v5/go/aws/s3"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		myBucket, err := s3.NewBucket(ctx, "myBucket", &s3.BucketArgs{
			Website: &s3.BucketWebsiteArgs{
				IndexDocument: khulnasoft.String("index.html"),
			},
		})
		if err != nil {
			return err
		}
		ownershipControls, err := s3.NewBucketOwnershipControls(ctx, "ownershipControls", &s3.BucketOwnershipControlsArgs{
			Bucket: myBucket.ID(),
			Rule: &s3.BucketOwnershipControlsRuleArgs{
				ObjectOwnership: khulnasoft.String("ObjectWriter"),
			},
		})
		if err != nil {
			return err
		}
		publicAccessBlock, err := s3.NewBucketPublicAccessBlock(ctx, "publicAccessBlock", &s3.BucketPublicAccessBlockArgs{
			Bucket:          myBucket.ID(),
			BlockPublicAcls: khulnasoft.Bool(false),
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucketObject(ctx, "index.html", &s3.BucketObjectArgs{
			Bucket:      myBucket.ID(),
			Source:      khulnasoft.NewFileAsset("./index.html"),
			ContentType: khulnasoft.String("text/html"),
			Acl:         khulnasoft.String("public-read"),
		}, khulnasoft.DependsOn([]khulnasoft.Resource{
			publicAccessBlock,
			ownershipControls,
		}))
		if err != nil {
			return err
		}
		ctx.Export("bucketName", myBucket.ID())
		ctx.Export("bucketEndpoint", myBucket.WebsiteEndpoint.ApplyT(func(websiteEndpoint string) (string, error) {
			return fmt.Sprintf("http://%v", websiteEndpoint), nil
		}).(khulnasoft.StringOutput))
		return nil
	})
}
