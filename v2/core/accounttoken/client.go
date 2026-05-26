//
//
// File generated from our OpenAPI spec
//
//

// Package accounttoken provides the accounttoken related APIs
package accounttoken

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke accounttoken related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates an Account Token.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2CoreAccountTokenParams) (*stripe.V2CoreAccountToken, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves an Account Token.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2CoreAccountTokenParams) (*stripe.V2CoreAccountToken, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
