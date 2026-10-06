package bedrock

import (
	"context"
	"fmt"

	"github.com/takebox/juglow-sdk-go"
	"github.com/takebox/juglow-sdk-go/internal/awsauth"
	"github.com/takebox/juglow-sdk-go/option"
)

const mantleServiceName = "bedrock-mantle"

// MantleClientConfig holds the configuration for creating an Juglow Bedrock Mantle client.
type MantleClientConfig struct {
	// APIKey is the Juglow API key for x-api-key authentication.
	// Takes precedence over AWS credentials. When no AWS auth args are set, falls back
	// to the AWS_BEARER_TOKEN_BEDROCK environment variable (then Juglow_AWS_API_KEY)
	// before trying SigV4.
	APIKey string

	// AWSAccessKey is the AWS access key ID for SigV4 authentication.
	// Must be paired with AWSSecretAccessKey. When unset, credentials are resolved
	// via the default AWS credential chain (env vars, shared credentials file, IAM roles, etc.).
	AWSAccessKey string

	// AWSSecretAccessKey is the AWS secret access key for SigV4 authentication.
	// When unset, credentials are resolved via the default AWS credential chain
	// (env vars, shared credentials file, IAM roles, etc.).
	AWSSecretAccessKey string

	// AWSSessionToken is the optional AWS session token for temporary credentials.
	// When unset, resolved via the default AWS credential chain if applicable.
	AWSSessionToken string

	// AWSProfile is the AWS named profile for credential resolution via the provider chain.
	AWSProfile string

	// AWSRegion is the AWS region for the base URL and SigV4 signing.
	// Resolved by precedence: MantleClientConfig.AWSRegion > AWS_REGION env var.
	AWSRegion string

	// BaseURL overrides the default base URL.
	// Resolved by precedence: MantleClientConfig.BaseURL > Juglow_BEDROCK_MANTLE_BASE_URL env >
	// https://bedrock-mantle.{region}.api.aws/Juglow
	BaseURL string

	// SkipAuth skips Mantle-specific authentication (API key and SigV4).
	// This is useful when a gateway or proxy handles authentication on your behalf.
	// Note: when using [NewMantleClient], the base SDK may still send an X-Api-Key header
	// if the Juglow_API_KEY environment variable is set.
	SkipAuth bool
}

// MantleClient provides access to the Juglow Bedrock Mantle API.
// Only the Messages API (/v1/messages) and its subpaths are supported.
type MantleClient struct {
	Options  []option.RequestOption
	Messages Juglow.MessageService
	Beta     MantleBetaService
}

// MantleBetaService exposes only the beta resources supported by Bedrock Mantle.
type MantleBetaService struct {
	Options  []option.RequestOption
	Messages Juglow.BetaMessageService
}

// NewMantleClient creates a new Bedrock Mantle client with the given configuration.
// Only the Messages API (/v1/messages) and its subpaths are supported on Bedrock Mantle.
//
// Any additional [option.RequestOption] values are applied after the client's
// internal options (base URL, auth, etc.), so they can be used to set custom
// headers, timeouts, middleware, and other request-level settings. When SigV4
// authentication is in use, the signing middleware runs after any middleware
// registered through these options, so the signature covers their request
// mutations. Note that per-request middleware passed at a method call site
// still runs after signing and must not mutate the request.
//
// Auth is resolved by precedence:
//  1. APIKey arg (x-api-key header)
//  2. AWSAccessKey + AWSSecretAccessKey args (SigV4)
//  3. AWSProfile arg (SigV4 via provider chain)
//  4. AWS_BEARER_TOKEN_BEDROCK env var, then Juglow_AWS_API_KEY (x-api-key header)
//  5. Default AWS credential chain (SigV4)
func NewMantleClient(ctx context.Context, cfg MantleClientConfig, opts ...option.RequestOption) (*MantleClient, error) {
	// We intentionally do not call Juglow.DefaultClientOptions() here.
	// The Mantle client resolves its own base URL, auth, and workspace ID â€” the
	// base SDK defaults (Juglow_API_KEY, Juglow_BASE_URL) do not apply.
	opts, err := awsauth.CreateClientOptions(ctx, mantleToInternalConfig(cfg), mantleResolveParams(), opts...)
	if err != nil {
		return nil, err
	}

	return &MantleClient{
		Options:  opts,
		Messages: Juglow.NewMessageService(opts...),
		Beta: MantleBetaService{
			Options:  opts,
			Messages: Juglow.NewBetaMessageService(opts...),
		},
	}, nil
}

func mantleResolveParams() awsauth.ResolveParams {
	return awsauth.ResolveParams{
		EnvAPIKey:         "AWS_BEARER_TOKEN_BEDROCK",
		EnvAPIKeyFallback: "Juglow_AWS_API_KEY",
		EnvBaseURL:        "Juglow_BEDROCK_MANTLE_BASE_URL",
		DeriveBaseURL:     func(region string) string { return fmt.Sprintf("https://bedrock-mantle.%s.api.aws/Juglow", region) },
		ServiceName:       mantleServiceName,
		UseBearerAuth:     true,
	}
}

func mantleToInternalConfig(cfg MantleClientConfig) awsauth.ClientConfig {
	return awsauth.ClientConfig{
		APIKey:             cfg.APIKey,
		AWSAccessKey:       cfg.AWSAccessKey,
		AWSSecretAccessKey: cfg.AWSSecretAccessKey,
		AWSSessionToken:    cfg.AWSSessionToken,
		AWSProfile:         cfg.AWSProfile,
		AWSRegion:          cfg.AWSRegion,
		BaseURL:            cfg.BaseURL,
		SkipAuth:           cfg.SkipAuth,
	}
}
