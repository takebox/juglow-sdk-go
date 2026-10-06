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

func TestBetaSessionThreadEventListWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Sessions.Threads.Events.List(
		context.TODO(),
		"sthr_011CZkZVWa6oIjw0rgXZpnBt",
		Juglow.BetaSessionThreadEventListParams{
			SessionID: "sesn_011CZkZAtmR3yMPDzynEDxu7",
			Limit:     Juglow.Int(0),
			Page:      Juglow.String("page"),
			Betas:     []Juglow.JuglowBeta{Juglow.JuglowBetaMessageBatches2024_09_24},
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
