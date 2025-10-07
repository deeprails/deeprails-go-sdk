// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomdeeprailsdeeprailsgosdk_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/deeprails/deeprails-go-sdk"
	"github.com/deeprails/deeprails-go-sdk/internal/testutil"
	"github.com/deeprails/deeprails-go-sdk/option"
)

func TestDefendNewWorkflowWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomdeeprailsdeeprailsgosdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Defend.NewWorkflow(context.TODO(), githubcomdeeprailsdeeprailsgosdk.DefendNewWorkflowParams{
		ImprovementAction: githubcomdeeprailsdeeprailsgosdk.DefendNewWorkflowParamsImprovementActionRegenerate,
		Metrics: map[string]float64{
			"foo": 0,
		},
		Name:               "name",
		Type:               githubcomdeeprailsdeeprailsgosdk.DefendNewWorkflowParamsTypeAutomatic,
		AutomaticTolerance: githubcomdeeprailsdeeprailsgosdk.DefendNewWorkflowParamsAutomaticToleranceLow,
		Description:        githubcomdeeprailsdeeprailsgosdk.String("description"),
		MaxRetries:         githubcomdeeprailsdeeprailsgosdk.Int(0),
	})
	if err != nil {
		var apierr *githubcomdeeprailsdeeprailsgosdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDefendGetEvent(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomdeeprailsdeeprailsgosdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Defend.GetEvent(
		context.TODO(),
		"event_id",
		githubcomdeeprailsdeeprailsgosdk.DefendGetEventParams{
			WorkflowID: "workflow_id",
		},
	)
	if err != nil {
		var apierr *githubcomdeeprailsdeeprailsgosdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDefendGetWorkflow(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomdeeprailsdeeprailsgosdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Defend.GetWorkflow(context.TODO(), "workflow_id")
	if err != nil {
		var apierr *githubcomdeeprailsdeeprailsgosdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDefendSubmitEventWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomdeeprailsdeeprailsgosdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Defend.SubmitEvent(
		context.TODO(),
		"workflow_id",
		githubcomdeeprailsdeeprailsgosdk.DefendSubmitEventParams{
			ModelInput: githubcomdeeprailsdeeprailsgosdk.DefendSubmitEventParamsModelInput{
				UserPrompt: "user_prompt",
				Context:    githubcomdeeprailsdeeprailsgosdk.String("context"),
			},
			ModelOutput: "model_output",
			ModelUsed:   "model_used",
			Nametag:     "nametag",
			RunMode:     githubcomdeeprailsdeeprailsgosdk.DefendSubmitEventParamsRunModePrecisionPlus,
		},
	)
	if err != nil {
		var apierr *githubcomdeeprailsdeeprailsgosdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDefendUpdateWorkflowWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomdeeprailsdeeprailsgosdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Defend.UpdateWorkflow(
		context.TODO(),
		"workflow_id",
		githubcomdeeprailsdeeprailsgosdk.DefendUpdateWorkflowParams{
			Description: githubcomdeeprailsdeeprailsgosdk.String("description"),
			Name:        githubcomdeeprailsdeeprailsgosdk.String("name"),
			Type:        githubcomdeeprailsdeeprailsgosdk.DefendUpdateWorkflowParamsTypeAutomatic,
		},
	)
	if err != nil {
		var apierr *githubcomdeeprailsdeeprailsgosdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
