//
//
// File generated from our OpenAPI spec
//
//

// Package taxid provides the /v1/tax_ids APIs
package taxid

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/tax_ids APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new tax_id object for a customer.
func New(params *stripe.TaxIDParams) (*stripe.TaxID, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates a new tax_id object for a customer.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.TaxIDParams) (*stripe.TaxID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the tax_id object with the given identifier.
func Get(id string, params *stripe.TaxIDParams) (*stripe.TaxID, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves the tax_id object with the given identifier.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.TaxIDParams) (*stripe.TaxID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deletes an existing tax_id object.
func Del(id string, params *stripe.TaxIDParams) (*stripe.TaxID, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Deletes an existing tax_id object.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Del(id string, params *stripe.TaxIDParams) (*stripe.TaxID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of tax IDs for a customer.
func List(params *stripe.TaxIDListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Returns a list of tax IDs for a customer.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.TaxIDListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for tax ids.
type Iter struct {
	*stripe.Iter
}

// TaxID returns the tax id which the iterator is currently pointing to.
func (i *Iter) TaxID() *stripe.TaxID { _ = "STUB: not implemented"; return nil }

// TaxIDList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TaxIDList() *stripe.TaxIDList { _ = "STUB: not implemented"; return nil }

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
