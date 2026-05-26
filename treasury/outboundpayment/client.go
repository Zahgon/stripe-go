//
//
// File generated from our OpenAPI spec
//
//

// Package outboundpayment provides the /v1/treasury/outbound_payments APIs
package outboundpayment

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/treasury/outbound_payments APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates an OutboundPayment.
func New(params *stripe.TreasuryOutboundPaymentParams) (*stripe.TreasuryOutboundPayment, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates an OutboundPayment.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.TreasuryOutboundPaymentParams) (*stripe.TreasuryOutboundPayment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the details of an existing OutboundPayment by passing the unique OutboundPayment ID from either the OutboundPayment creation request or OutboundPayment list.
func Get(id string, params *stripe.TreasuryOutboundPaymentParams) (*stripe.TreasuryOutboundPayment, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves the details of an existing OutboundPayment by passing the unique OutboundPayment ID from either the OutboundPayment creation request or OutboundPayment list.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.TreasuryOutboundPaymentParams) (*stripe.TreasuryOutboundPayment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Cancel an OutboundPayment.
func Cancel(id string, params *stripe.TreasuryOutboundPaymentCancelParams) (*stripe.TreasuryOutboundPayment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Cancel an OutboundPayment.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Cancel(id string, params *stripe.TreasuryOutboundPaymentCancelParams) (*stripe.TreasuryOutboundPayment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of OutboundPayments sent from the specified FinancialAccount.
func List(params *stripe.TreasuryOutboundPaymentListParams) *Iter {
	_ = "STUB: not implemented"
	return nil

	// Returns a list of OutboundPayments sent from the specified FinancialAccount.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) List(listParams *stripe.TreasuryOutboundPaymentListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for treasury outbound payments.
type Iter struct {
	*stripe.Iter
}

// TreasuryOutboundPayment returns the treasury outbound payment which the iterator is currently pointing to.
func (i *Iter) TreasuryOutboundPayment() *stripe.TreasuryOutboundPayment {
	_ = "STUB: not implemented"
	return nil
}

// TreasuryOutboundPaymentList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TreasuryOutboundPaymentList() *stripe.TreasuryOutboundPaymentList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
