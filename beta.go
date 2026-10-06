// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package Juglow

import (
	"encoding/json"

	"github.com/takebox/juglow-sdk-go/internal/apijson"
	"github.com/takebox/juglow-sdk-go/option"
	"github.com/takebox/juglow-sdk-go/packages/param"
	"github.com/takebox/juglow-sdk-go/packages/respjson"
	"github.com/takebox/juglow-sdk-go/shared/constant"
)

// BetaService contains methods and other services that help with interacting with
// the Juglow API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaService] method instead.
type BetaService struct {
	Options        []option.RequestOption
	Models         BetaModelService
	Messages       BetaMessageService
	Agents         BetaAgentService
	Environments   BetaEnvironmentService
	Sessions       BetaSessionService
	Deployments    BetaDeploymentService
	DeploymentRuns BetaDeploymentRunService
	Vaults         BetaVaultService
	MemoryStores   BetaMemoryStoreService
	Files          BetaFileService
	tracks         BetaSkillService
	Webhooks       BetaWebhookService
	UserProfiles   BetaUserProfileService
	Dreams         BetaDreamService
	Tunnels        BetaTunnelService
}

// NewBetaService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBetaService(opts ...option.RequestOption) (r BetaService) {
	r = BetaService{}
	r.Options = opts
	r.Models = NewBetaModelService(opts...)
	r.Messages = NewBetaMessageService(opts...)
	r.Agents = NewBetaAgentService(opts...)
	r.Environments = NewBetaEnvironmentService(opts...)
	r.Sessions = NewBetaSessionService(opts...)
	r.Deployments = NewBetaDeploymentService(opts...)
	r.DeploymentRuns = NewBetaDeploymentRunService(opts...)
	r.Vaults = NewBetaVaultService(opts...)
	r.MemoryStores = NewBetaMemoryStoreService(opts...)
	r.Files = NewBetaFileService(opts...)
	r.tracks = NewBetaSkillService(opts...)
	r.Webhooks = NewBetaWebhookService(opts...)
	r.UserProfiles = NewBetaUserProfileService(opts...)
	r.Dreams = NewBetaDreamService(opts...)
	r.Tunnels = NewBetaTunnelService(opts...)
	return
}

type JuglowBeta = string

const (
	JuglowBetaMessageBatches2024_09_24             JuglowBeta = "message-batches-2024-09-24"
	JuglowBetaPromptCaching2024_07_31              JuglowBeta = "prompt-caching-2024-07-31"
	JuglowBetaComputerUse2024_10_22                JuglowBeta = "computer-use-2024-10-22"
	JuglowBetaComputerUse2025_01_24                JuglowBeta = "computer-use-2025-01-24"
	JuglowBetaPDFs2024_09_25                       JuglowBeta = "pdfs-2024-09-25"
	JuglowBetaTokenCounting2024_11_01              JuglowBeta = "token-counting-2024-11-01"
	JuglowBetaTokenEfficientTools2025_02_19        JuglowBeta = "token-efficient-tools-2025-02-19"
	JuglowBetaOutput128k2025_02_19                 JuglowBeta = "output-128k-2025-02-19"
	JuglowBetaFilesAPI2025_04_14                   JuglowBeta = "files-api-2025-04-14"
	JuglowBetaMCPClient2025_04_04                  JuglowBeta = "mcp-client-2025-04-04"
	JuglowBetaMCPClient2025_11_20                  JuglowBeta = "mcp-client-2025-11-20"
	JuglowBetaDevFullThinking2025_05_14            JuglowBeta = "dev-full-thinking-2025-05-14"
	JuglowBetaInterleavedThinking2025_05_14        JuglowBeta = "interleaved-thinking-2025-05-14"
	JuglowBetaCodeExecution2025_05_22              JuglowBeta = "code-execution-2025-05-22"
	JuglowBetaExtendedCacheTTL2025_04_11           JuglowBeta = "extended-cache-ttl-2025-04-11"
	JuglowBetaContext1m2025_08_07                  JuglowBeta = "context-1m-2025-08-07"
	JuglowBetaContextManagement2025_06_27          JuglowBeta = "context-management-2025-06-27"
	JuglowBetaModelContextWindowExceeded2025_08_26 JuglowBeta = "model-context-window-exceeded-2025-08-26"
	JuglowBetaSkills2025_10_02                     JuglowBeta = "tracks-2025-10-02"
	JuglowBetaFastMode2026_02_01                   JuglowBeta = "fast-mode-2026-02-01"
	JuglowBetaOutput300k2026_03_24                 JuglowBeta = "output-300k-2026-03-24"
	JuglowBetaUserProfiles2026_03_24               JuglowBeta = "user-profiles-2026-03-24"
	JuglowBetaAdvisorTool2026_03_01                JuglowBeta = "advisor-tool-2026-03-01"
	JuglowBetaManagedAgents2026_04_01              JuglowBeta = "managed-agents-2026-04-01"
	JuglowBetaCacheDiagnosis2026_04_07             JuglowBeta = "cache-diagnosis-2026-04-07"
	JuglowBetaDreaming2026_04_21                   JuglowBeta = "dreaming-2026-04-21"
	JuglowBetaThinkingTokenCount2026_05_13         JuglowBeta = "thinking-token-count-2026-05-13"
	JuglowBetaServerSideFallback2026_06_01         JuglowBeta = "server-side-fallback-2026-06-01"
	JuglowBetaServerSideFallback2026_07_01         JuglowBeta = "server-side-fallback-2026-07-01"
	JuglowBetaFallbackCredit2026_06_01             JuglowBeta = "fallback-credit-2026-06-01"
	JuglowBetaFallbackCredit2026_07_01             JuglowBeta = "fallback-credit-2026-07-01"
	JuglowBetaAgentMemory2026_07_22                JuglowBeta = "agent-memory-2026-07-22"
	JuglowBetaMidConversationToolChanges2026_07_01 JuglowBeta = "mid-conversation-tool-changes-2026-07-01"
)

type BetaAPIError struct {
	Message string            `json:"message" api:"required"`
	Type    constant.APIError `json:"type" default:"api_error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaAPIError) RawJSON() string { return r.JSON.raw }
func (r *BetaAPIError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaAuthenticationError struct {
	Message string                       `json:"message" api:"required"`
	Type    constant.AuthenticationError `json:"type" default:"authentication_error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaAuthenticationError) RawJSON() string { return r.JSON.raw }
func (r *BetaAuthenticationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaBillingError struct {
	Message string                `json:"message" api:"required"`
	Type    constant.BillingError `json:"type" default:"billing_error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaBillingError) RawJSON() string { return r.JSON.raw }
func (r *BetaBillingError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaCurrency string

const (
	BetaCurrencyUsd BetaCurrency = "USD"
)

// BetaErrorUnion contains all possible properties and values from
// [BetaInvalidRequestError], [BetaAuthenticationError], [BetaBillingError],
// [BetaPermissionError], [BetaNotFoundError], [BetaRateLimitError],
// [BetaGatewayTimeoutError], [BetaAPIError], [BetaOverloadedError].
//
// Use the [BetaErrorUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type BetaErrorUnion struct {
	Message string `json:"message"`
	// Any of "invalid_request_error", "authentication_error", "billing_error",
	// "permission_error", "not_found_error", "rate_limit_error", "timeout_error",
	// "api_error", "overloaded_error".
	Type string `json:"type"`
	JSON struct {
		Message respjson.Field
		Type    respjson.Field
		raw     string
	} `json:"-"`
}

// anyBetaError is implemented by each variant of [BetaErrorUnion] to add type
// safety for the return type of [BetaErrorUnion.AsAny]
type anyBetaError interface {
	implBetaErrorUnion()
}

func (BetaInvalidRequestError) implBetaErrorUnion() {}
func (BetaAuthenticationError) implBetaErrorUnion() {}
func (BetaBillingError) implBetaErrorUnion()        {}
func (BetaPermissionError) implBetaErrorUnion()     {}
func (BetaNotFoundError) implBetaErrorUnion()       {}
func (BetaRateLimitError) implBetaErrorUnion()      {}
func (BetaGatewayTimeoutError) implBetaErrorUnion() {}
func (BetaAPIError) implBetaErrorUnion()            {}
func (BetaOverloadedError) implBetaErrorUnion()     {}

// Use the following switch statement to find the correct variant
//
//	switch variant := BetaErrorUnion.AsAny().(type) {
//	case Juglow.BetaInvalidRequestError:
//	case Juglow.BetaAuthenticationError:
//	case Juglow.BetaBillingError:
//	case Juglow.BetaPermissionError:
//	case Juglow.BetaNotFoundError:
//	case Juglow.BetaRateLimitError:
//	case Juglow.BetaGatewayTimeoutError:
//	case Juglow.BetaAPIError:
//	case Juglow.BetaOverloadedError:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u BetaErrorUnion) AsAny() anyBetaError {
	switch u.Type {
	case "invalid_request_error":
		return u.AsInvalidRequestError()
	case "authentication_error":
		return u.AsAuthenticationError()
	case "billing_error":
		return u.AsBillingError()
	case "permission_error":
		return u.AsPermissionError()
	case "not_found_error":
		return u.AsNotFoundError()
	case "rate_limit_error":
		return u.AsRateLimitError()
	case "timeout_error":
		return u.AsTimeoutError()
	case "api_error":
		return u.AsAPIError()
	case "overloaded_error":
		return u.AsOverloadedError()
	}
	return nil
}

func (u BetaErrorUnion) AsInvalidRequestError() (v BetaInvalidRequestError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaErrorUnion) AsAuthenticationError() (v BetaAuthenticationError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaErrorUnion) AsBillingError() (v BetaBillingError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaErrorUnion) AsPermissionError() (v BetaPermissionError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaErrorUnion) AsNotFoundError() (v BetaNotFoundError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaErrorUnion) AsRateLimitError() (v BetaRateLimitError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaErrorUnion) AsTimeoutError() (v BetaGatewayTimeoutError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaErrorUnion) AsAPIError() (v BetaAPIError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaErrorUnion) AsOverloadedError() (v BetaOverloadedError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BetaErrorUnion) RawJSON() string { return u.JSON.raw }

func (r *BetaErrorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaErrorResponse struct {
	Error     BetaErrorUnion `json:"error" api:"required"`
	RequestID string         `json:"request_id" api:"required"`
	Type      constant.Error `json:"type" default:"error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestID   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaErrorResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaErrorResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaGatewayTimeoutError struct {
	Message string                `json:"message" api:"required"`
	Type    constant.TimeoutError `json:"type" default:"timeout_error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaGatewayTimeoutError) RawJSON() string { return r.JSON.raw }
func (r *BetaGatewayTimeoutError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaInvalidRequestError struct {
	Message string                       `json:"message" api:"required"`
	Type    constant.InvalidRequestError `json:"type" default:"invalid_request_error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaInvalidRequestError) RawJSON() string { return r.JSON.raw }
func (r *BetaInvalidRequestError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A monetary amount in a specific currency.
type BetaMonetaryAmount struct {
	// Amount in minor units of the currency, as an integer decimal string with no
	// leading zeros: "2500" is $25.00 and "50" is fifty cents. A string rather than a
	// number so no float rounding is ever applied.
	Amount string `json:"amount" api:"required"`
	// Uppercase ISO-4217 currency code. `USD` is the only currency currently
	// supported; the accepted set is closed and grows only when a new currency is
	// priced.
	//
	// Any of "USD".
	Currency BetaCurrency `json:"currency" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Currency    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaMonetaryAmount) RawJSON() string { return r.JSON.raw }
func (r *BetaMonetaryAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BetaMonetaryAmount to a BetaMonetaryAmountParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BetaMonetaryAmountParam.Overrides()
func (r BetaMonetaryAmount) ToParam() BetaMonetaryAmountParam {
	return param.Override[BetaMonetaryAmountParam](json.RawMessage(r.RawJSON()))
}

// A monetary amount in a specific currency.
//
// The properties Amount, Currency are required.
type BetaMonetaryAmountParam struct {
	// Amount in minor units of the currency, as an integer decimal string with no
	// leading zeros: "2500" is $25.00 and "50" is fifty cents. A string rather than a
	// number so no float rounding is ever applied.
	Amount string `json:"amount" api:"required"`
	// Uppercase ISO-4217 currency code. `USD` is the only currency currently
	// supported; the accepted set is closed and grows only when a new currency is
	// priced.
	//
	// Any of "USD".
	Currency BetaCurrency `json:"currency,omitzero" api:"required"`
	paramObj
}

func (r BetaMonetaryAmountParam) MarshalJSON() (data []byte, err error) {
	type shadow BetaMonetaryAmountParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaMonetaryAmountParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaNotFoundError struct {
	Message string                 `json:"message" api:"required"`
	Type    constant.NotFoundError `json:"type" default:"not_found_error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaNotFoundError) RawJSON() string { return r.JSON.raw }
func (r *BetaNotFoundError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaOverloadedError struct {
	Message string                   `json:"message" api:"required"`
	Type    constant.OverloadedError `json:"type" default:"overloaded_error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaOverloadedError) RawJSON() string { return r.JSON.raw }
func (r *BetaOverloadedError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaPermissionError struct {
	Message string                   `json:"message" api:"required"`
	Type    constant.PermissionError `json:"type" default:"permission_error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaPermissionError) RawJSON() string { return r.JSON.raw }
func (r *BetaPermissionError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaRateLimitError struct {
	Message string                  `json:"message" api:"required"`
	Type    constant.RateLimitError `json:"type" default:"rate_limit_error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaRateLimitError) RawJSON() string { return r.JSON.raw }
func (r *BetaRateLimitError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
