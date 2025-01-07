package dotcom

import (
	"context"
	"net/http"

	"github.com/graph-gophers/graphql-go"

	"github.com/khulnasoft/khulnasoft/cmd/frontend/enterprise"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/graphqlbackend"
	"github.com/khulnasoft/khulnasoft/cmd/frontend/internal/dotcom/productsubscription"
	"github.com/khulnasoft/khulnasoft/internal/codeintel"
	"github.com/khulnasoft/khulnasoft/internal/conf/conftypes"
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/dotcom"
	"github.com/khulnasoft/khulnasoft/internal/env"
	"github.com/khulnasoft/khulnasoft/internal/observation"
)

var (
	enableOnlineLicenseChecks = env.MustGetBool("DOTCOM_ENABLE_ONLINE_LICENSE_CHECKS", true,
		"If false, online license checks from instances always return successfully.")
)

// dotcomRootResolver implements the GraphQL types DotcomMutation and DotcomQuery.
type dotcomRootResolver struct {
	productsubscription.ProductSubscriptionLicensingResolver
	productsubscription.CodyGatewayDotcomUserResolver
}

func (d dotcomRootResolver) Dotcom() graphqlbackend.DotcomResolver {
	return d
}

func (d dotcomRootResolver) NodeResolvers() map[string]graphqlbackend.NodeByIDFunc {
	return map[string]graphqlbackend.NodeByIDFunc{
		productsubscription.ProductLicenseIDKind: func(ctx context.Context, id graphql.ID) (graphqlbackend.Node, error) {
			return d.ProductLicenseByID(ctx, id)
		},
		productsubscription.ProductSubscriptionIDKind: func(ctx context.Context, id graphql.ID) (graphqlbackend.Node, error) {
			return d.ProductSubscriptionByID(ctx, id)
		},
	}
}

var _ graphqlbackend.DotcomRootResolver = dotcomRootResolver{}

func Init(
	ctx context.Context,
	observationCtx *observation.Context,
	db database.DB,
	_ codeintel.Services,
	_ conftypes.UnifiedWatchable,
	enterpriseServices *enterprise.Services,
) error {
	// Only enabled on Khulnasoft.com.
	if dotcom.KhulnasoftDotComMode() {
		enterpriseServices.DotcomRootResolver = dotcomRootResolver{
			ProductSubscriptionLicensingResolver: productsubscription.ProductSubscriptionLicensingResolver{
				Logger: observationCtx.Logger.Scoped("productsubscriptions"),
				DB:     db,
			},
			CodyGatewayDotcomUserResolver: productsubscription.CodyGatewayDotcomUserResolver{
				Logger: observationCtx.Logger.Scoped("codygatewayuser"),
				DB:     db,
			},
		}
		enterpriseServices.NewDotcomLicenseCheckHandler = func() http.Handler {
			return productsubscription.NewLicenseCheckHandler(db, enableOnlineLicenseChecks)
		}
	}
	return nil
}
