// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomdeeprailsdeeprailsgosdk_test

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
	client := githubcomdeeprailsdeeprailsgosdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	defendResponse, err := client.Defend.NewWorkflow(context.TODO(), githubcomdeeprailsdeeprailsgosdk.DefendNewWorkflowParams{
		ImprovementAction: githubcomdeeprailsdeeprailsgosdk.DefendNewWorkflowParamsImprovementActionFixit,
		Metrics: map[string]float64{
			"completeness":          0.8,
			"instruction_adherence": 0.75,
		},
		Name: "Push Alert Workflow",
		Type: githubcomdeeprailsdeeprailsgosdk.DefendNewWorkflowParamsTypeCustom,
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", defendResponse.WorkflowID)
}
