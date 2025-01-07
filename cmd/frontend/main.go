// Command frontend is a service that serves the web frontend and API.
package main

import (
	"github.com/khulnasoft/khulnasoft/cmd/frontend/shared"
	"github.com/khulnasoft/khulnasoft/internal/sanitycheck"
	"github.com/khulnasoft/khulnasoft/ui/assets"
)

func main() {
	assets.Init()
	sanitycheck.Pass()
	shared.FrontendMain(nil)
}
