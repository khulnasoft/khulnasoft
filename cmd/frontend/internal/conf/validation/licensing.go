package validation

import (
	"github.com/khulnasoft/khulnasoft/internal/conf"
	"github.com/khulnasoft/khulnasoft/internal/conf/conftypes"
	"github.com/khulnasoft/khulnasoft/internal/licensing"
)

func init() {
	conf.ContributeValidator(func(cfg conftypes.SiteConfigQuerier) conf.Problems {
		if cfg.SiteConfig().LicenseKey != "" {
			info, _, err := licensing.ParseProductLicenseKeyWithBuiltinOrGenerationKey(cfg.SiteConfig().LicenseKey)
			if err != nil {
				return conf.NewSiteProblems("should provide a valid license key")
			} else if err = info.HasUnknownPlan(); err != nil {
				return conf.NewSiteProblems(err.Error())
			}
		}
		return nil
	})
}
