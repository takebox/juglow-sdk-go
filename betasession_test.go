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

func TestBetaSessionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Sessions.New(context.TODO(), Juglow.BetaSessionNewParams{
		Agent: Juglow.BetaSessionNewParamsAgentUnion{
			OfString: Juglow.String("agent_011CZkYpogX7uDKUyvBTophP"),
		},
		EnvironmentID: "env_011CZkZ9X2dpNyB7HsEFoRfW",
		Budget: Juglow.BetaManagedAgentsBudgetLimitParam{
			MaxListCost: Juglow.BetaMonetaryAmountParam{
				Amount:   "2500",
				Currency: Juglow.BetaCurrencyUsd,
			},
			Type: Juglow.BetaManagedAgentsBudgetLimitTypeLimit,
		},
		InitialEvents: []Juglow.BetaSessionNewParamsInitialEventUnion{{
			OfUserMessage: &Juglow.BetaManagedAgentsUserMessageEventParams{
				Content: []Juglow.BetaManagedAgentsUserMessageEventParamsContentUnion{{
					OfText: &Juglow.BetaManagedAgentsTextBlockParam{
						Text: "Where is my order #1234?",
						Type: Juglow.BetaManagedAgentsTextBlockTypeText,
					},
				}},
				Type: Juglow.BetaManagedAgentsUserMessageEventParamsTypeUserMessage,
			},
		}},
		Metadata: map[string]string{
			"foo": "string",
		},
		Resources: []Juglow.BetaSessionNewParamsResourceUnion{{
			OfFile: &Juglow.BetaManagedAgentsFileResourceParams{
				FileID:    "file_011CNha8iCJcU1wXNR6q4V8w",
				Type:      Juglow.BetaManagedAgentsFileResourceParamsTypeFile,
				MountPath: Juglow.String("/uploads/receipt.pdf"),
			},
		}},
		Title:    Juglow.String("Order #1234 inquiry"),
		VaultIDs: []string{"string"},
		Betas:    []Juglow.JuglowBeta{Juglow.JuglowBetaMessageBatches2024_09_24},
	})
	if err != nil {
		var apierr *Juglow.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaSessionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Sessions.Get(
		context.TODO(),
		"sesn_011CZkZAtmR3yMPDzynEDxu7",
		Juglow.BetaSessionGetParams{
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

func TestBetaSessionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Sessions.Update(
		context.TODO(),
		"sesn_011CZkZAtmR3yMPDzynEDxu7",
		Juglow.BetaSessionUpdateParams{
			Agent: Juglow.BetaManagedAgentsSessionAgentUpdateParam{
				MCPServers: []Juglow.BetaManagedAgentsURLMCPServerParams{{
					Name: "example-mcp",
					Type: Juglow.BetaManagedAgentsURLMCPServerParamsTypeURL,
					URL:  "https://example-server.modelcontextprotocol.io/sse",
				}},
				Tools: []Juglow.BetaManagedAgentsSessionAgentUpdateToolUnionParam{{
					OfAgentToolset20260401: &Juglow.BetaManagedAgentsAgentToolset20260401Params{
						Type: Juglow.BetaManagedAgentsAgentToolset20260401ParamsTypeAgentToolset20260401,
						Configs: []Juglow.BetaManagedAgentsAgentToolConfigParams{{
							Name:    Juglow.BetaManagedAgentsAgentToolConfigParamsNameBash,
							Enabled: Juglow.Bool(true),
							PermissionPolicy: Juglow.BetaManagedAgentsAgentToolConfigParamsPermissionPolicyUnion{
								OfAlwaysAllow: &Juglow.BetaManagedAgentsAlwaysAllowPolicyParam{
									Type: Juglow.BetaManagedAgentsAlwaysAllowPolicyTypeAlwaysAllow,
								},
							},
						}},
						DefaultConfig: Juglow.BetaManagedAgentsAgentToolsetDefaultConfigParams{
							Enabled: Juglow.Bool(true),
							PermissionPolicy: Juglow.BetaManagedAgentsAgentToolsetDefaultConfigParamsPermissionPolicyUnion{
								OfAlwaysAllow: &Juglow.BetaManagedAgentsAlwaysAllowPolicyParam{
									Type: Juglow.BetaManagedAgentsAlwaysAllowPolicyTypeAlwaysAllow,
								},
							},
						},
					},
				}},
			},
			Budget: Juglow.BetaManagedAgentsBudgetLimitParam{
				MaxListCost: Juglow.BetaMonetaryAmountParam{
					Amount:   "2500",
					Currency: Juglow.BetaCurrencyUsd,
				},
				Type: Juglow.BetaManagedAgentsBudgetLimitTypeLimit,
			},
			Metadata: map[string]string{
				"foo": "string",
			},
			Title:    Juglow.String("Order #1234 inquiry"),
			VaultIDs: []string{"string"},
			Betas:    []Juglow.JuglowBeta{Juglow.JuglowBetaMessageBatches2024_09_24},
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

func TestBetaSessionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Sessions.List(context.TODO(), Juglow.BetaSessionListParams{
		AgentID:         Juglow.String("agent_id"),
		AgentVersion:    Juglow.Int(0),
		CreatedAtGt:     Juglow.Time(time.Now()),
		CreatedAtGte:    Juglow.Time(time.Now()),
		CreatedAtLt:     Juglow.Time(time.Now()),
		CreatedAtLte:    Juglow.Time(time.Now()),
		DeploymentID:    Juglow.String("deployment_id"),
		IncludeArchived: Juglow.Bool(true),
		Limit:           Juglow.Int(0),
		MemoryStoreID:   Juglow.String("memory_store_id"),
		Order:           Juglow.BetaSessionListParamsOrderAsc,
		Page:            Juglow.String("page"),
		Statuses:        []string{"rescheduling"},
		Betas:           []Juglow.JuglowBeta{Juglow.JuglowBetaMessageBatches2024_09_24},
	})
	if err != nil {
		var apierr *Juglow.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaSessionDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Sessions.Delete(
		context.TODO(),
		"sesn_011CZkZAtmR3yMPDzynEDxu7",
		Juglow.BetaSessionDeleteParams{
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

func TestBetaSessionArchiveWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Sessions.Archive(
		context.TODO(),
		"sesn_011CZkZAtmR3yMPDzynEDxu7",
		Juglow.BetaSessionArchiveParams{
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
