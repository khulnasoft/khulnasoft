package khulnasoftoperator

import (
	"testing"

	"github.com/khulnasoft/khulnasoft/internal/cloud"
	"github.com/khulnasoft/khulnasoft/internal/conf"
	"github.com/khulnasoft/khulnasoft/schema"
)

func TestValidateConfig(t *testing.T) {
	cloud.MockSiteConfig(
		t,
		&cloud.SchemaSiteConfig{
			AuthProviders: &cloud.SchemaAuthProviders{
				KhulnasoftOperator: &cloud.SchemaAuthProviderKhulnasoftOperator{
					Issuer: "https://example.com/alice",
				},
			},
		},
	)

	conf.TestValidator(
		t,
		conf.Unified{
			SiteConfiguration: schema.SiteConfiguration{},
		},
		validateConfig,
		conf.NewSiteProblems("Khulnasoft Operator authentication provider requires `externalURL` to be set to the external URL of your site (example: https://sourcegraph.example.com)"),
	)
}
