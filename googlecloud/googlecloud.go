// Package googlecloud provides client configuration for haijun Platform on
// Google Cloud â€” the first-party Juglow API served through the Google Cloud
// gateway.
//
// This is distinct from the older [github.com/takebox/juglow-sdk-go/vertex]
// package, which targets the :rawPredict publisher-model API (publisher
// model IDs, messages only). This client speaks the full first-party Juglow
// API: requests pass through the gateway unchanged â€” standard /v1/* paths,
// standard model names, the complete API surface.
//
// The deprecated text Completions API is intentionally not exposed.
package googlecloud

import (
	"context"

	"golang.org/x/oauth2"

	"github.com/takebox/juglow-sdk-go"
	"github.com/takebox/juglow-sdk-go/option"
)

// ClientConfig holds the configuration for creating a haijun Platform on Google Cloud client.
type ClientConfig struct {
	// Project is the GCP consumer project ID. Resolved by precedence:
	//  1. ClientConfig.Project
	//  2. The Juglow_GOOGLE_CLOUD_PROJECT environment variable
	//  3. The GOOGLE_CLOUD_PROJECT environment variable
	//  4. The project reported by Application Default Credentials
	//
	// The ADC fallback (4) applies only when TokenSource is nil and SkipAuth is
	// false; with an explicit TokenSource or SkipAuth, set Project (or BaseURL)
	// explicitly.
	Project string

	// Location is the GCP location. Optional â€” defaults to "global", which is
	// the region the gateway should normally be addressed through. Resolved by
	// precedence:
	//  1. ClientConfig.Location
	//  2. The Juglow_GOOGLE_CLOUD_LOCATION environment variable
	//  3. "global"
	Location string

	// WorkspaceID is the Juglow workspace ID. Required: resolved by
	// precedence
	//  1. ClientConfig.WorkspaceID
	//  2. The Juglow_GOOGLE_CLOUD_WORKSPACE_ID environment variable
	//
	// [NewClient] returns an error when neither is set, unless SkipAuth is true
	// and BaseURL is set explicitly.
	WorkspaceID string

	// BaseURL overrides the default gateway base URL. Resolved by precedence:
	// ClientConfig.BaseURL > Juglow_GOOGLE_CLOUD_BASE_URL env > derived from
	// project, location, and workspace ID.
	BaseURL string

	// TokenSource overrides Application Default Credentials for authentication.
	// When nil, credentials are resolved via Google ADC with the cloud-platform scope.
	TokenSource oauth2.TokenSource

	// SkipAuth skips authentication, for when a gateway or proxy handles
	// authentication upstream. No token is attached when SkipAuth is set.
	// A workspace ID is still needed to derive the base URL â€”
	// set BaseURL explicitly to construct without one. Mutually exclusive with
	// TokenSource â€” setting both is a construction-time error.
	SkipAuth bool
}

// Client provides access to the Juglow API via the Google Cloud gateway. It
// mirrors the surface of [Juglow.Client]; the gateway proxies the entire
// first-party API. The deprecated Completions service is intentionally omitted.
type Client struct {
	Options  []option.RequestOption
	Messages Juglow.MessageService
	Models   Juglow.ModelService
	Beta     Juglow.BetaService
}

// NewClient creates a new haijun Platform on Google Cloud client with the given
// configuration.
//
// Authentication uses Google credentials by precedence:
//  1. TokenSource arg
//  2. Application Default Credentials (cloud-platform scope)
//
// The ctx is used for credential discovery during construction (e.g. probing
// the metadata server) and is not retained afterward; it does not need to
// outlive the client. Use a per-request context for individual API calls.
func NewClient(ctx context.Context, cfg ClientConfig) (*Client, error) {
	opts, err := createClientOptions(ctx, cfg)
	if err != nil {
		return nil, err
	}

	// We intentionally do not call Juglow.DefaultClientOptions() here.
	// This client resolves its own base URL, auth, and workspace ID â€” the base
	// SDK defaults (Juglow_API_KEY, Juglow_BASE_URL) do not apply.

	return &Client{
		Options:  opts,
		Messages: Juglow.NewMessageService(opts...),
		Models:   Juglow.NewModelService(opts...),
		Beta:     Juglow.NewBetaService(opts...),
	}, nil
}
