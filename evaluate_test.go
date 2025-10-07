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

func TestEvaluateNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Evaluate.New(context.TODO(), githubcomdeeprailsdeeprailsgosdk.EvaluateNewParams{
		ModelInput: githubcomdeeprailsdeeprailsgosdk.EvaluateNewParamsModelInput{
			UserPrompt: "user_prompt",
			Context:    githubcomdeeprailsdeeprailsgosdk.String("context"),
		},
		ModelOutput:      "model_output",
		RunMode:          githubcomdeeprailsdeeprailsgosdk.EvaluateNewParamsRunModePrecisionPlus,
		GuardrailMetrics: []string{"correctness"},
		ModelUsed:        githubcomdeeprailsdeeprailsgosdk.String("model_used"),
		Nametag:          githubcomdeeprailsdeeprailsgosdk.String("nametag"),
	})
	if err != nil {
		var apierr *githubcomdeeprailsdeeprailsgosdk.Error
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
	client := githubcomdeeprailsdeeprailsgosdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Evaluate.Get(context.TODO(), "eval_id")
	if err != nil {
		var apierr *githubcomdeeprailsdeeprailsgosdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
