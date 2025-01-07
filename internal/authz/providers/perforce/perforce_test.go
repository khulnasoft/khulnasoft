package perforce

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	jsoniter "github.com/json-iterator/go"
	"github.com/sourcegraph/log/logtest"

	"github.com/khulnasoft/khulnasoft/internal/api"
	"github.com/khulnasoft/khulnasoft/internal/authz"
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/database/dbmocks"
	et "github.com/khulnasoft/khulnasoft/internal/encryption/testing"
	"github.com/khulnasoft/khulnasoft/internal/extsvc"
	"github.com/khulnasoft/khulnasoft/internal/extsvc/perforce"
	"github.com/khulnasoft/khulnasoft/internal/gitserver"
	"github.com/khulnasoft/khulnasoft/internal/gitserver/protocol"
	p4types "github.com/khulnasoft/khulnasoft/internal/perforce"
	"github.com/khulnasoft/khulnasoft/internal/types"
	"github.com/khulnasoft/khulnasoft/lib/errors"
)

func TestProvider_FetchAccount(t *testing.T) {
	logger := logtest.Scoped(t)
	ctx := context.Background()
	user := &types.User{
		ID:       1,
		Username: "alice",
	}

	db := dbmocks.NewMockDB()
	mockUserEmails := dbmocks.NewMockUserEmailsStore()
	db.UserEmailsFunc.SetDefaultReturn(mockUserEmails)

	gitserverClient := gitserver.NewStrictMockClient()
	gitserverClient.PerforceUsersFunc.SetDefaultReturn([]*p4types.User{
		{Username: "alice", Email: "Alice@Example.com"},
		{Username: "cindy", Email: "cindy@example.com"},
	}, nil)

	t.Run("no matching account", func(t *testing.T) {
		mockUserEmails.ListByUserFunc.SetDefaultReturn([]*database.UserEmail{{Email: "bob@example.com"}}, nil)
		p := NewProvider(logger, db, gitserverClient, "", "ssl:111.222.333.444:1666", "admin", "password", []extsvc.RepoID{}, false)
		got, err := p.FetchAccount(ctx, user)
		if err != nil {
			t.Fatal(err)
		}

		if got != nil {
			t.Fatalf("Want nil but got %v", got)
		}
	})

	t.Run("found matching account", func(t *testing.T) {
		p := NewProvider(logger, db, gitserverClient, "", "ssl:111.222.333.444:1666", "admin", "password", []extsvc.RepoID{}, false)
		mockUserEmails.ListByUserFunc.SetDefaultReturn([]*database.UserEmail{{Email: "alice@example.com"}}, nil)
		got, err := p.FetchAccount(ctx, user)
		if err != nil {
			t.Fatal(err)
		}

		accountData, err := jsoniter.Marshal(
			perforce.AccountData{
				Username: "alice",
				Email:    "alice@example.com",
			},
		)
		if err != nil {
			t.Fatal(err)
		}

		want := &extsvc.Account{
			UserID: user.ID,
			AccountSpec: extsvc.AccountSpec{
				ServiceType: p.codeHost.ServiceType,
				ServiceID:   p.codeHost.ServiceID,
				AccountID:   "alice@example.com",
			},
			AccountData: extsvc.AccountData{
				Data: extsvc.NewUnencryptedData(accountData),
			},
		}
		if diff := cmp.Diff(want, got, et.CompareEncryptable); diff != "" {
			t.Fatalf("Mismatch (-want got):\n%s", diff)
		}
	})

	t.Run("found matching account case insensitive", func(t *testing.T) {
		p := NewProvider(logger, db, gitserverClient, "", "ssl:111.222.333.444:1666", "admin", "password", []extsvc.RepoID{}, false)
		mockUserEmails.ListByUserFunc.SetDefaultReturn([]*database.UserEmail{{Email: "Alice@example.com"}}, nil)
		got, err := p.FetchAccount(ctx, user)
		if err != nil {
			t.Fatal(err)
		}

		accountData, err := jsoniter.Marshal(
			perforce.AccountData{
				Username: "alice",
				Email:    "alice@example.com",
			},
		)
		if err != nil {
			t.Fatal(err)
		}

		want := &extsvc.Account{
			UserID: user.ID,
			AccountSpec: extsvc.AccountSpec{
				ServiceType: p.codeHost.ServiceType,
				ServiceID:   p.codeHost.ServiceID,
				AccountID:   "alice@example.com",
			},
			AccountData: extsvc.AccountData{
				Data: extsvc.NewUnencryptedData(accountData),
			},
		}
		if diff := cmp.Diff(want, got, et.CompareEncryptable); diff != "" {
			t.Fatalf("Mismatch (-want got):\n%s", diff)
		}
	})
}

func TestProvider_FetchUserPerms(t *testing.T) {
	ctx := context.Background()

	db := dbmocks.NewMockDB()

	t.Run("nil account", func(t *testing.T) {
		logger := logtest.Scoped(t)
		p := NewProvider(logger, db, gitserver.NewTestClient(t), "", "ssl:111.222.333.444:1666", "admin", "password", []extsvc.RepoID{}, false)
		_, err := p.FetchUserPerms(ctx, nil, authz.FetchPermsOptions{})
		want := "no account provided"
		got := fmt.Sprintf("%v", err)
		if got != want {
			t.Fatalf("err: want %q but got %q", want, got)
		}
	})

	t.Run("not the code host of the account", func(t *testing.T) {
		logger := logtest.Scoped(t)
		p := NewProvider(logger, db, gitserver.NewTestClient(t), "", "ssl:111.222.333.444:1666", "admin", "password", []extsvc.RepoID{}, false)
		_, err := p.FetchUserPerms(context.Background(),
			&extsvc.Account{
				AccountSpec: extsvc.AccountSpec{
					ServiceType: extsvc.TypeGitLab,
					ServiceID:   "https://gitlab.com/",
				},
			},
			authz.FetchPermsOptions{},
		)
		want := `not a code host of the account: want "https://gitlab.com/" but have "ssl:111.222.333.444:1666"`
		got := fmt.Sprintf("%v", err)
		if got != want {
			t.Fatalf("err: want %q but got %q", want, got)
		}
	})

	t.Run("no user found in account data", func(t *testing.T) {
		logger := logtest.Scoped(t)
		p := NewProvider(logger, db, gitserver.NewTestClient(t), "", "ssl:111.222.333.444:1666", "admin", "password", []extsvc.RepoID{}, false)
		_, err := p.FetchUserPerms(ctx,
			&extsvc.Account{
				AccountSpec: extsvc.AccountSpec{
					ServiceType: extsvc.TypePerforce,
					ServiceID:   "ssl:111.222.333.444:1666",
				},
				AccountData: extsvc.AccountData{},
			},
			authz.FetchPermsOptions{},
		)
		want := `no user found in the external account data`
		got := fmt.Sprintf("%v", err)
		if got != want {
			t.Fatalf("err: want %q but got %q", want, got)
		}
	})

	accountData, err := jsoniter.Marshal(
		perforce.AccountData{
			Username: "alice",
			Email:    "alice@example.com",
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name      string
		protects  []*p4types.Protect
		wantPerms *authz.ExternalUserPermissions
	}{
		{
			name: "include only",
			protects: testParseP4ProtectsRaw(t, strings.NewReader(`
list user alice * //Khulnasoft/Security/... ## "list" can't grant read access
read user alice * //Khulnasoft/Engineering/...
owner user alice * //Khulnasoft/Engineering/Backend/...
open user alice * //Khulnasoft/Engineering/Frontend/...
review user alice * //Khulnasoft/Handbook/...
review user alice * //Khulnasoft/*/Handbook/...
review user alice * //Khulnasoft/.../Handbook/...
`)),
			wantPerms: &authz.ExternalUserPermissions{
				IncludeContains: []extsvc.RepoID{
					"//Khulnasoft/Engineering/%",
					"//Khulnasoft/Engineering/Backend/%",
					"//Khulnasoft/Engineering/Frontend/%",
					"//Khulnasoft/Handbook/%",
					"//Khulnasoft/[^/]+/Handbook/%",
					"//Khulnasoft/%/Handbook/%",
				},
			},
		},
		{
			name: "exclude only",
			protects: testParseP4ProtectsRaw(t, strings.NewReader(`
list user alice * -//Khulnasoft/Security/...
read user alice * -//Khulnasoft/Engineering/...
owner user alice * -//Khulnasoft/Engineering/Backend/...
open user alice * -//Khulnasoft/Engineering/Frontend/...
review user alice * -//Khulnasoft/Handbook/...
review user alice * -//Khulnasoft/*/Handbook/...
review user alice * -//Khulnasoft/.../Handbook/...
`)), wantPerms: &authz.ExternalUserPermissions{
				ExcludeContains: []extsvc.RepoID{
					"//Khulnasoft/[^/]+/Handbook/%",
					"//Khulnasoft/%/Handbook/%",
				},
			},
		},
		{
			name: "include and exclude",
			protects: testParseP4ProtectsRaw(t, strings.NewReader(`
read user alice * //Khulnasoft/Security/...
read user alice * //Khulnasoft/Engineering/...
owner user alice * //Khulnasoft/Engineering/Backend/...
open user alice * //Khulnasoft/Engineering/Frontend/...
review user alice * //Khulnasoft/Handbook/...
open user alice * //Khulnasoft/Engineering/.../Frontend/...
open user alice * //Khulnasoft/.../Handbook/...  ## wildcard A

list user alice * -//Khulnasoft/Security/...                        ## "list" can revoke read access
=read user alice * -//Khulnasoft/Engineering/Frontend/...           ## exact match of a previous include
open user alice * -//Khulnasoft/Engineering/Backend/Credentials/... ## sub-match of a previous include
open user alice * -//Khulnasoft/Engineering/*/Frontend/Folder/...   ## sub-match of a previous include
open user alice * -//Khulnasoft/*/Handbook/...                      ## sub-match of wildcard A include
`)),
			wantPerms: &authz.ExternalUserPermissions{
				IncludeContains: []extsvc.RepoID{
					"//Khulnasoft/Engineering/%",
					"//Khulnasoft/Engineering/Backend/%",
					"//Khulnasoft/Engineering/Frontend/%",
					"//Khulnasoft/Handbook/%",
					"//Khulnasoft/Engineering/%/Frontend/%",
					"//Khulnasoft/%/Handbook/%",
				},
				ExcludeContains: []extsvc.RepoID{
					"//Khulnasoft/Engineering/Frontend/%",
					"//Khulnasoft/Engineering/Backend/Credentials/%",
					"//Khulnasoft/Engineering/[^/]+/Frontend/Folder/%",
					"//Khulnasoft/[^/]+/Handbook/%",
				},
			},
		},
		{
			name: "include and exclude, then include again",
			protects: testParseP4ProtectsRaw(t, strings.NewReader(`
read user alice * //Khulnasoft/Security/...
read user alice * //Khulnasoft/Engineering/...
owner user alice * //Khulnasoft/Engineering/Backend/...
open user alice * //Khulnasoft/Engineering/Frontend/...
review user alice * //Khulnasoft/Handbook/...
open user alice * //Khulnasoft/Engineering/.../Frontend/...
open user alice * //Khulnasoft/.../Handbook/...  ## wildcard A

list user alice * -//Khulnasoft/Security/...                        ## "list" can revoke read access
=read user alice * -//Khulnasoft/Engineering/Frontend/...           ## exact match of a previous include
open user alice * -//Khulnasoft/Engineering/Backend/Credentials/... ## sub-match of a previous include
open user alice * -//Khulnasoft/Engineering/*/Frontend/Folder/...   ## sub-match of a previous include
open user alice * -//Khulnasoft/*/Handbook/...                      ## sub-match of wildcard A include

read user alice * //Khulnasoft/Security/... 						 ## give access to alice again after revoking
`)),
			wantPerms: &authz.ExternalUserPermissions{
				IncludeContains: []extsvc.RepoID{
					"//Khulnasoft/Engineering/%",
					"//Khulnasoft/Engineering/Backend/%",
					"//Khulnasoft/Engineering/Frontend/%",
					"//Khulnasoft/Handbook/%",
					"//Khulnasoft/Engineering/%/Frontend/%",
					"//Khulnasoft/%/Handbook/%",
					"//Khulnasoft/Security/%",
				},
				ExcludeContains: []extsvc.RepoID{
					"//Khulnasoft/Engineering/Frontend/%",
					"//Khulnasoft/Engineering/Backend/Credentials/%",
					"//Khulnasoft/Engineering/[^/]+/Frontend/Folder/%",
					"//Khulnasoft/[^/]+/Handbook/%",
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			logger := logtest.Scoped(t)
			gc := gitserver.NewStrictMockClient()
			gc.PerforceProtectsForUserFunc.SetDefaultReturn(test.protects, nil)

			p := NewProvider(logger, db, gc, "", "ssl:111.222.333.444:1666", "admin", "password", []extsvc.RepoID{}, false)
			got, err := p.FetchUserPerms(ctx,
				&extsvc.Account{
					AccountSpec: extsvc.AccountSpec{
						ServiceType: extsvc.TypePerforce,
						ServiceID:   "ssl:111.222.333.444:1666",
					},
					AccountData: extsvc.AccountData{
						Data: extsvc.NewUnencryptedData(accountData),
					},
				},
				authz.FetchPermsOptions{},
			)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(test.wantPerms, got); diff != "" {
				t.Fatalf("Mismatch (-want +got):\n%s", diff)
			}
		})
	}

	// Specific behaviour is tested in TestScanFullRepoPermissions
	t.Run("SubRepoPermissions", func(t *testing.T) {
		for _, test := range []struct {
			name     string
			input    []*p4types.Protect
			expected *authz.ExternalUserPermissions
		}{
			{
				name: "normal",
				input: []*p4types.Protect{
					{Level: "read", EntityType: "user", EntityName: "alice", Host: "*", Match: "//Khulnasoft/Engineering/..."},
					{Level: "read", EntityType: "user", EntityName: "alice", Host: "*", Match: "//Khulnasoft/Security/...", IsExclusion: true},
				},
				expected: &authz.ExternalUserPermissions{
					Exacts: []extsvc.RepoID{"//Khulnasoft/"},
					SubRepoPermissions: map[extsvc.RepoID]*authz.SubRepoPermissionsWithIPs{
						"//Khulnasoft/": {
							Paths: []authz.PathWithIP{
								{Path: mustGlobPattern(t, "/Engineering/..."), IP: "*"},
								{Path: mustGlobPattern(t, "-/Security/..."), IP: "*"},
							},
						},
					},
				},
			},
			{
				name: "with ips",

				input: []*p4types.Protect{
					{Level: "read", EntityType: "user", EntityName: "alice", Host: "1.2.3.6", Match: "//Khulnasoft/Engineering/..."},
					{Level: "read", EntityType: "user", EntityName: "alice", Host: "1.2.3.4", Match: "//Khulnasoft/Security/...", IsExclusion: true},
				},
				expected: &authz.ExternalUserPermissions{
					Exacts: []extsvc.RepoID{"//Khulnasoft/"},
					SubRepoPermissions: map[extsvc.RepoID]*authz.SubRepoPermissionsWithIPs{
						"//Khulnasoft/": {
							Paths: []authz.PathWithIP{
								{Path: mustGlobPattern(t, "/Engineering/..."), IP: "1.2.3.6"},
								{Path: mustGlobPattern(t, "-/Security/..."), IP: "1.2.3.4"},
							},
						},
					},
				},
			},
		} {
			t.Run(test.name, func(t *testing.T) {
				logger := logtest.Scoped(t)

				gitserverClient := gitserver.NewStrictMockClient()

				gitserverClient.PerforceProtectsForDepotFunc.SetDefaultReturn(test.input, nil)
				gitserverClient.PerforceProtectsForUserFunc.SetDefaultReturn(test.input, nil)

				p := NewProvider(logger, db, gitserverClient, "", "ssl:111.222.333.444:1666", "admin", "password", []extsvc.RepoID{}, false)
				p.depots = append(p.depots, "//Khulnasoft/")

				got, err := p.FetchUserPerms(ctx,
					&extsvc.Account{
						AccountSpec: extsvc.AccountSpec{
							ServiceType: extsvc.TypePerforce,
							ServiceID:   "ssl:111.222.333.444:1666",
						},
						AccountData: extsvc.AccountData{
							Data: extsvc.NewUnencryptedData(accountData),
						},
					},
					authz.FetchPermsOptions{},
				)
				if err != nil {
					t.Fatal(err)
				}

				if diff := cmp.Diff(test.expected, got); diff != "" {
					t.Fatalf("Mismatch (-want +got):\n%s", diff)
				}
			})
		}
	})
}

func TestProvider_FetchRepoPerms(t *testing.T) {
	logger := logtest.Scoped(t)
	ctx := context.Background()
	db := dbmocks.NewMockDB()

	t.Run("nil repository", func(t *testing.T) {
		p := NewProvider(logger, db, gitserver.NewTestClient(t), "", "ssl:111.222.333.444:1666", "admin", "password", []extsvc.RepoID{}, false)
		_, err := p.FetchRepoPerms(ctx, nil, authz.FetchPermsOptions{})
		want := "no repository provided"
		got := fmt.Sprintf("%v", err)
		if got != want {
			t.Fatalf("err: want %q but got %q", want, got)
		}
	})

	t.Run("not the code host of the repository", func(t *testing.T) {
		p := NewProvider(logger, db, gitserver.NewTestClient(t), "", "ssl:111.222.333.444:1666", "admin", "password", []extsvc.RepoID{}, false)
		_, err := p.FetchRepoPerms(ctx,
			&extsvc.Repository{
				URI: "gitlab.com/user/repo",
				ExternalRepoSpec: api.ExternalRepoSpec{
					ServiceType: extsvc.TypeGitLab,
					ServiceID:   "https://gitlab.com/",
				},
			},
			authz.FetchPermsOptions{},
		)
		want := `not a code host of the repository: want "https://gitlab.com/" but have "ssl:111.222.333.444:1666"`
		got := fmt.Sprintf("%v", err)
		if got != want {
			t.Fatalf("err: want %q but got %q", want, got)
		}
	})

	gitserverClient := gitserver.NewStrictMockClient()
	gitserverClient.PerforceUsersFunc.SetDefaultReturn([]*p4types.User{
		{Username: "alice", Email: "alice@example.com"},
		{Username: "bob", Email: "bob@example.com"},
		{Username: "cindy", Email: "cindy@example.com"},
		{Username: "david", Email: "david@example.com"},
		{Username: "frank", Email: "frank@example.com"},
	}, nil)
	gitserverClient.PerforceGroupMembersFunc.SetDefaultHook(func(ctx context.Context, conn protocol.PerforceConnectionDetails, group string) ([]string, error) {
		switch group {
		case "Backend":
			return []string{"alice", "cindy"}, nil
		case "Frontend":
			return []string{"bob", "david", "frank"}, nil
		default:
			return nil, errors.New("invalid group")
		}
	})
	gitserverClient.PerforceProtectsForDepotFunc.SetDefaultReturn([]*p4types.Protect{
		{Level: "list", EntityType: "user", EntityName: "*", Host: "*", Match: "//...", IsExclusion: true},
		{Level: "write", EntityType: "user", EntityName: "alice", Host: "*", Match: "//Khulnasoft/..."},
		{Level: "write", EntityType: "user", EntityName: "bob", Host: "*", Match: "//Khulnasoft/..."},
		{Level: "admin", EntityType: "group", EntityName: "Backend", Host: "*", Match: "//Khulnasoft/..."},                     // includes "alice" and "cindy"
		{Level: "admin", EntityType: "group", EntityName: "Frontend", Host: "*", Match: "//Khulnasoft/...", IsExclusion: true}, // excludes "bob", "david" and "frank"
		{Level: "read", EntityType: "user", EntityName: "cindy", Host: "*", Match: "//Khulnasoft/...", IsExclusion: true},
		{Level: "list", EntityType: "user", EntityName: "david", Host: "*", Match: "//Khulnasoft/..."}, // "list" can't grant read access
	}, nil)

	p := NewProvider(logger, db, gitserverClient, "", "ssl:111.222.333.444:1666", "admin", "password", []extsvc.RepoID{}, false)
	got, err := p.FetchRepoPerms(ctx,
		&extsvc.Repository{
			URI: "gitlab.com/user/repo",
			ExternalRepoSpec: api.ExternalRepoSpec{
				ServiceType: extsvc.TypePerforce,
				ServiceID:   "ssl:111.222.333.444:1666",
			},
		},
		authz.FetchPermsOptions{},
	)
	if err != nil {
		t.Fatal(err)
	}

	want := []extsvc.AccountID{"alice@example.com"}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Fatalf("Mismatch (-want +got):\n%s", diff)
	}
}
