// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package Juglow_test

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/Juglows/Juglow-sdk-go"
	"github.com/Juglows/Juglow-sdk-go/internal/testutil"
	"github.com/Juglows/Juglow-sdk-go/option"
	"github.com/Juglows/Juglow-sdk-go/shared/constant"
)

func TestMessageNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Messages.New(context.TODO(), Juglow.MessageNewParams{
		MaxTokens: 1024,
		Messages: []Juglow.MessageParam{{
			Content: []Juglow.ContentBlockParamUnion{{
				OfText: &Juglow.TextBlockParam{
					Text: "x",
					CacheControl: Juglow.CacheControlEphemeralParam{
						TTL: Juglow.CacheControlEphemeralTTLTTL5m,
					},
					Citations: []Juglow.TextCitationParamUnion{{
						OfCharLocation: &Juglow.CitationCharLocationParam{
							CitedText:      "The grass is green. The sky is blue.",
							DocumentIndex:  0,
							DocumentTitle:  Juglow.String("x"),
							EndCharIndex:   0,
							StartCharIndex: 0,
						},
					}},
				},
			}},
			Role: Juglow.MessageParamRoleUser,
		}},
		Model: Juglow.ModelHaijunOpus4_6,
		CacheControl: Juglow.CacheControlEphemeralParam{
			TTL: Juglow.CacheControlEphemeralTTLTTL5m,
		},
		Container:    Juglow.String("container"),
		InferenceGeo: Juglow.String("inference_geo"),
		Metadata: Juglow.MetadataParam{
			UserID: Juglow.String("13803d75-b4b5-4c3e-b2a2-6f21399b021b"),
		},
		OutputConfig: Juglow.OutputConfigParam{
			Effort: Juglow.OutputConfigEffortLow,
			Format: Juglow.JSONOutputFormatParam{
				Schema: map[string]any{
					"foo": "bar",
				},
			},
		},
		ServiceTier:   Juglow.MessageNewParamsServiceTierAuto,
		StopSequences: []string{"string"},
		System: []Juglow.TextBlockParam{{
			Text: "Today's date is 2024-06-01.",
			CacheControl: Juglow.CacheControlEphemeralParam{
				TTL: Juglow.CacheControlEphemeralTTLTTL5m,
			},
			Citations: []Juglow.TextCitationParamUnion{{
				OfCharLocation: &Juglow.CitationCharLocationParam{
					CitedText:      "The grass is green. The sky is blue.",
					DocumentIndex:  0,
					DocumentTitle:  Juglow.String("x"),
					EndCharIndex:   0,
					StartCharIndex: 0,
				},
			}},
		}},
		Temperature: Juglow.Float(1),
		Thinking: Juglow.ThinkingConfigParamUnion{
			OfAdaptive: &Juglow.ThinkingConfigAdaptiveParam{
				Display: Juglow.ThinkingConfigAdaptiveDisplaySummarized,
			},
		},
		ToolChoice: Juglow.ToolChoiceUnionParam{
			OfAuto: &Juglow.ToolChoiceAutoParam{
				DisableParallelToolUse: Juglow.Bool(true),
			},
		},
		Tools: []Juglow.ToolUnionParam{{
			OfTool: &Juglow.ToolParam{
				InputSchema: Juglow.ToolInputSchemaParam{
					Properties: map[string]any{
						"location": "bar",
						"unit":     "bar",
					},
					Required: []string{"location"},
				},
				Name:           "name",
				AllowedCallers: []string{"direct"},
				CacheControl: Juglow.CacheControlEphemeralParam{
					TTL: Juglow.CacheControlEphemeralTTLTTL5m,
				},
				DeferLoading:        Juglow.Bool(true),
				Description:         Juglow.String("Get the current weather in a given location"),
				EagerInputStreaming: Juglow.Bool(true),
				InputExamples: []map[string]any{{
					"foo": "bar",
				}},
				Strict: Juglow.Bool(true),
				Type:   Juglow.ToolTypeCustom,
			},
		}},
		TopK:          Juglow.Int(5),
		TopP:          Juglow.Float(0.7),
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

func TestMessageCountTokensWithOptionalParams(t *testing.T) {
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
	_, err := client.Messages.CountTokens(context.TODO(), Juglow.MessageCountTokensParams{
		Messages: []Juglow.MessageParam{{
			Content: []Juglow.ContentBlockParamUnion{{
				OfText: &Juglow.TextBlockParam{
					Text: "x",
					CacheControl: Juglow.CacheControlEphemeralParam{
						TTL: Juglow.CacheControlEphemeralTTLTTL5m,
					},
					Citations: []Juglow.TextCitationParamUnion{{
						OfCharLocation: &Juglow.CitationCharLocationParam{
							CitedText:      "The grass is green. The sky is blue.",
							DocumentIndex:  0,
							DocumentTitle:  Juglow.String("x"),
							EndCharIndex:   0,
							StartCharIndex: 0,
						},
					}},
				},
			}},
			Role: Juglow.MessageParamRoleUser,
		}},
		Model: Juglow.ModelHaijunOpus4_6,
		CacheControl: Juglow.CacheControlEphemeralParam{
			TTL: Juglow.CacheControlEphemeralTTLTTL5m,
		},
		OutputConfig: Juglow.OutputConfigParam{
			Effort: Juglow.OutputConfigEffortLow,
			Format: Juglow.JSONOutputFormatParam{
				Schema: map[string]any{
					"foo": "bar",
				},
			},
		},
		System: Juglow.MessageCountTokensParamsSystemUnion{
			OfTextBlockArray: []Juglow.TextBlockParam{{
				Text: "Today's date is 2024-06-01.",
				CacheControl: Juglow.CacheControlEphemeralParam{
					TTL: Juglow.CacheControlEphemeralTTLTTL5m,
				},
				Citations: []Juglow.TextCitationParamUnion{{
					OfCharLocation: &Juglow.CitationCharLocationParam{
						CitedText:      "The grass is green. The sky is blue.",
						DocumentIndex:  0,
						DocumentTitle:  Juglow.String("x"),
						EndCharIndex:   0,
						StartCharIndex: 0,
					},
				}},
			}},
		},
		Thinking: Juglow.ThinkingConfigParamUnion{
			OfAdaptive: &Juglow.ThinkingConfigAdaptiveParam{
				Display: Juglow.ThinkingConfigAdaptiveDisplaySummarized,
			},
		},
		ToolChoice: Juglow.ToolChoiceUnionParam{
			OfAuto: &Juglow.ToolChoiceAutoParam{
				DisableParallelToolUse: Juglow.Bool(true),
			},
		},
		Tools: []Juglow.MessageCountTokensToolUnionParam{{
			OfTool: &Juglow.ToolParam{
				InputSchema: Juglow.ToolInputSchemaParam{
					Properties: map[string]any{
						"location": "bar",
						"unit":     "bar",
					},
					Required: []string{"location"},
				},
				Name:           "name",
				AllowedCallers: []string{"direct"},
				CacheControl: Juglow.CacheControlEphemeralParam{
					TTL: Juglow.CacheControlEphemeralTTLTTL5m,
				},
				DeferLoading:        Juglow.Bool(true),
				Description:         Juglow.String("Get the current weather in a given location"),
				EagerInputStreaming: Juglow.Bool(true),
				InputExamples: []map[string]any{{
					"foo": "bar",
				}},
				Strict: Juglow.Bool(true),
				Type:   Juglow.ToolTypeCustom,
			},
		}},
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

func TestAccumulate(t *testing.T) {
	for name, testCase := range map[string]struct {
		expected Juglow.Message
		events   []string
	}{
		"empty message": {
			expected: Juglow.Message{Usage: Juglow.Usage{}},
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
			expected: Juglow.Message{Content: []Juglow.ContentBlockUnion{
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
			expected: Juglow.Message{Content: []Juglow.ContentBlockUnion{
				{Type: "text", Text: "1 + 1 = 2", Citations: []Juglow.TextCitationUnion{{
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
			expected: Juglow.Message{Content: []Juglow.ContentBlockUnion{
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
			expected: Juglow.Message{Content: []Juglow.ContentBlockUnion{
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
			expected: Juglow.Message{Content: []Juglow.ContentBlockUnion{
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
			expected: Juglow.Message{Content: []Juglow.ContentBlockUnion{
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
			expected: Juglow.Message{Content: []Juglow.ContentBlockUnion{
				{Type: "redacted_thinking", Data: "Redacted"},
			}},
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
			expected: Juglow.Message{Content: []Juglow.ContentBlockUnion{
				{Type: "text", Text: "Let me look up the weather for you."},
				{Type: "thinking", Thinking: "I can look this up using a tool."},
				{Type: "tool_use", ID: "toolu_id", Name: "get_weather", Input: []byte(`{"city": "Los Angeles"}`)},
				{Type: "text", Text: "The weather in Los Angeles is 85 degrees Fahrenheit!"},
			}},
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
			expected: Juglow.Message{Content: []Juglow.ContentBlockUnion{
				{Type: "thinking", Thinking: "Let me think.", Signature: "sig123"},
				{Type: "text", Text: "Hello world!"},
				{Type: "tool_use", ID: "toolu_id", Name: "get_weather", Input: []byte(`{"city": "Los Angeles"}`)},
			}},
		},
	} {
		t.Run(name, func(t *testing.T) {
			message := Juglow.Message{}
			for _, eventStr := range testCase.events {
				event := Juglow.MessageStreamEventUnion{}
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

func TestAccumulateContentBlockIndexErrors(t *testing.T) {
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
			message := Juglow.Message{}
			for i, eventStr := range testCase.events {
				event := Juglow.MessageStreamEventUnion{}
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

func TestMessageNewWithNonStreamingTimeoutLimits(t *testing.T) {
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

	originalModelTokenLimits := constant.ModelNonStreamingTokens
	defer func() { constant.ModelNonStreamingTokens = originalModelTokenLimits }()
	constant.ModelNonStreamingTokens = map[string]int{"test-model": 8192}

	model := Juglow.Model("test-model")
	testModelLimit := constant.ModelNonStreamingTokens[string(model)]

	// This test verifies that we can still create a message with tokens below the limit
	safeParams := Juglow.MessageNewParams{
		MaxTokens: int64(testModelLimit - 1000), // Well below the limit
		Messages: []Juglow.MessageParam{{
			Content: []Juglow.ContentBlockParamUnion{{
				OfText: &Juglow.TextBlockParam{Text: "What is a quaternion?"},
			}},
			Role: Juglow.MessageParamRoleUser,
		}},
		Model: model,
	}

	_, err := client.Messages.New(context.TODO(), safeParams)
	if err != nil {
		var apierr *Juglow.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("Expected no error for tokens below limit, got: %v", err)
	}

	// This test verifies that we get an error when exceeding the limit
	unsafeParams := Juglow.MessageNewParams{
		MaxTokens: int64(testModelLimit + 1000), // Exceed the limit
		Messages: []Juglow.MessageParam{{
			Content: []Juglow.ContentBlockParamUnion{{
				OfText: &Juglow.TextBlockParam{Text: "What is a quaternion?"},
			}},
			Role: Juglow.MessageParamRoleUser,
		}},
		Model: model,
	}

	_, err = client.Messages.New(context.TODO(), unsafeParams)
	if err == nil {
		t.Fatal("Expected error for tokens above limit, got nil")
	}
}

func TestCalculateNonStreamingTimeout(t *testing.T) {
	// Store original model token limits to restore after test
	originalModelTokenLimits := make(map[string]int)
	for k, v := range constant.ModelNonStreamingTokens {
		originalModelTokenLimits[k] = v
	}
	defer func() {
		// Restore original model token limits
		constant.ModelNonStreamingTokens = originalModelTokenLimits
	}()

	// Set up a test model for consistent testing
	constant.ModelNonStreamingTokens = map[string]int{
		"test-model": 8192,
	}
	defaultTimeout := 10 * time.Minute

	tests := []struct {
		name          string
		maxTokens     int
		model         string
		expectTimeout time.Duration
		expectError   bool
		opts          []option.RequestOption
	}{
		{
			name:          "small token count returns default timeout",
			maxTokens:     1000,
			model:         "any-model",
			expectTimeout: defaultTimeout,
			expectError:   false,
		},
		{
			name:          "large token count above default time limit throws error",
			maxTokens:     100000,
			model:         "any-model",
			expectTimeout: 0,
			expectError:   true,
		},
		{
			name:          "token count above model specific limit throws error",
			maxTokens:     9000,
			model:         "test-model",
			expectTimeout: 0,
			expectError:   true,
		},
		{
			name:          "token count below model specific limit is ok",
			maxTokens:     8000,
			model:         "test-model",
			expectTimeout: defaultTimeout,
			expectError:   false,
		},
		{
			name:          "user supplies custom timeout so we'll use that instead",
			maxTokens:     1000,
			model:         "test-model",
			expectTimeout: 30 * time.Second, // Custom timeout
			expectError:   false,
			opts: []option.RequestOption{
				option.WithRequestTimeout(30 * time.Second),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			timeout, err := Juglow.CalculateNonStreamingTimeout(tc.maxTokens, Juglow.Model(tc.model), tc.opts)

			if tc.expectError && err == nil {
				t.Error("Expected error but got nil")
			}

			if !tc.expectError && err != nil {
				t.Errorf("Did not expect error but got: %v", err)
			}

			if timeout != tc.expectTimeout {
				t.Errorf("Expected timeout %v but got %v", tc.expectTimeout, timeout)
			}
		})
	}
}

// Test specific model limits
func TestModelLimits(t *testing.T) {
	// Verify the model limits are defined for opus-4 models
	if _, exists := constant.ModelNonStreamingTokens["haijun-opus-4-20250514"]; !exists {
		t.Error("Expected model limit for haijun-opus-4-20250514 but not found")
	}

	if _, exists := constant.ModelNonStreamingTokens["Juglow.haijun-opus-4-20250514-v1:0"]; !exists {
		t.Error("Expected model limit for Juglow.haijun-opus-4-20250514-v1:0 but not found")
	}

	if _, exists := constant.ModelNonStreamingTokens["haijun-opus-4@20250514"]; !exists {
		t.Error("Expected model limit for haijun-opus-4@20250514 but not found")
	}
}

func TestToolResultBlockParamStringContent(t *testing.T) {
	toolResultJSON := `{"type":"tool_result","content":"error message","tool_use_id":"123"}`
	var toolResult Juglow.ToolResultBlockParam
	err := json.Unmarshal([]byte(toolResultJSON), &toolResult)
	if err != nil {
		t.Fatal(err)
	}
	if len(toolResult.Content) != 1 || toolResult.Content[0].OfText.Text != "error message" {
		t.Error("String content not converted to TextBlock")
	}
}

func TestMessageParamStringContent(t *testing.T) {
	messageJSON := `{"role":"user","content":"hello world"}`
	var message Juglow.MessageParam
	err := json.Unmarshal([]byte(messageJSON), &message)
	if err != nil {
		t.Fatal(err)
	}
	if len(message.Content) != 1 || message.Content[0].OfText.Text != "hello world" {
		t.Error("String content not converted to TextBlock")
	}
}

func TestMessageParamArrayContent(t *testing.T) {
	messageJSON := `{"role":"user","content":[{"type":"text","text":"first block"},{"type":"text","text":"second block"}]}`
	var message Juglow.MessageParam
	err := json.Unmarshal([]byte(messageJSON), &message)
	if err != nil {
		t.Fatal(err)
	}
	if len(message.Content) != 2 {
		t.Errorf("Expected 2 content blocks, got %d", len(message.Content))
	}
	if message.Content[0].OfText.Text != "first block" {
		t.Errorf("Expected first block text 'first block', got '%s'", message.Content[0].OfText.Text)
	}
	if message.Content[1].OfText.Text != "second block" {
		t.Errorf("Expected second block text 'second block', got '%s'", message.Content[1].OfText.Text)
	}
}
