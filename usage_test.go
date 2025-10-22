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
	defendResponse, err := client.Defend.NewWorkflow(context.TODO(), deeprails.DefendNewWorkflowParams{
		ImprovementAction: deeprails.F(deeprails.DefendNewWorkflowParamsImprovementActionFixit),
		Name:              deeprails.F("Push Alert Workflow"),
		Type:              deeprails.F(deeprails.DefendNewWorkflowParamsTypeCustom),
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", defendResponse.WorkflowID)
}
