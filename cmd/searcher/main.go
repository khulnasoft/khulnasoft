// Command searcher is a simple service which exposes an API to text search a
// repo at a specific commit. See the searcher package for more information.
package main

import (
	"github.com/khulnasoft/khulnasoft/cmd/searcher/shared"
	"github.com/khulnasoft/khulnasoft/internal/sanitycheck"
	"github.com/khulnasoft/khulnasoft/internal/service/svcmain"
)

func main() {
	sanitycheck.Pass()
	svcmain.SingleServiceMain(shared.Service)
}
