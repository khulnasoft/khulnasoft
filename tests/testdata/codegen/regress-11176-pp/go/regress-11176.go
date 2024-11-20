package main

import (
	"github.com/khulnasoft/khulnasoft-aws/sdk/v5/go/aws/ecs"
	awsxecs "github.com/khulnasoft/khulnasoft-awsx/sdk/go/awsx/ecs"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		cluster, err := ecs.NewCluster(ctx, "cluster", nil)
		if err != nil {
			return err
		}
		_, err = awsxecs.NewFargateService(ctx, "nginx", &awsxecs.FargateServiceArgs{
			Cluster: cluster.Arn,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
