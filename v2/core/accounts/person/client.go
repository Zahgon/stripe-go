//
//
// File generated from our OpenAPI spec
//
//

// Package person provides the person related APIs
package person

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke person related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a Person. Adds an individual to an Account's identity. You can set relationship attributes and identity information at creation.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2CoreAccountsPersonParams) (*stripe.V2CoreAccountPerson, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a Person associated with an Account.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2CoreAccountsPersonParams) (*stripe.V2CoreAccountPerson, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates a Person associated with an Account.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.V2CoreAccountsPersonParams) (*stripe.V2CoreAccountPerson, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete a Person associated with an Account.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Del(id string, params *stripe.V2CoreAccountsPersonParams) (*stripe.V2DeletedObject, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a paginated list of Persons associated with an Account.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2CoreAccountsPersonListParams) stripe.Seq2[*stripe.V2CoreAccountPerson, error] {
	_ = "STUB: not implemented"
	return nil
}
