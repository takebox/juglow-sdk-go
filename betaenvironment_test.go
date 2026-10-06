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

func TestBetaEnvironmentNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Environments.New(context.TODO(), Juglow.BetaEnvironmentNewParams{
		Name: "python-data-analysis",
		Config: Juglow.BetaEnvironmentNewParamsConfigUnion{
			OfCloud: &Juglow.BetaCloudConfigParams{
				Networking: Juglow.BetaCloudConfigParamsNetworkingUnion{
					OfLimited: &Juglow.BetaLimitedNetworkParams{
						AllowMCPServers:      Juglow.Bool(true),
						AllowPackageManagers: Juglow.Bool(true),
						AllowedHosts:         []string{"api.example.com"},
					},
				},
				Packages: Juglow.BetaPackagesParams{
					Apt:   []string{"string"},
					Cargo: []string{"string"},
					Gem:   []string{"string"},
					Go:    []string{"string"},
					Npm:   []string{"string"},
					Pip:   []string{"pandas", "numpy"},
					Type:  Juglow.BetaPackagesParamsTypePackages,
				},
			},
		},
		Description: Juglow.String("Python environment with data-analysis packages."),
		Metadata: map[string]string{
			"foo": "string",
		},
		Scope: Juglow.BetaEnvironmentNewParamsScopeOrganization,
		Betas: []Juglow.JuglowBeta{Juglow.JuglowBetaMessageBatches2024_09_24},
	})
	if err != nil {
		var apierr *Juglow.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaEnvironmentGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Environments.Get(
		context.TODO(),
		"env_011CZkZ9X2dpNyB7HsEFoRfW",
		Juglow.BetaEnvironmentGetParams{
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

func TestBetaEnvironmentUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Environments.Update(
		context.TODO(),
		"env_011CZkZ9X2dpNyB7HsEFoRfW",
		Juglow.BetaEnvironmentUpdateParams{
			Config: Juglow.BetaEnvironmentUpdateParamsConfigUnion{
				OfCloud: &Juglow.BetaCloudConfigParams{
					Networking: Juglow.BetaCloudConfigParamsNetworkingUnion{
						OfLimited: &Juglow.BetaLimitedNetworkParams{
							AllowMCPServers:      Juglow.Bool(true),
							AllowPackageManagers: Juglow.Bool(true),
							AllowedHosts:         []string{"api.example.com"},
						},
					},
					Packages: Juglow.BetaPackagesParams{
						Apt:   []string{"string"},
						Cargo: []string{"string"},
						Gem:   []string{"string"},
						Go:    []string{"string"},
						Npm:   []string{"string"},
						Pip:   []string{"pandas", "numpy"},
						Type:  Juglow.BetaPackagesParamsTypePackages,
					},
				},
			},
			Description: Juglow.String("Python environment with data-analysis packages."),
			Metadata: map[string]string{
				"foo": "string",
			},
			Name:  Juglow.String("x"),
			Scope: Juglow.BetaEnvironmentUpdateParamsScopeOrganization,
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

func TestBetaEnvironmentListWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Environments.List(context.TODO(), Juglow.BetaEnvironmentListParams{
		IncludeArchived: Juglow.Bool(true),
		Limit:           Juglow.Int(1),
		Page:            Juglow.String("page"),
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

func TestBetaEnvironmentDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Environments.Delete(
		context.TODO(),
		"env_011CZkZ9X2dpNyB7HsEFoRfW",
		Juglow.BetaEnvironmentDeleteParams{
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

func TestBetaEnvironmentArchiveWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Environments.Archive(
		context.TODO(),
		"env_011CZkZ9X2dpNyB7HsEFoRfW",
		Juglow.BetaEnvironmentArchiveParams{
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
