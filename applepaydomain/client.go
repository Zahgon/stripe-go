//
//
// File generated from our OpenAPI spec
//
//

// Package applepaydomain provides the /v1/apple_pay/domains APIs
package applepaydomain

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/apple_pay/domains APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create an apple pay domain.
func New(params *stripe.ApplePayDomainParams) (*stripe.ApplePayDomain, error) {
	_ = "STUB: not implemented"
	return nil,

		// Create an apple pay domain.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.ApplePayDomainParams) (*stripe.ApplePayDomain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieve an apple pay domain.
func Get(id string, params *stripe.ApplePayDomainParams) (*stripe.ApplePayDomain, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieve an apple pay domain.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.ApplePayDomainParams) (*stripe.ApplePayDomain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete an apple pay domain.
func Del(id string, params *stripe.ApplePayDomainParams) (*stripe.ApplePayDomain, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Delete an apple pay domain.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Del(id string, params *stripe.ApplePayDomainParams) (*stripe.ApplePayDomain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// List apple pay domains.
func List(params *stripe.ApplePayDomainListParams) *Iter { _ = "STUB: not implemented"; return nil }

// List apple pay domains.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.ApplePayDomainListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for apple pay domains.
type Iter struct {
	*stripe.Iter
}

// ApplePayDomain returns the apple pay domain which the iterator is currently pointing to.
func (i *Iter) ApplePayDomain() *stripe.ApplePayDomain { _ = "STUB: not implemented"; return nil }

// ApplePayDomainList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) ApplePayDomainList() *stripe.ApplePayDomainList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
