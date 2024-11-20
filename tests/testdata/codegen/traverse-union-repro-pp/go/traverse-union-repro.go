package main

import (
	"github.com/khulnasoft/khulnasoft-aws/sdk/v5/go/aws/fsx"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		_, err := fsx.NewOpenZfsFileSystem(ctx, "test", &fsx.OpenZfsFileSystemArgs{
			StorageCapacity: khulnasoft.Int(64),
			SubnetIds: khulnasoft.String{
				aws_subnet.Test1.Id,
			},
			DeploymentType:     khulnasoft.String("SINGLE_AZ_1"),
			ThroughputCapacity: khulnasoft.Int(64),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
