package main

import (
	"github.com/khulnasoft/khulnasoft-aws/sdk/v5/go/aws/rds"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		_, err := rds.NewCluster(ctx, "dbCluster", &rds.ClusterArgs{
			MasterPassword: khulnasoft.ToSecret("foobar").(khulnasoft.StringOutput),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
