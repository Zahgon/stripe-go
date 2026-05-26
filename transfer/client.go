//
//
// File generated from our OpenAPI spec
//
//

// Package transfer provides the /v1/transfers APIs
package transfer

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/transfers APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// To send funds from your Stripe account to a connected account, you create a new transfer object. Your [Stripe balance](https://docs.stripe.com/api#balance) must be able to cover the transfer amount, or you'll receive an “Insufficient Funds” error.
func New(params *stripe.TransferParams) (*stripe.Transfer, error) {
	_ = "STUB: not implemented"
	return nil,

		// To send funds from your Stripe account to a connected account, you create a new transfer object. Your [Stripe balance](https://docs.stripe.com/api#balance) must be able to cover the transfer amount, or you'll receive an “Insufficient Funds” error.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.TransferParams) (*stripe.Transfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the details of an existing transfer. Supply the unique transfer ID from either a transfer creation request or the transfer list, and Stripe will return the corresponding transfer information.
func Get(id string, params *stripe.TransferParams) (*stripe.Transfer, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves the details of an existing transfer. Supply the unique transfer ID from either a transfer creation request or the transfer list, and Stripe will return the corresponding transfer information.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.TransferParams) (*stripe.Transfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specified transfer by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
//
// This request accepts only metadata as an argument.
func Update(id string, params *stripe.TransferParams) (*stripe.Transfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specified transfer by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
//
// This request accepts only metadata as an argument.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.TransferParams) (*stripe.Transfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of existing transfers sent to connected accounts. The transfers are returned in sorted order, with the most recently created transfers appearing first.
func List(params *stripe.TransferListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Returns a list of existing transfers sent to connected accounts. The transfers are returned in sorted order, with the most recently created transfers appearing first.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.TransferListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for transfers.
type Iter struct {
	*stripe.Iter
}

// Transfer returns the transfer which the iterator is currently pointing to.
func (i *Iter) Transfer() *stripe.Transfer { _ = "STUB: not implemented"; return nil }

// TransferList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TransferList() *stripe.TransferList { _ = "STUB: not implemented"; return nil }

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
