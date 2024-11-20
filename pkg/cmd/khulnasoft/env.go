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
	"fmt"
	"io"

	"github.com/spf13/cobra"

	"github.com/khulnasoft/esc/cmd/esc/cli"
	escWorkspace "github.com/khulnasoft/esc/cmd/esc/cli/workspace"
	"github.com/khulnasoft/khulnasoft/pkg/v3/backend/httpstate"
	"github.com/khulnasoft/khulnasoft/pkg/v3/backend/httpstate/client"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/common/apitype"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/common/util/cmdutil"
)

func newEnvCmd() *cobra.Command {
	escCLI := cli.New(&cli.Options{
		ParentPath:      "khulnasoft",
		Colors:          cmdutil.GetGlobalColorization(),
		Login:           httpstate.NewLoginManager(),
		PulumiWorkspace: escWorkspace.DefaultPulumiWorkspace(),
		UserAgent:       client.UserAgent(),
	})

	// Add the `env` command to the root.
	envCommand := escCLI.Commands()[0]
	return envCommand
}

func printESCDiagnostics(out io.Writer, diags []apitype.EnvironmentDiagnostic) {
	for _, d := range diags {
		if d.Range != nil {
			fmt.Fprintf(out, "%v:", d.Range.Environment)
			if d.Range.Begin.Line != 0 {
				fmt.Fprintf(out, "%v:%v:", d.Range.Begin.Line, d.Range.Begin.Column)
			}
			fmt.Fprintf(out, " ")
		}
		fmt.Fprintf(out, "%v\n", d.Summary)
	}
}