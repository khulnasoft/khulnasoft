package main

import (
	"github.com/khulnasoft/khulnasoft/internal/sanitycheck"
	"github.com/khulnasoft/khulnasoft/internal/service"

	frontend_shared "github.com/khulnasoft/khulnasoft/cmd/frontend/shared"
	gitserver_shared "github.com/khulnasoft/khulnasoft/cmd/gitserver/shared"
	repoupdater_shared "github.com/khulnasoft/khulnasoft/cmd/repo-updater/shared"
	searcher_shared "github.com/khulnasoft/khulnasoft/cmd/searcher/shared"
	worker_shared "github.com/khulnasoft/khulnasoft/cmd/worker/shared"
)

func main() {
	sanitycheck.Pass()

	// Other services to run (in addition to `frontend`).
	otherServices := []service.Service{
		gitserver_shared.Service,
		repoupdater_shared.Service,
		searcher_shared.Service,
		worker_shared.Service,
	}

	frontend_shared.FrontendMain(otherServices)
}
