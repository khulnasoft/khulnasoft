package ci

import (
	bk "github.com/khulnasoft/khulnasoft/dev/ci/internal/buildkite"
	"github.com/khulnasoft/khulnasoft/dev/ci/internal/ci/operations"
)

func BazelOperations(buildOpts bk.BuildOptions, opts CoreTestOperationsOptions) []operations.Operation {
	ops := []operations.Operation{bazelPrechecks()}
	ops = append(ops, triggerBackCompatTest(buildOpts), bazelGoModTidy())
	return ops
}
