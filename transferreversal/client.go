//
//
// File generated from our OpenAPI spec
//
//

// Package transferreversal provides the /v1/transfers/{id}/reversals APIs
package transferreversal

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/transfers/{id}/reversals APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// When you create a new reversal, you must specify a transfer to create it on.
//
// When reversing transfers, you can optionally reverse part of the transfer. You can do so as many times as you wish until the entire transfer has been reversed.
//
// Once entirely reversed, a transfer can't be reversed again. This method will return an error when called on an already-reversed transfer, or when trying to reverse more money than is left on a transfer.
func New(params *stripe.TransferReversalParams) (*stripe.TransferReversal, error) {
	_ = "STUB: not implemented"
	return nil,

		// When you create a new reversal, you must specify a transfer to create it on.
		//
		// When reversing transfers, you can optionally reverse part of the transfer. You can do so as many times as you wish until the entire transfer has been reversed.
		//
		// Once entirely reversed, a transfer can't be reversed again. This method will return an error when called on an already-reversed transfer, or when trying to reverse more money than is left on a transfer.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.TransferReversalParams) (*stripe.TransferReversal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// By default, you can see the 10 most recent reversals stored directly on the transfer object, but you can also retrieve details about a specific reversal stored on the transfer.
func Get(id string, params *stripe.TransferReversalParams) (*stripe.TransferReversal, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// By default, you can see the 10 most recent reversals stored directly on the transfer object, but you can also retrieve details about a specific reversal stored on the transfer.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.TransferReversalParams) (*stripe.TransferReversal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specified reversal by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
//
// This request only accepts metadata and description as arguments.
func Update(id string, params *stripe.TransferReversalParams) (*stripe.TransferReversal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specified reversal by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
//
// This request only accepts metadata and description as arguments.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.TransferReversalParams) (*stripe.TransferReversal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// You can see a list of the reversals belonging to a specific transfer. Note that the 10 most recent reversals are always available by default on the transfer object. If you need more than those 10, you can use this API method and the limit and starting_after parameters to page through additional reversals.
func List(params *stripe.TransferReversalListParams) *Iter { _ = "STUB: not implemented"; return nil }

// You can see a list of the reversals belonging to a specific transfer. Note that the 10 most recent reversals are always available by default on the transfer object. If you need more than those 10, you can use this API method and the limit and starting_after parameters to page through additional reversals.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.TransferReversalListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for transfer reversals.
type Iter struct {
	*stripe.Iter
}

// TransferReversal returns the transfer reversal which the iterator is currently pointing to.
func (i *Iter) TransferReversal() *stripe.TransferReversal { _ = "STUB: not implemented"; return nil }

// TransferReversalList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TransferReversalList() *stripe.TransferReversalList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
