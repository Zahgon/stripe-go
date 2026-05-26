package mock

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stripe/stripe-go/v85/client"
	. "github.com/stripe/stripe-go/v85/testing"
)

type Assertion func(*testing.T, *http.Request)

func Server[T any](t *testing.T, method, path string, req T, resp func(T) []byte, asserts ...Assertion) (*httptest.Server, *client.API) {
	_ = "STUB: not implemented"
	return nil, nil
}
