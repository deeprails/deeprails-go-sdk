// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package deeprails_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/deeprails-go"
	"github.com/stainless-sdks/deeprails-go/internal/testutil"
	"github.com/stainless-sdks/deeprails-go/option"
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
	client := deeprails.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Defend.NewWorkflow(context.TODO(), deeprails.DefendNewWorkflowParams{
		ImprovementAction: deeprails.DefendNewWorkflowParamsImprovementActionRegenerate,
		Metrics: map[string]float64{
			"foo": 0,
		},
		Name:               "name",
		Type:               deeprails.DefendNewWorkflowParamsTypeAutomatic,
		AutomaticTolerance: deeprails.DefendNewWorkflowParamsAutomaticToleranceLow,
		Description:        deeprails.String("description"),
		MaxRetries:         deeprails.Int(0),
	})
	if err != nil {
		var apierr *deeprails.Error
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
	client := deeprails.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Defend.GetEvent(
		context.TODO(),
		"event_id",
		deeprails.DefendGetEventParams{
			WorkflowID: "workflow_id",
		},
	)
	if err != nil {
		var apierr *deeprails.Error
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
	client := deeprails.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Defend.GetWorkflow(context.TODO(), "workflow_id")
	if err != nil {
		var apierr *deeprails.Error
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
	client := deeprails.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Defend.SubmitEvent(
		context.TODO(),
		"workflow_id",
		deeprails.DefendSubmitEventParams{
			ModelInput: deeprails.DefendSubmitEventParamsModelInput{
				UserPrompt: "user_prompt",
				Context:    deeprails.String("context"),
			},
			ModelOutput: "model_output",
			ModelUsed:   "model_used",
			Nametag:     "nametag",
			RunMode:     deeprails.DefendSubmitEventParamsRunModePrecisionPlus,
		},
	)
	if err != nil {
		var apierr *deeprails.Error
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
	client := deeprails.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Defend.UpdateWorkflow(
		context.TODO(),
		"workflow_id",
		deeprails.DefendUpdateWorkflowParams{
			Description: deeprails.String("description"),
			Name:        deeprails.String("name"),
			Type:        deeprails.DefendUpdateWorkflowParamsTypeAutomatic,
		},
	)
	if err != nil {
		var apierr *deeprails.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
