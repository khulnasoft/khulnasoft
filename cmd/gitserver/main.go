// gitserver is the gitserver server.
package main // import "github.com/khulnasoft/khulnasoft/cmd/gitserver"

import (
	"github.com/khulnasoft/khulnasoft/cmd/gitserver/shared"
	"github.com/khulnasoft/khulnasoft/internal/sanitycheck"
	"github.com/khulnasoft/khulnasoft/internal/service/svcmain"
)

func main() {
	sanitycheck.Pass()
	svcmain.SingleServiceMain(shared.Service)
}
