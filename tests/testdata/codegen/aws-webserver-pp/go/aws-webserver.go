package main

import (
	"github.com/khulnasoft/khulnasoft-aws/sdk/v5/go/aws"
	"github.com/khulnasoft/khulnasoft-aws/sdk/v5/go/aws/ec2"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		// Create a new security group for port 80.
		securityGroup, err := ec2.NewSecurityGroup(ctx, "securityGroup", &ec2.SecurityGroupArgs{
			Ingress: ec2.SecurityGroupIngressArray{
				&ec2.SecurityGroupIngressArgs{
					Protocol: khulnasoft.String("tcp"),
					FromPort: khulnasoft.Int(0),
					ToPort:   khulnasoft.Int(0),
					CidrBlocks: khulnasoft.StringArray{
						khulnasoft.String("0.0.0.0/0"),
					},
				},
			},
		})
		if err != nil {
			return err
		}
		// Get the ID for the latest Amazon Linux AMI.
		ami, err := aws.GetAmi(ctx, &aws.GetAmiArgs{
			Filters: []aws.GetAmiFilter{
				{
					Name: "name",
					Values: []string{
						"amzn-ami-hvm-*-x86_64-ebs",
					},
				},
			},
			Owners: []string{
				"137112412989",
			},
			MostRecent: khulnasoft.BoolRef(true),
		}, nil)
		if err != nil {
			return err
		}
		// Create a simple web server using the startup script for the instance.
		server, err := ec2.NewInstance(ctx, "server", &ec2.InstanceArgs{
			Tags: khulnasoft.StringMap{
				"Name": khulnasoft.String("web-server-www"),
			},
			InstanceType: khulnasoft.String(ec2.InstanceType_T2_Micro),
			SecurityGroups: khulnasoft.StringArray{
				securityGroup.Name,
			},
			Ami:      khulnasoft.String(ami.Id),
			UserData: khulnasoft.String("#!/bin/bash\necho \"Hello, World!\" > index.html\nnohup python -m SimpleHTTPServer 80 &\n"),
		})
		if err != nil {
			return err
		}
		ctx.Export("publicIp", server.PublicIp)
		ctx.Export("publicHostName", server.PublicDns)
		return nil
	})
}
