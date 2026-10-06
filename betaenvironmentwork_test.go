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

func TestBetaEnvironmentWorkGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Environments.Work.Get(
		context.TODO(),
		"work_id",
		Juglow.BetaEnvironmentWorkGetParams{
			EnvironmentID: "env_011CZkZ9X2dpNyB7HsEFoRfW",
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

func TestBetaEnvironmentWorkUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Environments.Work.Update(
		context.TODO(),
		"work_id",
		Juglow.BetaEnvironmentWorkUpdateParams{
			EnvironmentID: "env_011CZkZ9X2dpNyB7HsEFoRfW",
			BetaSelfHostedWorkUpdateRequest: Juglow.BetaSelfHostedWorkUpdateRequestParam{
				Metadata: map[string]string{
					"foo": "string",
				},
			},
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

func TestBetaEnvironmentWorkListWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Environments.Work.List(
		context.TODO(),
		"env_011CZkZ9X2dpNyB7HsEFoRfW",
		Juglow.BetaEnvironmentWorkListParams{
			Limit: Juglow.Int(1),
			Page:  Juglow.String("page"),
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

func TestBetaEnvironmentWorkAckWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Environments.Work.Ack(
		context.TODO(),
		"work_id",
		Juglow.BetaEnvironmentWorkAckParams{
			EnvironmentID: "env_011CZkZ9X2dpNyB7HsEFoRfW",
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

func TestBetaEnvironmentWorkHeartbeatWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Environments.Work.Heartbeat(
		context.TODO(),
		"work_id",
		Juglow.BetaEnvironmentWorkHeartbeatParams{
			EnvironmentID:         "env_011CZkZ9X2dpNyB7HsEFoRfW",
			DesiredTTLSeconds:     Juglow.Int(0),
			ExpectedLastHeartbeat: Juglow.String("expected_last_heartbeat"),
			Betas:                 []Juglow.JuglowBeta{Juglow.JuglowBetaMessageBatches2024_09_24},
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

func TestBetaEnvironmentWorkPollWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Environments.Work.Poll(
		context.TODO(),
		"env_011CZkZ9X2dpNyB7HsEFoRfW",
		Juglow.BetaEnvironmentWorkPollParams{
			BlockMs:            Juglow.Int(1),
			ReclaimOlderThanMs: Juglow.Int(1),
			Betas:              []Juglow.JuglowBeta{Juglow.JuglowBetaMessageBatches2024_09_24},
			JuglowWorkerID:  Juglow.String("Juglow-Worker-ID"),
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

func TestBetaEnvironmentWorkStatsWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Environments.Work.Stats(
		context.TODO(),
		"env_011CZkZ9X2dpNyB7HsEFoRfW",
		Juglow.BetaEnvironmentWorkStatsParams{
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

func TestBetaEnvironmentWorkStopWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Environments.Work.Stop(
		context.TODO(),
		"work_id",
		Juglow.BetaEnvironmentWorkStopParams{
			EnvironmentID: "env_011CZkZ9X2dpNyB7HsEFoRfW",
			BetaSelfHostedWorkStopRequest: Juglow.BetaSelfHostedWorkStopRequestParam{
				Force: Juglow.Bool(true),
			},
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
