package main

import (
	awsstaticwebsite "github.com/khulnasoft/khulnasoft-aws-static-website/sdk/go/aws-static-website"
	"github.com/khulnasoft/khulnasoft-aws/sdk/v5/go/aws/cloudfront"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		_, err := awsstaticwebsite.NewWebsite(ctx, "websiteResource", &awsstaticwebsite.WebsiteArgs{
			SitePath:  khulnasoft.String("string"),
			IndexHTML: khulnasoft.String("string"),
			CacheTTL:  khulnasoft.Float64(0),
			CdnArgs: &awsstaticwebsite.CDNArgsArgs{
				CloudfrontFunctionAssociations: cloudfront.DistributionOrderedCacheBehaviorFunctionAssociationArray{
					&cloudfront.DistributionOrderedCacheBehaviorFunctionAssociationArgs{
						EventType:   khulnasoft.String("string"),
						FunctionArn: khulnasoft.String("string"),
					},
				},
				ForwardedValues: &cloudfront.DistributionDefaultCacheBehaviorForwardedValuesArgs{
					Cookies: &cloudfront.DistributionDefaultCacheBehaviorForwardedValuesCookiesArgs{
						Forward: khulnasoft.String("string"),
						WhitelistedNames: khulnasoft.StringArray{
							khulnasoft.String("string"),
						},
					},
					QueryString: khulnasoft.Bool(false),
					Headers: khulnasoft.StringArray{
						khulnasoft.String("string"),
					},
					QueryStringCacheKeys: khulnasoft.StringArray{
						khulnasoft.String("string"),
					},
				},
				LambdaFunctionAssociations: cloudfront.DistributionOrderedCacheBehaviorLambdaFunctionAssociationArray{
					&cloudfront.DistributionOrderedCacheBehaviorLambdaFunctionAssociationArgs{
						EventType:   khulnasoft.String("string"),
						LambdaArn:   khulnasoft.String("string"),
						IncludeBody: khulnasoft.Bool(false),
					},
				},
			},
			CertificateARN:          khulnasoft.String("string"),
			Error404:                khulnasoft.String("string"),
			AddWebsiteVersionHeader: khulnasoft.Bool(false),
			PriceClass:              khulnasoft.String("string"),
			AtomicDeployments:       khulnasoft.Bool(false),
			Subdomain:               khulnasoft.String("string"),
			TargetDomain:            khulnasoft.String("string"),
			WithCDN:                 khulnasoft.Bool(false),
			WithLogs:                khulnasoft.Bool(false),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
