// blobstore is the blobstore server.
package main // import "github.com/khulnasoft/khulnasoft/cmd/blobstore"

import (
	"github.com/khulnasoft/khulnasoft/cmd/blobstore/shared"
	"github.com/khulnasoft/khulnasoft/internal/sanitycheck"
	"github.com/khulnasoft/khulnasoft/internal/service/svcmain"
)

func main() {
	sanitycheck.Pass()
	svcmain.SingleServiceMain(shared.Service)
}
