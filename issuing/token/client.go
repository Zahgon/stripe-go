//
//
// File generated from our OpenAPI spec
//
//

// Package token provides the /v1/issuing/tokens APIs
package token

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/issuing/tokens APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves an Issuing Token object.
func Get(id string, params *stripe.IssuingTokenParams) (*stripe.IssuingToken, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves an Issuing Token object.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.IssuingTokenParams) (*stripe.IssuingToken, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Attempts to update the specified Issuing Token object to the status specified.
func Update(id string, params *stripe.IssuingTokenParams) (*stripe.IssuingToken, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Attempts to update the specified Issuing Token object to the status specified.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.IssuingTokenParams) (*stripe.IssuingToken, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Lists all Issuing Token objects for a given card.
func List(params *stripe.IssuingTokenListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Lists all Issuing Token objects for a given card.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.IssuingTokenListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for issuing tokens.
type Iter struct {
	*stripe.Iter
}

// IssuingToken returns the issuing token which the iterator is currently pointing to.
func (i *Iter) IssuingToken() *stripe.IssuingToken { _ = "STUB: not implemented"; return nil }

// IssuingTokenList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) IssuingTokenList() *stripe.IssuingTokenList { _ = "STUB: not implemented"; return nil }

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
