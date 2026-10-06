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
)

func TestBetaSessionResourceGetWithOptionalParams(t *testing.T) {
	t.Skip("prism can't find endpoint with beta only tag")
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
	_, err := client.Beta.Sessions.Resources.Get(
		context.TODO(),
		"sesrsc_011CZkZBJq5dWxk9fVLNcPht",
		Juglow.BetaSessionResourceGetParams{
			SessionID: "sesn_011CZkZAtmR3yMPDzynEDxu7",
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

func TestBetaSessionResourceUpdateWithOptionalParams(t *testing.T) {
	t.Skip("prism can't find endpoint with beta only tag")
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
	_, err := client.Beta.Sessions.Resources.Update(
		context.TODO(),
		"sesrsc_011CZkZBJq5dWxk9fVLNcPht",
		Juglow.BetaSessionResourceUpdateParams{
			SessionID:          "sesn_011CZkZAtmR3yMPDzynEDxu7",
			AuthorizationToken: "ghp_exampletoken",
			Betas:              []Juglow.JuglowBeta{Juglow.JuglowBetaMessageBatches2024_09_24},
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

func TestBetaSessionResourceListWithOptionalParams(t *testing.T) {
	t.Skip("prism can't find endpoint with beta only tag")
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
	_, err := client.Beta.Sessions.Resources.List(
		context.TODO(),
		"sesn_011CZkZAtmR3yMPDzynEDxu7",
		Juglow.BetaSessionResourceListParams{
			Limit: Juglow.Int(0),
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

func TestBetaSessionResourceDeleteWithOptionalParams(t *testing.T) {
	t.Skip("prism can't find endpoint with beta only tag")
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
	_, err := client.Beta.Sessions.Resources.Delete(
		context.TODO(),
		"sesrsc_011CZkZBJq5dWxk9fVLNcPht",
		Juglow.BetaSessionResourceDeleteParams{
			SessionID: "sesn_011CZkZAtmR3yMPDzynEDxu7",
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

func TestBetaSessionResourceAddWithOptionalParams(t *testing.T) {
	t.Skip("prism can't find endpoint with beta only tag")
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
	_, err := client.Beta.Sessions.Resources.Add(
		context.TODO(),
		"sesn_011CZkZAtmR3yMPDzynEDxu7",
		Juglow.BetaSessionResourceAddParams{
			BetaManagedAgentsFileResourceParams: Juglow.BetaManagedAgentsFileResourceParams{
				FileID:    "file_011CNha8iCJcU1wXNR6q4V8w",
				Type:      Juglow.BetaManagedAgentsFileResourceParamsTypeFile,
				MountPath: Juglow.String("/uploads/receipt.pdf"),
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
