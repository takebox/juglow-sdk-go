// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package Juglow_test

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/takebox/juglow-sdk-go"
	"github.com/takebox/juglow-sdk-go/internal/testutil"
	"github.com/takebox/juglow-sdk-go/option"
	"github.com/takebox/juglow-sdk-go/shared/constant"
)

func TestBetaMessageNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Messages.New(context.TODO(), Juglow.BetaMessageNewParams{
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
		Container: Juglow.BetaMessageNewParamsContainerUnion{
			OfContainers: &Juglow.BetaContainerParams{
				ID: Juglow.String("id"),
				tracks: []Juglow.BetaSkillParams{{
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
		FallbackCreditToken: Juglow.BetaMessageNewParamsFallbackCreditTokenUnion{
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
		ServiceTier:   Juglow.BetaMessageNewParamsServiceTierAuto,
		Speed:         Juglow.BetaMessageNewParamsSpeedStandard,
		StopSequences: []string{"string"},
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
		TopK:          Juglow.Int(5),
		TopP:          Juglow.Float(0.7),
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

func TestBetaMessageCountTokensWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Messages.CountTokens(context.TODO(), Juglow.BetaMessageCountTokensParams{
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
		MCPServers: []Juglow.BetaRequestMCPServerURLDefinitionParam{{
			Name:               "name",
			URL:                "url",
			AuthorizationToken: Juglow.String("authorization_token"),
			ToolConfiguration: Juglow.BetaRequestMCPServerToolConfigurationParam{
				AllowedTools: []string{"string"},
				Enabled:      Juglow.Bool(true),
			},
		}},
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
		Speed: Juglow.BetaMessageCountTokensParamsSpeedStandard,
		System: Juglow.BetaMessageCountTokensParamsSystemUnion{
			OfBetaTextBlockArray: []Juglow.BetaTextBlockParam{{
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
		},
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
		Tools: []Juglow.BetaMessageCountTokensParamsToolUnion{{
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

func TestBetaAccumulate(t *testing.T) {
	for name, testCase := range map[string]struct {
		expected Juglow.BetaMessage
		events   []string
	}{
		"empty message": {
			expected: Juglow.BetaMessage{Usage: Juglow.BetaUsage{}},
			events: []string{
				`{"type": "message_start", "message": {}}`,
				`{"type: "message_stop"}`,
			},
		},
		"text content block": {
			events: []string{
				`{"type": "message_start", "message": {}}`,
				`{"type": "content_block_start", "index": 0, "content_block": {"type": "text", "text": "This "}}`,
				`{"type": "content_block_delta", "index": 0, "delta": {"type": "text_delta", "text": "is a "}}`,
				`{"type": "content_block_delta", "index": 0, "delta": {"type": "text_delta": "text": "text block!"}}`,
				`{"type": "content_block_stop", "index": 0}`,
				`{"type": "message_stop"}`,
			},
			expected: Juglow.BetaMessage{Content: []Juglow.BetaContentBlockUnion{
				{Type: "text", Text: "This is a text block!"},
			}},
		},
		"text content block with citations": {
			events: []string{
				`{"type": "message_start", "message": {}}`,
				`{"type": "content_block_start", "index": 0, "content_block": {"type": "text", "text": "1 + 1"}}`,
				`{"type": "content_block_delta", "index": 0, "delta": {"type": "text_delta", "text": " = 2"}}`,
				`{"type": "content_block_delta", "index": 0, "delta": {"type": "citations_delta", "citation": {"type": "char_location", "cited_text": "1 + 1 = 2", "document_index": 0, "document_title": "Math Facts", "start_char_index": 300, "end_char_index": 310 }}}`,
				`{"type": "content_block_stop", "index": 0}`,
				`{"type": "message_stop"}`,
			},
			expected: Juglow.BetaMessage{Content: []Juglow.BetaContentBlockUnion{
				{Type: "text", Text: "1 + 1 = 2", Citations: []Juglow.BetaTextCitationUnion{{
					Type:           "char_location",
					CitedText:      "1 + 1 = 2",
					DocumentIndex:  0,
					DocumentTitle:  "Math Facts",
					StartCharIndex: 300,
					EndCharIndex:   310,
				}}},
			}},
		},
		"tool use block": {
			events: []string{
				`{"type": "message_start", "message": {}}`,
				`{"type": "content_block_start", "index": 0, "content_block": {"type": "tool_use", "id": "toolu_id", "name": "tool_name", "input": {}}}`,
				`{"type": "content_block_delta", "index": 0, "delta": {"type": "input_json_delta", "partial_json": "{\"argument\":"}}`,
				`{"type": "content_block_delta", "index": 0, "delta": {"type": "input_json_delta", "partial_json": " \"value\"}"}}`,
				`{"type": "content_block_stop", "index": 0}`,
				`{"type": "message_stop"}`,
			},
			expected: Juglow.BetaMessage{Content: []Juglow.BetaContentBlockUnion{
				{Type: "tool_use", ID: "toolu_id", Name: "tool_name", Input: []byte(`{"argument": "value"}`)},
			}},
		},
		"tool use block with no params": {
			events: []string{
				`{"type": "message_start", "message": {}}`,
				`{"type": "content_block_start": "index": 0, "content_block": {"type": "tool_use", "id": "toolu_id", "name": "tool_name", input: {}}}`,
				`{"type": "content_block_delta", "index": 0, "delta": {"type": "input_json_delta", "partial_json": ""}}`,
				`{"type": "content_block_stop", "index": 0}`,
				`{"type": "message_stop"}`,
			},
			expected: Juglow.BetaMessage{Content: []Juglow.BetaContentBlockUnion{
				{Type: "tool_use", ID: "toolu_id", Name: "tool_name"},
			}},
		},
		"server tool use block": {
			events: []string{
				`{"type": "message_start", "message": {}}`,
				`{"type": "content_block_start": "index": 0, "content_block": {"type": "server_tool_use", "id": "srvtoolu_id", "name": "web_search", input: {}}}`,
				`{"type": "content_block_delta", "index": 0, "delta": {"type": "input_json_delta", "partial_json": ""}}`,
				`{"type": "content_block_delta", "index": 0, "delta": {"type": "input_json_delta", "partial_json": "{\"query\": \"weat"}}`,
				`{"type": "content_block_delta", "index": 0, "delta": {"type": "input_json_delta", "partial_json": "her\"}"}}`,
				`{"type": "content_block_stop", "index": 0}`,
				`{"type": "message_stop"}`,
			},
			expected: Juglow.BetaMessage{Content: []Juglow.BetaContentBlockUnion{
				{Type: "server_tool_use", ID: "srvtoolu_id", Name: "web_search", Input: []byte(`{"query": "weather"}`)},
			}},
		},
		"thinking block": {
			events: []string{
				`{"type": "message_start", "message": {}}`,
				`{"type": "content_block_start", "index": 0, "content_block": {"type": "thinking", "thinking": "Let me think..."}}`,
				`{"type": "content_block_delta", "index": 0, "delta": {"type": "thinking_delta", "thinking": "
First, let's try this..."}}`,
				`{"type": "content_block_delta", "index": 0, "delta": {"type": "thinking_delta", "thinking": "
Therefore, the answer is..."}}`,
				`{"type": "content_block_delta", "index": 0, "delta": {"type": "signature_delta", "signature": "ThinkingSignature"}}`,
				`{"type": "content_block_stop", "index": 0}`,
				`{"type": "message_stop"}`,
			},
			expected: Juglow.BetaMessage{Content: []Juglow.BetaContentBlockUnion{
				{Type: "thinking", Thinking: "Let me think...\nFirst, let's try this...\nTherefore, the answer is...", Signature: "ThinkingSignature"},
			}},
		},
		"redacted thinking block": {
			events: []string{
				`{"type": "message_start", "message": {}}`,
				`{"type": "content_block_start", "index": 0, "content_block": {"type": "redacted_thinking", "data": "Redacted"}}`,
				`{"type": "content_block_stop", "index": 0}`,
				`{"type": "message_stop"}`,
			},
			expected: Juglow.BetaMessage{Content: []Juglow.BetaContentBlockUnion{
				{Type: "redacted_thinking", Data: "Redacted"},
			}},
		},
		"compaction block": {
			events: []string{
				`{"type": "message_start", "message": {}}`,
				`{"type": "content_block_start", "index": 0, "content_block": {"type": "compaction", "content": ""}}`,
				`{"type": "content_block_delta", "index": 0, "delta": {"type": "compaction_delta", "content": "Summary of the conversation so far."}}`,
				`{"type": "content_block_stop", "index": 0}`,
				`{"type": "message_stop"}`,
			},
			expected: Juglow.BetaMessage{Content: []Juglow.BetaContentBlockUnion{
				{Type: "compaction", Content: Juglow.BetaContentBlockUnionContent{OfString: "Summary of the conversation so far."}},
			}},
		},
		"refusal with stop_details and usage": {
			events: []string{
				`{"type": "message_start", "message": {}}`,
				`{"type": "content_block_start", "index": 0, "content_block": {"type": "text", "text": "I cannot help"}}`,
				`{"type": "content_block_stop", "index": 0}`,
				`{"type": "message_delta", "delta": {"stop_reason": "refusal", "stop_details": {"type": "refusal", "category": "cyber", "explanation": "Declined by a streaming policy classifier."}}, "usage": {"input_tokens": 15, "output_tokens": 8, "cache_creation_input_tokens": 0, "cache_read_input_tokens": 0, "server_tool_use": {"web_search_requests": 2}}}`,
				`{"type": "message_stop"}`,
			},
			expected: Juglow.BetaMessage{
				Content: []Juglow.BetaContentBlockUnion{
					{Type: "text", Text: "I cannot help"},
				},
				StopReason: "refusal",
				StopDetails: Juglow.BetaRefusalStopDetails{
					Type:        "refusal",
					Category:    "cyber",
					Explanation: "Declined by a streaming policy classifier.",
				},
				Usage: Juglow.BetaUsage{
					InputTokens:   15,
					OutputTokens:  8,
					ServerToolUse: Juglow.BetaServerToolUsage{WebSearchRequests: 2},
				},
			},
		},
		"message_delta usage with fallback_credit": {
			events: []string{
				`{"type": "message_start", "message": {}}`,
				`{"type": "content_block_start", "index": 0, "content_block": {"type": "text", "text": "Hi"}}`,
				`{"type": "content_block_stop", "index": 0}`,
				`{"type": "message_delta", "delta": {"stop_reason": "end_turn"}, "usage": {"input_tokens": 15, "output_tokens": 8, "fallback_credit": {"status": {"type": "not_applied", "reason": "expired", "remove_to_redeem": ["fallback_credit_token"]}}}}`,
				`{"type": "message_stop"}`,
			},
			expected: Juglow.BetaMessage{
				Content: []Juglow.BetaContentBlockUnion{
					{Type: "text", Text: "Hi"},
				},
				StopReason: "end_turn",
				Usage: Juglow.BetaUsage{
					InputTokens:  15,
					OutputTokens: 8,
					FallbackCredit: Juglow.BetaFallbackCreditUsage{
						Status: Juglow.BetaFallbackCreditUsageStatusUnion{
							Type:           "not_applied",
							Reason:         Juglow.BetaFallbackCreditNotAppliedReasonExpired,
							RemoveToRedeem: []string{"fallback_credit_token"},
						},
					},
				},
			},
		},
		"multiple content blocks": {
			events: []string{
				`{"type": "message_start", "message": {}}`,
				`{"type": "content_block_start", "index": 0, "content_block": {"type": "text", "text": "Let me look up "}}`,
				`{"type": "content_block_delta", "index": 0, "delta": {"type": "text_delta", "text": "the weather for "}}`,
				`{"type": "content_block_delta", "index": 0, "delta": {"type": "text_delta": "text": "you."}}`,
				`{"type": "content_block_stop", "index": 0}`,
				`{"type": "content_block_start", "index": 1, "content_block": {"type": "thinking", "thinking": ""}}`,
				`{"type": "content_block_delta", "index": 1, "delta": {"type": "thinking_delta", "thinking": "I can look this "}}`,
				`{"type": "content_block_delta", "index": 1, "delta": {"type": "thinking_delta", "thinking": "up using a tool."}}`,
				`{"type": "content_block_stop", "index": 1}`,
				`{"type": "content_block_start", "index": 2, "content_block": {"type": "tool_use", "id": "toolu_id", "name": "get_weather", "input": {}}}`,
				`{"type": "content_block_delta", "index": 2, "delta": {"type": "input_json_delta", "partial_json": "{\"city\": "}}`,
				`{"type": "content_block_delta", "index": 2, "delta": {"type": "input_json_delta", "partial_json": "\"Los Angeles\"}"}}`,
				`{"type": "content_block_stop", "index": 2}`,
				`{"type": "content_block_start", "index": 3, "content_block": {"type": "text", "text": ""}}`,
				`{"type": "content_block_delta", "index": 3, "delta": {"type": "text_delta", "text": "The weather in Los Angeles"}}`,
				`{"type": "content_block_delta", "index": 3, "delta": {"type": "text_delta", "text": " is 85 degrees Fahrenheit!"}}`,
				`{"type": "content_block_stop", "index": 3"}`,
				`{"type": "message_stop"}`,
			},
			expected: Juglow.BetaMessage{Content: []Juglow.BetaContentBlockUnion{
				{Type: "text", Text: "Let me look up the weather for you."},
				{Type: "thinking", Thinking: "I can look this up using a tool."},
				{Type: "tool_use", ID: "toolu_id", Name: "get_weather", Input: []byte(`{"city": "Los Angeles"}`)},
				{Type: "text", Text: "The weather in Los Angeles is 85 degrees Fahrenheit!"},
			}},
		},
		"fallback block relabels accumulated model": {
			events: []string{
				`{"type": "message_start", "message": {"model": "model-a"}}`,
				`{"type": "content_block_start", "index": 0, "content_block": {"type": "fallback", "from": {"model": "model-a"}, "to": {"model": "model-b"}}}`,
				`{"type": "content_block_stop", "index": 0}`,
				`{"type": "message_delta", "delta": {"stop_reason": "end_turn"}, "usage": {"output_tokens": 5}}`,
				`{"type": "message_stop"}`,
			},
			expected: Juglow.BetaMessage{
				Model: "model-b",
				Content: []Juglow.BetaContentBlockUnion{
					{
						Type: "fallback",
						From: Juglow.BetaFallbackInfo{Model: "model-a"},
						To:   Juglow.BetaFallbackInfo{Model: "model-b"},
					},
				},
				StopReason: "end_turn",
				Usage:      Juglow.BetaUsage{OutputTokens: 5},
			},
		},
		"interleaved content blocks": {
			events: []string{
				`{"type": "message_start", "message": {}}`,
				`{"type": "content_block_start", "index": 0, "content_block": {"type": "thinking", "thinking": ""}}`,
				`{"type": "content_block_delta", "index": 0, "delta": {"type": "thinking_delta", "thinking": "Let me think."}}`,
				`{"type": "content_block_delta", "index": 0, "delta": {"type": "signature_delta", "signature": "sig123"}}`,
				`{"type": "content_block_stop", "index": 0}`,
				`{"type": "content_block_start", "index": 1, "content_block": {"type": "text", "text": ""}}`,
				`{"type": "content_block_delta", "index": 1, "delta": {"type": "text_delta", "text": "Hello"}}`,
				`{"type": "content_block_start", "index": 2, "content_block": {"type": "tool_use", "id": "toolu_id", "name": "get_weather", "input": {}}}`,
				`{"type": "content_block_delta", "index": 1, "delta": {"type": "text_delta", "text": " world"}}`,
				`{"type": "content_block_delta", "index": 2, "delta": {"type": "input_json_delta", "partial_json": "{\"city\": "}}`,
				`{"type": "content_block_delta", "index": 1, "delta": {"type": "text_delta", "text": "!"}}`,
				`{"type": "content_block_delta", "index": 2, "delta": {"type": "input_json_delta", "partial_json": "\"Los Angeles\"}"}}`,
				`{"type": "content_block_stop", "index": 1}`,
				`{"type": "content_block_stop", "index": 2}`,
				`{"type": "message_stop"}`,
			},
			expected: Juglow.BetaMessage{Content: []Juglow.BetaContentBlockUnion{
				{Type: "thinking", Thinking: "Let me think.", Signature: "sig123"},
				{Type: "text", Text: "Hello world!"},
				{Type: "tool_use", ID: "toolu_id", Name: "get_weather", Input: []byte(`{"city": "Los Angeles"}`)},
			}},
		},
	} {
		t.Run(name, func(t *testing.T) {
			message := Juglow.BetaMessage{}
			for _, eventStr := range testCase.events {
				event := Juglow.BetaRawMessageStreamEventUnion{}
				err := (&event).UnmarshalJSON([]byte(eventStr))
				if err != nil {
					t.Fatal(err)
				}
				(&message).Accumulate(event)
			}
			marshaledMessage, err := json.Marshal(message)
			if err != nil {
				t.Fatal(err)
			}
			marshaledExpectedMessage, err := json.Marshal(testCase.expected)
			if err != nil {
				t.Fatal(err)
			}
			if string(marshaledMessage) != string(marshaledExpectedMessage) {
				t.Fatalf("Mismatched message: expected %s but got %s", marshaledExpectedMessage, marshaledMessage)
			}
		})
	}
}

func TestBetaAccumulateContentBlockIndexErrors(t *testing.T) {
	for name, testCase := range map[string]struct {
		events  []string
		wantErr string
	}{
		"start with an index gap": {
			events: []string{
				`{"type": "message_start", "message": {}}`,
				`{"type": "content_block_start", "index": 1, "content_block": {"type": "text", "text": ""}}`,
			},
			wantErr: "expected index 0",
		},
		"delta for a block that never started": {
			events: []string{
				`{"type": "message_start", "message": {}}`,
				`{"type": "content_block_start", "index": 0, "content_block": {"type": "text", "text": ""}}`,
				`{"type": "content_block_delta", "index": 1, "delta": {"type": "text_delta", "text": "hi"}}`,
			},
			wantErr: "only 1 content blocks",
		},
		"delta with a negative index": {
			events: []string{
				`{"type": "message_start", "message": {}}`,
				`{"type": "content_block_start", "index": 0, "content_block": {"type": "text", "text": ""}}`,
				`{"type": "content_block_delta", "index": -1, "delta": {"type": "text_delta", "text": "hi"}}`,
			},
			wantErr: "index -1",
		},
		"stop for a block that never started": {
			events: []string{
				`{"type": "message_start", "message": {}}`,
				`{"type": "content_block_stop", "index": 0}`,
			},
			wantErr: "only 0 content blocks",
		},
	} {
		t.Run(name, func(t *testing.T) {
			message := Juglow.BetaMessage{}
			for i, eventStr := range testCase.events {
				event := Juglow.BetaRawMessageStreamEventUnion{}
				if err := (&event).UnmarshalJSON([]byte(eventStr)); err != nil {
					t.Fatal(err)
				}
				err := (&message).Accumulate(event)
				if i < len(testCase.events)-1 {
					if err != nil {
						t.Fatal(err)
					}
					continue
				}
				if err == nil {
					t.Fatalf("expected error containing %q, got nil", testCase.wantErr)
				}
				if !strings.Contains(err.Error(), testCase.wantErr) {
					t.Fatalf("expected error containing %q, got %q", testCase.wantErr, err.Error())
				}
			}
		})
	}
}
