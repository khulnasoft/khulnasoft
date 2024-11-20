package main

import (
	"github.com/khulnasoft/khulnasoft-aws/sdk/v4/go/aws"
	"github.com/khulnasoft/khulnasoft-aws/sdk/v4/go/aws/s3"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		provider, err := aws.NewProvider(ctx, "provider", &aws.ProviderArgs{
			Region: khulnasoft.String("us-west-2"),
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucket(ctx, "bucket1", nil, khulnasoft.Provider(provider), khulnasoft.DependsOn([]khulnasoft.Resource{
			provider,
		}), khulnasoft.Protect(true), khulnasoft.IgnoreChanges([]string{
			"bucket",
			"lifecycleRules[0]",
		}))
		if err != nil {
			return err
		}
		return nil
	})
}
