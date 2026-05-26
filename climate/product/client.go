//
//
// File generated from our OpenAPI spec
//
//

// Package product provides the /v1/climate/products APIs
package product

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/climate/products APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the details of a Climate product with the given ID.
func Get(id string, params *stripe.ClimateProductParams) (*stripe.ClimateProduct, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves the details of a Climate product with the given ID.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.ClimateProductParams) (*stripe.ClimateProduct, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Lists all available Climate product objects.
func List(params *stripe.ClimateProductListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Lists all available Climate product objects.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.ClimateProductListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for climate products.
type Iter struct {
	*stripe.Iter
}

// ClimateProduct returns the climate product which the iterator is currently pointing to.
func (i *Iter) ClimateProduct() *stripe.ClimateProduct { _ = "STUB: not implemented"; return nil }

// ClimateProductList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) ClimateProductList() *stripe.ClimateProductList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
