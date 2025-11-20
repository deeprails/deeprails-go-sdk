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

func TestFileUpload(t *testing.T) {
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
	_, err := client.Files.Upload(context.TODO(), deeprails.FileUploadParams{
		Files: deeprails.F([]string{"string"}),
	})
	if err != nil {
		var apierr *deeprails.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
