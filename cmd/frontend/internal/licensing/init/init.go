package init

import (
	"context"

	"github.com/sourcegraph/log"

	"github.com/khulnasoft/khulnasoft/cmd/frontend/enterprise"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/dotcom/productsubscription"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/licensing/enforcement"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/licensing/resolvers"
	"github.com/khulnasoft/khulnasoft/internal/codeintel"
	"github.com/khulnasoft/khulnasoft/internal/conf/conftypes"
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/dotcom"
	"github.com/khulnasoft/khulnasoft/internal/env"
	"github.com/khulnasoft/khulnasoft/internal/goroutine"
	"github.com/khulnasoft/khulnasoft/internal/observation"
)

var (
	enableUpcomingLicenseExpirationChecker = env.MustGetBool("DOTCOM_ENABLE_UPCOMING_LICENSE_EXPIRATION_CHECKER", true,
		"If false, we do not monitor for upcoming license expirations to post in Slack.")
	enableAnomalousLicenseChecker = env.MustGetBool("DOTCOM_ENABLE_ANOMALOUS_LICENSE_CHECKER", true,
		"If false, we do not monitor for anomalous license checks to post in Slack.")
)

func Init(
	ctx context.Context,
	observationCtx *observation.Context,
	db database.DB,
	codeIntelServices codeintel.Services,
	conf conftypes.UnifiedWatchable,
	enterpriseServices *enterprise.Services,
) error {
	// Enforce the license's max user count by preventing the creation of new users when the max is
	// reached.
	database.BeforeCreateUser = enforcement.NewBeforeCreateUserHook()

	// Enforce non-site admin roles in Free tier.
	database.AfterCreateUser = enforcement.NewAfterCreateUserHook()

	// Enforce site admin creation rules.
	database.BeforeSetUserIsSiteAdmin = enforcement.NewBeforeSetUserIsSiteAdmin()

	enterpriseServices.LicenseResolver = resolvers.LicenseResolver{}

	if dotcom.KhulnasoftDotComMode() {
		logger := log.Scoped("licensing")
		if enableUpcomingLicenseExpirationChecker {
			goroutine.Go(func() {
				productsubscription.StartCheckForUpcomingLicenseExpirations(logger, db)
			})
		}
		if enableAnomalousLicenseChecker {
			goroutine.Go(func() {
				productsubscription.StartCheckForAnomalousLicenseUsage(logger, db)
			})
		}
	}

	return nil
}
