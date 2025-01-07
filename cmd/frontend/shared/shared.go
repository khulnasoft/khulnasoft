// Package shared is the enterprise frontend program's shared main entrypoint.
//
// It lets the invoker of the OSS frontend shared entrypoint injects a few
// proprietary things into it via e.g. blank/underscore imports in this file
// which register side effects with the frontend package.
package shared

import (
	"context"
	"os"
	"strconv"

	"github.com/sourcegraph/log"

	// sourcegraph/internal
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/prompts"
	"github.com/khulnasoft/khulnasoft/internal/codeintel"
	codeintelshared "github.com/khulnasoft/khulnasoft/internal/codeintel/shared"
	"github.com/khulnasoft/khulnasoft/internal/conf"
	"github.com/khulnasoft/khulnasoft/internal/conf/conftypes"
	"github.com/khulnasoft/khulnasoft/internal/database"
	connections "github.com/khulnasoft/khulnasoft/internal/database/connections/live"
	"github.com/khulnasoft/khulnasoft/internal/httpcli"
	"github.com/khulnasoft/khulnasoft/internal/observation"
	"github.com/khulnasoft/khulnasoft/schema"

	// sourcegraph/cmd/frontend/internal
	"github.com/khulnasoft/khulnasoft/cmd/frontend/enterprise"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/auth"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/authz"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/batches"
	codeintelinit "github.com/khulnasoft/khulnasoft/cmd/frontend/internal/codeintel"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/codemonitors"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/completions"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/compute"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/contentlibrary"
	internalcontext "github.com/khulnasoft/khulnasoft/cmd/frontend/internal/context"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/dotcom"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/embeddings"
	executor "github.com/khulnasoft/khulnasoft/cmd/frontend/internal/executorqueue"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/githubapp"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/guardrails"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/insights"
	licensing "github.com/khulnasoft/khulnasoft/cmd/frontend/internal/licensing/init"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/modelconfig"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/notebooks"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/own"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/rbac"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/repos/webhooks"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/savedsearches"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/scim"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/search"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/searchcontexts"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/telemetry"
)

type EnterpriseInitializer = func(context.Context, *observation.Context, database.DB, codeintel.Services, conftypes.UnifiedWatchable, *enterprise.Services) error

var initFunctions = map[string]EnterpriseInitializer{
	"authz":          authz.Init,
	"batches":        batches.Init,
	"codeintel":      codeintelinit.Init,
	"codemonitors":   codemonitors.Init,
	"completions":    completions.Init,
	"compute":        compute.Init,
	"dotcom":         dotcom.Init,
	"embeddings":     embeddings.Init,
	"context":        internalcontext.Init,
	"githubapp":      githubapp.Init,
	"guardrails":     guardrails.Init,
	"insights":       insights.Init,
	"licensing":      licensing.Init,
	"modelconfig":    modelconfig.Init,
	"notebooks":      notebooks.Init,
	"own":            own.Init,
	"rbac":           rbac.Init,
	"repos.webhooks": webhooks.Init,
	"scim":           scim.Init,
	"searchcontexts": searchcontexts.Init,
	"savedsearches":  savedsearches.Init,
	"contentLibrary": contentlibrary.Init,
	"search":         search.Init,
	"telemetry":      telemetry.Init,
	"prompts":        prompts.Init,
}

func EnterpriseSetupHook(db database.DB, conf conftypes.UnifiedWatchable) enterprise.Services {
	logger := log.Scoped("enterprise")
	debug, _ := strconv.ParseBool(os.Getenv("DEBUG"))
	if debug {
		logger.Debug("enterprise edition")
	}

	auth.Init(logger, db)

	ctx := context.Background()
	enterpriseServices := enterprise.DefaultServices()

	observationCtx := observation.NewContext(logger)

	codeIntelServices, err := codeintel.NewServices(codeintel.ServiceDependencies{
		DB:             db,
		CodeIntelDB:    mustInitializeCodeIntelDB(logger),
		ObservationCtx: observationCtx,
	})
	if err != nil {
		logger.Fatal("failed to initialize code intelligence", log.Error(err))
	}

	for name, fn := range initFunctions {
		if err := fn(ctx, observationCtx, db, codeIntelServices, conf, &enterpriseServices); err != nil {
			logger.Fatal("failed to initialize", log.String("name", name), log.Error(err))
		}
	}

	// Inititalize executor last, as we require code intel and batch changes services to be
	// already populated on the enterpriseServices object.
	if err := executor.Init(observationCtx, db, conf, &enterpriseServices); err != nil {
		logger.Fatal("failed to initialize executor", log.Error(err))
	}

	return enterpriseServices
}

func mustInitializeCodeIntelDB(logger log.Logger) codeintelshared.CodeIntelDB {
	dsn := conf.GetServiceConnectionValueAndRestartOnChange(func(serviceConnections conftypes.ServiceConnections) string {
		return serviceConnections.CodeIntelPostgresDSN
	})

	db, err := connections.EnsureNewCodeIntelDB(observation.NewContext(logger), dsn, "frontend")
	if err != nil {
		logger.Fatal("Failed to connect to codeintel database", log.Error(err))
	}

	return codeintelshared.NewCodeIntelDB(logger, db)
}

func switchableSiteConfig() conftypes.WatchableSiteConfig {
	confClient := conf.DefaultClient()
	switchable := &switchingSiteConfig{
		watchers:            make([]func(), 0),
		WatchableSiteConfig: &noopSiteConfig{},
	}
	switchable.WatchableSiteConfig.(*noopSiteConfig).switcher = switchable

	go func() {
		<-AutoUpgradeDone
		httpcli.Configure(confClient)
		switchable.WatchableSiteConfig = confClient
		for _, watcher := range switchable.watchers {
			confClient.Watch(watcher)
		}
		switchable.watchers = nil
	}()

	return switchable
}

type switchingSiteConfig struct {
	watchers []func()
	conftypes.WatchableSiteConfig
}

type noopSiteConfig struct {
	switcher *switchingSiteConfig
}

func (n *noopSiteConfig) SiteConfig() schema.SiteConfiguration {
	return schema.SiteConfiguration{}
}

func (n *noopSiteConfig) Watch(f func()) {
	f()
	n.switcher.watchers = append(n.switcher.watchers, f)
}
