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

func TestCompletionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Completions.New(context.TODO(), Juglow.CompletionNewParams{
		MaxTokensToSample: 256,
		Model:             Juglow.ModelHaijunSonnet5,
		Prompt:            "\n\nHuman: Hello, world!\n\nAssistant:",
		Metadata: Juglow.MetadataParam{
			UserID: Juglow.String("13803d75-b4b5-4c3e-b2a2-6f21399b021b"),
		},
		StopSequences: []string{"string"},
		Temperature:   Juglow.Float(1),
		TopK:          Juglow.Int(5),
		TopP:          Juglow.Float(0.7),
		Betas:         []Juglow.JuglowBeta{Juglow.JuglowBetaMessageBatches2024_09_24},
	})
	if err != nil {
		var apierr *Juglow.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
