package main

import (
	"github.com/khulnasoft/khulnasoft-aws/sdk/v5/go/aws/ec2"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		invokeLookupVpc, err := ec2.LookupVpc(ctx, &ec2.LookupVpcArgs{
			Default: khulnasoft.BoolRef(true),
		}, nil)
		if err != nil {
			return err
		}
		_, err = ec2.NewSecurityGroup(ctx, "webSecurityGroup", &ec2.SecurityGroupArgs{
			VpcId: invokeLookupVpc.Id,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
