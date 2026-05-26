//
//
// File generated from our OpenAPI spec
//
//

// Package supplier provides the /v1/climate/suppliers APIs
package supplier

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/climate/suppliers APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves a Climate supplier object.
func Get(id string, params *stripe.ClimateSupplierParams) (*stripe.ClimateSupplier, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves a Climate supplier object.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.ClimateSupplierParams) (*stripe.ClimateSupplier, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Lists all available Climate supplier objects.
func List(params *stripe.ClimateSupplierListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Lists all available Climate supplier objects.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.ClimateSupplierListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for climate suppliers.
type Iter struct {
	*stripe.Iter
}

// ClimateSupplier returns the climate supplier which the iterator is currently pointing to.
func (i *Iter) ClimateSupplier() *stripe.ClimateSupplier { _ = "STUB: not implemented"; return nil }

// ClimateSupplierList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) ClimateSupplierList() *stripe.ClimateSupplierList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
