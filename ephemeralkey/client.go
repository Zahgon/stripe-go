//
//
// File generated from our OpenAPI spec
//
//

// Package ephemeralkey provides the /v1/ephemeral_keys APIs
package ephemeralkey

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/ephemeral_keys APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a short-lived API key for a given resource.
func New(params *stripe.EphemeralKeyParams) (*stripe.EphemeralKey, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates a short-lived API key for a given resource.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.EphemeralKeyParams) (*stripe.EphemeralKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Invalidates a short-lived API key for a given resource.
func Del(id string, params *stripe.EphemeralKeyParams) (*stripe.EphemeralKey, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Invalidates a short-lived API key for a given resource.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Del(id string, params *stripe.EphemeralKeyParams) (*stripe.EphemeralKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
