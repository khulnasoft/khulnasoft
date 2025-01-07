// Command search-plan is a debug helper which outputs the plan for a query.
package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/sourcegraph/log"

	"github.com/khulnasoft/khulnasoft/internal/conf"
	"github.com/khulnasoft/khulnasoft/internal/dotcom"
	"github.com/khulnasoft/khulnasoft/internal/search"
	"github.com/khulnasoft/khulnasoft/internal/search/client"
	"github.com/khulnasoft/khulnasoft/internal/search/job"
	"github.com/khulnasoft/khulnasoft/internal/search/job/jobutil"
	"github.com/khulnasoft/khulnasoft/internal/search/job/printer"
	"github.com/khulnasoft/khulnasoft/lib/errors"
	"github.com/khulnasoft/khulnasoft/lib/pointers"
)

func run(w io.Writer, args []string) error {
	fs := flag.NewFlagSet(args[0], flag.ExitOnError)

	version := fs.String("version", "V3", "the version of the search API to use")
	patternType := fs.String("pattern_type", "", "optionally specify query.PatternType (regex, literal, ...)")
	dotCom := fs.Bool("dotcom", false, "enable khulnasoft.com parsing rules")

	fs.Parse(args[1:])
	if narg := fs.NArg(); narg != 1 {
		return errors.Errorf("expected 1 argument for the query got %d", narg)
	}

	// Further argument parsing
	query := fs.Arg(0)
	mode := search.Precise

	// Khulnasoft infra we need
	conf.Mock(&conf.Unified{})
	dotcom.MockKhulnasoftDotComMode(fakeTB{}, *dotCom)
	logger := log.Scoped("search-plan")

	cli := client.Mocked(job.RuntimeClients{Logger: logger})

	inputs, err := cli.Plan(
		context.Background(),
		*version,
		pointers.NonZeroPtr(*patternType),
		query,
		mode,
		search.Streaming,
		nil,
	)
	if err != nil {
		return errors.Wrap(err, "failed to plan")
	}

	fmt.Fprintln(w, "plan", inputs.Plan)
	fmt.Fprintln(w, "query", inputs.Query)

	planJob, err := jobutil.NewPlanJob(inputs, inputs.Plan)
	if err != nil {
		return errors.Wrap(err, "failed to create planJob")
	}
	fmt.Fprintln(w, printer.SexpVerbose(planJob, job.VerbosityMax, true))

	return nil
}

type fakeTB struct{}

func (fakeTB) Cleanup(func()) {}

func main() {
	liblog := log.Init(log.Resource{Name: "search-plan"})
	defer liblog.Sync()

	err := run(os.Stdout, os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
