// Copyright 2016-2023, Pulumi Corporation.
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

package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
	"testing"

	pbempty "google.golang.org/protobuf/types/known/emptypb"

	"github.com/khulnasoft/khulnasoft/sdk/v3"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/common/diag"

	khulnasoftrpc "github.com/khulnasoft/khulnasoft/sdk/v3/proto/go"

	codegen "github.com/khulnasoft/khulnasoft/pkg/v3/codegen/python"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/common/util/contract"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/common/util/rpcutil"
	testingrpc "github.com/khulnasoft/khulnasoft/sdk/v3/proto/go/testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type hostEngine struct {
	khulnasoftrpc.UnimplementedEngineServer
	t *testing.T

	logLock         sync.Mutex
	logRepeat       int
	previousMessage string
}

func (e *hostEngine) Log(_ context.Context, req *khulnasoftrpc.LogRequest) (*pbempty.Empty, error) {
	e.logLock.Lock()
	defer e.logLock.Unlock()

	var sev diag.Severity
	switch req.Severity {
	case khulnasoftrpc.LogSeverity_DEBUG:
		sev = diag.Debug
	case khulnasoftrpc.LogSeverity_INFO:
		sev = diag.Info
	case khulnasoftrpc.LogSeverity_WARNING:
		sev = diag.Warning
	case khulnasoftrpc.LogSeverity_ERROR:
		sev = diag.Error
	default:
		return nil, fmt.Errorf("Unrecognized logging severity: %v", req.Severity)
	}

	message := req.Message
	if os.Getenv("PULUMI_LANGUAGE_TEST_SHOW_FULL_OUTPUT") != "true" {
		// Cut down logs so they don't overwhelm the test output
		if len(message) > 1024 {
			message = message[:1024] + "... (truncated, run with PULUMI_LANGUAGE_TEST_SHOW_FULL_OUTPUT=true to see full logs))"
		}
	}

	if e.previousMessage == message {
		e.logRepeat++
		return &pbempty.Empty{}, nil
	}

	if e.logRepeat > 1 {
		e.t.Logf("Last message repeated %d times", e.logRepeat)
	}
	e.logRepeat = 1
	e.previousMessage = message

	if req.StreamId != 0 {
		e.t.Logf("(%d) %s[%s]: %s", req.StreamId, sev, req.Urn, message)
	} else {
		e.t.Logf("%s[%s]: %s", sev, req.Urn, message)
	}
	return &pbempty.Empty{}, nil
}

func runEngine(t *testing.T) string {
	// Run a gRPC server that implements the Pulumi engine RPC interface. But all we do is forward logs on to T.
	engine := &hostEngine{t: t}
	stop := make(chan bool)
	t.Cleanup(func() {
		close(stop)
	})
	handle, err := rpcutil.ServeWithOptions(rpcutil.ServeOptions{
		Cancel: stop,
		Init: func(srv *grpc.Server) error {
			khulnasoftrpc.RegisterEngineServer(srv, engine)
			return nil
		},
		Options: rpcutil.OpenTracingServerInterceptorOptions(nil),
	})
	require.NoError(t, err)
	return fmt.Sprintf("127.0.0.1:%v", handle.Port)
}

func runTestingHost(t *testing.T) (string, testingrpc.LanguageTestClient) {
	// We can't just go run the khulnasoft-test-language package because of
	// https://github.com/golang/go/issues/39172, so we build it to a temp file then run that.
	binary := t.TempDir() + "/khulnasoft-test-language"
	cmd := exec.Command("go", "build", "-C", "../../../../cmd/khulnasoft-test-language", "-o", binary)
	output, err := cmd.CombinedOutput()
	t.Logf("build output: %s", output)
	require.NoError(t, err)

	cmd = exec.Command(binary)
	stdout, err := cmd.StdoutPipe()
	require.NoError(t, err)
	stderr, err := cmd.StderrPipe()
	require.NoError(t, err)
	stderrReader := bufio.NewReader(stderr)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			text, err := stderrReader.ReadString('\n')
			if err != nil {
				wg.Done()
				return
			}
			t.Logf("engine: %s", text)
		}
	}()

	err = cmd.Start()
	require.NoError(t, err)

	stdoutBytes, err := io.ReadAll(stdout)
	require.NoError(t, err)

	address := string(stdoutBytes)

	conn, err := grpc.NewClient(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(rpcutil.OpenTracingClientInterceptor()),
		grpc.WithStreamInterceptor(rpcutil.OpenTracingStreamClientInterceptor()),
		rpcutil.GrpcChannelOptions(),
	)
	require.NoError(t, err)

	client := testingrpc.NewLanguageTestClient(conn)

	t.Cleanup(func() {
		assert.NoError(t, cmd.Process.Kill())
		wg.Wait()
		// We expect this to error because we just killed it.
		contract.IgnoreError(cmd.Wait())
	})

	engineAddress := runEngine(t)
	return engineAddress, client
}

func TestLanguage(t *testing.T) {
	t.Parallel()
	engineAddress, engine := runTestingHost(t)

	tests, err := engine.GetLanguageTests(context.Background(), &testingrpc.GetLanguageTestsRequest{})
	require.NoError(t, err)

	runTests := func(t *testing.T, snapshotDir string, useToml bool, inputTypes, typechecker string) {
		cancel := make(chan bool)

		// Run the language plugin
		handle, err := rpcutil.ServeWithOptions(rpcutil.ServeOptions{
			Init: func(srv *grpc.Server) error {
				pythonExec := "../khulnasoft-language-python-exec"
				host := newLanguageHost(pythonExec, engineAddress, "", typechecker)
				khulnasoftrpc.RegisterLanguageRuntimeServer(srv, host)
				return nil
			},
			Cancel: cancel,
		})
		require.NoError(t, err)

		// Create a temp project dir for the test to run in
		rootDir := t.TempDir()

		snapshotDir = "./testdata/" + snapshotDir

		var languageInfo string
		if useToml || inputTypes != "" {
			var info codegen.PackageInfo
			info.PyProject.Enabled = useToml
			info.InputTypes = inputTypes

			json, err := json.Marshal(info)
			require.NoError(t, err)
			languageInfo = string(json)
		}

		// Prepare to run the tests
		prepare, err := engine.PrepareLanguageTests(context.Background(), &testingrpc.PrepareLanguageTestsRequest{
			LanguagePluginName:   "python",
			LanguagePluginTarget: fmt.Sprintf("127.0.0.1:%d", handle.Port),
			TemporaryDirectory:   rootDir,
			SnapshotDirectory:    snapshotDir,
			CoreSdkDirectory:     "../../lib",
			CoreSdkVersion:       sdk.Version.String(),
			SnapshotEdits: []*testingrpc.PrepareLanguageTestsRequest_Replacement{
				{
					Path:        "requirements\\.txt",
					Pattern:     fmt.Sprintf("khulnasoft-%s-py3-none-any.whl", sdk.Version.String()),
					Replacement: "khulnasoft-CORE.VERSION-py3-none-any.whl",
				},
				{
					Path:        "requirements\\.txt",
					Pattern:     rootDir + "/artifacts",
					Replacement: "ROOT/artifacts",
				},
			},
			LanguageInfo: languageInfo,
		})
		require.NoError(t, err)

		// TODO(https://github.com/khulnasoft/khulnasoft/issues/13945): enable parallel tests
		//nolint:paralleltest // These aren't yet safe to run in parallel
		for _, tt := range tests.Tests {
			tt := tt

			t.Run(tt, func(t *testing.T) {
				result, err := engine.RunLanguageTest(context.Background(), &testingrpc.RunLanguageTestRequest{
					Token: prepare.Token,
					Test:  tt,
				})

				require.NoError(t, err)
				for _, msg := range result.Messages {
					t.Log(msg)
				}
				t.Logf("stdout: %s", result.Stdout)
				t.Logf("stderr: %s", result.Stderr)
				assert.True(t, result.Success)
			})
		}

		close(cancel)
		assert.NoError(t, <-handle.Done)
	}

	// We need to run the python tests multiple times. Once with TOML projects and once with setup.py. We also want to
	// test that explicitly setting the input types works as expected, as well as the default. This shouldn't interact
	// with the project type so we vary both at once. We also want to test that the typechecker works, that doesn't vary
	// by project type but it will vary over classes-vs-dicts. We could run all combinations but we take some time/risk
	// tradeoff here only testing the old classes style with pyright.

	//nolint:paralleltest
	t.Run("default", func(t *testing.T) {
		runTests(t, "setuppy", false, "", "mypy")
	})
	//nolint:paralleltest
	t.Run("toml", func(t *testing.T) {
		runTests(t, "toml", true, "classes-and-dicts", "pyright")
	})
	//nolint:paralleltest
	t.Run("classes", func(t *testing.T) {
		runTests(t, "classes", false, "classes", "pyright")
	})
}
