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

func TestBetaDreamNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Dreams.New(context.TODO(), Juglow.BetaDreamNewParams{
		Inputs: []Juglow.BetaDreamInputUnionParam{{
			OfMemoryStore: &Juglow.BetaDreamMemoryStoreInputParam{
				MemoryStoreID: "x",
				Type:          Juglow.BetaDreamMemoryStoreInputTypeMemoryStore,
			},
		}},
		Model: Juglow.BetaDreamNewParamsModelUnion{
			OfString: Juglow.String("string"),
		},
		Instructions: Juglow.String("x"),
		Betas:        []Juglow.JuglowBeta{Juglow.JuglowBetaMessageBatches2024_09_24},
	})
	if err != nil {
		var apierr *Juglow.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaDreamGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Dreams.Get(
		context.TODO(),
		"dream_id",
		Juglow.BetaDreamGetParams{
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

func TestBetaDreamListWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Dreams.List(context.TODO(), Juglow.BetaDreamListParams{
		CreatedAtGt:     Juglow.Time(time.Now()),
		CreatedAtLt:     Juglow.Time(time.Now()),
		IncludeArchived: Juglow.Bool(true),
		Limit:           Juglow.Int(0),
		Page:            Juglow.String("page"),
		Statuses:        []Juglow.BetaDreamStatus{Juglow.BetaDreamStatusPending},
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

func TestBetaDreamArchiveWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Dreams.Archive(
		context.TODO(),
		"dream_id",
		Juglow.BetaDreamArchiveParams{
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

func TestBetaDreamCancelWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Dreams.Cancel(
		context.TODO(),
		"dream_id",
		Juglow.BetaDreamCancelParams{
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
