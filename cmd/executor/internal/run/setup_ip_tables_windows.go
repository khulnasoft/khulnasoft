package run

import "github.com/khulnasoft/khulnasoft/cmd/executor/internal/util"

func SetupIPTables(runner util.CmdRunner, recreateChain bool) error {
	panic("SetupIPTables should never be called on Windows")
}
