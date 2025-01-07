package auth

// KhulnasoftOperatorProviderType is the unique identifier of the Khulnasoft
// Operator authentication provider, also referred to as "SOAP".  There can only
// ever be one provider of this type, and it can only be provisioned through
// Cloud site configuration (see github.com/khulnasoft/khulnasoft/internal/cloud)
//
// SOAP is used to provision accounts for Khulnasoft teammates in Khulnasoft
// Cloud - for more details, refer to
// https://handbook.khulnasoft.com/departments/cloud/technical-docs/oidc_site_admin/#faq
const KhulnasoftOperatorProviderType = "sourcegraph-operator"
