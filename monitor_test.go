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

func TestMonitorNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Monitor.New(context.TODO(), githubcomdeeprailsdeeprailsgosdk.MonitorNewParams{
		Name:        "name",
		Description: githubcomdeeprailsdeeprailsgosdk.String("description"),
	})
	if err != nil {
		var apierr *githubcomdeeprailsdeeprailsgosdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMonitorGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Monitor.Get(
		context.TODO(),
		"monitor_id",
		githubcomdeeprailsdeeprailsgosdk.MonitorGetParams{
			Limit: githubcomdeeprailsdeeprailsgosdk.Int(0),
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

func TestMonitorUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Monitor.Update(
		context.TODO(),
		"monitor_id",
		githubcomdeeprailsdeeprailsgosdk.MonitorUpdateParams{
			Description:   githubcomdeeprailsdeeprailsgosdk.String("description"),
			MonitorStatus: githubcomdeeprailsdeeprailsgosdk.MonitorUpdateParamsMonitorStatusActive,
			Name:          githubcomdeeprailsdeeprailsgosdk.String("name"),
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

func TestMonitorSubmitEventWithOptionalParams(t *testing.T) {
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
	_, err := client.Monitor.SubmitEvent(
		context.TODO(),
		"monitor_id",
		githubcomdeeprailsdeeprailsgosdk.MonitorSubmitEventParams{
			GuardrailMetrics: []string{"correctness"},
			ModelInput: githubcomdeeprailsdeeprailsgosdk.MonitorSubmitEventParamsModelInput{
				UserPrompt: "user_prompt",
				Context:    githubcomdeeprailsdeeprailsgosdk.String("context"),
			},
			ModelOutput: "model_output",
			ModelUsed:   githubcomdeeprailsdeeprailsgosdk.String("model_used"),
			Nametag:     githubcomdeeprailsdeeprailsgosdk.String("nametag"),
			RunMode:     githubcomdeeprailsdeeprailsgosdk.MonitorSubmitEventParamsRunModePrecisionPlus,
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
