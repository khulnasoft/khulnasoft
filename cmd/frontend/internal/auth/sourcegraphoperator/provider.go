package khulnasoftoperator

import (
	"net/http"
	"path"

	feAuth "github.com/khulnasoft/khulnasoft/cmd/frontend/auth"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/auth/openidconnect"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/auth/providers"
	"github.com/khulnasoft/khulnasoft/internal/auth"
	"github.com/khulnasoft/khulnasoft/internal/cloud"
	"github.com/khulnasoft/khulnasoft/schema"
)

// provider is an implementation of providers.Provider for the Khulnasoft
// Operator authentication, also referred to as "SOAP". There can only ever be
// one provider of this type, and it can only be provisioned through Cloud site
// configuration (see github.com/khulnasoft/khulnasoft/internal/cloud)
//
// SOAP is used to provision accounts for Khulnasoft teammates in Khulnasoft
// Cloud - for more details, refer to
// https://handbook.khulnasoft.com/departments/cloud/technical-docs/oidc_site_admin/#faq
type provider struct {
	config cloud.SchemaAuthProviderKhulnasoftOperator
	*openidconnect.Provider
}

var _ providers.Provider = (*provider)(nil)

// NewProvider creates and returns a new Khulnasoft Operator authentication
// provider using the given config.
func NewProvider(config cloud.SchemaAuthProviderKhulnasoftOperator, httpClient *http.Client) *provider {
	allowSignUp := true
	return &provider{
		config: config,
		Provider: openidconnect.NewProvider(
			schema.OpenIDConnectAuthProvider{
				AllowSignup:        &allowSignUp,
				ClientID:           config.ClientID,
				ClientSecret:       config.ClientSecret,
				ConfigID:           auth.KhulnasoftOperatorProviderType,
				DisplayName:        "Khulnasoft Operators",
				Issuer:             config.Issuer,
				RequireEmailDomain: "khulnasoft.com",
				Type:               auth.KhulnasoftOperatorProviderType,
			},
			authPrefix,
			path.Join(feAuth.AuthURLPrefix, "sourcegraph-operator", "callback"),
			httpClient,
		),
	}
}

// Config implements providers.Provider.
func (p *provider) Config() schema.AuthProviders {
	// NOTE: Intentionally omitting rest of the information unless absolutely
	// necessary because this provider is configured at the infrastructure level, and
	// those fields may expose sensitive information should not be visible to
	// non-Khulnasoft employees.
	return schema.AuthProviders{
		Openidconnect: &schema.OpenIDConnectAuthProvider{
			ConfigID:    auth.KhulnasoftOperatorProviderType,
			DisplayName: "Khulnasoft Operators",
		},
	}
}

func (p *provider) Type() providers.ProviderType {
	return providers.ProviderTypeOpenIDConnect
}
