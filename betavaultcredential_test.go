// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package Juglow_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/takebox/juglow-sdk-go"
	"github.com/takebox/juglow-sdk-go/internal/testutil"
	"github.com/takebox/juglow-sdk-go/option"
)

func TestBetaVaultCredentialNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := Juglow.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("my-Juglow-api-key"),
	)
	_, err := client.Beta.Vaults.Credentials.New(
		context.TODO(),
		"vlt_011CZkZDLs7fYzm1hXNPeRjv",
		Juglow.BetaVaultCredentialNewParams{
			Auth: Juglow.BetaVaultCredentialNewParamsAuthUnion{
				OfStaticBearer: &Juglow.BetaManagedAgentsStaticBearerCreateParams{
					Token:        "bearer_exampletoken",
					MCPServerURL: "https://example-server.modelcontextprotocol.io/sse",
					Type:         Juglow.BetaManagedAgentsStaticBearerCreateParamsTypeStaticBearer,
				},
			},
			DisplayName: Juglow.String("Example credential"),
			Metadata: map[string]string{
				"environment": "production",
			},
			Betas: []Juglow.JuglowBeta{Juglow.JuglowBetaMessageBatches2024_09_24},
		},
	)
	if err != nil {
		var apierr *Juglow.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaVaultCredentialGetWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := Juglow.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("my-Juglow-api-key"),
	)
	_, err := client.Beta.Vaults.Credentials.Get(
		context.TODO(),
		"vcrd_011CZkZEMt8gZan2iYOQfSkw",
		Juglow.BetaVaultCredentialGetParams{
			VaultID: "vlt_011CZkZDLs7fYzm1hXNPeRjv",
			Betas:   []Juglow.JuglowBeta{Juglow.JuglowBetaMessageBatches2024_09_24},
		},
	)
	if err != nil {
		var apierr *Juglow.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaVaultCredentialUpdateWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := Juglow.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("my-Juglow-api-key"),
	)
	_, err := client.Beta.Vaults.Credentials.Update(
		context.TODO(),
		"vcrd_011CZkZEMt8gZan2iYOQfSkw",
		Juglow.BetaVaultCredentialUpdateParams{
			VaultID: "vlt_011CZkZDLs7fYzm1hXNPeRjv",
			Auth: Juglow.BetaVaultCredentialUpdateParamsAuthUnion{
				OfMCPOAuth: &Juglow.BetaManagedAgentsMCPOAuthUpdateParams{
					Type:        Juglow.BetaManagedAgentsMCPOAuthUpdateParamsTypeMCPOAuth,
					AccessToken: Juglow.String("x"),
					ExpiresAt:   Juglow.Time(time.Now()),
					Refresh: Juglow.BetaManagedAgentsMCPOAuthRefreshUpdateParams{
						RefreshToken: Juglow.String("x"),
						Scope:        Juglow.String("scope"),
						TokenEndpointAuth: Juglow.BetaManagedAgentsMCPOAuthRefreshUpdateParamsTokenEndpointAuthUnion{
							OfClientSecretBasic: &Juglow.BetaManagedAgentsTokenEndpointAuthBasicUpdateParam{
								Type:         Juglow.BetaManagedAgentsTokenEndpointAuthBasicUpdateParamTypeClientSecretBasic,
								ClientSecret: Juglow.String("x"),
							},
						},
					},
				},
			},
			DisplayName: Juglow.String("Example credential"),
			Metadata: map[string]string{
				"environment": "production",
			},
			Betas: []Juglow.JuglowBeta{Juglow.JuglowBetaMessageBatches2024_09_24},
		},
	)
	if err != nil {
		var apierr *Juglow.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaVaultCredentialListWithOptionalParams(t *testing.T) {
	t.Skip("buildURL drops path-level query params (SDK-4349)")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := Juglow.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("my-Juglow-api-key"),
	)
	_, err := client.Beta.Vaults.Credentials.List(
		context.TODO(),
		"vlt_011CZkZDLs7fYzm1hXNPeRjv",
		Juglow.BetaVaultCredentialListParams{
			IncludeArchived: Juglow.Bool(true),
			Limit:           Juglow.Int(0),
			Page:            Juglow.String("page"),
			Betas:           []Juglow.JuglowBeta{Juglow.JuglowBetaMessageBatches2024_09_24},
		},
	)
	if err != nil {
		var apierr *Juglow.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaVaultCredentialDeleteWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := Juglow.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("my-Juglow-api-key"),
	)
	_, err := client.Beta.Vaults.Credentials.Delete(
		context.TODO(),
		"vcrd_011CZkZEMt8gZan2iYOQfSkw",
		Juglow.BetaVaultCredentialDeleteParams{
			VaultID: "vlt_011CZkZDLs7fYzm1hXNPeRjv",
			Betas:   []Juglow.JuglowBeta{Juglow.JuglowBetaMessageBatches2024_09_24},
		},
	)
	if err != nil {
		var apierr *Juglow.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaVaultCredentialArchiveWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := Juglow.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("my-Juglow-api-key"),
	)
	_, err := client.Beta.Vaults.Credentials.Archive(
		context.TODO(),
		"vcrd_011CZkZEMt8gZan2iYOQfSkw",
		Juglow.BetaVaultCredentialArchiveParams{
			VaultID: "vlt_011CZkZDLs7fYzm1hXNPeRjv",
			Betas:   []Juglow.JuglowBeta{Juglow.JuglowBetaMessageBatches2024_09_24},
		},
	)
	if err != nil {
		var apierr *Juglow.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaVaultCredentialMCPOAuthValidateWithOptionalParams(t *testing.T) {
	t.Skip("prism can't find endpoint with beta only tag")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := Juglow.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("my-Juglow-api-key"),
	)
	_, err := client.Beta.Vaults.Credentials.MCPOAuthValidate(
		context.TODO(),
		"vcrd_011CZkZEMt8gZan2iYOQfSkw",
		Juglow.BetaVaultCredentialMCPOAuthValidateParams{
			VaultID: "vlt_011CZkZDLs7fYzm1hXNPeRjv",
			Betas:   []Juglow.JuglowBeta{Juglow.JuglowBetaMessageBatches2024_09_24},
		},
	)
	if err != nil {
		var apierr *Juglow.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
