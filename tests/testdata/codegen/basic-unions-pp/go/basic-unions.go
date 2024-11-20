package main

import (
	basicunions "github.com/khulnasoft/khulnasoft-basic-unions/sdk/v4/go/basic-unions"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		// properties field is bound to union case ServerPropertiesForReplica
		_, err := basicunions.NewExampleServer(ctx, "replica", &basicunions.ExampleServerArgs{
			Properties: &basicunions.ServerPropertiesForReplicaArgs{
				CreateMode: khulnasoft.String("Replica"),
				Version:    khulnasoft.String("0.1.0-dev"),
			},
		})
		if err != nil {
			return err
		}
		// properties field is bound to union case ServerPropertiesForRestore
		_, err = basicunions.NewExampleServer(ctx, "restore", &basicunions.ExampleServerArgs{
			Properties: &basicunions.ServerPropertiesForRestoreArgs{
				CreateMode:         khulnasoft.String("PointInTimeRestore"),
				RestorePointInTime: khulnasoft.String("example"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
