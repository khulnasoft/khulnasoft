package main

import (
	"github.com/khulnasoft/khulnasoft-aws-native/sdk/go/aws/iam"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		_, err := iam.NewRole(ctx, "role", &iam.RoleArgs{
			RoleName: khulnasoft.String("ScriptIAMRole"),
			AssumeRolePolicyDocument: khulnasoft.Any(map[string]interface{}{
				"Version": "2012-10-17",
				"Statement": []map[string]interface{}{
					map[string]interface{}{
						"Effect": "Allow",
						"Action": "sts:AssumeRole",
						"Principal": map[string]interface{}{
							"Service": []string{
								"cloudformation.amazonaws.com",
								"gamelift.amazonaws.com",
							},
						},
					},
				},
			}),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
