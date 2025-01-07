package authz

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/url"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/grafana/regexp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sourcegraph/log/logtest"

	"github.com/khulnasoft/khulnasoft/internal/actor"
	"github.com/khulnasoft/khulnasoft/internal/api"
	"github.com/khulnasoft/khulnasoft/internal/authz"
	authzBitbucketServer "github.com/khulnasoft/khulnasoft/internal/authz/providers/bitbucketserver"
	authzGitHub "github.com/khulnasoft/khulnasoft/internal/authz/providers/github"
	authzGitLab "github.com/khulnasoft/khulnasoft/internal/authz/providers/gitlab"
	"github.com/khulnasoft/khulnasoft/internal/conf"
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/database/dbtest"
	"github.com/khulnasoft/khulnasoft/internal/extsvc"
	"github.com/khulnasoft/khulnasoft/internal/extsvc/auth"
	"github.com/khulnasoft/khulnasoft/internal/extsvc/bitbucketserver"
	"github.com/khulnasoft/khulnasoft/internal/extsvc/github"
	extsvcGitHub "github.com/khulnasoft/khulnasoft/internal/extsvc/github"
	"github.com/khulnasoft/khulnasoft/internal/httptestutil"
	"github.com/khulnasoft/khulnasoft/internal/ratelimit"
	"github.com/khulnasoft/khulnasoft/internal/rcache"
	"github.com/khulnasoft/khulnasoft/internal/repos"
	"github.com/khulnasoft/khulnasoft/internal/timeutil"
	"github.com/khulnasoft/khulnasoft/internal/types"
	"github.com/khulnasoft/khulnasoft/schema"
)

var updateRegex = flag.String("update-integration", "", "Update testdata of tests matching the given regex")

func update(name string) bool {
	if updateRegex == nil || *updateRegex == "" {
		return false
	}
	return regexp.MustCompile(*updateRegex).MatchString(name)
}

func assertGitHubUserPermissions(t *testing.T, ctx context.Context, userID int32, ghURL string, syncer *permsSyncerImpl, permsStore database.PermsStore, wantIDs []int32) {
	t.Helper()

	_, providerStates, err := syncer.syncUserPerms(ctx, userID, false, authz.FetchPermsOptions{})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, database.CodeHostStatusesSet{{
		ProviderID:   ghURL,
		ProviderType: "github",
		Status:       database.CodeHostStatusSuccess,
		Message:      "FetchUserPerms",
	}}, providerStates)

	p, err := permsStore.LoadUserPermissions(ctx, userID)
	if err != nil {
		t.Fatal(err)
	}
	gotIDs := make([]int32, len(p))
	for i, perm := range p {
		gotIDs[i] = perm.RepoID
	}

	if diff := cmp.Diff(wantIDs, gotIDs); diff != "" {
		t.Fatalf("IDs mismatch (-want +got):\n%s", diff)
	}
}

func assertGitHubRepoPermissions(t *testing.T, ctx context.Context, repoID api.RepoID, userID int32, ghURL string, syncer *permsSyncerImpl, permsStore database.PermsStore, wantIDs []int32) {
	t.Helper()

	_, providerStates, err := syncer.syncRepoPerms(ctx, repoID, false, authz.FetchPermsOptions{})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, database.CodeHostStatusesSet{{
		ProviderID:   ghURL,
		ProviderType: "github",
		Status:       database.CodeHostStatusSuccess,
		Message:      "FetchRepoPerms",
	}}, providerStates)

	p, err := permsStore.LoadUserPermissions(ctx, userID)
	if err != nil {
		t.Fatal(err)
	}
	gotIDs := make([]int32, len(p))
	for i, perm := range p {
		gotIDs[i] = perm.RepoID
	}

	if diff := cmp.Diff(wantIDs, gotIDs); diff != "" {
		t.Fatalf("IDs mismatch (-want +got):\n%s", diff)
	}
}

// NOTE: To update VCR for these tests, please use the token of "sourcegraph-vcr"
// for GITHUB_TOKEN, which can be found in 1Password.
//
// We also recommend setting up a new token for "sourcegraph-vcr" using the auth scope
// guidelines https://khulnasoft.com/docs/admin/code_hosts/github#github-api-access
// to ensure everything works, in case of new scopes being required.
func TestIntegration_GitHubPermissions(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	conf.Mock(&conf.Unified{
		SiteConfiguration: schema.SiteConfiguration{
			AuthProviders: []schema.AuthProviders{
				{
					Github: &schema.GitHubAuthProvider{
						Url: "https://github.com",
					},
				},
			},
		},
	})
	t.Cleanup(func() {
		conf.Mock(nil)
	})

	ratelimit.SetupForTest(t)

	logger := logtest.Scoped(t)
	token := os.Getenv("GITHUB_TOKEN")

	spec := extsvc.AccountSpec{
		ServiceType: extsvc.TypeGitHub,
		ServiceID:   "https://github.com/",
		AccountID:   "66464926",
	}
	svc := types.ExternalService{
		Kind:      extsvc.KindGitHub,
		CreatedAt: timeutil.Now(),
		Config:    extsvc.NewUnencryptedConfig(`{"url": "https://github.com", "authorization": {}, "token": "abc", "repos": ["owner/name"]}`),
	}
	uri, err := url.Parse("https://github.com/")
	if err != nil {
		t.Fatal(err)
	}

	newUser := database.NewUser{
		Email:           "sourcegraph-vcr-bob@khulnasoft.com",
		Username:        "sourcegraph-vcr-bob",
		EmailIsVerified: true,
	}
	testDB := database.NewDB(logger, dbtest.NewDB(t))
	ctx := actor.WithInternalActor(context.Background())

	reposStore := repos.NewStore(logtest.Scoped(t), testDB)

	err = reposStore.ExternalServiceStore().Upsert(ctx, &svc)
	if err != nil {
		t.Fatal(err)
	}

	repo := types.Repo{
		Name:    "github.com/sourcegraph-vcr-repos/private-org-repo-1",
		Private: true,
		URI:     "github.com/sourcegraph-vcr-repos/private-org-repo-1",
		ExternalRepo: api.ExternalRepoSpec{
			ID:          "MDEwOlJlcG9zaXRvcnkzOTk4OTQyODY=",
			ServiceType: extsvc.TypeGitHub,
			ServiceID:   "https://github.com/",
		},
		Sources: map[string]*types.SourceInfo{
			svc.URN(): {
				ID: svc.URN(),
			},
		},
	}

	err = reposStore.RepoStore().Create(ctx, &repo)
	if err != nil {
		t.Fatal(err)
	}

	authData := json.RawMessage(fmt.Sprintf(`{"access_token": "%s"}`, token))
	user, err := testDB.Users().CreateWithExternalAccount(ctx, newUser,
		&extsvc.Account{
			AccountSpec: spec,
			AccountData: extsvc.AccountData{
				AuthData: extsvc.NewUnencryptedData(authData),
			},
		})
	if err != nil {
		t.Fatal(err)
	}

	permsStore := database.Perms(logger, testDB, timeutil.Now)
	syncer := newPermsSyncer(logger, testDB, reposStore, permsStore, timeutil.Now)

	// This integration tests performs a repository-centric permissions syncing against
	// https://github.com, then check if permissions are correctly granted for the test
	// user "sourcegraph-vcr-bob", who is a outside collaborator of the repository
	// "sourcegraph-vcr-repos/private-org-repo-1".
	t.Run("repo-centric", func(t *testing.T) {
		t.Run("no-groups", func(t *testing.T) {
			cli := newTestRecorderClient(t, svc.URN(), uri, token)

			provider := authzGitHub.NewProvider(svc.URN(), authzGitHub.ProviderOptions{
				GitHubClient:   cli,
				GitHubURL:      uri,
				BaseAuther:     &auth.OAuthBearerToken{Token: token},
				GroupsCacheTTL: -1,
				DB:             testDB,
			})

			prevProviderFactory := syncer.providerFactory
			syncer.providerFactory = func(context.Context) []authz.Provider {
				return []authz.Provider{provider}
			}
			defer func() {
				syncer.providerFactory = prevProviderFactory
			}()

			assertGitHubRepoPermissions(t, ctx, repo.ID, user.ID, uri.String(), syncer, permsStore, []int32{1})
		})

		t.Run("groups-enabled", func(t *testing.T) {
			cli := newTestRecorderClient(t, svc.URN(), uri, token)

			provider := authzGitHub.NewProvider(svc.URN(), authzGitHub.ProviderOptions{
				GitHubClient:   cli,
				GitHubURL:      uri,
				BaseAuther:     &auth.OAuthBearerToken{Token: token},
				GroupsCacheTTL: 72,
				DB:             testDB,
			})

			prevProviderFactory := syncer.providerFactory
			syncer.providerFactory = func(context.Context) []authz.Provider {
				return []authz.Provider{provider}
			}
			defer func() {
				syncer.providerFactory = prevProviderFactory
			}()

			assertGitHubRepoPermissions(t, ctx, repo.ID, user.ID, uri.String(), syncer, permsStore, []int32{1})
		})
	})

	// This integration tests performs a user-centric permissions syncing against
	// https://github.com, then check if permissions are correctly granted for the test
	// user "sourcegraph-vcr-bob", who is a collaborator of "sourcegraph-vcr-repos/private-org-repo-1".
	t.Run("user-centric", func(t *testing.T) {
		t.Run("no-groups", func(t *testing.T) {
			cli := newTestRecorderClient(t, svc.URN(), uri, token)

			provider := authzGitHub.NewProvider(svc.URN(), authzGitHub.ProviderOptions{
				GitHubClient:   cli,
				GitHubURL:      uri,
				BaseAuther:     &auth.OAuthBearerToken{Token: token},
				GroupsCacheTTL: -1,
				DB:             testDB,
			})

			prevProviderFactory := syncer.providerFactory
			syncer.providerFactory = func(context.Context) []authz.Provider {
				return []authz.Provider{provider}
			}
			defer func() {
				syncer.providerFactory = prevProviderFactory
			}()

			assertGitHubUserPermissions(t, ctx, user.ID, uri.String(), syncer, permsStore, []int32{1})
		})

		t.Run("groups-enabled", func(t *testing.T) {
			cli := newTestRecorderClient(t, svc.URN(), uri, token)

			provider := authzGitHub.NewProvider(svc.URN(), authzGitHub.ProviderOptions{
				GitHubClient:   cli,
				GitHubURL:      uri,
				BaseAuther:     &auth.OAuthBearerToken{Token: token},
				GroupsCacheTTL: 72,
				DB:             testDB,
			})

			prevProviderFactory := syncer.providerFactory
			syncer.providerFactory = func(context.Context) []authz.Provider {
				return []authz.Provider{provider}
			}
			defer func() {
				syncer.providerFactory = prevProviderFactory
			}()

			assertGitHubUserPermissions(t, ctx, user.ID, uri.String(), syncer, permsStore, []int32{1})
		})
	})
}

func newTestRecorderClient(t *testing.T, urn string, apiURL *url.URL, token string) *extsvcGitHub.V3Client {
	name := t.Name()

	cf, save := httptestutil.NewGitHubRecorderFactory(t, update(name), name)
	t.Cleanup(save)
	doer, err := cf.Doer()
	if err != nil {
		t.Fatal(err)
	}
	cli := extsvcGitHub.NewV3Client(logtest.Scoped(t), urn, apiURL, &auth.OAuthBearerToken{Token: token}, doer)
	return cli
}

// TestIntegration_GitHubInternalRepositories asserts that internal repositories of
// organizations that a user belongs to are fetched alongside user permission
// syncs.
//
// The test setup requires a user that belongs to an organization with internal repos.
// It is kept separate from the other integration test since it connects to
// ghe.sgdev.org instead of github.com
func TestIntegration_GitHubInternalRepositories(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	conf.Mock(&conf.Unified{
		SiteConfiguration: schema.SiteConfiguration{
			AuthProviders: []schema.AuthProviders{
				{
					Github: &schema.GitHubAuthProvider{
						Url: "https://ghe.sgdev.org",
					},
				},
			},
		},
	})
	t.Cleanup(func() {
		conf.Mock(nil)
	})

	ratelimit.SetupForTest(t)
	rcache.SetupForTest(t)

	logger := logtest.Scoped(t)
	token := os.Getenv("GITHUB_TOKEN")

	spec := extsvc.AccountSpec{
		ServiceType: extsvc.TypeGitHub,
		ServiceID:   "https://ghe.sgdev.org/",
		AccountID:   "3",
	}
	svc := types.ExternalService{
		Kind:      extsvc.KindGitHub,
		CreatedAt: timeutil.Now(),
		Config:    extsvc.NewUnencryptedConfig(`{"url": "https://ghe.sgdev.org/", "authorization": {}, "token": "abc", "repos": ["owner/name"]}`),
	}
	uri, err := url.Parse("https://ghe.sgdev.org/")
	if err != nil {
		t.Fatal(err)
	}
	apiURI, _ := github.APIRoot(uri)

	cli := newTestRecorderClient(t, uri.String(), apiURI, token)

	testDB := database.NewDB(logger, dbtest.NewDB(t))
	ctx := actor.WithInternalActor(context.Background())

	reposStore := repos.NewStore(logtest.Scoped(t), testDB)

	err = reposStore.ExternalServiceStore().Upsert(ctx, &svc)
	if err != nil {
		t.Fatal(err)
	}

	provider := authzGitHub.NewProvider(svc.URN(), authzGitHub.ProviderOptions{
		GitHubClient:   cli,
		GitHubURL:      uri,
		BaseAuther:     &auth.OAuthBearerToken{Token: token},
		GroupsCacheTTL: -1, // disable groups caching
		DB:             testDB,
	})

	repo := types.Repo{
		Name:    "ghe.sgdev.org/sourcegraph/sourcegraph_internal_repo",
		Private: true,
		URI:     "ghe.sgdev.org/sourcegraph/sourcegraph_internal_repo",
		ExternalRepo: api.ExternalRepoSpec{
			ID:          "MDEwOlJlcG9zaXRvcnkxMDU3MDMy",
			ServiceType: extsvc.TypeGitHub,
			ServiceID:   "https://ghe.sgdev.org/",
		},
		Sources: map[string]*types.SourceInfo{
			svc.URN(): {
				ID: svc.URN(),
			},
		},
	}
	err = reposStore.RepoStore().Create(ctx, &repo)
	if err != nil {
		t.Fatal(err)
	}

	newUser := database.NewUser{
		Email:           "sourcegraph-vcr@khulnasoft.com",
		Username:        "sourcegraph-vcr",
		EmailIsVerified: true,
	}

	authData := json.RawMessage(fmt.Sprintf(`{"access_token": "%s"}`, token))
	user, err := testDB.Users().CreateWithExternalAccount(ctx, newUser, &extsvc.Account{
		AccountSpec: spec,
		AccountData: extsvc.AccountData{
			AuthData: extsvc.NewUnencryptedData(authData),
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	permsStore := database.Perms(logger, testDB, timeutil.Now)
	syncer := newPermsSyncer(logger, testDB, reposStore, permsStore, timeutil.Now)

	syncer.providerFactory = func(context.Context) []authz.Provider {
		return []authz.Provider{provider}
	}

	assertGitHubUserPermissions(t, ctx, user.ID, uri.String(), syncer, permsStore, []int32{1})
}

func TestIntegration_GitLabPermissions(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	logger := logtest.Scoped(t)
	token := os.Getenv("GITLAB_TOKEN")

	spec := extsvc.AccountSpec{
		ServiceType: extsvc.TypeGitLab,
		ServiceID:   "https://gitlab.sgdev.org/",
		AccountID:   "107564",
	}
	svc := types.ExternalService{
		Kind:   extsvc.KindGitLab,
		Config: extsvc.NewUnencryptedConfig(`{"url": "https://gitlab.sgdev.org", "authorization": {"identityProvider": {"type": "oauth"}}, "token": "abc", "projectQuery": [ "projects?membership=true&archived=no" ]}`),
	}
	uri, err := url.Parse("https://gitlab.sgdev.org")
	if err != nil {
		t.Fatal(err)
	}

	newUser := database.NewUser{
		Email:           "sourcegraph-vcr@khulnasoft.com",
		Username:        "sourcegraph-vcr",
		EmailIsVerified: true,
	}

	// These tests require two repos to be set up:
	// Both schwifty2 and getschwifty are internal projects.
	// The user is an explicit collaborator on getschwifty, so
	// should have access to getschwifty regardless of the feature flag.
	// The user does not have explicit access to schwifty2, however
	// schwifty2 is configured so that anyone on the instance has read
	// access, so when the feature flag is enabled, the user should
	// see this repo as well.
	testRepos := []types.Repo{
		{
			Name:    "gitlab.sgdev.org/petrissupercoolgroup/schwifty2",
			Private: true,
			URI:     "gitlab.sgdev.org/petrissupercoolgroup/schwifty2",
			ExternalRepo: api.ExternalRepoSpec{
				ID:          "371335",
				ServiceType: extsvc.TypeGitLab,
				ServiceID:   "https://gitlab.sgdev.org/",
			},
			Sources: map[string]*types.SourceInfo{
				svc.URN(): {
					ID: svc.URN(),
				},
			},
		},
		{
			Name:    "gitlab.sgdev.org/petri.last/getschwifty",
			Private: true,
			URI:     "gitlab.sgdev.org/petri.last/getschwifty",
			ExternalRepo: api.ExternalRepoSpec{
				ID:          "371334",
				ServiceType: extsvc.TypeGitLab,
				ServiceID:   "https://gitlab.sgdev.org/",
			},
			Sources: map[string]*types.SourceInfo{
				svc.URN(): {
					ID: svc.URN(),
				},
			},
		},
	}

	authData := json.RawMessage(fmt.Sprintf(`{"access_token": "%s"}`, token))

	// This integration tests performs a user-centric permissions syncing against
	// https://gitlab.sgdev.org, then check if permissions are correctly granted for the test
	// user "sourcegraph-vcr".
	t.Run("test gitLabProjectVisibilityExperimental feature flag", func(t *testing.T) {
		name := t.Name()

		cf, save := httptestutil.NewRecorderFactory(t, update(name), name)
		defer save()
		doer, err := cf.Doer()
		require.NoError(t, err)

		testDB := database.NewDB(logger, dbtest.NewDB(t))

		ctx := actor.WithInternalActor(context.Background())

		reposStore := repos.NewStore(logtest.Scoped(t), testDB)

		err = reposStore.ExternalServiceStore().Upsert(ctx, &svc)
		require.NoError(t, err)

		provider := authzGitLab.NewOAuthProvider(authzGitLab.OAuthProviderOp{
			BaseURL:                     uri,
			DB:                          testDB,
			CLI:                         doer,
			SyncInternalRepoPermissions: true,
		})

		for _, repo := range testRepos {
			err = reposStore.RepoStore().Create(ctx, &repo)
			require.NoError(t, err)
		}

		user, err := testDB.Users().CreateWithExternalAccount(ctx, newUser, &extsvc.Account{
			AccountSpec: spec,
			AccountData: extsvc.AccountData{
				AuthData: extsvc.NewUnencryptedData(authData),
			},
		})
		require.NoError(t, err)

		permsStore := database.Perms(logger, testDB, timeutil.Now)
		syncer := newPermsSyncer(logger, testDB, reposStore, permsStore, timeutil.Now)

		syncer.providerFactory = func(context.Context) []authz.Provider {
			return []authz.Provider{provider}
		}

		assertUserPermissions := func(t *testing.T, wantIDs []int32) {
			t.Helper()
			_, providerStates, err := syncer.syncUserPerms(ctx, user.ID, false, authz.FetchPermsOptions{})
			require.NoError(t, err)

			assert.Equal(t, database.CodeHostStatusesSet{{
				ProviderID:   "https://gitlab.sgdev.org/",
				ProviderType: "gitlab",
				Status:       database.CodeHostStatusSuccess,
				Message:      "FetchUserPerms",
			}}, providerStates)

			p, err := permsStore.LoadUserPermissions(ctx, user.ID)
			require.NoError(t, err)

			gotIDs := make([]int32, len(p))
			for i, perm := range p {
				gotIDs[i] = perm.RepoID
			}

			if diff := cmp.Diff(wantIDs, gotIDs); diff != "" {
				t.Fatalf("IDs mismatch (-want +got):\n%s", diff)
			}
		}

		// With the feature flag disabled (default state) the user should only have access to one repo
		assertUserPermissions(t, []int32{2})

		// With the feature flag enabled the user should have access to both repositories
		_, err = testDB.FeatureFlags().CreateBool(ctx, "gitLabProjectVisibilityExperimental", true)
		require.NoError(t, err, "feature flag creation failed")

		assertUserPermissions(t, []int32{1, 2})
	})
}

func TestIntegration_BitbucketServerPermissions(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	logger := logtest.Scoped(t)
	// BITBUCKETSERVER_TOKEN should be a valid OAuth2 token for the bitbucket server instance
	token := os.Getenv("BITBUCKETSERVER_TOKEN")

	spec := extsvc.AccountSpec{
		ServiceType: extsvc.TypeBitbucketServer,
		ServiceID:   "https://bitbucket.sgdev.org/",
		AccountID:   "603",
	}
	svc := types.ExternalService{
		Kind:   extsvc.KindBitbucketServer,
		Config: extsvc.NewUnencryptedConfig(`{"url": "https://bitbucket.sgdev.org", "authorization": {"oauth2": true}, "token": "abc", "repos": [ "PRIVATE/test-repo" ], "username": "whatever"}`),
	}

	newUser := database.NewUser{
		Email:           "sourcegraph-vcr@khulnasoft.com",
		Username:        "sourcegraph-vcr",
		EmailIsVerified: true,
	}

	// This test requires the "PRIVATE/test-repo" repo to be set up
	testRepos := []types.Repo{
		{
			Name:    "bitbucket.sgdev.org/PRIVATE/test-repo",
			Private: true,
			URI:     "bitbucket.sgdev.org/PRIVATE/test-repo",
			ExternalRepo: api.ExternalRepoSpec{
				ID:          "10093",
				ServiceType: extsvc.TypeBitbucketServer,
				ServiceID:   "https://bitbucket.sgdev.org/",
			},
			Sources: map[string]*types.SourceInfo{
				svc.URN(): {
					ID: svc.URN(),
				},
			},
		},
	}

	authData := json.RawMessage(fmt.Sprintf(`{"access_token": "%s"}`, token))
	accountData := json.RawMessage(`{"name":"pjlast","emailAddress":"petri.last@khulnasoft.com","id":603,"displayName":"Petri-Johan Last","active":true,"slug":"pjlast","type":"NORMAL"}`)

	t.Run("test bitbucket server oauth permissions", func(t *testing.T) {
		name := t.Name()

		cf, save := httptestutil.NewRecorderFactory(t, update(name), name)
		defer save()
		doer, err := cf.Doer()
		require.NoError(t, err)

		testDB := database.NewDB(logger, dbtest.NewDB(t))

		ctx := actor.WithInternalActor(context.Background())

		reposStore := repos.NewStore(logtest.Scoped(t), testDB)

		err = reposStore.ExternalServiceStore().Upsert(ctx, &svc)
		require.NoError(t, err)

		cfg, err := extsvc.ParseEncryptableConfig(ctx, svc.Kind, svc.Config)
		require.NoError(t, err)

		conn := &types.BitbucketServerConnection{
			URN:                       svc.URN(),
			BitbucketServerConnection: cfg.(*schema.BitbucketServerConnection),
		}

		cli, err := bitbucketserver.NewClient("https://bitbucket.sgdev.org", conn.BitbucketServerConnection, doer)
		require.NoError(t, err)

		provider := authzBitbucketServer.NewOAuthProvider(testDB, conn, authzBitbucketServer.ProviderOptions{BitbucketServerClient: cli}, false)

		for _, repo := range testRepos {
			err = reposStore.RepoStore().Create(ctx, &repo)
			require.NoError(t, err)
		}

		user, err := testDB.Users().CreateWithExternalAccount(ctx, newUser, &extsvc.Account{
			AccountSpec: spec,
			AccountData: extsvc.AccountData{
				AuthData: extsvc.NewUnencryptedData(authData),
				Data:     extsvc.NewUnencryptedData(accountData),
			},
		})
		require.NoError(t, err)

		permsStore := database.Perms(logger, testDB, timeutil.Now)
		syncer := newPermsSyncer(logger, testDB, reposStore, permsStore, timeutil.Now)

		syncer.providerFactory = func(context.Context) []authz.Provider {
			return []authz.Provider{provider}
		}

		assertUserPermissions := func(t *testing.T, wantIDs []int32) {
			t.Helper()
			_, providerStates, err := syncer.syncUserPerms(ctx, user.ID, false, authz.FetchPermsOptions{})
			require.NoError(t, err)

			assert.Equal(t, database.CodeHostStatusesSet{{
				ProviderID:   "https://bitbucket.sgdev.org/",
				ProviderType: "bitbucketServer",
				Status:       database.CodeHostStatusSuccess,
				Message:      "FetchUserPerms",
			}}, providerStates)

			p, err := permsStore.LoadUserPermissions(ctx, user.ID)
			require.NoError(t, err)

			gotIDs := make([]int32, len(p))
			for i, perm := range p {
				gotIDs[i] = perm.RepoID
			}

			if diff := cmp.Diff(wantIDs, gotIDs); diff != "" {
				t.Fatalf("IDs mismatch (-want +got):\n%s", diff)
			}
		}

		assertUserPermissions(t, []int32{1})
	})
}
