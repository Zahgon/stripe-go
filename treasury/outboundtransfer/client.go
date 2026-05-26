//
//
// File generated from our OpenAPI spec
//
//

// Package outboundtransfer provides the /v1/treasury/outbound_transfers APIs
package outboundtransfer

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/treasury/outbound_transfers APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates an OutboundTransfer.
func New(params *stripe.TreasuryOutboundTransferParams) (*stripe.TreasuryOutboundTransfer, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates an OutboundTransfer.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.TreasuryOutboundTransferParams) (*stripe.TreasuryOutboundTransfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the details of an existing OutboundTransfer by passing the unique OutboundTransfer ID from either the OutboundTransfer creation request or OutboundTransfer list.
func Get(id string, params *stripe.TreasuryOutboundTransferParams) (*stripe.TreasuryOutboundTransfer, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves the details of an existing OutboundTransfer by passing the unique OutboundTransfer ID from either the OutboundTransfer creation request or OutboundTransfer list.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.TreasuryOutboundTransferParams) (*stripe.TreasuryOutboundTransfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// An OutboundTransfer can be canceled if the funds have not yet been paid out.
func Cancel(id string, params *stripe.TreasuryOutboundTransferCancelParams) (*stripe.TreasuryOutboundTransfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// An OutboundTransfer can be canceled if the funds have not yet been paid out.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Cancel(id string, params *stripe.TreasuryOutboundTransferCancelParams) (*stripe.TreasuryOutboundTransfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of OutboundTransfers sent from the specified FinancialAccount.
func List(params *stripe.TreasuryOutboundTransferListParams) *Iter {
	_ = "STUB: not implemented"
	return nil

	// Returns a list of OutboundTransfers sent from the specified FinancialAccount.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) List(listParams *stripe.TreasuryOutboundTransferListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for treasury outbound transfers.
type Iter struct {
	*stripe.Iter
}

// TreasuryOutboundTransfer returns the treasury outbound transfer which the iterator is currently pointing to.
func (i *Iter) TreasuryOutboundTransfer() *stripe.TreasuryOutboundTransfer {
	_ = "STUB: not implemented"
	return nil
}

// TreasuryOutboundTransferList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TreasuryOutboundTransferList() *stripe.TreasuryOutboundTransferList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
