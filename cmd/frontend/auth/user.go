package auth

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/khulnasoft/khulnasoft/internal/actor"
	"github.com/khulnasoft/khulnasoft/internal/auth"
	"github.com/khulnasoft/khulnasoft/internal/telemetry"

	sglog "github.com/sourcegraph/log"

	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/auth/userpasswd"
	sgactor "github.com/khulnasoft/khulnasoft/internal/actor"
	"github.com/khulnasoft/khulnasoft/internal/authz"
	"github.com/khulnasoft/khulnasoft/internal/authz/permssync"
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/deviceid"
	"github.com/khulnasoft/khulnasoft/internal/errcode"
	"github.com/khulnasoft/khulnasoft/internal/extsvc"
	"github.com/khulnasoft/khulnasoft/internal/featureflag"
	"github.com/khulnasoft/khulnasoft/internal/usagestats"
	"github.com/khulnasoft/khulnasoft/lib/errors"
)

var MockGetAndSaveUser func(ctx context.Context, op GetAndSaveUserOp) (newUserCreated bool, userID int32, safeErrMsg string, err error)

type GetAndSaveUserOp struct {
	UserProps           database.NewUser
	ExternalAccount     extsvc.AccountSpec
	ExternalAccountData extsvc.AccountData
	CreateIfNotExist    bool
	LookUpByUsername    bool
	// SingleIdentityPerUser indicates that the provider should only allow to
	// connect a single external identity per user.
	SingleIdentityPerUser bool
	// UserCreateEventProperties is a map of key-value pairs to be added to the
	// `ExternalAuthSignupSucceeded` (V1) or 'externalAuthSignup' (V2) telemetry
	// event that is logged when a new user is created.
	//
	// It must fulfil telemetry.EventMetadata requirements, i.e. not contain
	// any sensitive data/PII.
	UserCreateEventProperties telemetry.EventMetadata
}

const telemetryV2UserSignUpFeatureName = "externalAuthSignup"

// GetAndSaveUser accepts authentication information associated with a given user, validates and applies
// the necessary updates to the DB, and returns the user ID after the updates have been applied.
//
// At a high level, it does the following:
//  1. Determine the identity of the user by applying the following rules in order:
//     a. If ctx contains an authenticated Actor, the Actor's identity is the user identity.
//     b. Look up the user by external account ID.
//     c. If the email specified in op.UserProps is verified, Look up the user by verified email.
//     If op.LookUpByUsername is true, look up by username instead of verified email.
//     (Note: most clients should look up by email, as username is typically insecure.)
//     d. If op.CreateIfNotExist is true, attempt to create a new user with the properties
//     specified in op.UserProps. This may fail if the desired username is already taken.
//     e. If a new user is successfully created, attempt to grant pending permissions.
//  2. Ensure that the user is associated with the external account information. This means
//     creating the external account if it does not already exist or updating it if it
//     already does.
//  3. Update any user props that have changed.
//  4. Return the user ID.
//
// 🚨 SECURITY: It is the caller's responsibility to ensure the veracity of the information that
// op contains (e.g., by receiving it from the appropriate authentication mechanism). It must
// also ensure that the user identity implied by op is consistent. Specifically, the values used
// in step 1 above must be consistent:
// * The authenticated Actor, if it exists
// * op.ExternalAccount
// * op.UserProps, especially op.UserProps.Email
//
// 🚨 SECURITY: The safeErrMsg is an error message that can be shown to unauthenticated users to
// describe the problem. The err may contain sensitive information and should only be written to the
// server error logs, not to the HTTP response to shown to unauthenticated users.
func GetAndSaveUser(
	ctx context.Context,
	logger sglog.Logger, // accept logger for scoping from various callsites, and readability of test logs
	db database.DB,
	tr *telemetry.EventRecorder, // accept explicit recorder for ease of testing
	op GetAndSaveUserOp,
) (newUserCreated bool, userID int32, safeErrMsg string, err error) {
	if MockGetAndSaveUser != nil {
		return MockGetAndSaveUser(ctx, op)
	}

	logger = logger.Scoped("GetAndSaveUser")
	recorder := telemetry.NewBestEffortEventRecorder(logger, tr)

	externalAccountsStore := db.UserExternalAccounts()
	users := db.Users()

	acct := &extsvc.Account{
		AccountSpec: op.ExternalAccount,
		AccountData: op.ExternalAccountData,
	}

	userID, newUserSaved, extAcctSaved, safeErrMsg, err := func() (int32, bool, bool, string, error) {
		// First, check if the user is already logged in. If so, return that user.
		if actor := sgactor.FromContext(ctx); actor.IsAuthenticated() {
			return actor.UID, false, false, "", nil
		}

		// Second, check if the user account already exists. If so, return that user.
		extsvcAcct, lookupByExternalErr := externalAccountsStore.Update(ctx, acct)
		if lookupByExternalErr == nil {
			return extsvcAcct.UserID, false, true, "", nil
		}
		if !errcode.IsNotFound(lookupByExternalErr) {
			return 0, false, false, "Unexpected error looking up the Khulnasoft user account associated with the external account. Ask a site admin for help.", lookupByExternalErr
		}

		if op.LookUpByUsername {
			user, getByUsernameErr := db.Users().GetByUsername(ctx, op.UserProps.Username)
			if getByUsernameErr == nil {
				return user.ID, false, false, "", nil
			}
			if !errcode.IsNotFound(getByUsernameErr) {
				return 0, false, false, "Unexpected error looking up the Khulnasoft user by username. Ask a site admin for help.", getByUsernameErr
			}
			if !op.CreateIfNotExist {
				return 0, false, false, fmt.Sprintf("User account with username %q does not exist. Ask a site admin to create your account.", op.UserProps.Username), getByUsernameErr
			}
		} else if op.UserProps.EmailIsVerified {
			user, getByVerifiedEmailErr := db.Users().GetByVerifiedEmail(ctx, op.UserProps.Email)
			if getByVerifiedEmailErr == nil {
				return user.ID, false, false, "", nil
			}
			if !errcode.IsNotFound(getByVerifiedEmailErr) {
				return 0, false, false, "Unexpected error looking up the Khulnasoft user by verified email. Ask a site admin for help.", getByVerifiedEmailErr
			}
			if !op.CreateIfNotExist {
				return 0, false, false, fmt.Sprintf("User account with verified email %q does not exist. Ask a site admin to create your account and then verify your email.", op.UserProps.Email), getByVerifiedEmailErr
			}
		}

		// Third, return an error here if creating new users is disabled.
		if !op.CreateIfNotExist {
			return 0, false, false, "It looks like this is your first time signing in with this external identity. Khulnasoft couldn't link it to an existing user, because no verified email was provided. Ask your site admin to configure the auth provider to include the user's verified email on sign-in.", lookupByExternalErr
		}

		// Fourth and finally, create a new user account and return it.
		//
		// If CreateIfNotExist is true, create the new user, regardless of whether the
		// email was verified or not.
		//
		// NOTE: It is important to propagate the correct context that carries the
		// information of the actor, especially whether the actor is a Khulnasoft
		// operator or not.
		act := &sgactor.Actor{
			KhulnasoftOperator: acct.AccountSpec.ServiceType == auth.KhulnasoftOperatorProviderType,
		}
		ctx = sgactor.WithActor(ctx, act)
		user, err := users.CreateWithExternalAccount(ctx, op.UserProps, acct)
		// We re-try creation with a different username with random suffix if the first one is taken.
		if database.IsUsernameExists(err) {
			op.UserProps.Username, err = userpasswd.AddRandomSuffix(op.UserProps.Username)
			if err != nil {
				return 0, false, false, "Unable to create a new user account due to a unexpected error. Ask a site admin for help.", errors.Wrapf(err, "username: %q, email: %q", op.UserProps.Username, op.UserProps.Email)
			}
			logger.Info("encountered username conflict, creating user with randomized suffix")
			user, err = users.CreateWithExternalAccount(ctx, op.UserProps, acct)
		}
		switch {
		case database.IsUsernameExists(err):
			return 0, false, false, fmt.Sprintf("Username %q already exists, but no verified email matched %q", op.UserProps.Username, op.UserProps.Email), err
		case errcode.PresentationMessage(err) != "":
			return 0, false, false, errcode.PresentationMessage(err), err
		case err != nil:
			return 0, false, false, "Unable to create a new user account due to a unexpected error. Ask a site admin for help.", errors.Wrapf(err, "username: %q, email: %q", op.UserProps.Username, op.UserProps.Email)
		}
		act.UID = user.ID

		// Schedule a permission sync, since this is new user
		permssync.SchedulePermsSync(ctx, logger, db, permssync.ScheduleSyncOpts{
			UserIDs:           []int32{user.ID},
			Reason:            database.ReasonUserAdded,
			TriggeredByUserID: user.ID,
		})

		if err = db.Authz().GrantPendingPermissions(ctx, &database.GrantPendingPermissionsArgs{
			UserID: user.ID,
			Perm:   authz.Read,
			Type:   authz.PermRepos,
		}); err != nil {
			logger.Error(
				"failed to grant user pending permissions",
				sglog.Int32("userID", user.ID),
				sglog.Error(err),
			)
			// OK to continue, since this is a best-effort to improve the UX with some initial permissions available.
		}

		// Legacy event - new event is recorded outside this closure to cover
		// more scenarios. We retain the legacy event because it is still
		// exported by the legacy Cloud exporter, remove in future release.
		const legacyEventName = "ExternalAuthSignupSucceeded"

		// SECURITY: This args map is treated as a public argument in the LogEvent call below, so it must not contain
		// any sensitive data.
		argMap := map[string]any{}
		for k, v := range op.UserCreateEventProperties {
			argMap[string(k)] = v
		}
		// NOTE: The conventional name should be "service_type", but keeping as-is for
		// backwards capability.
		argMap["serviceType"] = acct.AccountSpec.ServiceType

		args, err := json.Marshal(argMap)
		if err != nil {
			logger.Error(
				"failed to marshal JSON for event log argument",
				sglog.String("eventName", legacyEventName),
				sglog.Error(err),
			)
			// OK to continue, we still want the event log to be created
		}

		// NOTE: It is important to propagate the correct context that carries the
		// information of the actor, especially whether the actor is a Khulnasoft
		// operator or not.
		//
		// TODO: Use EventRecorder from internal/telemetryrecorder instead.
		//lint:ignore SA1019 existing usage of deprecated functionality.
		err = usagestats.LogEvent(
			ctx,
			db,
			usagestats.Event{
				EventName:      legacyEventName,
				UserID:         act.UID,
				Argument:       args,
				PublicArgument: args,
				Source:         "BACKEND",
			},
		)
		if err != nil {
			logger.Error(
				"failed to log event",
				sglog.String("eventName", legacyEventName),
				sglog.Error(err),
			)
		}

		return user.ID, true, true, "", nil
	}()

	// Annotate the telemetry ctx early with the user ID we found if an actor
	// isn't already present in context, irrespective of successful auth - this
	// is only used for telemetry. We need this now because userID may later be
	// overwritten.
	//
	// 🚨 SECURITY: Only use telemetryCtx for telemetry operations!
	telemetryCtx := ctx
	if userID != 0 && !sgactor.FromContext(telemetryCtx).IsAuthenticated() {
		telemetryCtx = sgactor.WithActor(telemetryCtx, actor.FromUser(userID))
	}
	// We handle all V2 telemetry related to GetAndSaveUser within this defer
	// closure, to ensure we cover all exit paths correctly after the other mega
	// closure above.
	//
	// We only store the event if a new user was created.
	if newUserSaved {
		defer func() {
			action := telemetry.ActionSucceeded
			if err != nil { // check final error
				action = telemetry.ActionFailed
			}

			// Most auth providers services have an exstvc.Variant, so try and
			// extract that from the account spec. For ease of use in we also
			// preserve the raw value in the private metadata.
			serviceVariant, _ := extsvc.VariantValueOf(acct.AccountSpec.ServiceType)
			privateMetadata := map[string]any{"serviceType": acct.AccountSpec.ServiceType}

			// Include safe err if there is one for maybe-useful diagnostics
			if len(safeErrMsg) > 0 {
				privateMetadata["safeErrMsg"] = safeErrMsg
			}

			// Record our V2 event.
			recorder.Record(telemetryCtx, telemetryV2UserSignUpFeatureName, action, &telemetry.EventParameters{
				Version: 2, // We've significantly refactored telemetryV2UserSignUpFeatureName occurrences
				Metadata: telemetry.MergeMetadata(
					telemetry.EventMetadata{
						"serviceVariant": telemetry.Number(serviceVariant),
						// Track the various outcomes of the massive signup closure above.
						"extAcctSaved": telemetry.Bool(extAcctSaved),
					},
					op.UserCreateEventProperties,
				),
				PrivateMetadata: privateMetadata,
			})
		}()
	}

	if err != nil {
		// Legacy event - retain because it is still exported by the legacy
		// Cloud exporter, remove in future release.
		const legacyEventName = "ExternalAuthSignupFailed"
		serviceTypeArg := json.RawMessage(fmt.Sprintf(`{"serviceType": %q}`, acct.AccountSpec.ServiceType))
		// TODO: Use EventRecorder from internal/telemetryrecorder instead.
		//lint:ignore SA1019 existing usage of deprecated functionality.
		if logErr := usagestats.LogBackendEvent(db, sgactor.FromContext(telemetryCtx).UID, deviceid.FromContext(telemetryCtx), legacyEventName, serviceTypeArg, serviceTypeArg, featureflag.GetEvaluatedFlagSet(ctx), nil); logErr != nil {
			logger.Error(
				"failed to log event",
				sglog.String("eventName", legacyEventName),
				sglog.Error(err),
			)
		}
		return newUserSaved, 0, safeErrMsg, err
	}

	if op.SingleIdentityPerUser && !extAcctSaved {
		// If we're here, the user has already signed up before this request,
		// but has no entry for this exact accountID in the user external accounts
		// table yet.
		// In single identity mode, we want to attach a new external account
		// only if no other identity of the same provider for the same user already
		// exists.
		other, err := externalAccountsStore.List(ctx, database.ExternalAccountsListOptions{
			UserID:      userID,
			ServiceType: acct.ServiceType,
			ServiceID:   acct.ServiceID,
			ClientID:    acct.ClientID,
		})
		if err != nil {
			return newUserSaved, 0, "Failed to list user identities. Ask a site admin for help.", err
		}

		// Confirm that the user already has a different identity from the same
		// provider. In cases where the user is already logged in (on an existing
		// browser session) extAcctSaved would be false, even though they logged in with
		// the same external identity. So as long as is the same identity, we are OK.
		found := false
		for _, eac := range other {
			if eac.ServiceType == acct.ServiceType &&
				eac.ServiceID == acct.ServiceID &&
				eac.AccountID == acct.AccountID &&
				eac.ClientID == acct.ClientID {
				found = true
				break
			}
		}
		if !found && len(other) >= 1 {
			return newUserSaved, 0, "Another identity for this user from this provider already exists. Remove the link to the other identity from your account.", errors.New("duplicate identity for single identity provider")
		}
	}

	// Update user properties, if they've changed
	if !newUserSaved {
		// Update user in our DB if their profile info changed on the issuer. (Except username and
		// email, which the user is somewhat likely to want to control separately on Khulnasoft.)
		user, err := db.Users().GetByID(ctx, userID)
		if err != nil {
			return newUserSaved, 0, "Unexpected error getting the Khulnasoft user account. Ask a site admin for help.", err
		}
		var userUpdate database.UserUpdate
		if user.DisplayName == "" && op.UserProps.DisplayName != "" {
			userUpdate.DisplayName = &op.UserProps.DisplayName
		}
		if user.AvatarURL == "" && op.UserProps.AvatarURL != "" {
			userUpdate.AvatarURL = &op.UserProps.AvatarURL
		}
		if userUpdate != (database.UserUpdate{}) {
			if err := db.Users().Update(ctx, user.ID, userUpdate); err != nil {
				return newUserSaved, 0, "Unexpected error updating the Khulnasoft user account with new user profile information from the external account. Ask a site admin for help.", err
			}
		}
	}

	// Create/update the external account and ensure it's associated with the user ID
	if !extAcctSaved {
		acct.UserID = userID
		_, err := externalAccountsStore.Upsert(ctx, acct)
		if err != nil {
			return newUserSaved, 0, "Unexpected error associating the external account with your Khulnasoft user. The most likely cause for this problem is that another Khulnasoft user is already linked with this external account. A site admin or the other user can unlink the account to fix this problem.", err
		}

		// Schedule a permission sync, since this is probably a new external account for the user
		permssync.SchedulePermsSync(ctx, logger, db, permssync.ScheduleSyncOpts{
			UserIDs:           []int32{userID},
			Reason:            database.ReasonExternalAccountAdded,
			TriggeredByUserID: userID,
		})

		if err = db.Authz().GrantPendingPermissions(ctx, &database.GrantPendingPermissionsArgs{
			UserID: userID,
			Perm:   authz.Read,
			Type:   authz.PermRepos,
		}); err != nil {
			logger.Error(
				"failed to grant user pending permissions",
				sglog.Int32("userID", userID),
				sglog.Error(err),
			)
			// OK to continue, since this is a best-effort to improve the UX with some initial permissions available.
		}
	}

	return newUserSaved, userID, "", nil
}
