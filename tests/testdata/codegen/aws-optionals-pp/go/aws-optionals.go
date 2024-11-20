package main

import (
	"github.com/khulnasoft/khulnasoft-aws/sdk/v5/go/aws/iam"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		policyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Sid: khulnasoft.StringRef("1"),
					Actions: []string{
						"s3:ListAllMyBuckets",
						"s3:GetBucketLocation",
					},
					Resources: []string{
						"arn:aws:s3:::*",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		_, err = iam.NewPolicy(ctx, "example", &iam.PolicyArgs{
			Name:   khulnasoft.String("example_policy"),
			Path:   khulnasoft.String("/"),
			Policy: khulnasoft.String(policyDocument.Json),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
