package bodyclose

import (
	"github.com/timakin/bodyclose/passes/bodyclose"

	"github.com/khulnasoft/khulnasoft/dev/linters/nolint"
)

var Analyzer = nolint.Wrap(bodyclose.Analyzer)
