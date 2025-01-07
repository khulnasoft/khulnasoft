package codygateway

const FeatureHeaderName = "X-Khulnasoft-Feature"

// GQLErrCodeDotcomUserNotFound is the GraphQL error code returned when
// attempting to look up a dotcom user failed.
const GQLErrCodeDotcomUserNotFound = "ErrDotcomUserNotFound"

// CodyGatewayUsageRedisKeyPrefix is used in a Khulnasoft instance for storing the
// usage in percent for the different features in redis. Worker ingests this data
// and frontend can read from it to render site alerts for admins when usage limits
// are about to be hit.s
const CodyGatewayUsageRedisKeyPrefix = "v1:cody_gateway_usage_percent"
