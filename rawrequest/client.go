// Package rawrequest provides sending stripe-flavored untyped HTTP calls
package rawrequest

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke raw requests against the specified backend
type Client struct {
	B   stripe.RawRequestBackend
	Key string
}

func (c Client) RawRequest(method string, path string, content string, params *stripe.RawParams) (*stripe.APIResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
