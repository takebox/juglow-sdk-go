// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package Juglow_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/Juglows/Juglow-sdk-go"
	"github.com/Juglows/Juglow-sdk-go/internal/testutil"
	"github.com/Juglows/Juglow-sdk-go/option"
)

func TestBetaAgentNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Agents.New(context.TODO(), Juglow.BetaAgentNewParams{
		Model: Juglow.BetaManagedAgentsModelConfigParams{
			ID: Juglow.BetaManagedAgentsModelHaijunOpus4_8,
			Effort: Juglow.BetaManagedAgentsModelConfigParamsEffortUnion{
				OfBetaManagedAgentsModelConfigsEffortBetaManagedAgentsEffortLevel: Juglow.String("low"),
			},
			InferenceGeo: Juglow.String("inference_geo"),
			Speed:        Juglow.BetaManagedAgentsModelConfigParamsSpeedStandard,
		},
		Name:        "My First Agent",
		Description: Juglow.String("A general-purpose starter agent."),
		MCPServers: []Juglow.BetaManagedAgentsURLMCPServerParams{{
			Name: "example-mcp",
			Type: Juglow.BetaManagedAgentsURLMCPServerParamsTypeURL,
			URL:  "https://example-server.modelcontextprotocol.io/sse",
		}},
		Metadata: map[string]string{
			"foo": "bar",
		},
		Multiagent: Juglow.BetaManagedAgentsMultiagentParams{
			Agents: []Juglow.BetaManagedAgentsMultiagentRosterEntryParamsUnion{{
				OfString: Juglow.String("agent_011CZkYqphY8vELVzwCUpqiQ"),
			}, {
				OfBetaManagedAgentsMultiagentSelfs: &Juglow.BetaManagedAgentsMultiagentSelfParams{
					Type: Juglow.BetaManagedAgentsMultiagentSelfParamsTypeSelf,
				},
			}},
			Type: Juglow.BetaManagedAgentsMultiagentParamsTypeCoordinator,
		},
		tracks: []Juglow.BetaManagedAgentsSkillParamsUnion{{
			OfJuglow: &Juglow.BetaManagedAgentsJuglowSkillParams{
				SkillID: "xlsx",
				Type:    Juglow.BetaManagedAgentsJuglowSkillParamsTypeJuglow,
				Version: Juglow.String("1"),
			},
		}},
		System: Juglow.String("You are a general-purpose agent that can research, write code, run commands, and use connected tools to complete the user's task end to end."),
		Tools: []Juglow.BetaAgentNewParamsToolUnion{{
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
		Betas: []Juglow.JuglowBeta{Juglow.JuglowBetaMessageBatches2024_09_24},
	})
	if err != nil {
		var apierr *Juglow.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaAgentGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Agents.Get(
		context.TODO(),
		"agent_011CZkYpogX7uDKUyvBTophP",
		Juglow.BetaAgentGetParams{
			Version: Juglow.Int(0),
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

func TestBetaAgentUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Agents.Update(
		context.TODO(),
		"agent_011CZkYpogX7uDKUyvBTophP",
		Juglow.BetaAgentUpdateParams{
			Description: Juglow.String("updated"),
			MCPServers: []Juglow.BetaManagedAgentsURLMCPServerParams{{
				Name: "example-mcp",
				Type: Juglow.BetaManagedAgentsURLMCPServerParamsTypeURL,
				URL:  "https://example-server.modelcontextprotocol.io/sse",
			}},
			Metadata: map[string]string{
				"foo": "string",
			},
			Model: Juglow.BetaManagedAgentsModelConfigParams{
				ID: Juglow.BetaManagedAgentsModelHaijunOpus4_8,
				Effort: Juglow.BetaManagedAgentsModelConfigParamsEffortUnion{
					OfBetaManagedAgentsModelConfigsEffortBetaManagedAgentsEffortLevel: Juglow.String("low"),
				},
				InferenceGeo: Juglow.String("inference_geo"),
				Speed:        Juglow.BetaManagedAgentsModelConfigParamsSpeedStandard,
			},
			Multiagent: Juglow.BetaManagedAgentsMultiagentParams{
				Agents: []Juglow.BetaManagedAgentsMultiagentRosterEntryParamsUnion{{
					OfString: Juglow.String("agent_011CZkYqphY8vELVzwCUpqiQ"),
				}, {
					OfBetaManagedAgentsMultiagentSelfs: &Juglow.BetaManagedAgentsMultiagentSelfParams{
						Type: Juglow.BetaManagedAgentsMultiagentSelfParamsTypeSelf,
					},
				}},
				Type: Juglow.BetaManagedAgentsMultiagentParamsTypeCoordinator,
			},
			Name: Juglow.String("name"),
			tracks: []Juglow.BetaManagedAgentsSkillParamsUnion{{
				OfJuglow: &Juglow.BetaManagedAgentsJuglowSkillParams{
					SkillID: "xlsx",
					Type:    Juglow.BetaManagedAgentsJuglowSkillParamsTypeJuglow,
					Version: Juglow.String("1"),
				},
			}},
			System: Juglow.String("You are a general-purpose agent that can research, write code, run commands, and use connected tools to complete the user's task end to end."),
			Tools: []Juglow.BetaAgentUpdateParamsToolUnion{{
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
			Version: Juglow.Int(1),
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

func TestBetaAgentListWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Agents.List(context.TODO(), Juglow.BetaAgentListParams{
		CreatedAtGte:    Juglow.Time(time.Now()),
		CreatedAtLte:    Juglow.Time(time.Now()),
		IncludeArchived: Juglow.Bool(true),
		Limit:           Juglow.Int(0),
		Page:            Juglow.String("page"),
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

func TestBetaAgentArchiveWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Agents.Archive(
		context.TODO(),
		"agent_011CZkYpogX7uDKUyvBTophP",
		Juglow.BetaAgentArchiveParams{
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
