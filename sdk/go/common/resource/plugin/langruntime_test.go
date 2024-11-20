// Copyright 2024, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package plugin

import (
	"context"
	"errors"
	"testing"

	khulnasoftrpc "github.com/khulnasoft/khulnasoft/sdk/v3/proto/go"
	"github.com/stretchr/testify/require"

	"google.golang.org/protobuf/types/known/emptypb"

	grpc "google.golang.org/grpc"
)

func TestMakeExecutablePromptChoices(t *testing.T) {
	t.Parallel()

	// Not found executables come after the found ones, and have a [not found] suffix.
	choices := MakeExecutablePromptChoices("executable_that_does_not_exist_in_path", "ls")
	require.Equal(t, 2, len(choices))
	require.Equal(t, choices[0].StringValue, "ls")
	require.Equal(t, choices[0].DisplayName, "ls")
	require.Equal(t, choices[1].StringValue, "executable_that_does_not_exist_in_path")
	require.Equal(t, choices[1].DisplayName, "executable_that_does_not_exist_in_path [not found]")

	// Executables are not reordered within their group.
	choices = MakeExecutablePromptChoices("ls", "cat", "zzz_does_not_exist", "aaa_does_not_exist")
	require.Equal(t, choices[0].StringValue, "ls")
	require.Equal(t, choices[0].DisplayName, "ls")
	require.Equal(t, choices[1].StringValue, "cat")
	require.Equal(t, choices[1].DisplayName, "cat")
	require.Equal(t, choices[2].StringValue, "zzz_does_not_exist")
	require.Equal(t, choices[2].DisplayName, "zzz_does_not_exist [not found]")
	require.Equal(t, choices[3].StringValue, "aaa_does_not_exist")
	require.Equal(t, choices[3].DisplayName, "aaa_does_not_exist [not found]")
}

type MockLanguageRuntimeClient struct {
	RunPluginF (func(ctx context.Context, info *khulnasoftrpc.RunPluginRequest,
	) (khulnasoftrpc.LanguageRuntime_RunPluginClient, error))
}

func (m *MockLanguageRuntimeClient) RunPlugin(
	ctx context.Context, info *khulnasoftrpc.RunPluginRequest, _ ...grpc.CallOption,
) (khulnasoftrpc.LanguageRuntime_RunPluginClient, error) {
	return m.RunPluginF(ctx, info)
}

func (m *MockLanguageRuntimeClient) GetRequiredPlugins(
	ctx context.Context, in *khulnasoftrpc.GetRequiredPluginsRequest, opts ...grpc.CallOption,
) (*khulnasoftrpc.GetRequiredPluginsResponse, error) {
	panic("not implemented")
}

func (m *MockLanguageRuntimeClient) Run(
	ctx context.Context, in *khulnasoftrpc.RunRequest, opts ...grpc.CallOption,
) (*khulnasoftrpc.RunResponse, error) {
	panic("not implemented")
}

func (m *MockLanguageRuntimeClient) GetPluginInfo(
	ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption,
) (*khulnasoftrpc.PluginInfo, error) {
	panic("not implemented")
}

func (m *MockLanguageRuntimeClient) InstallDependencies(
	ctx context.Context, in *khulnasoftrpc.InstallDependenciesRequest, opts ...grpc.CallOption,
) (khulnasoftrpc.LanguageRuntime_InstallDependenciesClient, error) {
	panic("not implemented")
}

func (m *MockLanguageRuntimeClient) RuntimeOptionsPrompts(
	ctx context.Context, in *khulnasoftrpc.RuntimeOptionsRequest, opts ...grpc.CallOption,
) (*khulnasoftrpc.RuntimeOptionsResponse, error) {
	panic("not implemented")
}

func (m *MockLanguageRuntimeClient) About(
	ctx context.Context, in *khulnasoftrpc.AboutRequest, opts ...grpc.CallOption,
) (*khulnasoftrpc.AboutResponse, error) {
	panic("not implemented")
}

func (m *MockLanguageRuntimeClient) GetProgramDependencies(
	ctx context.Context, in *khulnasoftrpc.GetProgramDependenciesRequest, opts ...grpc.CallOption,
) (*khulnasoftrpc.GetProgramDependenciesResponse, error) {
	panic("not implemented")
}

func (m *MockLanguageRuntimeClient) GenerateProgram(
	ctx context.Context, in *khulnasoftrpc.GenerateProgramRequest, opts ...grpc.CallOption,
) (*khulnasoftrpc.GenerateProgramResponse, error) {
	panic("not implemented")
}

func (m *MockLanguageRuntimeClient) GenerateProject(
	ctx context.Context, in *khulnasoftrpc.GenerateProjectRequest, opts ...grpc.CallOption,
) (*khulnasoftrpc.GenerateProjectResponse, error) {
	panic("not implemented")
}

func (m *MockLanguageRuntimeClient) GeneratePackage(
	ctx context.Context, in *khulnasoftrpc.GeneratePackageRequest, opts ...grpc.CallOption,
) (*khulnasoftrpc.GeneratePackageResponse, error) {
	panic("not implemented")
}

func (m *MockLanguageRuntimeClient) Pack(
	ctx context.Context, in *khulnasoftrpc.PackRequest, opts ...grpc.CallOption,
) (*khulnasoftrpc.PackResponse, error) {
	panic("not implemented")
}

func TestRunPluginPassesCorrectPwd(t *testing.T) {
	t.Parallel()

	returnErr := errors.New("erroring so we don't need to implement the whole thing")
	mockLanguageRuntime := &MockLanguageRuntimeClient{
		RunPluginF: func(
			ctx context.Context, info *khulnasoftrpc.RunPluginRequest,
		) (khulnasoftrpc.LanguageRuntime_RunPluginClient, error) {
			require.Equal(t, "/tmp", info.Pwd)
			return nil, returnErr
		},
	}

	pCtx, err := NewContext(nil, nil, nil, nil, "", nil, false, nil)
	require.NoError(t, err)
	host := &langhost{
		ctx:     pCtx,
		runtime: "go",
		plug:    nil,
		client:  mockLanguageRuntime,
	}

	// Test that the plugin is run with the correct working directory.
	_, _, _, err = host.RunPlugin(RunPluginInfo{
		WorkingDirectory: "/tmp",
	})
	require.Equal(t, returnErr, err)
}
