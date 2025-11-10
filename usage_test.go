// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package deeprails_test

import (
	"context"
	"os"
	"testing"

	"github.com/deeprails/deeprails-go-sdk"
	"github.com/deeprails/deeprails-go-sdk/internal/testutil"
	"github.com/deeprails/deeprails-go-sdk/option"
)

func TestUsage(t *testing.T) {
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
	defendCreateResponse, err := client.Defend.NewWorkflow(context.TODO(), deeprails.DefendNewWorkflowParams{
		ImprovementAction: deeprails.F(deeprails.DefendNewWorkflowParamsImprovementActionFixit),
		Name:              deeprails.F("Push Alert Workflow"),
		ThresholdType:     deeprails.F(deeprails.DefendNewWorkflowParamsThresholdTypeAutomatic),
		CustomHallucinationThresholdValues: deeprails.F(map[string]float64{
			"completeness":          0.700000,
			"instruction_adherence": 0.750000,
		}),
		WebSearch: deeprails.F(true),
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", defendCreateResponse.WorkflowID)
}
