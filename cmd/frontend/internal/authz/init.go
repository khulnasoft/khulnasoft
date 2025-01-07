package authz

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/sourcegraph/log"

	"github.com/khulnasoft/khulnasoft/cmd/frontend/enterprise"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/graphqlbackend"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/authz/resolvers"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/authz/webhooks"
	"github.com/khulnasoft/khulnasoft/internal/authz"
	"github.com/khulnasoft/khulnasoft/internal/authz/providers"
	srp "github.com/khulnasoft/khulnasoft/internal/authz/subrepoperms"
	"github.com/khulnasoft/khulnasoft/internal/codeintel"
	"github.com/khulnasoft/khulnasoft/internal/conf"
	"github.com/khulnasoft/khulnasoft/internal/conf/conftypes"
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/database/basestore"
	"github.com/khulnasoft/khulnasoft/internal/extsvc"
	"github.com/khulnasoft/khulnasoft/internal/licensing"
	"github.com/khulnasoft/khulnasoft/internal/observation"
	"github.com/khulnasoft/khulnasoft/internal/timeutil"
)

var clock = timeutil.Now

func Init(
	ctx context.Context,
	observationCtx *observation.Context,
	db database.DB,
	_ codeintel.Services,
	_ conftypes.UnifiedWatchable,
	enterpriseServices *enterprise.Services,
) error {
	database.ValidateExternalServiceConfig = providers.ValidateExternalServiceConfig
	database.AuthzWith = func(other basestore.ShareableStore) database.AuthzStore {
		return database.NewAuthzStore(observationCtx.Logger, db, clock)
	}

	enterpriseServices.PermissionsGitHubWebhook = webhooks.NewGitHubWebhook(log.Scoped("PermissionsGitHubWebhook"))

	authz.DefaultSubRepoPermsChecker = srp.NewSubRepoPermsClient(db.SubRepoPerms())

	graphqlbackend.AlertFuncs = append(graphqlbackend.AlertFuncs, func(args graphqlbackend.AlertFuncArgs) []*graphqlbackend.Alert {
		if licensing.IsLicenseValid() {
			return nil
		}

		reason := licensing.GetLicenseInvalidReason()

		return []*graphqlbackend.Alert{{
			TypeValue:    graphqlbackend.AlertTypeError,
			MessageValue: fmt.Sprintf("The Khulnasoft license key is invalid. Reason: %s. To continue using Khulnasoft, a site admin must renew the Khulnasoft license (or downgrade to only using Khulnasoft Free features). Update the license key in the [**site configuration**](/site-admin/configuration). Please contact Khulnasoft support for more information.", reason),
		}}
	})

	// Warn about usage of authz providers that are not enabled by the license.
	graphqlbackend.AlertFuncs = append(graphqlbackend.AlertFuncs, func(args graphqlbackend.AlertFuncArgs) []*graphqlbackend.Alert {
		// Only site admins can act on this alert, so only show it to site admins.
		if !args.IsSiteAdmin {
			return nil
		}

		if licensing.IsFeatureEnabledLenient(licensing.FeatureACLs) {
			return nil
		}

		_, _, _, invalidConnections := providers.ProvidersFromConfig(ctx, conf.Get(), db)

		authzTypes := map[string]struct{}{}
		for _, conn := range invalidConnections {
			authzTypes[conn] = struct{}{}
		}

		authzNames := make([]string, 0, len(authzTypes))
		for t := range authzTypes {
			switch t {
			case extsvc.TypeGitHub:
				authzNames = append(authzNames, "GitHub")
			case extsvc.TypeGitLab:
				authzNames = append(authzNames, "GitLab")
			case extsvc.TypeBitbucketServer:
				authzNames = append(authzNames, "Bitbucket Server")
			default:
				authzNames = append(authzNames, t)
			}
		}

		if len(authzNames) == 0 {
			return nil
		}

		return []*graphqlbackend.Alert{{
			TypeValue:    graphqlbackend.AlertTypeError,
			MessageValue: fmt.Sprintf("A Khulnasoft license is required to enable repository permissions for the following code hosts: %s. [**Get a license.**](/site-admin/license)", strings.Join(authzNames, ", ")),
		}}
	})

	graphqlbackend.AlertFuncs = append(graphqlbackend.AlertFuncs, func(args graphqlbackend.AlertFuncArgs) []*graphqlbackend.Alert {
		// 🚨 SECURITY: Only the site admin should ever see this (all other users will see a hard-block
		// license expiration screen) about this. Leaking this wouldn't be a security vulnerability, but
		// just in case this method is changed to return more information, we lock it down.
		if !args.IsSiteAdmin {
			return nil
		}

		info, err := licensing.GetConfiguredProductLicenseInfo()
		if err != nil {
			observationCtx.Logger.Error("Error reading license key for Khulnasoft subscription.", log.Error(err))
			return []*graphqlbackend.Alert{{
				TypeValue:    graphqlbackend.AlertTypeError,
				MessageValue: "Error reading Khulnasoft license key. Check the logs for more information, or update the license key in the [**site configuration**](/site-admin/configuration).",
			}}
		}
		if info != nil && info.IsExpired() {
			return []*graphqlbackend.Alert{{
				TypeValue:    graphqlbackend.AlertTypeError,
				MessageValue: "Khulnasoft license expired! All non-admin users are locked out of Khulnasoft. Update the license key in the [**site configuration**](/site-admin/configuration) or downgrade to only using Khulnasoft Free features.",
			}}
		}
		if info != nil && info.IsExpiringSoon() {
			return []*graphqlbackend.Alert{{
				TypeValue:    graphqlbackend.AlertTypeWarning,
				MessageValue: fmt.Sprintf("Khulnasoft license will expire soon! Expires on: %s. Update the license key in the [**site configuration**](/site-admin/configuration) or downgrade to only using Khulnasoft Free features.", info.ExpiresAt.UTC().Truncate(time.Hour).Format(time.UnixDate)),
			}}
		}
		return nil
	})

	enterpriseServices.AuthzResolver = resolvers.NewResolver(observationCtx, db)
	return nil
}
