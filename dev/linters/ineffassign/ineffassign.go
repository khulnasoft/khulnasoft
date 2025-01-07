package ineffassign

import (
	"github.com/gordonklaus/ineffassign/pkg/ineffassign"

	"github.com/khulnasoft/khulnasoft/dev/linters/nolint"
)

var Analyzer = nolint.Wrap(ineffassign.Analyzer)
