//
//
// File generated from our OpenAPI spec
//
//

// Package shippingrate provides the /v1/shipping_rates APIs
package shippingrate

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/shipping_rates APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new shipping rate object.
func New(params *stripe.ShippingRateParams) (*stripe.ShippingRate, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates a new shipping rate object.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.ShippingRateParams) (*stripe.ShippingRate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns the shipping rate object with the given ID.
func Get(id string, params *stripe.ShippingRateParams) (*stripe.ShippingRate, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Returns the shipping rate object with the given ID.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.ShippingRateParams) (*stripe.ShippingRate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates an existing shipping rate object.
func Update(id string, params *stripe.ShippingRateParams) (*stripe.ShippingRate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates an existing shipping rate object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.ShippingRateParams) (*stripe.ShippingRate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of your shipping rates.
func List(params *stripe.ShippingRateListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Returns a list of your shipping rates.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.ShippingRateListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for shipping rates.
type Iter struct {
	*stripe.Iter
}

// ShippingRate returns the shipping rate which the iterator is currently pointing to.
func (i *Iter) ShippingRate() *stripe.ShippingRate { _ = "STUB: not implemented"; return nil }

// ShippingRateList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) ShippingRateList() *stripe.ShippingRateList { _ = "STUB: not implemented"; return nil }

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
