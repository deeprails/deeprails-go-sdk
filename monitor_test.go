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
		Name:        deeprails.F("name"),
		Description: deeprails.F("description"),
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
			Limit: deeprails.F(int64(0)),
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
			Description:   deeprails.F("description"),
			MonitorStatus: deeprails.F(deeprails.MonitorUpdateParamsMonitorStatusActive),
			Name:          deeprails.F("name"),
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
			GuardrailMetrics: deeprails.F([]deeprails.MonitorSubmitEventParamsGuardrailMetric{deeprails.MonitorSubmitEventParamsGuardrailMetricCorrectness}),
			ModelInput: deeprails.F(deeprails.MonitorSubmitEventParamsModelInput{
				GroundTruth:  deeprails.F("ground_truth"),
				SystemPrompt: deeprails.F("system_prompt"),
				UserPrompt:   deeprails.F("user_prompt"),
			}),
			ModelOutput: deeprails.F("model_output"),
			ModelUsed:   deeprails.F("model_used"),
			Nametag:     deeprails.F("nametag"),
			RunMode:     deeprails.F(deeprails.MonitorSubmitEventParamsRunModePrecisionPlus),
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
