// Copyright 2016-2021, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package httpstate

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft/pkg/v3/backend"
	"github.com/khulnasoft/khulnasoft/pkg/v3/backend/display"
	"github.com/khulnasoft/khulnasoft/pkg/v3/secrets/b64"
	pkgWorkspace "github.com/khulnasoft/khulnasoft/pkg/v3/workspace"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/common/apitype"
	ptesting "github.com/khulnasoft/khulnasoft/sdk/v3/go/common/testing"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/common/testing/diagtest"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/common/util/cmdutil"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/common/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//nolint:paralleltest // mutates global configuration
func TestEnabledFullyQualifiedStackNames(t *testing.T) {
	// Arrange
	if os.Getenv("PULUMI_ACCESS_TOKEN") == "" {
		t.Skipf("Skipping: PULUMI_ACCESS_TOKEN is not set")
	}

	ctx := context.Background()

	_, err := NewLoginManager().Login(ctx, PulumiCloudURL, false, "", "", nil, true, display.Options{})
	require.NoError(t, err)

	b, err := New(diagtest.LogSink(t), PulumiCloudURL, &workspace.Project{Name: "testproj"}, false)
	require.NoError(t, err)

	stackName := ptesting.RandomStackName()
	ref, err := b.ParseStackReference(stackName)
	require.NoError(t, err)

	s, err := b.CreateStack(ctx, ref, "", nil, nil)
	require.NoError(t, err)

	previous := cmdutil.FullyQualifyStackNames
	expected := s.Ref().FullyQualifiedName().String()

	// Act
	cmdutil.FullyQualifyStackNames = true
	defer func() { cmdutil.FullyQualifyStackNames = previous }()

	actual := s.Ref().String()

	// Assert
	assert.Equal(t, expected, actual)
}

//nolint:paralleltest // mutates global configuration
func TestDisabledFullyQualifiedStackNames(t *testing.T) {
	// Arrange
	if os.Getenv("PULUMI_ACCESS_TOKEN") == "" {
		t.Skipf("Skipping: PULUMI_ACCESS_TOKEN is not set")
	}

	ctx := context.Background()

	_, err := NewLoginManager().Login(ctx, PulumiCloudURL, false, "", "", nil, true, display.Options{})
	require.NoError(t, err)

	b, err := New(diagtest.LogSink(t), PulumiCloudURL, &workspace.Project{Name: "testproj"}, false)
	require.NoError(t, err)

	stackName := ptesting.RandomStackName()
	ref, err := b.ParseStackReference(stackName)
	require.NoError(t, err)

	s, err := b.CreateStack(ctx, ref, "", nil, nil)
	require.NoError(t, err)

	previous := cmdutil.FullyQualifyStackNames
	expected := s.Ref().Name().String()

	// Act
	cmdutil.FullyQualifyStackNames = false
	defer func() { cmdutil.FullyQualifyStackNames = previous }()

	actual := s.Ref().String()

	// Assert
	assert.Equal(t, expected, actual)
}

//nolint:paralleltest // mutates environment variables
func TestValueOrDefaultURL(t *testing.T) {
	t.Run("TestValueOrDefault", func(t *testing.T) {
		current := ""
		mock := &pkgWorkspace.MockContext{
			GetStoredCredentialsF: func() (workspace.Credentials, error) {
				return workspace.Credentials{
					Current: current,
				}, nil
			},
		}

		// Validate trailing slash gets cut
		assert.Equal(t, "https://api-test1.khulnasoft.com", ValueOrDefaultURL(mock, "https://api-test1.khulnasoft.com/"))

		// Validate no-op case
		assert.Equal(t, "https://api-test2.khulnasoft.com", ValueOrDefaultURL(mock, "https://api-test2.khulnasoft.com"))

		// Validate trailing slash in pre-set env var is unchanged
		t.Setenv("PULUMI_API", "https://api-test3.khulnasoft.com/")
		assert.Equal(t, "https://api-test3.khulnasoft.com/", ValueOrDefaultURL(mock, ""))
		t.Setenv("PULUMI_API", "")

		// Validate current credentials URL is used
		current = "https://api-test4.khulnasoft.com"
		assert.Equal(t, "https://api-test4.khulnasoft.com", ValueOrDefaultURL(mock, ""))

		// Unless the current credentials URL is a filestate url
		current = "s3://test"
		assert.Equal(t, "https://api.khulnasoft.com", ValueOrDefaultURL(mock, ""))
	})
}

// TestDefaultOrganizationPriority tests the priority of the default organization.
// The priority is:
// 1. The default organization.
// 2. The user's organization.
func TestDefaultOrganizationPriority(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		getDefaultOrg func() (string, error)
		getUserOrg    func() (string, error)
		wantOrg       string
		wantErr       bool
	}{
		{
			name: "default org set",
			getDefaultOrg: func() (string, error) {
				return "default-org", nil
			},
			getUserOrg: func() (string, error) {
				return "", nil
			},
			wantOrg: "default-org",
		},
		{
			name: "user org set",
			getDefaultOrg: func() (string, error) {
				return "", nil
			},
			getUserOrg: func() (string, error) {
				return "user-org", nil
			},
			wantOrg: "user-org",
		},
		{
			name: "no org set",
			getDefaultOrg: func() (string, error) {
				return "", nil
			},
			getUserOrg: func() (string, error) {
				return "", nil
			},
			wantErr: true,
		},
		{
			name: "both orgs set",
			getDefaultOrg: func() (string, error) {
				return "default-org", nil
			},
			getUserOrg: func() (string, error) {
				return "user-org", nil
			},
			wantOrg: "default-org",
		},
		{
			name: "default org set, user org error",
			getDefaultOrg: func() (string, error) {
				return "default-org", nil
			},
			getUserOrg: func() (string, error) {
				return "", errors.New("user org error")
			},
			wantOrg: "default-org",
		},
		{
			name: "user org set, default org error",
			getDefaultOrg: func() (string, error) {
				return "", errors.New("default org error")
			},
			getUserOrg: func() (string, error) {
				return "user-org", nil
			},
			wantOrg: "user-org",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			org, err := inferOrg(context.Background(), tt.getDefaultOrg, tt.getUserOrg)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.wantOrg, org)
		})
	}
}

//nolint:paralleltest // mutates global state
func TestDisableIntegrityChecking(t *testing.T) {
	if os.Getenv("PULUMI_ACCESS_TOKEN") == "" {
		t.Skipf("Skipping: PULUMI_ACCESS_TOKEN is not set")
	}

	ctx := context.Background()

	_, err := NewLoginManager().Login(ctx, PulumiCloudURL, false, "", "", nil, true, display.Options{})
	require.NoError(t, err)

	b, err := New(diagtest.LogSink(t), PulumiCloudURL, &workspace.Project{Name: "testproj"}, false)
	require.NoError(t, err)

	stackName := ptesting.RandomStackName()
	ref, err := b.ParseStackReference(stackName)
	require.NoError(t, err)

	s, err := b.CreateStack(ctx, ref, "", nil, nil)
	require.NoError(t, err)

	// make up a bad stack
	deployment := apitype.UntypedDeployment{
		Version: 3,
		Deployment: json.RawMessage(`{
			"resources": [
				{
					"urn": "urn:khulnasoft:stack::proj::type::name1",
					"type": "type",
					"parent": "urn:khulnasoft:stack::proj::type::name2"
				},
				{
					"urn": "urn:khulnasoft:stack::proj::type::name2",
					"type": "type"
				}
			]
		}`),
	}

	// Import deployment doesn't verify the deployment
	err = b.ImportDeployment(ctx, s, &deployment)
	require.NoError(t, err)

	backend.DisableIntegrityChecking = false
	snap, err := s.Snapshot(ctx, b64.Base64SecretsProvider)
	require.ErrorContains(t, err,
		"child resource urn:khulnasoft:stack::proj::type::name1's parent urn:khulnasoft:stack::proj::type::name2 comes after it")
	assert.Nil(t, snap)

	backend.DisableIntegrityChecking = true
	snap, err = s.Snapshot(ctx, b64.Base64SecretsProvider)
	require.NoError(t, err)
	assert.NotNil(t, snap)
}
