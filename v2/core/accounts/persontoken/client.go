//
//
// File generated from our OpenAPI spec
//
//

// Package persontoken provides the persontoken related APIs
package persontoken

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke persontoken related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a Person Token associated with an Account.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2CoreAccountsPersonTokenParams) (*stripe.V2CoreAccountPersonToken, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a Person Token associated with an Account.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2CoreAccountsPersonTokenParams) (*stripe.V2CoreAccountPersonToken, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
