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

func TestMonitorNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Monitor.New(context.TODO(), deeprails.MonitorNewParams{
		Name:        "name",
		Description: deeprails.String("description"),
	})
	if err != nil {
		var apierr *deeprails.Error
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
	client := deeprails.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Monitor.Get(
		context.TODO(),
		"monitor_id",
		deeprails.MonitorGetParams{
			Limit: deeprails.Int(0),
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

func TestMonitorUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Monitor.Update(
		context.TODO(),
		"monitor_id",
		deeprails.MonitorUpdateParams{
			Description:   deeprails.String("description"),
			MonitorStatus: deeprails.MonitorUpdateParamsMonitorStatusActive,
			Name:          deeprails.String("name"),
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

func TestMonitorSubmitEventWithOptionalParams(t *testing.T) {
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
	_, err := client.Monitor.SubmitEvent(
		context.TODO(),
		"monitor_id",
		deeprails.MonitorSubmitEventParams{
			GuardrailMetrics: []string{"correctness"},
			ModelInput: deeprails.MonitorSubmitEventParamsModelInput{
				UserPrompt: "user_prompt",
				Context:    deeprails.String("context"),
			},
			ModelOutput: "model_output",
			ModelUsed:   deeprails.String("model_used"),
			Nametag:     deeprails.String("nametag"),
			RunMode:     deeprails.MonitorSubmitEventParamsRunModePrecisionPlus,
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
