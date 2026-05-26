//
//
// File generated from our OpenAPI spec
//
//

// Package confirmationtoken provides the /v1/confirmation_tokens APIs
package confirmationtoken

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/confirmation_tokens APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a test mode Confirmation Token server side for your integration tests.
func New(params *stripe.TestHelpersConfirmationTokenParams) (*stripe.ConfirmationToken, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates a test mode Confirmation Token server side for your integration tests.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.TestHelpersConfirmationTokenParams) (*stripe.ConfirmationToken, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
