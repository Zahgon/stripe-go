//
//
// File generated from our OpenAPI spec
//
//

// Package inboundtransfer provides the /v1/treasury/inbound_transfers APIs
package inboundtransfer

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/treasury/inbound_transfers APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates an InboundTransfer.
func New(params *stripe.TreasuryInboundTransferParams) (*stripe.TreasuryInboundTransfer, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates an InboundTransfer.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.TreasuryInboundTransferParams) (*stripe.TreasuryInboundTransfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the details of an existing InboundTransfer.
func Get(id string, params *stripe.TreasuryInboundTransferParams) (*stripe.TreasuryInboundTransfer, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves the details of an existing InboundTransfer.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.TreasuryInboundTransferParams) (*stripe.TreasuryInboundTransfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Cancels an InboundTransfer.
func Cancel(id string, params *stripe.TreasuryInboundTransferCancelParams) (*stripe.TreasuryInboundTransfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Cancels an InboundTransfer.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Cancel(id string, params *stripe.TreasuryInboundTransferCancelParams) (*stripe.TreasuryInboundTransfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of InboundTransfers sent from the specified FinancialAccount.
func List(params *stripe.TreasuryInboundTransferListParams) *Iter {
	_ = "STUB: not implemented"
	return nil

	// Returns a list of InboundTransfers sent from the specified FinancialAccount.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) List(listParams *stripe.TreasuryInboundTransferListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for treasury inbound transfers.
type Iter struct {
	*stripe.Iter
}

// TreasuryInboundTransfer returns the treasury inbound transfer which the iterator is currently pointing to.
func (i *Iter) TreasuryInboundTransfer() *stripe.TreasuryInboundTransfer {
	_ = "STUB: not implemented"
	return nil
}

// TreasuryInboundTransferList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TreasuryInboundTransferList() *stripe.TreasuryInboundTransferList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
