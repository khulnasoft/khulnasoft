package cody

import (
	"context"

	"github.com/sourcegraph/log"

	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/backend"
	"github.com/khulnasoft/khulnasoft/internal/actor"
	"github.com/khulnasoft/khulnasoft/internal/auth"
	"github.com/khulnasoft/khulnasoft/internal/conf"
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/dotcom"
	"github.com/khulnasoft/khulnasoft/internal/featureflag"
	"github.com/khulnasoft/khulnasoft/internal/licensing"
	"github.com/khulnasoft/khulnasoft/internal/rbac"
	"github.com/khulnasoft/khulnasoft/lib/errors"
)

// IsCodyEnabled determines if cody is enabled for the actor in the given context.
// If it is an unauthenticated request, cody is disabled.
// If authenticated it checks if cody is enabled for the deployment type
func IsCodyEnabled(ctx context.Context, db database.DB) (enabled bool, reason string) {
	a := actor.FromContext(ctx)
	if !a.IsAuthenticated() {
		return false, "not authenticated"
	}
	return isCodyEnabled(ctx, db)
}

// isCodyEnabled determines if cody is enabled for the actor in the given context.
// If the license does not have the Cody feature, cody is disabled.
// If Completions aren't configured, cody is disabled.
// If Completions are not enabled, cody is disabled
// If CodyRestrictUsersFeatureFlag is set, the cody featureflag
// will determine access.
// If CodyPermissions is enabled, RBAC will determine access.
// Otherwise, all authenticated users are granted access.
func isCodyEnabled(ctx context.Context, db database.DB) (enabled bool, reason string) {
	if err := licensing.Check(licensing.FeatureCody); err != nil {
		return false, "instance license does not allow cody"
	}

	if !conf.CodyEnabled() {
		return false, "cody is disabled"
	}

	// Note: we respect the deprecated feature flag, which was in use before
	// we had proper RBAC implemented.
	if conf.CodyRestrictUsersFeatureFlag() {
		enabled = featureflag.FromContext(ctx).GetBoolOr("cody", false)
		if enabled {
			return true, ""
		}
		return false, "cody is restricted to feature flag but feature flag is not enabled"
	}

	if conf.CodyPermissionsEnabled() {
		// Check if user has cody permission via RBAC
		err := rbac.CheckCurrentUserHasPermission(ctx, db, rbac.CodyAccessPermission)
		if err != nil {
			return false, "user does not have permission " + rbac.CodyAccessPermission
		}
		return true, ""
	}
	return true, ""
}

var ErrRequiresVerifiedEmailAddress = errors.New("cody requires a verified email address")

func CheckVerifiedEmailRequirement(ctx context.Context, db database.DB, logger log.Logger) error {
	// Only check on dotcom
	if !dotcom.SourcegraphDotComMode() {
		return nil
	}

	// Do not require if user is site-admin
	if err := auth.CheckCurrentUserIsSiteAdmin(ctx, db); err == nil {
		return nil
	}

	verified, err := backend.NewUserEmailsService(db, logger).CurrentActorHasVerifiedEmail(ctx)
	if err != nil {
		return err
	}
	if verified {
		return nil
	}

	return ErrRequiresVerifiedEmailAddress
}
