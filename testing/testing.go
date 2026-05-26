package testing

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	stripe "github.com/stripe/stripe-go/v85"
	"github.com/stripe/stripe-go/v85/form"
)

// This file should contain any testing helpers that should be commonly
// available across all tests in the Stripe package.
//
// There's not much in here because it' a relatively recent addition to the
// package, but should be used as appropriate for any new changes.

const (
	// MockMinimumVersion is the minimum acceptable version for stripe-mock.
	// It's here so that if the library depends on new endpoints or features
	// added in a more recent version of stripe-mock, we can show people a
	// better error message instead of the test suite crashing with a bunch of
	// confusing 404 errors or the like.
	MockMinimumVersion = "0.109.0"

	// TestMerchantID is a token that can be used to represent a merchant ID in
	// simple tests.
	TestMerchantID = "acct_123"

	// TestAPIKey is an API key that can be used against stripe-mock in tests.
	TestAPIKey = "sk_test_myTestKey"
)

func init() {
	// Enable strict mode on form encoding so that we'll panic if any kind of
	// malformed param struct is detected
	form.Strict = true

	port := os.Getenv("STRIPE_MOCK_PORT")
	if port == "" {
		port = "12112"
	}

	// stripe-mock's certificate for localhost is self-signed so configure a
	// specialized client that skips the certificate authority check.
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		// ForceAttemptHTTP2 is needed because Go disables HTTP/2 when TLSClientConfig is set
		ForceAttemptHTTP2: true,
	}

	httpClient := &http.Client{
		Transport: transport,
	}

	resp, err := httpClient.Get("https://localhost:" + port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't reach stripe-mock at `localhost:%s` (%v). Is "+
			"it running? Please see README for setup instructions.\n", port, err)
		os.Exit(1)
	}

	if resp.ProtoMajor != 2 {
		fmt.Fprintf(os.Stderr, "Expected HTTP/2 connection to stripe-mock, but got %s. "+
			"This may indicate ForceAttemptHTTP2 is not set correctly.\n", resp.Proto)
		os.Exit(1)
	}

	version := resp.Header.Get("Stripe-Mock-Version")
	if version != "master" && compareVersions(version, MockMinimumVersion) > 0 {
		fmt.Fprintf(os.Stderr, "Your version of stripe-mock (%s) is too old. The "+
			"minimum version to run this test suite is %s. Please see its "+
			"repository for upgrade instructions.\n", version, MockMinimumVersion)
		os.Exit(1)
	}

	stripe.Key = TestAPIKey

	// Configure a backend for stripe-mock and set it for both the API and
	// Uploads (unlike the real Stripe API, stripe-mock supports both these
	// backends).
	stripeMockBackend := stripe.GetBackendWithConfig(
		stripe.APIBackend,
		&stripe.BackendConfig{
			URL:           stripe.String("https://localhost:" + port),
			HTTPClient:    httpClient,
			LeveledLogger: stripe.DefaultLeveledLogger,
		},
	)
	stripe.SetBackend(stripe.APIBackend, stripeMockBackend)
	stripe.SetBackend(stripe.UploadsBackend, stripeMockBackend)
}

func MockServer(t *testing.T, method, path string, params interface{}, resp string) *httptest.Server {
	_ = "STUB: not implemented"
	return nil
}

func MockServerWithStripeContext(t *testing.T, method, path, stripeContext string, params interface{}, resp string) *httptest.Server {
	_ = "STUB: not implemented"
	return nil
}

// compareVersions compares two semantic version strings. We need this because
// with more complex double-digit numbers, lexical comparison breaks down.
func compareVersions(a, b string) (ret int) { _ = "STUB: not implemented"; return 0 }
