// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package dotcom

import (
	"context"
	"encoding/json"

	"github.com/Khan/genqlient/graphql"
	"github.com/khulnasoft/khulnasoft/cmd/cody-gateway/internal/dotcom/genhelper"
)

// CheckDotcomUserAccessTokenDotcomDotcomQuery includes the requested fields of the GraphQL type DotcomQuery.
// The GraphQL type's documentation follows.
//
// Mutations that are only used on Khulnasoft.com.
// FOR INTERNAL USE ONLY.
type CheckDotcomUserAccessTokenDotcomDotcomQuery struct {
	// A dotcom user for purposes of connecting to the Cody Gateway.
	// Only Khulnasoft.com site admins or service accounts may perform this query.
	// Token is a Cody Gateway token, not a Khulnasoft instance access token.
	// FOR INTERNAL USE ONLY.
	CodyGatewayDotcomUserByToken *CheckDotcomUserAccessTokenDotcomDotcomQueryCodyGatewayDotcomUserByTokenCodyGatewayDotcomUser `json:"codyGatewayDotcomUserByToken"`
}

// GetCodyGatewayDotcomUserByToken returns CheckDotcomUserAccessTokenDotcomDotcomQuery.CodyGatewayDotcomUserByToken, and is useful for accessing the field via an interface.
func (v *CheckDotcomUserAccessTokenDotcomDotcomQuery) GetCodyGatewayDotcomUserByToken() *CheckDotcomUserAccessTokenDotcomDotcomQueryCodyGatewayDotcomUserByTokenCodyGatewayDotcomUser {
	return v.CodyGatewayDotcomUserByToken
}

// CheckDotcomUserAccessTokenDotcomDotcomQueryCodyGatewayDotcomUserByTokenCodyGatewayDotcomUser includes the requested fields of the GraphQL type CodyGatewayDotcomUser.
// The GraphQL type's documentation follows.
//
// A dotcom user allowed to access the Cody Gateway
// FOR INTERNAL USE ONLY.
type CheckDotcomUserAccessTokenDotcomDotcomQueryCodyGatewayDotcomUserByTokenCodyGatewayDotcomUser struct {
	DotcomUserState `json:"-"`
}

// GetId returns CheckDotcomUserAccessTokenDotcomDotcomQueryCodyGatewayDotcomUserByTokenCodyGatewayDotcomUser.Id, and is useful for accessing the field via an interface.
func (v *CheckDotcomUserAccessTokenDotcomDotcomQueryCodyGatewayDotcomUserByTokenCodyGatewayDotcomUser) GetId() string {
	return v.DotcomUserState.Id
}

// GetUsername returns CheckDotcomUserAccessTokenDotcomDotcomQueryCodyGatewayDotcomUserByTokenCodyGatewayDotcomUser.Username, and is useful for accessing the field via an interface.
func (v *CheckDotcomUserAccessTokenDotcomDotcomQueryCodyGatewayDotcomUserByTokenCodyGatewayDotcomUser) GetUsername() string {
	return v.DotcomUserState.Username
}

// GetCodyGatewayAccess returns CheckDotcomUserAccessTokenDotcomDotcomQueryCodyGatewayDotcomUserByTokenCodyGatewayDotcomUser.CodyGatewayAccess, and is useful for accessing the field via an interface.
func (v *CheckDotcomUserAccessTokenDotcomDotcomQueryCodyGatewayDotcomUserByTokenCodyGatewayDotcomUser) GetCodyGatewayAccess() DotcomUserStateCodyGatewayAccess {
	return v.DotcomUserState.CodyGatewayAccess
}

func (v *CheckDotcomUserAccessTokenDotcomDotcomQueryCodyGatewayDotcomUserByTokenCodyGatewayDotcomUser) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*CheckDotcomUserAccessTokenDotcomDotcomQueryCodyGatewayDotcomUserByTokenCodyGatewayDotcomUser
		graphql.NoUnmarshalJSON
	}
	firstPass.CheckDotcomUserAccessTokenDotcomDotcomQueryCodyGatewayDotcomUserByTokenCodyGatewayDotcomUser = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	err = json.Unmarshal(
		b, &v.DotcomUserState)
	if err != nil {
		return err
	}
	return nil
}

type __premarshalCheckDotcomUserAccessTokenDotcomDotcomQueryCodyGatewayDotcomUserByTokenCodyGatewayDotcomUser struct {
	Id string `json:"id"`

	Username string `json:"username"`

	CodyGatewayAccess DotcomUserStateCodyGatewayAccess `json:"codyGatewayAccess"`
}

func (v *CheckDotcomUserAccessTokenDotcomDotcomQueryCodyGatewayDotcomUserByTokenCodyGatewayDotcomUser) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *CheckDotcomUserAccessTokenDotcomDotcomQueryCodyGatewayDotcomUserByTokenCodyGatewayDotcomUser) __premarshalJSON() (*__premarshalCheckDotcomUserAccessTokenDotcomDotcomQueryCodyGatewayDotcomUserByTokenCodyGatewayDotcomUser, error) {
	var retval __premarshalCheckDotcomUserAccessTokenDotcomDotcomQueryCodyGatewayDotcomUserByTokenCodyGatewayDotcomUser

	retval.Id = v.DotcomUserState.Id
	retval.Username = v.DotcomUserState.Username
	retval.CodyGatewayAccess = v.DotcomUserState.CodyGatewayAccess
	return &retval, nil
}

// CheckDotcomUserAccessTokenResponse is returned by CheckDotcomUserAccessToken on success.
type CheckDotcomUserAccessTokenResponse struct {
	// Queries that are only used on Khulnasoft.com.
	//
	// FOR INTERNAL USE ONLY.
	Dotcom CheckDotcomUserAccessTokenDotcomDotcomQuery `json:"dotcom"`
}

// GetDotcom returns CheckDotcomUserAccessTokenResponse.Dotcom, and is useful for accessing the field via an interface.
func (v *CheckDotcomUserAccessTokenResponse) GetDotcom() CheckDotcomUserAccessTokenDotcomDotcomQuery {
	return v.Dotcom
}

// CodyGatewayAccessFields includes the GraphQL fields of CodyGatewayAccess requested by the fragment CodyGatewayAccessFields.
// The GraphQL type's documentation follows.
//
// Cody Gateway access granted to a subscription.
// FOR INTERNAL USE ONLY.
type CodyGatewayAccessFields struct {
	// Whether or not a subscription has Cody Gateway access.
	Enabled bool `json:"enabled"`
	// Rate limit for chat completions access, or null if not enabled.
	ChatCompletionsRateLimit *CodyGatewayAccessFieldsChatCompletionsRateLimitCodyGatewayRateLimit `json:"chatCompletionsRateLimit"`
	// Rate limit for code completions access, or null if not enabled.
	CodeCompletionsRateLimit *CodyGatewayAccessFieldsCodeCompletionsRateLimitCodyGatewayRateLimit `json:"codeCompletionsRateLimit"`
	// Rate limit for embedding text chunks, or null if not enabled.
	EmbeddingsRateLimit *CodyGatewayAccessFieldsEmbeddingsRateLimitCodyGatewayRateLimit `json:"embeddingsRateLimit"`
}

// GetEnabled returns CodyGatewayAccessFields.Enabled, and is useful for accessing the field via an interface.
func (v *CodyGatewayAccessFields) GetEnabled() bool { return v.Enabled }

// GetChatCompletionsRateLimit returns CodyGatewayAccessFields.ChatCompletionsRateLimit, and is useful for accessing the field via an interface.
func (v *CodyGatewayAccessFields) GetChatCompletionsRateLimit() *CodyGatewayAccessFieldsChatCompletionsRateLimitCodyGatewayRateLimit {
	return v.ChatCompletionsRateLimit
}

// GetCodeCompletionsRateLimit returns CodyGatewayAccessFields.CodeCompletionsRateLimit, and is useful for accessing the field via an interface.
func (v *CodyGatewayAccessFields) GetCodeCompletionsRateLimit() *CodyGatewayAccessFieldsCodeCompletionsRateLimitCodyGatewayRateLimit {
	return v.CodeCompletionsRateLimit
}

// GetEmbeddingsRateLimit returns CodyGatewayAccessFields.EmbeddingsRateLimit, and is useful for accessing the field via an interface.
func (v *CodyGatewayAccessFields) GetEmbeddingsRateLimit() *CodyGatewayAccessFieldsEmbeddingsRateLimitCodyGatewayRateLimit {
	return v.EmbeddingsRateLimit
}

// CodyGatewayAccessFieldsChatCompletionsRateLimitCodyGatewayRateLimit includes the requested fields of the GraphQL type CodyGatewayRateLimit.
// The GraphQL type's documentation follows.
//
// Cody Gateway access rate limits for a subscription.
// FOR INTERNAL USE ONLY.
type CodyGatewayAccessFieldsChatCompletionsRateLimitCodyGatewayRateLimit struct {
	RateLimitFields `json:"-"`
}

// GetAllowedModels returns CodyGatewayAccessFieldsChatCompletionsRateLimitCodyGatewayRateLimit.AllowedModels, and is useful for accessing the field via an interface.
func (v *CodyGatewayAccessFieldsChatCompletionsRateLimitCodyGatewayRateLimit) GetAllowedModels() []string {
	return v.RateLimitFields.AllowedModels
}

// GetSource returns CodyGatewayAccessFieldsChatCompletionsRateLimitCodyGatewayRateLimit.Source, and is useful for accessing the field via an interface.
func (v *CodyGatewayAccessFieldsChatCompletionsRateLimitCodyGatewayRateLimit) GetSource() CodyGatewayRateLimitSource {
	return v.RateLimitFields.Source
}

// GetLimit returns CodyGatewayAccessFieldsChatCompletionsRateLimitCodyGatewayRateLimit.Limit, and is useful for accessing the field via an interface.
func (v *CodyGatewayAccessFieldsChatCompletionsRateLimitCodyGatewayRateLimit) GetLimit() genhelper.BigInt {
	return v.RateLimitFields.Limit
}

// GetIntervalSeconds returns CodyGatewayAccessFieldsChatCompletionsRateLimitCodyGatewayRateLimit.IntervalSeconds, and is useful for accessing the field via an interface.
func (v *CodyGatewayAccessFieldsChatCompletionsRateLimitCodyGatewayRateLimit) GetIntervalSeconds() int {
	return v.RateLimitFields.IntervalSeconds
}

func (v *CodyGatewayAccessFieldsChatCompletionsRateLimitCodyGatewayRateLimit) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*CodyGatewayAccessFieldsChatCompletionsRateLimitCodyGatewayRateLimit
		graphql.NoUnmarshalJSON
	}
	firstPass.CodyGatewayAccessFieldsChatCompletionsRateLimitCodyGatewayRateLimit = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	err = json.Unmarshal(
		b, &v.RateLimitFields)
	if err != nil {
		return err
	}
	return nil
}

type __premarshalCodyGatewayAccessFieldsChatCompletionsRateLimitCodyGatewayRateLimit struct {
	AllowedModels []string `json:"allowedModels"`

	Source CodyGatewayRateLimitSource `json:"source"`

	Limit genhelper.BigInt `json:"limit"`

	IntervalSeconds int `json:"intervalSeconds"`
}

func (v *CodyGatewayAccessFieldsChatCompletionsRateLimitCodyGatewayRateLimit) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *CodyGatewayAccessFieldsChatCompletionsRateLimitCodyGatewayRateLimit) __premarshalJSON() (*__premarshalCodyGatewayAccessFieldsChatCompletionsRateLimitCodyGatewayRateLimit, error) {
	var retval __premarshalCodyGatewayAccessFieldsChatCompletionsRateLimitCodyGatewayRateLimit

	retval.AllowedModels = v.RateLimitFields.AllowedModels
	retval.Source = v.RateLimitFields.Source
	retval.Limit = v.RateLimitFields.Limit
	retval.IntervalSeconds = v.RateLimitFields.IntervalSeconds
	return &retval, nil
}

// CodyGatewayAccessFieldsCodeCompletionsRateLimitCodyGatewayRateLimit includes the requested fields of the GraphQL type CodyGatewayRateLimit.
// The GraphQL type's documentation follows.
//
// Cody Gateway access rate limits for a subscription.
// FOR INTERNAL USE ONLY.
type CodyGatewayAccessFieldsCodeCompletionsRateLimitCodyGatewayRateLimit struct {
	RateLimitFields `json:"-"`
}

// GetAllowedModels returns CodyGatewayAccessFieldsCodeCompletionsRateLimitCodyGatewayRateLimit.AllowedModels, and is useful for accessing the field via an interface.
func (v *CodyGatewayAccessFieldsCodeCompletionsRateLimitCodyGatewayRateLimit) GetAllowedModels() []string {
	return v.RateLimitFields.AllowedModels
}

// GetSource returns CodyGatewayAccessFieldsCodeCompletionsRateLimitCodyGatewayRateLimit.Source, and is useful for accessing the field via an interface.
func (v *CodyGatewayAccessFieldsCodeCompletionsRateLimitCodyGatewayRateLimit) GetSource() CodyGatewayRateLimitSource {
	return v.RateLimitFields.Source
}

// GetLimit returns CodyGatewayAccessFieldsCodeCompletionsRateLimitCodyGatewayRateLimit.Limit, and is useful for accessing the field via an interface.
func (v *CodyGatewayAccessFieldsCodeCompletionsRateLimitCodyGatewayRateLimit) GetLimit() genhelper.BigInt {
	return v.RateLimitFields.Limit
}

// GetIntervalSeconds returns CodyGatewayAccessFieldsCodeCompletionsRateLimitCodyGatewayRateLimit.IntervalSeconds, and is useful for accessing the field via an interface.
func (v *CodyGatewayAccessFieldsCodeCompletionsRateLimitCodyGatewayRateLimit) GetIntervalSeconds() int {
	return v.RateLimitFields.IntervalSeconds
}

func (v *CodyGatewayAccessFieldsCodeCompletionsRateLimitCodyGatewayRateLimit) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*CodyGatewayAccessFieldsCodeCompletionsRateLimitCodyGatewayRateLimit
		graphql.NoUnmarshalJSON
	}
	firstPass.CodyGatewayAccessFieldsCodeCompletionsRateLimitCodyGatewayRateLimit = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	err = json.Unmarshal(
		b, &v.RateLimitFields)
	if err != nil {
		return err
	}
	return nil
}

type __premarshalCodyGatewayAccessFieldsCodeCompletionsRateLimitCodyGatewayRateLimit struct {
	AllowedModels []string `json:"allowedModels"`

	Source CodyGatewayRateLimitSource `json:"source"`

	Limit genhelper.BigInt `json:"limit"`

	IntervalSeconds int `json:"intervalSeconds"`
}

func (v *CodyGatewayAccessFieldsCodeCompletionsRateLimitCodyGatewayRateLimit) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *CodyGatewayAccessFieldsCodeCompletionsRateLimitCodyGatewayRateLimit) __premarshalJSON() (*__premarshalCodyGatewayAccessFieldsCodeCompletionsRateLimitCodyGatewayRateLimit, error) {
	var retval __premarshalCodyGatewayAccessFieldsCodeCompletionsRateLimitCodyGatewayRateLimit

	retval.AllowedModels = v.RateLimitFields.AllowedModels
	retval.Source = v.RateLimitFields.Source
	retval.Limit = v.RateLimitFields.Limit
	retval.IntervalSeconds = v.RateLimitFields.IntervalSeconds
	return &retval, nil
}

// CodyGatewayAccessFieldsEmbeddingsRateLimitCodyGatewayRateLimit includes the requested fields of the GraphQL type CodyGatewayRateLimit.
// The GraphQL type's documentation follows.
//
// Cody Gateway access rate limits for a subscription.
// FOR INTERNAL USE ONLY.
type CodyGatewayAccessFieldsEmbeddingsRateLimitCodyGatewayRateLimit struct {
	RateLimitFields `json:"-"`
}

// GetAllowedModels returns CodyGatewayAccessFieldsEmbeddingsRateLimitCodyGatewayRateLimit.AllowedModels, and is useful for accessing the field via an interface.
func (v *CodyGatewayAccessFieldsEmbeddingsRateLimitCodyGatewayRateLimit) GetAllowedModels() []string {
	return v.RateLimitFields.AllowedModels
}

// GetSource returns CodyGatewayAccessFieldsEmbeddingsRateLimitCodyGatewayRateLimit.Source, and is useful for accessing the field via an interface.
func (v *CodyGatewayAccessFieldsEmbeddingsRateLimitCodyGatewayRateLimit) GetSource() CodyGatewayRateLimitSource {
	return v.RateLimitFields.Source
}

// GetLimit returns CodyGatewayAccessFieldsEmbeddingsRateLimitCodyGatewayRateLimit.Limit, and is useful for accessing the field via an interface.
func (v *CodyGatewayAccessFieldsEmbeddingsRateLimitCodyGatewayRateLimit) GetLimit() genhelper.BigInt {
	return v.RateLimitFields.Limit
}

// GetIntervalSeconds returns CodyGatewayAccessFieldsEmbeddingsRateLimitCodyGatewayRateLimit.IntervalSeconds, and is useful for accessing the field via an interface.
func (v *CodyGatewayAccessFieldsEmbeddingsRateLimitCodyGatewayRateLimit) GetIntervalSeconds() int {
	return v.RateLimitFields.IntervalSeconds
}

func (v *CodyGatewayAccessFieldsEmbeddingsRateLimitCodyGatewayRateLimit) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*CodyGatewayAccessFieldsEmbeddingsRateLimitCodyGatewayRateLimit
		graphql.NoUnmarshalJSON
	}
	firstPass.CodyGatewayAccessFieldsEmbeddingsRateLimitCodyGatewayRateLimit = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	err = json.Unmarshal(
		b, &v.RateLimitFields)
	if err != nil {
		return err
	}
	return nil
}

type __premarshalCodyGatewayAccessFieldsEmbeddingsRateLimitCodyGatewayRateLimit struct {
	AllowedModels []string `json:"allowedModels"`

	Source CodyGatewayRateLimitSource `json:"source"`

	Limit genhelper.BigInt `json:"limit"`

	IntervalSeconds int `json:"intervalSeconds"`
}

func (v *CodyGatewayAccessFieldsEmbeddingsRateLimitCodyGatewayRateLimit) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *CodyGatewayAccessFieldsEmbeddingsRateLimitCodyGatewayRateLimit) __premarshalJSON() (*__premarshalCodyGatewayAccessFieldsEmbeddingsRateLimitCodyGatewayRateLimit, error) {
	var retval __premarshalCodyGatewayAccessFieldsEmbeddingsRateLimitCodyGatewayRateLimit

	retval.AllowedModels = v.RateLimitFields.AllowedModels
	retval.Source = v.RateLimitFields.Source
	retval.Limit = v.RateLimitFields.Limit
	retval.IntervalSeconds = v.RateLimitFields.IntervalSeconds
	return &retval, nil
}

// The source of the rate limit returned.
// FOR INTERNAL USE ONLY.
type CodyGatewayRateLimitSource string

const (
	// Indicates that a custom override for the rate limit has been stored.
	CodyGatewayRateLimitSourceOverride CodyGatewayRateLimitSource = "OVERRIDE"
	// Indicates that the rate limit is inferred by the subscriptions active plan.
	CodyGatewayRateLimitSourcePlan CodyGatewayRateLimitSource = "PLAN"
)

// DotcomUserState includes the GraphQL fields of CodyGatewayDotcomUser requested by the fragment DotcomUserState.
// The GraphQL type's documentation follows.
//
// A dotcom user allowed to access the Cody Gateway
// FOR INTERNAL USE ONLY.
type DotcomUserState struct {
	// The id of the user
	Id string `json:"id"`
	// The user name of the user
	Username string `json:"username"`
	// Cody Gateway access granted to this user. Properties may be inferred from dotcom site config, or be defined in overrides on the user.
	CodyGatewayAccess DotcomUserStateCodyGatewayAccess `json:"codyGatewayAccess"`
}

// GetId returns DotcomUserState.Id, and is useful for accessing the field via an interface.
func (v *DotcomUserState) GetId() string { return v.Id }

// GetUsername returns DotcomUserState.Username, and is useful for accessing the field via an interface.
func (v *DotcomUserState) GetUsername() string { return v.Username }

// GetCodyGatewayAccess returns DotcomUserState.CodyGatewayAccess, and is useful for accessing the field via an interface.
func (v *DotcomUserState) GetCodyGatewayAccess() DotcomUserStateCodyGatewayAccess {
	return v.CodyGatewayAccess
}

// DotcomUserStateCodyGatewayAccess includes the requested fields of the GraphQL type CodyGatewayAccess.
// The GraphQL type's documentation follows.
//
// Cody Gateway access granted to a subscription.
// FOR INTERNAL USE ONLY.
type DotcomUserStateCodyGatewayAccess struct {
	CodyGatewayAccessFields `json:"-"`
}

// GetEnabled returns DotcomUserStateCodyGatewayAccess.Enabled, and is useful for accessing the field via an interface.
func (v *DotcomUserStateCodyGatewayAccess) GetEnabled() bool {
	return v.CodyGatewayAccessFields.Enabled
}

// GetChatCompletionsRateLimit returns DotcomUserStateCodyGatewayAccess.ChatCompletionsRateLimit, and is useful for accessing the field via an interface.
func (v *DotcomUserStateCodyGatewayAccess) GetChatCompletionsRateLimit() *CodyGatewayAccessFieldsChatCompletionsRateLimitCodyGatewayRateLimit {
	return v.CodyGatewayAccessFields.ChatCompletionsRateLimit
}

// GetCodeCompletionsRateLimit returns DotcomUserStateCodyGatewayAccess.CodeCompletionsRateLimit, and is useful for accessing the field via an interface.
func (v *DotcomUserStateCodyGatewayAccess) GetCodeCompletionsRateLimit() *CodyGatewayAccessFieldsCodeCompletionsRateLimitCodyGatewayRateLimit {
	return v.CodyGatewayAccessFields.CodeCompletionsRateLimit
}

// GetEmbeddingsRateLimit returns DotcomUserStateCodyGatewayAccess.EmbeddingsRateLimit, and is useful for accessing the field via an interface.
func (v *DotcomUserStateCodyGatewayAccess) GetEmbeddingsRateLimit() *CodyGatewayAccessFieldsEmbeddingsRateLimitCodyGatewayRateLimit {
	return v.CodyGatewayAccessFields.EmbeddingsRateLimit
}

func (v *DotcomUserStateCodyGatewayAccess) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*DotcomUserStateCodyGatewayAccess
		graphql.NoUnmarshalJSON
	}
	firstPass.DotcomUserStateCodyGatewayAccess = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	err = json.Unmarshal(
		b, &v.CodyGatewayAccessFields)
	if err != nil {
		return err
	}
	return nil
}

type __premarshalDotcomUserStateCodyGatewayAccess struct {
	Enabled bool `json:"enabled"`

	ChatCompletionsRateLimit *CodyGatewayAccessFieldsChatCompletionsRateLimitCodyGatewayRateLimit `json:"chatCompletionsRateLimit"`

	CodeCompletionsRateLimit *CodyGatewayAccessFieldsCodeCompletionsRateLimitCodyGatewayRateLimit `json:"codeCompletionsRateLimit"`

	EmbeddingsRateLimit *CodyGatewayAccessFieldsEmbeddingsRateLimitCodyGatewayRateLimit `json:"embeddingsRateLimit"`
}

func (v *DotcomUserStateCodyGatewayAccess) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *DotcomUserStateCodyGatewayAccess) __premarshalJSON() (*__premarshalDotcomUserStateCodyGatewayAccess, error) {
	var retval __premarshalDotcomUserStateCodyGatewayAccess

	retval.Enabled = v.CodyGatewayAccessFields.Enabled
	retval.ChatCompletionsRateLimit = v.CodyGatewayAccessFields.ChatCompletionsRateLimit
	retval.CodeCompletionsRateLimit = v.CodyGatewayAccessFields.CodeCompletionsRateLimit
	retval.EmbeddingsRateLimit = v.CodyGatewayAccessFields.EmbeddingsRateLimit
	return &retval, nil
}

// RateLimitFields includes the GraphQL fields of CodyGatewayRateLimit requested by the fragment RateLimitFields.
// The GraphQL type's documentation follows.
//
// Cody Gateway access rate limits for a subscription.
// FOR INTERNAL USE ONLY.
type RateLimitFields struct {
	// The models that are allowed for this rate limit bucket.
	// Usually, customers will have two separate rate limits, one
	// for chat completions and one for code completions. A usual
	// config could include:
	//
	// chatCompletionsRateLimit: {
	// allowedModels: [anthropic/claude-v1, anthropic/claude-v1.3]
	// },
	// codeCompletionsRateLimit: {
	// allowedModels: [anthropic/claude-instant-v1]
	// }
	//
	// In general, the model names are of the format "$PROVIDER/$MODEL_NAME".
	AllowedModels []string `json:"allowedModels"`
	// The source of the rate limit configuration.
	Source CodyGatewayRateLimitSource `json:"source"`
	// Requests per time interval.
	Limit genhelper.BigInt `json:"limit"`
	// Interval for rate limiting.
	IntervalSeconds int `json:"intervalSeconds"`
}

// GetAllowedModels returns RateLimitFields.AllowedModels, and is useful for accessing the field via an interface.
func (v *RateLimitFields) GetAllowedModels() []string { return v.AllowedModels }

// GetSource returns RateLimitFields.Source, and is useful for accessing the field via an interface.
func (v *RateLimitFields) GetSource() CodyGatewayRateLimitSource { return v.Source }

// GetLimit returns RateLimitFields.Limit, and is useful for accessing the field via an interface.
func (v *RateLimitFields) GetLimit() genhelper.BigInt { return v.Limit }

// GetIntervalSeconds returns RateLimitFields.IntervalSeconds, and is useful for accessing the field via an interface.
func (v *RateLimitFields) GetIntervalSeconds() int { return v.IntervalSeconds }

// SnippetAttributionResponse is returned by SnippetAttribution on success.
type SnippetAttributionResponse struct {
	// EXPERIMENTAL: Searches the instances indexed code for code matching snippet.
	SnippetAttribution SnippetAttributionSnippetAttributionSnippetAttributionConnection `json:"snippetAttribution"`
}

// GetSnippetAttribution returns SnippetAttributionResponse.SnippetAttribution, and is useful for accessing the field via an interface.
func (v *SnippetAttributionResponse) GetSnippetAttribution() SnippetAttributionSnippetAttributionSnippetAttributionConnection {
	return v.SnippetAttribution
}

// SnippetAttributionSnippetAttributionSnippetAttributionConnection includes the requested fields of the GraphQL type SnippetAttributionConnection.
// The GraphQL type's documentation follows.
//
// EXPERIMENTAL: A list of snippet attributions.
type SnippetAttributionSnippetAttributionSnippetAttributionConnection struct {
	// totalCount is the total number of repository attributions we found before
	// stopping the search.
	//
	// Note: if we didn't finish searching the full corpus then limitHit will be
	// true. For filtering use case this means if limitHit is true you need to be
	// conservative with TotalCount and assume it could be higher.
	TotalCount int `json:"totalCount"`
	// limitHit is true if we stopped searching before looking into the full
	// corpus. If limitHit is true then it is possible there are more than
	// totalCount attributions.
	LimitHit bool `json:"limitHit"`
	// The page set of SnippetAttribution entries in this connection.
	Nodes []SnippetAttributionSnippetAttributionSnippetAttributionConnectionNodesSnippetAttribution `json:"nodes"`
}

// GetTotalCount returns SnippetAttributionSnippetAttributionSnippetAttributionConnection.TotalCount, and is useful for accessing the field via an interface.
func (v *SnippetAttributionSnippetAttributionSnippetAttributionConnection) GetTotalCount() int {
	return v.TotalCount
}

// GetLimitHit returns SnippetAttributionSnippetAttributionSnippetAttributionConnection.LimitHit, and is useful for accessing the field via an interface.
func (v *SnippetAttributionSnippetAttributionSnippetAttributionConnection) GetLimitHit() bool {
	return v.LimitHit
}

// GetNodes returns SnippetAttributionSnippetAttributionSnippetAttributionConnection.Nodes, and is useful for accessing the field via an interface.
func (v *SnippetAttributionSnippetAttributionSnippetAttributionConnection) GetNodes() []SnippetAttributionSnippetAttributionSnippetAttributionConnectionNodesSnippetAttribution {
	return v.Nodes
}

// SnippetAttributionSnippetAttributionSnippetAttributionConnectionNodesSnippetAttribution includes the requested fields of the GraphQL type SnippetAttribution.
// The GraphQL type's documentation follows.
//
// EXPERIMENTAL: Attribution result from snippetAttribution.
type SnippetAttributionSnippetAttributionSnippetAttributionConnectionNodesSnippetAttribution struct {
	// The name of the repository containing the snippet.
	//
	// Note: we do not return a type Repository since repositoryName may
	// represent a repository not on this instance. eg a match from the
	// khulnasoft.com open source corpus.
	RepositoryName string `json:"repositoryName"`
}

// GetRepositoryName returns SnippetAttributionSnippetAttributionSnippetAttributionConnectionNodesSnippetAttribution.RepositoryName, and is useful for accessing the field via an interface.
func (v *SnippetAttributionSnippetAttributionSnippetAttributionConnectionNodesSnippetAttribution) GetRepositoryName() string {
	return v.RepositoryName
}

// __CheckDotcomUserAccessTokenInput is used internally by genqlient
type __CheckDotcomUserAccessTokenInput struct {
	Token string `json:"token"`
}

// GetToken returns __CheckDotcomUserAccessTokenInput.Token, and is useful for accessing the field via an interface.
func (v *__CheckDotcomUserAccessTokenInput) GetToken() string { return v.Token }

// __SnippetAttributionInput is used internally by genqlient
type __SnippetAttributionInput struct {
	Snippet string `json:"snippet"`
	First   int    `json:"first"`
}

// GetSnippet returns __SnippetAttributionInput.Snippet, and is useful for accessing the field via an interface.
func (v *__SnippetAttributionInput) GetSnippet() string { return v.Snippet }

// GetFirst returns __SnippetAttributionInput.First, and is useful for accessing the field via an interface.
func (v *__SnippetAttributionInput) GetFirst() int { return v.First }

// CheckDotcomUserAccessToken returns traits of the product subscription associated with
// the given access token.
func CheckDotcomUserAccessToken(
	ctx context.Context,
	client graphql.Client,
	token string,
) (*CheckDotcomUserAccessTokenResponse, error) {
	req := &graphql.Request{
		OpName: "CheckDotcomUserAccessToken",
		Query: `
query CheckDotcomUserAccessToken ($token: String!) {
	dotcom {
		codyGatewayDotcomUserByToken(token: $token) {
			... DotcomUserState
		}
	}
}
fragment DotcomUserState on CodyGatewayDotcomUser {
	id
	username
	codyGatewayAccess {
		... CodyGatewayAccessFields
	}
}
fragment CodyGatewayAccessFields on CodyGatewayAccess {
	enabled
	chatCompletionsRateLimit {
		... RateLimitFields
	}
	codeCompletionsRateLimit {
		... RateLimitFields
	}
	embeddingsRateLimit {
		... RateLimitFields
	}
}
fragment RateLimitFields on CodyGatewayRateLimit {
	allowedModels
	source
	limit
	intervalSeconds
}
`,
		Variables: &__CheckDotcomUserAccessTokenInput{
			Token: token,
		},
	}
	var err error

	var data CheckDotcomUserAccessTokenResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

// Searches the instances indexed code for code matching snippet.
func SnippetAttribution(
	ctx context.Context,
	client graphql.Client,
	snippet string,
	first int,
) (*SnippetAttributionResponse, error) {
	req := &graphql.Request{
		OpName: "SnippetAttribution",
		Query: `
query SnippetAttribution ($snippet: String!, $first: Int!) {
	snippetAttribution(snippet: $snippet, first: $first) {
		totalCount
		limitHit
		nodes {
			repositoryName
		}
	}
}
`,
		Variables: &__SnippetAttributionInput{
			Snippet: snippet,
			First:   first,
		},
	}
	var err error

	var data SnippetAttributionResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
