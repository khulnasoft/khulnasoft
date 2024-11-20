package main

import (
	"github.com/khulnasoft/khulnasoft-azure-native/sdk/go/azure/servicebus"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		_, err := servicebus.NewNamespaceAuthorizationRule(ctx, "namespaceAuthorizationRule", &servicebus.NamespaceAuthorizationRuleArgs{
			AuthorizationRuleName: khulnasoft.String("sdk-AuthRules-1788"),
			NamespaceName:         khulnasoft.String("sdk-Namespace-6914"),
			ResourceGroupName:     khulnasoft.String("ArunMonocle"),
			Rights: khulnasoft.StringArray{
				khulnasoft.String(servicebus.AccessRightsListen),
				khulnasoft.String(servicebus.AccessRightsSend),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
