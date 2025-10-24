// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package deeprails_test

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
	client := deeprails.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Defend.NewWorkflow(context.TODO(), deeprails.DefendNewWorkflowParams{
		ImprovementAction: deeprails.F(deeprails.DefendNewWorkflowParamsImprovementActionRegen),
		Name:              deeprails.F("name"),
		Type:              deeprails.F(deeprails.DefendNewWorkflowParamsTypeAutomatic),
		AutomaticHallucinationToleranceLevels: deeprails.F(map[string]deeprails.DefendNewWorkflowParamsAutomaticHallucinationToleranceLevels{
			"foo": deeprails.DefendNewWorkflowParamsAutomaticHallucinationToleranceLevelsLow,
		}),
		Description:            deeprails.F("description"),
		MaxImprovementAttempts: deeprails.F(int64(3)),
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
		"workflow_id",
		"event_id",
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
			ModelInput: deeprails.F(deeprails.DefendSubmitEventParamsModelInput{
				GroundTruth:  deeprails.F("ground_truth"),
				SystemPrompt: deeprails.F("system_prompt"),
				UserPrompt:   deeprails.F("user_prompt"),
			}),
			ModelOutput: deeprails.F("model_output"),
			ModelUsed:   deeprails.F("model_used"),
			RunMode:     deeprails.F(deeprails.DefendSubmitEventParamsRunModePrecisionPlus),
			Nametag:     deeprails.F("nametag"),
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
			Description: deeprails.F("description"),
			Name:        deeprails.F("name"),
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
