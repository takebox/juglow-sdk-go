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

func TestBetaSessionEventListWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Sessions.Events.List(
		context.TODO(),
		"sesn_011CZkZAtmR3yMPDzynEDxu7",
		Juglow.BetaSessionEventListParams{
			CreatedAtGt:  Juglow.Time(time.Now()),
			CreatedAtGte: Juglow.Time(time.Now()),
			CreatedAtLt:  Juglow.Time(time.Now()),
			CreatedAtLte: Juglow.Time(time.Now()),
			Limit:        Juglow.Int(0),
			Order:        Juglow.BetaSessionEventListParamsOrderAsc,
			Page:         Juglow.String("page"),
			Types:        []string{"string"},
			Betas:        []Juglow.JuglowBeta{Juglow.JuglowBetaMessageBatches2024_09_24},
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

func TestBetaSessionEventSendWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Sessions.Events.Send(
		context.TODO(),
		"sesn_011CZkZAtmR3yMPDzynEDxu7",
		Juglow.BetaSessionEventSendParams{
			Events: []Juglow.BetaManagedAgentsEventParamsUnion{{
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
