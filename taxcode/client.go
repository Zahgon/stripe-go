//
//
// File generated from our OpenAPI spec
//
//

// Package taxcode provides the /v1/tax_codes APIs
package taxcode

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/tax_codes APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the details of an existing tax code. Supply the unique tax code ID and Stripe will return the corresponding tax code information.
func Get(id string, params *stripe.TaxCodeParams) (*stripe.TaxCode, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves the details of an existing tax code. Supply the unique tax code ID and Stripe will return the corresponding tax code information.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.TaxCodeParams) (*stripe.TaxCode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// A list of [all tax codes available](https://stripe.com/docs/tax/tax-categories) to add to Products in order to allow specific tax calculations.
func List(params *stripe.TaxCodeListParams) *Iter { _ = "STUB: not implemented"; return nil }

// A list of [all tax codes available](https://stripe.com/docs/tax/tax-categories) to add to Products in order to allow specific tax calculations.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.TaxCodeListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for tax codes.
type Iter struct {
	*stripe.Iter
}

// TaxCode returns the tax code which the iterator is currently pointing to.
func (i *Iter) TaxCode() *stripe.TaxCode { _ = "STUB: not implemented"; return nil }

// TaxCodeList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TaxCodeList() *stripe.TaxCodeList { _ = "STUB: not implemented"; return nil }

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
