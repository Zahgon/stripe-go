//
//
// File generated from our OpenAPI spec
//
//

// Package paymentsource provides the /v1/customers/{customer}/sources APIs
package paymentsource

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/customers/{customer}/sources APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// When you create a new credit card, you must specify a customer or recipient on which to create it.
//
// If the card's owner has no default card, then the new card will become the default.
// However, if the owner already has a default, then it will not change.
// To change the default, you should [update the customer](https://docs.stripe.com/api/customers/update) to have a new default_source.
func New(params *stripe.PaymentSourceParams) (*stripe.PaymentSource, error) {
	_ = "STUB: not implemented"
	return nil,

		// When you create a new credit card, you must specify a customer or recipient on which to create it.
		//
		// If the card's owner has no default card, then the new card will become the default.
		// However, if the owner already has a default, then it will not change.
		// To change the default, you should [update the customer](https://docs.stripe.com/api/customers/update) to have a new default_source.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.PaymentSourceParams) (*stripe.PaymentSource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieve a specified source for a given customer.
func Get(id string, params *stripe.PaymentSourceParams) (*stripe.PaymentSource, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieve a specified source for a given customer.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.PaymentSourceParams) (*stripe.PaymentSource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update a specified source for a given customer.
func Update(id string, params *stripe.PaymentSourceParams) (*stripe.PaymentSource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update a specified source for a given customer.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.PaymentSourceParams) (*stripe.PaymentSource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete a specified source for a given customer.
func Del(id string, params *stripe.PaymentSourceParams) (*stripe.PaymentSource, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Delete a specified source for a given customer.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Del(id string, params *stripe.PaymentSourceParams) (*stripe.PaymentSource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Verify verifies a source which is used for bank accounts.
// Verify a specified bank account for a given customer.
func Verify(id string, params *stripe.PaymentSourceVerifyParams) (*stripe.PaymentSource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Verify verifies a source which is used for bank accounts.
// Verify a specified bank account for a given customer.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Verify(id string, params *stripe.PaymentSourceVerifyParams) (*stripe.PaymentSource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// List sources for a specified customer.
func List(params *stripe.PaymentSourceListParams) *Iter { _ = "STUB: not implemented"; return nil }

// List sources for a specified customer.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.PaymentSourceListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for payment sources.
type Iter struct {
	*stripe.Iter
}

// PaymentSource returns the payment source which the iterator is currently pointing to.
func (i *Iter) PaymentSource() *stripe.PaymentSource { _ = "STUB: not implemented"; return nil }

// PaymentSourceList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) PaymentSourceList() *stripe.PaymentSourceList { _ = "STUB: not implemented"; return nil }

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
