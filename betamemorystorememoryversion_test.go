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

func TestBetaMemoryStoreMemoryVersionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.MemoryStores.MemoryVersions.Get(
		context.TODO(),
		"memory_version_id",
		Juglow.BetaMemoryStoreMemoryVersionGetParams{
			MemoryStoreID: "memory_store_id",
			View:          Juglow.BetaManagedAgentsMemoryViewBasic,
			Betas:         []Juglow.JuglowBeta{Juglow.JuglowBetaMessageBatches2024_09_24},
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

func TestBetaMemoryStoreMemoryVersionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.MemoryStores.MemoryVersions.List(
		context.TODO(),
		"memory_store_id",
		Juglow.BetaMemoryStoreMemoryVersionListParams{
			APIKeyID:     Juglow.String("api_key_id"),
			CreatedAtGte: Juglow.Time(time.Now()),
			CreatedAtLte: Juglow.Time(time.Now()),
			Limit:        Juglow.Int(0),
			MemoryID:     Juglow.String("memory_id"),
			Operation:    Juglow.BetaManagedAgentsMemoryVersionOperationCreated,
			Page:         Juglow.String("page"),
			SessionID:    Juglow.String("session_id"),
			View:         Juglow.BetaManagedAgentsMemoryViewBasic,
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

func TestBetaMemoryStoreMemoryVersionRedactWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.MemoryStores.MemoryVersions.Redact(
		context.TODO(),
		"memory_version_id",
		Juglow.BetaMemoryStoreMemoryVersionRedactParams{
			MemoryStoreID: "memory_store_id",
			Betas:         []Juglow.JuglowBeta{Juglow.JuglowBetaMessageBatches2024_09_24},
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
