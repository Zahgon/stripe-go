//
//
// File generated from our OpenAPI spec
//
//

// Package transaction provides the /v1/issuing/transactions APIs
package transaction

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/issuing/transactions APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves an Issuing Transaction object.
func Get(id string, params *stripe.IssuingTransactionParams) (*stripe.IssuingTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves an Issuing Transaction object.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.IssuingTransactionParams) (*stripe.IssuingTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specified Issuing Transaction object by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
func Update(id string, params *stripe.IssuingTransactionParams) (*stripe.IssuingTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specified Issuing Transaction object by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.IssuingTransactionParams) (*stripe.IssuingTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of Issuing Transaction objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
func List(params *stripe.IssuingTransactionListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Returns a list of Issuing Transaction objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.IssuingTransactionListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for issuing transactions.
type Iter struct {
	*stripe.Iter
}

// IssuingTransaction returns the issuing transaction which the iterator is currently pointing to.
func (i *Iter) IssuingTransaction() *stripe.IssuingTransaction {
	_ = "STUB: not implemented"
	return nil
}

// IssuingTransactionList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) IssuingTransactionList() *stripe.IssuingTransactionList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
