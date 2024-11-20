package main

import (
	eventgrid "github.com/khulnasoft/khulnasoft-azure-native-sdk/eventgrid/v2"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		_, err := eventgrid.NewEventSubscription(ctx, "example", &eventgrid.EventSubscriptionArgs{
			Destination: &eventgrid.EventHubEventSubscriptionDestinationArgs{
				EndpointType: khulnasoft.String("EventHub"),
				ResourceId:   khulnasoft.String("example"),
			},
			ExpirationTimeUtc: khulnasoft.String("example"),
			Scope:             khulnasoft.String("example"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
