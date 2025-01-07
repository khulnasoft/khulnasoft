package repos

import (
	"context"
	"fmt"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/khulnasoft/khulnasoft/internal/extsvc"
	"github.com/khulnasoft/khulnasoft/internal/gitserver"
	"github.com/khulnasoft/khulnasoft/internal/gitserver/protocol"
	"github.com/khulnasoft/khulnasoft/internal/testutil"
	"github.com/khulnasoft/khulnasoft/internal/types"
	"github.com/khulnasoft/khulnasoft/internal/types/typestest"
	"github.com/khulnasoft/khulnasoft/lib/errors"
	"github.com/khulnasoft/khulnasoft/schema"
)

func TestPerforceSource_ListRepos(t *testing.T) {
	assertAllReposListed := func(want []string) typestest.ReposAssertion {
		return func(t testing.TB, rs types.Repos) {
			t.Helper()

			have := rs.Names()
			sort.Strings(have)
			sort.Strings(want)

			if diff := cmp.Diff(want, have); diff != "" {
				t.Errorf("Mismatch (-want +got):\n%s", diff)
			}
		}
	}

	testCases := []struct {
		name   string
		assert typestest.ReposAssertion
		conf   *schema.PerforceConnection
		err    string
	}{
		{
			name: "list",
			assert: assertAllReposListed([]string{
				"Khulnasoft",
				"Engineering/Cloud",
			}),
			conf: &schema.PerforceConnection{
				P4Port:   "ssl:111.222.333.444:1666",
				P4User:   "admin",
				P4Passwd: "pa$$word",
				Depots: []string{
					"//Khulnasoft",
					"//Engineering/Cloud",
				},
			},
			err: "<nil>",
		},
		{
			name: "unknown depot among existing",
			assert: assertAllReposListed([]string{
				"Khulnasoft",
			}),
			conf: &schema.PerforceConnection{
				P4Port:   "ssl:111.222.333.444:1666",
				P4User:   "admin",
				P4Passwd: "pa$$word",
				Depots: []string{
					"//Khulnasoft",
					"//NotFound",
				},
			},
			err: "checking if perforce path is cloneable: unknown depot",
		},
	}

	for _, tc := range testCases {
		tc := tc
		tc.name = "PERFORCE-LIST-REPOS/" + tc.name
		t.Run(tc.name, func(t *testing.T) {
			svc := typestest.MakeExternalService(t, extsvc.VariantPerforce, tc.conf)

			gc := gitserver.NewMockClient()
			gc.IsPerforcePathCloneableFunc.SetDefaultHook(func(ctx context.Context, _ protocol.PerforceConnectionDetails, depotPath string) error {
				if depotPath == "//Khulnasoft" || depotPath == "//Engineering/Cloud" {
					return nil
				}
				return errors.New("unknown depot")
			})

			perforceSrc, err := newPerforceSource(gc, svc, tc.conf)
			if err != nil {
				t.Fatal(err)
			}

			repos, err := ListAll(context.Background(), perforceSrc)

			if have, want := fmt.Sprint(err), tc.err; have != want {
				t.Errorf("error:\nhave: %q\nwant: %q", have, want)
			}

			if tc.assert != nil {
				tc.assert(t, repos)
			}
		})
	}
}

func TestPerforceSource_makeRepo(t *testing.T) {
	depots := []string{
		"//Khulnasoft",
		"//Engineering/Cloud",
	}

	svc := types.ExternalService{
		ID:     1,
		Kind:   extsvc.KindPerforce,
		Config: extsvc.NewEmptyConfig(),
	}

	tests := []struct {
		name   string
		schema *schema.PerforceConnection
	}{
		{
			name: "simple",
			schema: &schema.PerforceConnection{
				P4Port:   "ssl:111.222.333.444:1666",
				P4User:   "admin",
				P4Passwd: "pa$$word",
			},
		}, {
			name: "path-pattern",
			schema: &schema.PerforceConnection{
				P4Port:                "ssl:111.222.333.444:1666",
				P4User:                "admin",
				P4Passwd:              "pa$$word",
				RepositoryPathPattern: "perforce/{depot}",
			},
		},
	}
	for _, test := range tests {
		test.name = "PerforceSource_makeRepo_" + test.name
		t.Run(test.name, func(t *testing.T) {
			s, err := newPerforceSource(gitserver.NewMockClient(), &svc, test.schema)
			if err != nil {
				t.Fatal(err)
			}

			var got []*types.Repo
			for _, depot := range depots {
				got = append(got, s.makeRepo(depot))
			}

			testutil.AssertGolden(t, "testdata/golden/"+test.name, Update(test.name), got)
		})
	}
}
