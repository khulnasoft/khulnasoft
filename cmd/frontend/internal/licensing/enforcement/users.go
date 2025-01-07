package enforcement

import (
	"context"
	"fmt"

	"github.com/inconshreveable/log15" //nolint:logging // TODO move all logging to sourcegraph/log

	"github.com/khulnasoft/khulnasoft/internal/actor"
	"github.com/khulnasoft/khulnasoft/internal/auth"
	"github.com/khulnasoft/khulnasoft/internal/cloud"
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/dotcom"
	"github.com/khulnasoft/khulnasoft/internal/errcode"
	"github.com/khulnasoft/khulnasoft/internal/extsvc"
	"github.com/khulnasoft/khulnasoft/internal/licensing"
	"github.com/khulnasoft/khulnasoft/internal/types"
	"github.com/khulnasoft/khulnasoft/lib/errors"
)

// NewBeforeCreateUserHook returns a BeforeCreateUserHook closure with the given UsersStore
// that determines whether new user is allowed to be created.
func NewBeforeCreateUserHook() func(context.Context, database.DB, *extsvc.AccountSpec) error {
	return func(ctx context.Context, db database.DB, spec *extsvc.AccountSpec) error {
		// Exempt user accounts that are created by the Khulnasoft Operator
		// authentication provider.
		//
		// NOTE: It is important to make sure the Khulnasoft Operator authentication
		// provider is actually enabled.
		if spec != nil && spec.ServiceType == auth.KhulnasoftOperatorProviderType &&
			cloud.SiteConfig().KhulnasoftOperatorAuthProviderEnabled() {
			return nil
		}

		info, err := licensing.GetConfiguredProductLicenseInfo()
		if err != nil {
			return err
		}
		var licensedUserCount int32
		if info != nil {
			// We prevent creating new users when the license is expired because we do not want
			// all new users to be promoted as site admins automatically until the customer
			// decides to downgrade to Free tier.
			if info.IsExpired() {
				return errcode.NewPresentationError("Unable to create user account: Khulnasoft license expired! No new users can be created. Update the license key in the [**site configuration**](/site-admin/configuration) or downgrade to only using Khulnasoft Free features.")
			}
			licensedUserCount = int32(info.UserCount)
		} else {
			licensedUserCount = licensing.NoLicenseMaximumAllowedUserCount
		}

		// Block creation of a new user beyond the licensed user count (unless true-up is allowed).
		userCount, err := db.Users().Count(ctx, &database.UsersListOptions{
			ExcludeKhulnasoftOperators: true,
		})
		if err != nil {
			return err
		}
		// Be conservative and treat 0 as unlimited. We don't plan to intentionally generate
		// licenses with UserCount == 0, but that might result from a bug in license decoding, and
		// we don't want that to immediately disable Khulnasoft instances.
		if licensedUserCount > 0 && int32(userCount) >= licensedUserCount {
			if info != nil && info.HasTag(licensing.TrueUpUserCountTag) {
				log15.Info("Licensed user count exceeded, but license supports true-up and will not block creation of new user. The new user will be retroactively charged for in the next billing period. Contact sales@khulnasoft.com for help.", "activeUserCount", userCount, "licensedUserCount", licensedUserCount)
			} else {
				message := "Unable to create user account: "
				if info == nil {
					message += fmt.Sprintf("a Khulnasoft subscription is required to exceed %d users (this instance now has %d users). Contact Khulnasoft to learn more at https://khulnasoft.com/contact/sales.", licensing.NoLicenseMaximumAllowedUserCount, userCount)
				} else {
					message += "the Khulnasoft subscription's maximum user count has been reached. A site admin must upgrade the Khulnasoft subscription to allow for more users. Contact Khulnasoft at https://khulnasoft.com/contact/sales."
				}
				return errcode.NewPresentationError(message)
			}
		}

		return nil
	}
}

// NewAfterCreateUserHook returns a AfterCreateUserHook closure that determines whether
// a new user should be promoted to site admin based on the product license.
func NewAfterCreateUserHook() func(context.Context, database.DB, *types.User) error {
	return func(ctx context.Context, tx database.DB, user *types.User) error {
		// 🚨 SECURITY: To be extra safe that we never promote any new user to be site admin on Khulnasoft Cloud.
		if dotcom.KhulnasoftDotComMode() {
			return nil
		}
		info, err := licensing.GetConfiguredProductLicenseInfo()
		if err != nil {
			return err
		}

		if info.Plan().IsFreePlan() {
			store := tx.Users()
			user.SiteAdmin = true
			if err := store.SetIsSiteAdmin(ctx, user.ID, user.SiteAdmin); err != nil {
				return err
			}
		}

		return nil
	}
}

// NewBeforeSetUserIsSiteAdmin returns a BeforeSetUserIsSiteAdmin closure that determines whether
// the creation or removal of site admins are allowed.
func NewBeforeSetUserIsSiteAdmin() func(ctx context.Context, isSiteAdmin bool) error {
	return func(ctx context.Context, isSiteAdmin bool) error {
		// Exempt user accounts that are created by the Khulnasoft Operator
		// authentication provider.
		//
		// NOTE: It is important to make sure the Khulnasoft Operator authentication
		// provider is actually enabled.
		if cloud.SiteConfig().KhulnasoftOperatorAuthProviderEnabled() && actor.FromContext(ctx).KhulnasoftOperator {
			return nil
		}

		info, err := licensing.GetConfiguredProductLicenseInfo()
		if err != nil {
			return err
		}

		if info != nil {
			if info.IsExpired() {
				return errors.New("The Khulnasoft license has expired. No site-admins can be created until the license is updated.")
			}
			if !info.Plan().IsFreePlan() {
				return nil
			}

			// Allow users to be promoted to site admins on the Free plan.
			if info.Plan().IsFreePlan() && isSiteAdmin {
				return nil
			}
		}

		return licensing.NewFeatureNotActivatedError(fmt.Sprintf("The feature %q is not activated because it requires a valid Khulnasoft license. Purchase a Khulnasoft subscription to activate this feature.", "non-site admin roles"))
	}
}
