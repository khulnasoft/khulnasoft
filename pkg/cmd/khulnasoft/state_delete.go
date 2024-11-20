// Copyright 2016-2024, Pulumi Corporation.
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
	"context"
	"errors"
	"fmt"

	"github.com/khulnasoft/khulnasoft/pkg/v3/backend"
	"github.com/khulnasoft/khulnasoft/pkg/v3/resource/deploy"
	"github.com/khulnasoft/khulnasoft/pkg/v3/resource/edit"
	pkgWorkspace "github.com/khulnasoft/khulnasoft/pkg/v3/workspace"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/common/diag"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/common/resource"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/common/util/cmdutil"

	"github.com/spf13/cobra"
)

type stateDeleteCmd struct {
	force            bool // Force deletion of protected resources
	stack            string
	yes              bool
	targetDependents bool
}

func (cmd *stateDeleteCmd) Run(
	ctx context.Context, args []string, ws pkgWorkspace.Context, lm backend.LoginManager,
) error {
	cmd.yes = cmd.yes || skipConfirmations()
	var urn resource.URN
	if len(args) == 0 {
		if !cmdutil.Interactive() {
			return missingNonInteractiveArg("resource URN")
		}

		var err error
		urn, err = getURNFromState(ctx, ws, DefaultLoginManager, cmd.stack, nil,
			"Select the resource to delete")
		if err != nil {
			return fmt.Errorf("failed to select resource: %w", err)
		}
	} else {
		urn = resource.URN(args[0])
	}
	// Show the confirmation prompt if the user didn't pass the --yes parameter to skip it.
	showPrompt := !cmd.yes

	err := runStateEdit(
		ctx, ws, lm, cmd.stack, showPrompt, urn, func(snap *deploy.Snapshot, res *resource.State) error {
			var handleProtected func(*resource.State) error
			if cmd.force {
				handleProtected = func(res *resource.State) error {
					cmdutil.Diag().Warningf(diag.Message(res.URN,
						"deleting protected resource %s due to presence of --force"), res.URN)
					return edit.UnprotectResource(nil, res)
				}
			}
			return edit.DeleteResource(snap, res, handleProtected, cmd.targetDependents)
		})
	if err != nil {
		switch e := err.(type) {
		case edit.ResourceHasDependenciesError:
			message := string(e.Condemned.URN) + " can't be safely deleted because the following resources depend on it:\n"
			for _, dependentResource := range e.Dependencies {
				depUrn := dependentResource.URN
				message += fmt.Sprintf(" * %-15q (%s)\n", depUrn.Name(), depUrn)
			}

			message += "\nDelete those resources first or pass --target-dependents."
			return errors.New(message)
		case edit.ResourceProtectedError:
			return fmt.Errorf(
				"%s can't be safely deleted because it is protected. "+
					"Re-run this command with --force to force deletion", string(e.Condemned.URN))
		default:
			return err
		}
	}
	fmt.Println("Resource deleted")
	return nil
}

func newStateDeleteCommand(ws pkgWorkspace.Context, lm backend.LoginManager) *cobra.Command {
	sdcmd := &stateDeleteCmd{}

	cmd := &cobra.Command{
		Use:   "delete [resource URN]",
		Short: "Deletes a resource from a stack's state",
		Long: `Deletes a resource from a stack's state

This command deletes a resource from a stack's state, as long as it is safe to do so. The resource is specified
by its Pulumi URN. If the URN is omitted, this command will prompt for it.

Resources can't be deleted if other resources depend on it or are parented to it. Protected resources
will not be deleted unless specifically requested using the --force flag.

Make sure that URNs are single-quoted to avoid having characters unexpectedly interpreted by the shell.

To see the list of URNs in a stack, use ` + "`khulnasoft stack --show-urns`" + `.
`,
		Example: "khulnasoft state delete 'urn:khulnasoft:stage::demo::eks:index:Cluster$khulnasoft:providers:kubernetes::eks-provider'",
		Args:    cmdutil.MaximumNArgs(1),
		Run: runCmdFunc(func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			return sdcmd.Run(ctx, args, ws, lm)
		}),
	}

	cmd.PersistentFlags().StringVarP(
		&sdcmd.stack, "stack", "s", "",
		"The name of the stack to operate on. Defaults to the current stack")
	cmd.Flags().BoolVar(&sdcmd.force, "force", false, "Force deletion of protected resources")
	cmd.Flags().BoolVarP(&sdcmd.yes, "yes", "y", false, "Skip confirmation prompts")
	cmd.Flags().BoolVar(&sdcmd.targetDependents, "target-dependents", false, "Delete the URN and all its dependents")
	return cmd
}