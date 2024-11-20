package main

import (
	"github.com/khulnasoft/khulnasoft-azure-native/sdk/go/azure/cdn"
	"github.com/khulnasoft/khulnasoft-azure-native/sdk/go/azure/network"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		_, err := network.NewFrontDoor(ctx, "frontDoor", &network.FrontDoorArgs{
			ResourceGroupName: khulnasoft.String("someGroupName"),
			RoutingRules: network.RoutingRuleArray{
				&network.RoutingRuleArgs{
					RouteConfiguration: network.ForwardingConfiguration{
						OdataType: "#Microsoft.Azure.FrontDoor.Models.FrontdoorForwardingConfiguration",
						BackendPool: network.SubResource{
							Id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoors/frontDoor1/backendPools/backendPool1",
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = cdn.NewEndpoint(ctx, "endpoint", &cdn.EndpointArgs{
			Origins: cdn.DeepCreatedOriginArray{},
			DeliveryPolicy: &cdn.EndpointPropertiesUpdateParametersDeliveryPolicyArgs{
				Rules: cdn.DeliveryRuleArray{
					&cdn.DeliveryRuleArgs{
						Actions: khulnasoft.Array{
							cdn.DeliveryRuleCacheExpirationAction{
								Name: "CacheExpiration",
								Parameters: cdn.CacheExpirationActionParameters{
									CacheBehavior: cdn.CacheBehaviorOverride,
									CacheDuration: "10:10:09",
									CacheType:     cdn.CacheTypeAll,
									OdataType:     "#Microsoft.Azure.Cdn.Models.DeliveryRuleCacheExpirationActionParameters",
								},
							},
							cdn.DeliveryRuleResponseHeaderAction{
								Name: "ModifyResponseHeader",
								Parameters: cdn.HeaderActionParameters{
									HeaderAction: cdn.HeaderActionOverwrite,
									HeaderName:   "Access-Control-Allow-Origin",
									OdataType:    "#Microsoft.Azure.Cdn.Models.DeliveryRuleHeaderActionParameters",
									Value:        "*",
								},
							},
							cdn.DeliveryRuleRequestHeaderAction{
								Name: "ModifyRequestHeader",
								Parameters: cdn.HeaderActionParameters{
									HeaderAction: cdn.HeaderActionOverwrite,
									HeaderName:   "Accept-Encoding",
									OdataType:    "#Microsoft.Azure.Cdn.Models.DeliveryRuleHeaderActionParameters",
									Value:        "gzip",
								},
							},
						},
						Conditions: khulnasoft.Array{
							cdn.DeliveryRuleRemoteAddressCondition{
								Name: "RemoteAddress",
								Parameters: cdn.RemoteAddressMatchConditionParameters{
									MatchValues: []string{
										"192.168.1.0/24",
										"10.0.0.0/24",
									},
									NegateCondition: true,
									OdataType:       "#Microsoft.Azure.Cdn.Models.DeliveryRuleRemoteAddressConditionParameters",
									Operator:        cdn.RemoteAddressOperatorIPMatch,
								},
							},
						},
						Name:  khulnasoft.String("rule1"),
						Order: khulnasoft.Int(1),
					},
				},
			},
			EndpointName:         khulnasoft.String("endpoint1"),
			IsCompressionEnabled: khulnasoft.Bool(true),
			IsHttpAllowed:        khulnasoft.Bool(true),
			IsHttpsAllowed:       khulnasoft.Bool(true),
			Location:             khulnasoft.String("WestUs"),
			ProfileName:          khulnasoft.String("profileName"),
			ResourceGroupName:    khulnasoft.String("resourceGroupName"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
