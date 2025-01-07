package httpheader

import (
	"github.com/sourcegraph/log"

	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/auth/providers"
	"github.com/khulnasoft/khulnasoft/internal/conf"
	"github.com/khulnasoft/khulnasoft/internal/conf/conftypes"
	"github.com/khulnasoft/khulnasoft/internal/licensing"
	"github.com/khulnasoft/khulnasoft/schema"
)

// getProviderConfig returns the HTTP header auth provider config. At most 1 can be specified in
// site config; if there is more than 1, it returns multiple == true (which the caller should handle
// by returning an error and refusing to proceed with auth).
func getProviderConfig() (pc *schema.HTTPHeaderAuthProvider, multiple bool) {
	for _, p := range conf.Get().AuthProviders {
		if p.HttpHeader != nil {
			if pc != nil {
				return pc, true // multiple http-header auth providers
			}
			pc = p.HttpHeader
		}
	}
	return pc, false
}

const pkgName = "httpheader"

func Init() {
	conf.ContributeValidator(validateConfig)

	logger := log.Scoped(pkgName)
	go func() {
		conf.Watch(func() {
			newPC, _ := getProviderConfig()
			if newPC == nil {
				providers.Update(pkgName, nil)
				return
			}

			if err := licensing.Check(licensing.FeatureSSO); err != nil {
				logger.Error("Check license for SSO (HTTP header)", log.Error(err))
				providers.Update(pkgName, nil)
				return
			}
			providers.Update(pkgName, []providers.Provider{&provider{c: newPC}})
		})
	}()
}

func validateConfig(c conftypes.SiteConfigQuerier) (problems conf.Problems) {
	var httpHeaderAuthProviders int
	for _, p := range c.SiteConfig().AuthProviders {
		if p.HttpHeader != nil {
			httpHeaderAuthProviders++
		}
	}
	if httpHeaderAuthProviders >= 2 {
		problems = append(problems, conf.NewSiteProblem(`at most 1 HTTP header auth provider may be set in site config`))
	}
	return problems
}
