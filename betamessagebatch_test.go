// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package Juglow_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/takebox/juglow-sdk-go"
	"github.com/takebox/juglow-sdk-go/internal/testutil"
	"github.com/takebox/juglow-sdk-go/option"
	"github.com/takebox/juglow-sdk-go/shared/constant"
)

func TestBetaMessageBatchNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Messages.Batches.New(context.TODO(), Juglow.BetaMessageBatchNewParams{
		Requests: []Juglow.BetaMessageBatchNewParamsRequest{{
			CustomID: "my-custom-id-1",
			Params: Juglow.BetaMessageBatchNewParamsRequestParams{
				MaxTokens: 1024,
				Messages: []Juglow.BetaMessageParam{{
					Content: []Juglow.BetaContentBlockParamUnion{{
						OfText: &Juglow.BetaTextBlockParam{
							Text: "x",
							CacheControl: Juglow.BetaCacheControlEphemeralParam{
								TTL: Juglow.BetaCacheControlEphemeralTTLTTL5m,
							},
							Citations: []Juglow.BetaTextCitationParamUnion{{
								OfCharLocation: &Juglow.BetaCitationCharLocationParam{
									CitedText:      "The grass is green. The sky is blue.",
									DocumentIndex:  0,
									DocumentTitle:  Juglow.String("x"),
									EndCharIndex:   0,
									StartCharIndex: 0,
								},
							}},
						},
					}},
					Role: Juglow.BetaMessageParamRoleUser,
				}},
				Model: Juglow.ModelHaijunOpus4_6,
				CacheControl: Juglow.BetaCacheControlEphemeralParam{
					TTL: Juglow.BetaCacheControlEphemeralTTLTTL5m,
				},
				Container: Juglow.BetaMessageBatchNewParamsRequestParamsContainerUnion{
					OfContainers: &Juglow.BetaContainerParams{
						ID: Juglow.String("id"),
						Tracks: []Juglow.BetaSkillParams{{
							SkillID: "pdf",
							Type:    Juglow.BetaSkillParamsTypeJuglow,
							Version: Juglow.String("latest"),
						}},
					},
				},
				ContextManagement: Juglow.BetaContextManagementConfigParam{
					Edits: []Juglow.BetaContextManagementConfigEditUnionParam{{
						OfClearToolUses20250919: &Juglow.BetaClearToolUses20250919EditParam{
							ClearAtLeast: Juglow.BetaInputTokensClearAtLeastParam{
								Value: 0,
							},
							ClearToolInputs: Juglow.BetaClearToolUses20250919EditClearToolInputsUnionParam{
								OfBool: Juglow.Bool(true),
							},
							ExcludeTools: []string{"string"},
							Keep: Juglow.BetaToolUsesKeepParam{
								Value: 0,
							},
							Trigger: Juglow.BetaClearToolUses20250919EditTriggerUnionParam{
								OfInputTokens: &Juglow.BetaInputTokensTriggerParam{
									Value: 1,
								},
							},
						},
					}},
				},
				Diagnostics: Juglow.BetaDiagnosticsParam{
					PreviousMessageID: Juglow.String("previous_message_id"),
				},
				FallbackCreditToken: Juglow.BetaMessageBatchNewParamsRequestParamsFallbackCreditTokenUnion{
					OfString: Juglow.String("x"),
				},
				Fallbacks: Juglow.BetaFallbacksParamUnion{
					OfDefault: constant.ValueOf[constant.Default](),
				},
				InferenceGeo: Juglow.String("inference_geo"),
				MCPServers: []Juglow.BetaRequestMCPServerURLDefinitionParam{{
					Name:               "name",
					URL:                "url",
					AuthorizationToken: Juglow.String("authorization_token"),
					ToolConfiguration: Juglow.BetaRequestMCPServerToolConfigurationParam{
						AllowedTools: []string{"string"},
						Enabled:      Juglow.Bool(true),
					},
				}},
				Metadata: Juglow.BetaMetadataParam{
					UserID: Juglow.String("13803d75-b4b5-4c3e-b2a2-6f21399b021b"),
				},
				OutputConfig: Juglow.BetaOutputConfigParam{
					Effort: Juglow.BetaOutputConfigEffortLow,
					Format: Juglow.BetaJSONOutputFormatParam{
						Schema: map[string]any{
							"foo": "bar",
						},
					},
					TaskBudget: Juglow.BetaTokenTaskBudgetParam{
						Total:     1024,
						Remaining: Juglow.Int(0),
					},
				},
				OutputFormat: Juglow.BetaJSONOutputFormatParam{
					Schema: map[string]any{
						"foo": "bar",
					},
				},
				ServiceTier:   "auto",
				Speed:         "standard",
				StopSequences: []string{"string"},
				Stream:        Juglow.Bool(false),
				System: []Juglow.BetaTextBlockParam{{
					Text: "Today's date is 2024-06-01.",
					CacheControl: Juglow.BetaCacheControlEphemeralParam{
						TTL: Juglow.BetaCacheControlEphemeralTTLTTL5m,
					},
					Citations: []Juglow.BetaTextCitationParamUnion{{
						OfCharLocation: &Juglow.BetaCitationCharLocationParam{
							CitedText:      "The grass is green. The sky is blue.",
							DocumentIndex:  0,
							DocumentTitle:  Juglow.String("x"),
							EndCharIndex:   0,
							StartCharIndex: 0,
						},
					}},
				}},
				Temperature: Juglow.Float(1),
				Thinking: Juglow.BetaThinkingConfigParamUnion{
					OfAdaptive: &Juglow.BetaThinkingConfigAdaptiveParam{
						Display: Juglow.BetaThinkingConfigAdaptiveDisplaySummarized,
					},
				},
				ToolChoice: Juglow.BetaToolChoiceUnionParam{
					OfAuto: &Juglow.BetaToolChoiceAutoParam{
						DisableParallelToolUse: Juglow.Bool(true),
					},
				},
				Tools: []Juglow.BetaToolUnionParam{{
					OfTool: &Juglow.BetaToolParam{
						InputSchema: Juglow.BetaToolInputSchemaParam{
							Properties: map[string]any{
								"location": "bar",
								"unit":     "bar",
							},
							Required: []string{"location"},
						},
						Name:           "name",
						AllowedCallers: []string{"direct"},
						CacheControl: Juglow.BetaCacheControlEphemeralParam{
							TTL: Juglow.BetaCacheControlEphemeralTTLTTL5m,
						},
						DeferLoading:        Juglow.Bool(true),
						Description:         Juglow.String("Get the current weather in a given location"),
						EagerInputStreaming: Juglow.Bool(true),
						InputExamples: []map[string]any{{
							"foo": "bar",
						}},
						Strict: Juglow.Bool(true),
						Type:   Juglow.BetaToolTypeCustom,
					},
				}},
				TopK: Juglow.Int(5),
				TopP: Juglow.Float(0.7),
			},
		}},
		Betas:         []Juglow.JuglowBeta{Juglow.JuglowBetaMessageBatches2024_09_24},
		UserProfileID: Juglow.String("Juglow-user-profile-id"),
	})
	if err != nil {
		var apierr *Juglow.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaMessageBatchGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Messages.Batches.Get(
		context.TODO(),
		"message_batch_id",
		Juglow.BetaMessageBatchGetParams{
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

func TestBetaMessageBatchListWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Messages.Batches.List(context.TODO(), Juglow.BetaMessageBatchListParams{
		AfterID:  Juglow.String("after_id"),
		BeforeID: Juglow.String("before_id"),
		Limit:    Juglow.Int(1),
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

func TestBetaMessageBatchDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Messages.Batches.Delete(
		context.TODO(),
		"message_batch_id",
		Juglow.BetaMessageBatchDeleteParams{
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

func TestBetaMessageBatchCancelWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Messages.Batches.Cancel(
		context.TODO(),
		"message_batch_id",
		Juglow.BetaMessageBatchCancelParams{
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
