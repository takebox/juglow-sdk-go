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

func TestBetaDeploymentNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Deployments.New(context.TODO(), Juglow.BetaDeploymentNewParams{
		Agent: Juglow.BetaDeploymentNewParamsAgentUnion{
			OfString: Juglow.String("string"),
		},
		EnvironmentID: "x",
		InitialEvents: []Juglow.BetaManagedAgentsDeploymentInitialEventParamsUnion{{
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
		Name: "x",
		Budget: Juglow.BetaManagedAgentsBudgetLimitParam{
			MaxListCost: Juglow.BetaMonetaryAmountParam{
				Amount:   "2500",
				Currency: Juglow.BetaCurrencyUsd,
			},
			Type: Juglow.BetaManagedAgentsBudgetLimitTypeLimit,
		},
		Description: Juglow.String("description"),
		Metadata: map[string]string{
			"foo": "string",
		},
		Resources: []Juglow.BetaDeploymentNewParamsResourceUnion{{
			OfFile: &Juglow.BetaManagedAgentsFileResourceParams{
				FileID:    "file_011CNha8iCJcU1wXNR6q4V8w",
				Type:      Juglow.BetaManagedAgentsFileResourceParamsTypeFile,
				MountPath: Juglow.String("/uploads/receipt.pdf"),
			},
		}},
		Schedule: Juglow.BetaManagedAgentsScheduleParams{
			Expression: "0 9 * * 1-5",
			Timezone:   "America/Los_Angeles",
			Type:       Juglow.BetaManagedAgentsScheduleParamsTypeCron,
		},
		VaultIDs: []string{"string"},
		Betas:    []Juglow.JuglowBeta{Juglow.JuglowBetaMessageBatches2024_09_24},
	})
	if err != nil {
		var apierr *Juglow.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaDeploymentGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Deployments.Get(
		context.TODO(),
		"depl_011CZkZcDH3vPqd7xnEfwTai",
		Juglow.BetaDeploymentGetParams{
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

func TestBetaDeploymentUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Deployments.Update(
		context.TODO(),
		"depl_011CZkZcDH3vPqd7xnEfwTai",
		Juglow.BetaDeploymentUpdateParams{
			Agent: Juglow.BetaDeploymentUpdateParamsAgentUnion{
				OfString: Juglow.String("string"),
			},
			Budget: Juglow.BetaManagedAgentsBudgetLimitParam{
				MaxListCost: Juglow.BetaMonetaryAmountParam{
					Amount:   "2500",
					Currency: Juglow.BetaCurrencyUsd,
				},
				Type: Juglow.BetaManagedAgentsBudgetLimitTypeLimit,
			},
			Description:   Juglow.String("description"),
			EnvironmentID: Juglow.String("environment_id"),
			InitialEvents: []Juglow.BetaManagedAgentsDeploymentInitialEventParamsUnion{{
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
			Metadata: map[string]string{
				"foo": "string",
			},
			Name: Juglow.String("name"),
			Resources: []Juglow.BetaDeploymentUpdateParamsResourceUnion{{
				OfFile: &Juglow.BetaManagedAgentsFileResourceParams{
					FileID:    "file_011CNha8iCJcU1wXNR6q4V8w",
					Type:      Juglow.BetaManagedAgentsFileResourceParamsTypeFile,
					MountPath: Juglow.String("/uploads/receipt.pdf"),
				},
			}},
			Schedule: Juglow.BetaManagedAgentsScheduleParams{
				Expression: "0 9 * * 1-5",
				Timezone:   "America/Los_Angeles",
				Type:       Juglow.BetaManagedAgentsScheduleParamsTypeCron,
			},
			VaultIDs: []string{"string"},
			Betas:    []Juglow.JuglowBeta{Juglow.JuglowBetaMessageBatches2024_09_24},
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

func TestBetaDeploymentListWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Deployments.List(context.TODO(), Juglow.BetaDeploymentListParams{
		AgentID:         Juglow.String("agent_id"),
		CreatedAtGte:    Juglow.Time(time.Now()),
		CreatedAtLte:    Juglow.Time(time.Now()),
		IncludeArchived: Juglow.Bool(true),
		Limit:           Juglow.Int(0),
		Page:            Juglow.String("page"),
		Status:          Juglow.BetaManagedAgentsDeploymentStatusActive,
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

func TestBetaDeploymentArchiveWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Deployments.Archive(
		context.TODO(),
		"depl_011CZkZcDH3vPqd7xnEfwTai",
		Juglow.BetaDeploymentArchiveParams{
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

func TestBetaDeploymentPauseWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Deployments.Pause(
		context.TODO(),
		"depl_011CZkZcDH3vPqd7xnEfwTai",
		Juglow.BetaDeploymentPauseParams{
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

func TestBetaDeploymentRunWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Deployments.Run(
		context.TODO(),
		"depl_011CZkZcDH3vPqd7xnEfwTai",
		Juglow.BetaDeploymentRunParams{
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

func TestBetaDeploymentUnpauseWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Deployments.Unpause(
		context.TODO(),
		"depl_011CZkZcDH3vPqd7xnEfwTai",
		Juglow.BetaDeploymentUnpauseParams{
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
