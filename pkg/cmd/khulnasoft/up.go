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
	"context"
	"errors"
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"

	"github.com/khulnasoft/khulnasoft/pkg/v3/backend"
	"github.com/khulnasoft/khulnasoft/pkg/v3/backend/display"
	"github.com/khulnasoft/khulnasoft/pkg/v3/engine"
	"github.com/khulnasoft/khulnasoft/pkg/v3/resource/deploy"
	"github.com/khulnasoft/khulnasoft/pkg/v3/resource/stack"
	pkgWorkspace "github.com/khulnasoft/khulnasoft/pkg/v3/workspace"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/common/apitype"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/common/env"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/common/resource/config"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/common/tokens"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/common/util/cmdutil"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/common/util/contract"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/common/util/logging"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/common/workspace"
)

// The default number of parallel resource operations to run at once during an update, if --parallel is unset.
// See https://github.com/khulnasoft/khulnasoft/issues/14989 for context around the cpu * 4 choice.
var defaultParallel = int32(runtime.NumCPU()) * 4 //nolint:gosec // NumCPU is an int32 internally,
//                                                                  but the NumCPU function returns an int.

// intentionally disabling here for cleaner err declaration/assignment.
//
//nolint:vetshadow
func newUpCmd() *cobra.Command {
	var debug bool
	var expectNop bool
	var message string
	var execKind string
	var execAgent string
	var stackName string
	var configArray []string
	var path bool
	var client string

	// Flags for remote operations.
	remoteArgs := RemoteArgs{}

	// Flags for engine.UpdateOptions.
	var jsonDisplay bool
	var policyPackPaths []string
	var policyPackConfigPaths []string
	var diffDisplay bool
	var eventLogPath string
	var parallel int32
	var refresh string
	var showConfig bool
	var showPolicyRemediations bool
	var showReplacementSteps bool
	var showSames bool
	var showReads bool
	var skipPreview bool
	var showFullOutput bool
	var suppressOutputs bool
	var suppressProgress bool
	var continueOnError bool
	var suppressPermalink string
	var yes bool
	var secretsProvider string
	var targets []string
	var replaces []string
	var targetReplaces []string
	var targetDependents bool
	var planFilePath string
	var attachDebugger bool

	// up implementation used when the source of the Pulumi program is in the current working directory.
	upWorkingDirectory := func(
		ctx context.Context,
		ssml stackSecretsManagerLoader,
		ws pkgWorkspace.Context,
		lm backend.LoginManager,
		opts backend.UpdateOptions,
		cmd *cobra.Command,
	) error {
		s, err := requireStack(ctx, ws, lm, stackName, stackOfferNew, opts.Display)
		if err != nil {
			return err
		}

		// Save any config values passed via flags.
		if err := parseAndSaveConfigArray(ws, s, configArray, path); err != nil {
			return err
		}

		proj, root, err := readProjectForUpdate(ws, client)
		if err != nil {
			return err
		}

		cfg, sm, err := getStackConfiguration(ctx, ssml, s, proj)
		if err != nil {
			return fmt.Errorf("getting stack configuration: %w", err)
		}

		m, err := getUpdateMetadata(message, root, execKind, execAgent, planFilePath != "", cfg, cmd.Flags())
		if err != nil {
			return fmt.Errorf("gathering environment metadata: %w", err)
		}

		decrypter, err := sm.Decrypter()
		if err != nil {
			return fmt.Errorf("getting stack decrypter: %w", err)
		}
		encrypter, err := sm.Encrypter()
		if err != nil {
			return fmt.Errorf("getting stack encrypter: %w", err)
		}

		stackName := s.Ref().Name().String()
		configErr := workspace.ValidateStackConfigAndApplyProjectConfig(
			ctx,
			stackName,
			proj,
			cfg.Environment,
			cfg.Config,
			encrypter,
			decrypter)
		if configErr != nil {
			return fmt.Errorf("validating stack config: %w", configErr)
		}

		targetURNs, replaceURNs := []string{}, []string{}
		targetURNs = append(targetURNs, targets...)
		replaceURNs = append(replaceURNs, replaces...)

		for _, tr := range targetReplaces {
			targetURNs = append(targetURNs, tr)
			replaceURNs = append(replaceURNs, tr)
		}

		refreshOption, err := getRefreshOption(proj, refresh)
		if err != nil {
			return err
		}
		opts.Engine = engine.UpdateOptions{
			LocalPolicyPacks:          engine.MakeLocalPolicyPacks(policyPackPaths, policyPackConfigPaths),
			Parallel:                  parallel,
			Debug:                     debug,
			Refresh:                   refreshOption,
			ReplaceTargets:            deploy.NewUrnTargets(replaceURNs),
			UseLegacyDiff:             useLegacyDiff(),
			UseLegacyRefreshDiff:      useLegacyRefreshDiff(),
			DisableProviderPreview:    disableProviderPreview(),
			DisableResourceReferences: disableResourceReferences(),
			DisableOutputValues:       disableOutputValues(),
			Targets:                   deploy.NewUrnTargets(targetURNs),
			TargetDependents:          targetDependents,
			// Trigger a plan to be generated during the preview phase which can be constrained to during the
			// update phase.
			GeneratePlan:    true,
			Experimental:    hasExperimentalCommands(),
			ContinueOnError: continueOnError,
			AttachDebugger:  attachDebugger,
		}

		if planFilePath != "" {
			dec, err := sm.Decrypter()
			if err != nil {
				return err
			}
			enc, err := sm.Encrypter()
			if err != nil {
				return err
			}
			plan, err := readPlan(planFilePath, dec, enc)
			if err != nil {
				return err
			}
			opts.Engine.Plan = plan
		}

		changes, err := s.Update(ctx, backend.UpdateOperation{
			Proj:               proj,
			Root:               root,
			M:                  m,
			Opts:               opts,
			StackConfiguration: cfg,
			SecretsManager:     sm,
			SecretsProvider:    stack.DefaultSecretsProvider,
			Scopes:             backend.CancellationScopes,
		})
		switch {
		case err == context.Canceled:
			return errors.New("update cancelled")
		case err != nil:
			return err
		case expectNop && changes != nil && engine.HasChanges(changes):
			return errors.New("no changes were expected but changes occurred")
		default:
			return nil
		}
	}

	// up implementation used when the source of the Pulumi program is a template name or a URL to a template.
	upTemplateNameOrURL := func(
		ctx context.Context,
		ssml stackSecretsManagerLoader,
		ws pkgWorkspace.Context,
		lm backend.LoginManager,
		templateNameOrURL string,
		opts backend.UpdateOptions,
		cmd *cobra.Command,
	) error {
		// Retrieve the template repo.
		repo, err := workspace.RetrieveTemplates(ctx, templateNameOrURL, false, workspace.TemplateKindPulumiProject)
		if err != nil {
			return err
		}
		defer func() {
			contract.IgnoreError(repo.Delete())
		}()

		// List the templates from the repo.
		templates, err := repo.Templates()
		if err != nil {
			return err
		}

		var template workspace.Template
		if len(templates) == 0 {
			return errors.New("no template found")
		} else if len(templates) == 1 {
			template = templates[0]
		} else {
			if template, err = chooseTemplate(templates, opts.Display); err != nil {
				return err
			}
		}

		// Validate secrets provider type
		if err := validateSecretsProvider(secretsProvider); err != nil {
			return err
		}

		// Create temp directory for the "virtual workspace".
		temp, err := os.MkdirTemp("", "khulnasoft-up-")
		if err != nil {
			return err
		}
		defer func() {
			contract.IgnoreError(os.RemoveAll(temp))
		}()

		// Change the working directory to the "virtual workspace" directory.
		if err = os.Chdir(temp); err != nil {
			return fmt.Errorf("changing the working directory: %w", err)
		}

		// There is no current project at this point to pass into currentBackend
		b, err := currentBackend(ctx, ws, lm, nil, opts.Display)
		if err != nil {
			return err
		}

		// If a stack was specified via --stack, see if it already exists.
		var name string
		var description string
		var s backend.Stack
		if stackName != "" {
			if s, name, description, err = getStack(ctx, b, stackName, opts.Display); err != nil {
				return err
			}
		}

		// Prompt for the project name, if we don't already have one from an existing stack.
		if name == "" {
			defaultValue := pkgWorkspace.ValueOrSanitizedDefaultProjectName(name, template.ProjectName, template.Name)
			name, err = promptForValue(
				yes, "project name", defaultValue, false, pkgWorkspace.ValidateProjectName, opts.Display)
			if err != nil {
				return err
			}
		}

		// Prompt for the project description, if we don't already have one from an existing stack.
		if description == "" {
			defaultValue := pkgWorkspace.ValueOrDefaultProjectDescription(
				description, template.ProjectDescription, template.Description)
			description, err = promptForValue(
				yes, "project description", defaultValue, false, pkgWorkspace.ValidateProjectDescription, opts.Display)
			if err != nil {
				return err
			}
		}

		// Copy the template files from the repo to the temporary "virtual workspace" directory.
		if err = workspace.CopyTemplateFiles(template.Dir, temp, true, name, description); err != nil {
			return err
		}

		// Load the project, update the name & description, remove the template section, and save it.
		proj, root, err := ws.ReadProject()
		if err != nil {
			return err
		}
		proj.Name = tokens.PackageName(name)
		proj.Description = &description
		proj.Template = nil
		if err = workspace.SaveProject(proj); err != nil {
			return fmt.Errorf("saving project: %w", err)
		}

		// Create the stack, if needed.
		if s == nil {
			if s, err = promptAndCreateStack(ctx, ws, b, promptForValue, stackName, root, false /*setCurrent*/, yes,
				opts.Display, secretsProvider); err != nil {
				return err
			}
			// The backend will print "Created stack '<stack>'." on success.
		}

		// Prompt for config values (if needed) and save.
		if err = handleConfig(
			ctx,
			ssml,
			ws,
			promptForValue,
			proj,
			s,
			templateNameOrURL,
			template,
			configArray,
			yes,
			path,
			opts.Display,
		); err != nil {
			return err
		}

		// Install dependencies.

		projinfo := &engine.Projinfo{Proj: proj, Root: root}
		_, main, pctx, err := engine.ProjectInfoContext(projinfo, nil, cmdutil.Diag(), cmdutil.Diag(), nil, false, nil, nil)
		if err != nil {
			return fmt.Errorf("building project context: %w", err)
		}

		defer pctx.Close()

		if err = installDependencies(pctx, &proj.Runtime, main); err != nil {
			return err
		}

		cfg, sm, err := getStackConfiguration(ctx, ssml, s, proj)
		if err != nil {
			return fmt.Errorf("getting stack configuration: %w", err)
		}

		m, err := getUpdateMetadata(message, root, execKind, execAgent, planFilePath != "", cfg, cmd.Flags())
		if err != nil {
			return fmt.Errorf("gathering environment metadata: %w", err)
		}

		decrypter, err := sm.Decrypter()
		if err != nil {
			return fmt.Errorf("getting stack decrypter: %w", err)
		}
		encrypter, err := sm.Encrypter()
		if err != nil {
			return fmt.Errorf("getting stack encrypter: %w", err)
		}

		stackName := s.Ref().String()
		configErr := workspace.ValidateStackConfigAndApplyProjectConfig(
			ctx,
			stackName,
			proj,
			cfg.Environment,
			cfg.Config,
			encrypter,
			decrypter)
		if configErr != nil {
			return fmt.Errorf("validating stack config: %w", configErr)
		}

		refreshOption, err := getRefreshOption(proj, refresh)
		if err != nil {
			return err
		}

		opts.Engine = engine.UpdateOptions{
			LocalPolicyPacks: engine.MakeLocalPolicyPacks(policyPackPaths, policyPackConfigPaths),
			Parallel:         parallel,
			Debug:            debug,
			Refresh:          refreshOption,

			// If we're in experimental mode then we trigger a plan to be generated during the preview phase
			// which will be constrained to during the update phase.
			GeneratePlan: hasExperimentalCommands(),
			Experimental: hasExperimentalCommands(),

			UseLegacyRefreshDiff: useLegacyRefreshDiff(),
			ContinueOnError:      continueOnError,

			AttachDebugger: attachDebugger,
		}

		// TODO for the URL case:
		// - suppress preview display/prompt unless error.
		// - attempt `destroy` on any update errors.
		// - show template.Quickstart?

		changes, err := s.Update(ctx, backend.UpdateOperation{
			Proj:               proj,
			Root:               root,
			M:                  m,
			Opts:               opts,
			StackConfiguration: cfg,
			SecretsManager:     sm,
			SecretsProvider:    stack.DefaultSecretsProvider,
			Scopes:             backend.CancellationScopes,
		})
		switch {
		case err == context.Canceled:
			return errors.New("update cancelled")
		case err != nil:
			return err
		case expectNop && changes != nil && engine.HasChanges(changes):
			return errors.New("no changes were expected but changes occurred")
		default:
			return nil
		}
	}

	cmd := &cobra.Command{
		Use:        "up [template|url]",
		Aliases:    []string{"update"},
		SuggestFor: []string{"apply", "deploy", "push"},
		Short:      "Create or update the resources in a stack",
		Long: "Create or update the resources in a stack.\n" +
			"\n" +
			"This command creates or updates resources in a stack. The new desired goal state for the target stack\n" +
			"is computed by running the current Pulumi program and observing all resource allocations to produce a\n" +
			"resource graph. This goal state is then compared against the existing state to determine what create,\n" +
			"read, update, and/or delete operations must take place to achieve the desired goal state, in the most\n" +
			"minimally disruptive way. This command records a full transactional snapshot of the stack's new state\n" +
			"afterwards so that the stack may be updated incrementally again later on.\n" +
			"\n" +
			"The program to run is loaded from the project in the current directory by default. Use the `-C` or\n" +
			"`--cwd` flag to use a different directory.",
		Args: cmdutil.MaximumNArgs(1),
		Run: runCmdFunc(func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			ssml := newStackSecretsManagerLoaderFromEnv()
			ws := pkgWorkspace.Instance

			// Remote implies we're skipping previews.
			if remoteArgs.remote {
				skipPreview = true
			}

			yes = yes || skipPreview || skipConfirmations()

			interactive := cmdutil.Interactive()
			if !interactive && !yes {
				return errors.New(
					"--yes or --skip-preview must be passed in to proceed when running in non-interactive mode",
				)
			}

			opts, err := updateFlagsToOptions(interactive, skipPreview, yes, false /* previewOnly */)
			if err != nil {
				return err
			}

			if err = validatePolicyPackConfig(policyPackPaths, policyPackConfigPaths); err != nil {
				return err
			}

			displayType := display.DisplayProgress
			if diffDisplay {
				displayType = display.DisplayDiff
			}

			opts.Display = display.Options{
				Color:                  cmdutil.GetGlobalColorization(),
				ShowConfig:             showConfig,
				ShowPolicyRemediations: showPolicyRemediations,
				ShowReplacementSteps:   showReplacementSteps,
				ShowSameResources:      showSames,
				ShowReads:              showReads,
				SuppressOutputs:        suppressOutputs,
				SuppressProgress:       suppressProgress,
				TruncateOutput:         !showFullOutput,
				IsInteractive:          interactive,
				Type:                   displayType,
				EventLogPath:           eventLogPath,
				Debug:                  debug,
				JSONDisplay:            jsonDisplay,
			}

			// we only suppress permalinks if the user passes true. the default is an empty string
			// which we pass as 'false'
			if suppressPermalink == "true" {
				opts.Display.SuppressPermalink = true
			} else {
				opts.Display.SuppressPermalink = false
			}

			if remoteArgs.remote {
				err = validateUnsupportedRemoteFlags(expectNop, configArray, path, client, jsonDisplay, policyPackPaths,
					policyPackConfigPaths, refresh, showConfig, showPolicyRemediations, showReplacementSteps, showSames,
					showReads, suppressOutputs, secretsProvider, &targets, replaces, targetReplaces,
					targetDependents, planFilePath, stackConfigFile)
				if err != nil {
					return err
				}

				var url string
				if len(args) > 0 {
					url = args[0]
				}

				if err = validateRemoteDeploymentFlags(url, remoteArgs); err != nil {
					return err
				}

				return runDeployment(ctx, ws, cmd, opts.Display, apitype.Update, stackName, url, remoteArgs)
			}

			isDIYBackend, err := isDIYBackend(ws, opts.Display)
			if err != nil {
				return err
			}

			// by default, we are going to suppress the permalink when using DIY backends
			// this can be re-enabled by explicitly passing "false" to the `suppress-permalink` flag
			if suppressPermalink != "false" && isDIYBackend {
				opts.Display.SuppressPermalink = true
			}

			// Link to Copilot will be shown for orgs that have Copilot enabled, unless the user explicitly suppressed it.
			logging.V(7).Infof("PULUMI_SUPPRESS_COPILOT_LINK=%v", env.SuppressCopilotLink.Value())
			opts.Display.ShowLinkToCopilot = !env.SuppressCopilotLink.Value()

			if len(args) > 0 {
				return upTemplateNameOrURL(
					ctx,
					ssml,
					ws,
					DefaultLoginManager,
					args[0],
					opts,
					cmd,
				)
			}

			return upWorkingDirectory(
				ctx,
				ssml,
				ws,
				DefaultLoginManager,
				opts,
				cmd,
			)
		}),
	}

	cmd.PersistentFlags().BoolVarP(
		&debug, "debug", "d", false,
		"Print detailed debugging output during resource operations")
	cmd.PersistentFlags().BoolVar(
		&expectNop, "expect-no-changes", false,
		"Return an error if any changes occur during this update")
	cmd.PersistentFlags().StringVarP(
		&stackName, "stack", "s", "",
		"The name of the stack to operate on. Defaults to the current stack")
	cmd.PersistentFlags().StringVar(
		&stackConfigFile, "config-file", "",
		"Use the configuration values in the specified file rather than detecting the file name")
	cmd.PersistentFlags().StringArrayVarP(
		&configArray, "config", "c", []string{},
		"Config to use during the update and save to the stack config file")
	cmd.PersistentFlags().BoolVar(
		&path, "config-path", false,
		"Config keys contain a path to a property in a map or list to set")
	cmd.PersistentFlags().StringVar(
		&secretsProvider, "secrets-provider", "default", "The type of the provider that should be used to encrypt and "+
			"decrypt secrets (possible choices: default, passphrase, awskms, azurekeyvault, gcpkms, hashivault). Only "+
			"used when creating a new stack from an existing template")

	cmd.PersistentFlags().StringVar(
		&client, "client", "", "The address of an existing language runtime host to connect to")
	_ = cmd.PersistentFlags().MarkHidden("client")

	cmd.PersistentFlags().StringVarP(
		&message, "message", "m", "",
		"Optional message to associate with the update operation")

	cmd.PersistentFlags().StringArrayVarP(
		&targets, "target", "t", []string{},
		"Specify a single resource URN to update. Other resources will not be updated."+
			" Multiple resources can be specified using --target urn1 --target urn2."+
			" Wildcards (*, **) are also supported")
	cmd.PersistentFlags().StringArrayVar(
		&replaces, "replace", []string{},
		"Specify a single resource URN to replace. Multiple resources can be specified using --replace urn1 --replace urn2."+
			" Wildcards (*, **) are also supported")
	cmd.PersistentFlags().StringArrayVar(
		&targetReplaces, "target-replace", []string{},
		"Specify a single resource URN to replace. Other resources will not be updated."+
			" Shorthand for --target urn --replace urn.")
	cmd.PersistentFlags().BoolVar(
		&targetDependents, "target-dependents", false,
		"Allows updating of dependent targets discovered but not specified in --target list")

	// Flags for engine.UpdateOptions.
	cmd.PersistentFlags().StringSliceVar(
		&policyPackPaths, "policy-pack", []string{},
		"Run one or more policy packs as part of this update")
	cmd.PersistentFlags().StringSliceVar(
		&policyPackConfigPaths, "policy-pack-config", []string{},
		`Path to JSON file containing the config for the policy pack of the corresponding "--policy-pack" flag`)
	cmd.PersistentFlags().BoolVar(
		&diffDisplay, "diff", false,
		"Display operation as a rich diff showing the overall change")
	cmd.Flags().BoolVarP(
		&jsonDisplay, "json", "j", false,
		"Serialize the update diffs, operations, and overall output as JSON")
	cmd.PersistentFlags().Int32VarP(
		&parallel, "parallel", "p", defaultParallel,
		"Allow P resource operations to run in parallel at once (1 for no parallelism).")
	cmd.PersistentFlags().StringVarP(
		&refresh, "refresh", "r", "",
		"Refresh the state of the stack's resources before this update")
	cmd.PersistentFlags().Lookup("refresh").NoOptDefVal = "true"
	cmd.PersistentFlags().BoolVar(
		&showConfig, "show-config", false,
		"Show configuration keys and variables")
	cmd.PersistentFlags().BoolVar(
		&showPolicyRemediations, "show-policy-remediations", false,
		"Show per-resource policy remediation details instead of a summary")
	cmd.PersistentFlags().BoolVar(
		&showReplacementSteps, "show-replacement-steps", false,
		"Show detailed resource replacement creates and deletes instead of a single step")

	cmd.PersistentFlags().BoolVar(
		&showSames, "show-sames", false,
		"Show resources that don't need be updated because they haven't changed, alongside those that do")
	cmd.PersistentFlags().BoolVar(
		&showReads, "show-reads", false,
		"Show resources that are being read in, alongside those being managed directly in the stack")

	cmd.PersistentFlags().BoolVarP(
		&skipPreview, "skip-preview", "f", false,
		"Do not calculate a preview before performing the update")
	cmd.PersistentFlags().BoolVar(
		&suppressOutputs, "suppress-outputs", false,
		"Suppress display of stack outputs (in case they contain sensitive values)")
	cmd.PersistentFlags().BoolVar(
		&suppressProgress, "suppress-progress", false,
		"Suppress display of periodic progress dots")
	cmd.PersistentFlags().BoolVar(
		&showFullOutput, "show-full-output", true,
		"Display full length of stack outputs")
	cmd.PersistentFlags().StringVar(
		&suppressPermalink, "suppress-permalink", "",
		"Suppress display of the state permalink")
	cmd.Flag("suppress-permalink").NoOptDefVal = "false"
	cmd.PersistentFlags().BoolVarP(
		&yes, "yes", "y", false,
		"Automatically approve and perform the update after previewing it")
	cmd.PersistentFlags().BoolVar(
		&continueOnError, "continue-on-error", env.ContinueOnError.Value(),
		"Continue updating resources even if an error is encountered "+
			"(can also be set with PULUMI_CONTINUE_ON_ERROR environment variable)")
	cmd.PersistentFlags().BoolVar(
		&attachDebugger, "attach-debugger", false,
		"Enable the ability to attach a debugger to the program being executed")

	cmd.PersistentFlags().StringVar(
		&planFilePath, "plan", "",
		"[EXPERIMENTAL] Path to a plan file to use for the update. The update will not "+
			"perform operations that exceed its plan (e.g. replacements instead of updates, or updates instead"+
			"of sames).")
	if !hasExperimentalCommands() {
		contract.AssertNoErrorf(cmd.PersistentFlags().MarkHidden("plan"), `Could not mark "plan" as hidden`)
	}

	// Remote flags
	remoteArgs.applyFlags(cmd)

	if hasDebugCommands() {
		cmd.PersistentFlags().StringVar(
			&eventLogPath, "event-log", "",
			"Log events to a file at this path")
	}

	// internal flags
	cmd.PersistentFlags().StringVar(&execKind, "exec-kind", "", "")
	// ignore err, only happens if flag does not exist
	_ = cmd.PersistentFlags().MarkHidden("exec-kind")
	cmd.PersistentFlags().StringVar(&execAgent, "exec-agent", "", "")
	// ignore err, only happens if flag does not exist
	_ = cmd.PersistentFlags().MarkHidden("exec-agent")

	return cmd
}

// validatePolicyPackConfig validates the `--policy-pack-config` and `--policy-pack` flags. These two flags are
// order-dependent, e.g., the first `--policy-pack-config` flag value corresponds to the first `--policy-pack`
// flag value, and so on for the second, third, etc. An error is returned if `--policy-pack-config` is specified
// and there isn't a `--policy-pack-config` for every `--policy-pack` that was set.
func validatePolicyPackConfig(policyPackPaths []string, policyPackConfigPaths []string) error {
	if len(policyPackConfigPaths) > 0 {
		if len(policyPackPaths) == 0 {
			return errors.New(`"--policy-pack-config" must be specified with "--policy-pack"`)
		}
		if len(policyPackConfigPaths) != len(policyPackPaths) {
			return errors.New(
				`the number of "--policy-pack-config" flags must match the number of "--policy-pack" flags`)
		}
	}
	return nil
}

// handleConfig handles prompting for config values (as needed) and saving config.
func handleConfig(
	ctx context.Context,
	ssml stackSecretsManagerLoader,
	ws pkgWorkspace.Context,
	prompt promptForValueFunc,
	project *workspace.Project,
	s backend.Stack,
	templateNameOrURL string,
	template workspace.Template,
	configArray []string,
	yes bool,
	path bool,
	opts display.Options,
) error {
	// Get the existing config. stackConfig will be nil if there wasn't a previous deployment.
	stackConfig, err := backend.GetLatestConfiguration(ctx, s)
	if err != nil && err != backend.ErrNoPreviousDeployment {
		return err
	}

	// Get the existing snapshot.
	snap, err := s.Snapshot(ctx, stack.DefaultSecretsProvider)
	if err != nil {
		return err
	}

	// Handle config.
	// If this is an initial preconfigured empty stack (i.e. configured in the Pulumi Console),
	// use its config without prompting.
	// Otherwise, use the values specified on the command line and prompt for new values.
	// If the stack already existed and had previous config, those values will be used as the defaults.
	var c config.Map
	if isPreconfiguredEmptyStack(templateNameOrURL, template.Config, stackConfig, snap) {
		c = stackConfig
		// TODO[khulnasoft/khulnasoft#1894] consider warning if templateNameOrURL is different from
		// the stack's `khulnasoft:template` config value.
	} else {
		// Get config values passed on the command line.
		commandLineConfig, parseErr := parseConfig(configArray, path)
		if parseErr != nil {
			return parseErr
		}

		// Prompt for config as needed.
		c, err = promptForConfig(
			ctx,
			ssml,
			prompt,
			project,
			s,
			template.Config,
			commandLineConfig,
			stackConfig,
			yes,
			opts,
		)
		if err != nil {
			return err
		}
	}

	// Save the config.
	if len(c) > 0 {
		if err = saveConfig(ws, s, c); err != nil {
			return fmt.Errorf("saving config: %w", err)
		}

		fmt.Println("Saved config")
		fmt.Println()
	}

	return nil
}

var templateKey = config.MustMakeKey("khulnasoft", "template")

// isPreconfiguredEmptyStack returns true if the url matches the value of `khulnasoft:template` in stackConfig,
// the stackConfig values satisfy the config requirements of templateConfig, and the snapshot is empty.
// This is the state of an initial preconfigured empty stack (i.e. a stack that's been created and configured
// in the Pulumi Console).
func isPreconfiguredEmptyStack(
	url string,
	templateConfig map[string]workspace.ProjectTemplateConfigValue,
	stackConfig config.Map,
	snap *deploy.Snapshot,
) bool {
	// Does stackConfig have a `khulnasoft:template` value and does it match url?
	if stackConfig == nil {
		return false
	}
	templateURLValue, hasTemplateKey := stackConfig[templateKey]
	if !hasTemplateKey {
		return false
	}
	templateURL, err := templateURLValue.Value(nil)
	if err != nil {
		contract.IgnoreError(err)
		return false
	}
	if templateURL != url {
		return false
	}

	// Does the snapshot only contain a single root resource?
	if len(snap.Resources) != 1 {
		return false
	}
	stackResource, err := stack.GetRootStackResource(snap)
	if err != nil || stackResource == nil {
		return false
	}

	// Can stackConfig satisfy the config requirements of templateConfig?
	for templateKey, templateVal := range templateConfig {
		parsedTemplateKey, parseErr := parseConfigKey(templateKey)
		if parseErr != nil {
			contract.IgnoreError(parseErr)
			return false
		}

		stackVal, ok := stackConfig[parsedTemplateKey]
		if !ok {
			return false
		}

		if templateVal.Secret != stackVal.Secure() {
			return false
		}
	}

	return true
}