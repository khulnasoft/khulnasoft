package tables

import (
	"gorm.io/gorm/schema"

	"github.com/khulnasoft/khulnasoft/cmd/enterprise-portal/internal/database/codyaccess"
	"github.com/khulnasoft/khulnasoft/cmd/enterprise-portal/internal/database/subscriptions"
)

// All tables provisioned for the Enterprise Portal database are defined here.
//
// ⚠️ WARNING: This list is meant to be read-only.
func All() []schema.Tabler {
	return []schema.Tabler{
		&subscriptions.TableSubscription{},
		&subscriptions.SubscriptionCondition{},
		&subscriptions.TableSubscriptionLicense{},
		&subscriptions.SubscriptionLicenseCondition{},

		&codyaccess.TableCodyGatewayAccess{},
	}
}
