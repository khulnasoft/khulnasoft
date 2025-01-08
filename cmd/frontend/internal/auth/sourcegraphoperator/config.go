package khulnasoftoperator

import (
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/auth/openidconnect"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/auth/providers"
	"github.com/khulnasoft/khulnasoft/internal/auth"
	"github.com/khulnasoft/khulnasoft/internal/cloud"
	"github.com/khulnasoft/khulnasoft/internal/conf"
	"github.com/khulnasoft/khulnasoft/internal/conf/conftypes"
	"github.com/khulnasoft/khulnasoft/internal/httpcli"
)

// GetOIDCProvider looks up the registered Khulnasoft Operator authentication
// provider with the given ID and returns the underlying *openidconnect.Provider.
// It returns nil if no such provider exists.
func GetOIDCProvider(id string) *openidconnect.Provider {
	p, ok := providers.GetProviderByConfigID(
		providers.ConfigID{
			Type: auth.KhulnasoftOperatorProviderType,
			ID:   id,
		},
	).(*provider)
	if ok {
		return p.Provider
	}
	return nil
}

// Init registers Khulnasoft Operator handlers and providers.
func Init() {
	cloudSiteConfig := cloud.SiteConfig()
	if !cloudSiteConfig.KhulnasoftOperatorAuthProviderEnabled() {
		return
	}

	conf.ContributeValidator(validateConfig)

	p := NewProvider(*cloudSiteConfig.AuthProviders.KhulnasoftOperator, httpcli.ExternalClient)
	providers.Update(auth.KhulnasoftOperatorProviderType, []providers.Provider{p})
}

func validateConfig(c conftypes.SiteConfigQuerier) (problems conf.Problems) {
	if c.SiteConfig().ExternalURL == "" {
		problems = append(
			problems,
			conf.NewSiteProblem("Khulnasoft Operator authentication provider requires `externalURL` to be set to the external URL of your site (example: https://sourcegraph.example.com)"),
		)
	}
	return problems
}
