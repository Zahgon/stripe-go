//
//
// File generated from our OpenAPI spec
//
//

// Package taxrate provides the /v1/tax_rates APIs
package taxrate

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/tax_rates APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new tax rate.
func New(params *stripe.TaxRateParams) (*stripe.TaxRate, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates a new tax rate.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.TaxRateParams) (*stripe.TaxRate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a tax rate with the given ID
func Get(id string, params *stripe.TaxRateParams) (*stripe.TaxRate, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves a tax rate with the given ID
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.TaxRateParams) (*stripe.TaxRate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates an existing tax rate.
func Update(id string, params *stripe.TaxRateParams) (*stripe.TaxRate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates an existing tax rate.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.TaxRateParams) (*stripe.TaxRate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of your tax rates. Tax rates are returned sorted by creation date, with the most recently created tax rates appearing first.
func List(params *stripe.TaxRateListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Returns a list of your tax rates. Tax rates are returned sorted by creation date, with the most recently created tax rates appearing first.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.TaxRateListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for tax rates.
type Iter struct {
	*stripe.Iter
}

// TaxRate returns the tax rate which the iterator is currently pointing to.
func (i *Iter) TaxRate() *stripe.TaxRate { _ = "STUB: not implemented"; return nil }

// TaxRateList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TaxRateList() *stripe.TaxRateList { _ = "STUB: not implemented"; return nil }

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
