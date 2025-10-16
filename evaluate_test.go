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

func TestEvaluateNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Evaluate.New(context.TODO(), deeprails.EvaluateNewParams{
		ModelInput: deeprails.F(deeprails.EvaluateNewParamsModelInput{
			GroundTruth:  deeprails.F("ground_truth"),
			SystemPrompt: deeprails.F("system_prompt"),
			UserPrompt:   deeprails.F("user_prompt"),
		}),
		ModelOutput:      deeprails.F("model_output"),
		RunMode:          deeprails.F(deeprails.EvaluateNewParamsRunModePrecisionPlus),
		GuardrailMetrics: deeprails.F([]deeprails.EvaluateNewParamsGuardrailMetric{deeprails.EvaluateNewParamsGuardrailMetricCorrectness}),
		ModelUsed:        deeprails.F("model_used"),
		Nametag:          deeprails.F("nametag"),
	})
	if err != nil {
		var apierr *deeprails.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEvaluateGet(t *testing.T) {
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
	_, err := client.Evaluate.Get(context.TODO(), "eval_id")
	if err != nil {
		var apierr *deeprails.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
