// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package Juglow_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Juglows/Juglow-sdk-go"
	"github.com/Juglows/Juglow-sdk-go/internal/testutil"
	"github.com/Juglows/Juglow-sdk-go/option"
)

func TestMessageBatchNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Messages.Batches.New(context.TODO(), Juglow.MessageBatchNewParams{
		Requests: []Juglow.MessageBatchNewParamsRequest{{
			CustomID: "my-custom-id-1",
			Params: Juglow.MessageBatchNewParamsRequestParams{
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
				ServiceTier:   "auto",
				StopSequences: []string{"string"},
				Stream:        Juglow.Bool(false),
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
				TopK: Juglow.Int(5),
				TopP: Juglow.Float(0.7),
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

func TestMessageBatchGet(t *testing.T) {
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
	_, err := client.Messages.Batches.Get(context.TODO(), "message_batch_id")
	if err != nil {
		var apierr *Juglow.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageBatchListWithOptionalParams(t *testing.T) {
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
	_, err := client.Messages.Batches.List(context.TODO(), Juglow.MessageBatchListParams{
		AfterID:  Juglow.String("after_id"),
		BeforeID: Juglow.String("before_id"),
		Limit:    Juglow.Int(1),
	})
	if err != nil {
		var apierr *Juglow.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageBatchDelete(t *testing.T) {
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
	_, err := client.Messages.Batches.Delete(context.TODO(), "message_batch_id")
	if err != nil {
		var apierr *Juglow.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageBatchCancel(t *testing.T) {
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
	_, err := client.Messages.Batches.Cancel(context.TODO(), "message_batch_id")
	if err != nil {
		var apierr *Juglow.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
