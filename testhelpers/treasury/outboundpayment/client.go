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

// Updates a test mode created OutboundPayment with tracking details. The OutboundPayment must not be cancelable, and cannot be in the canceled or failed states.
func Update(id string, params *stripe.TestHelpersTreasuryOutboundPaymentParams) (*stripe.TreasuryOutboundPayment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates a test mode created OutboundPayment with tracking details. The OutboundPayment must not be cancelable, and cannot be in the canceled or failed states.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.TestHelpersTreasuryOutboundPaymentParams) (*stripe.TreasuryOutboundPayment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Transitions a test mode created OutboundPayment to the failed status. The OutboundPayment must already be in the processing state.
func Fail(id string, params *stripe.TestHelpersTreasuryOutboundPaymentFailParams) (*stripe.TreasuryOutboundPayment, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Transitions a test mode created OutboundPayment to the failed status. The OutboundPayment must already be in the processing state.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Fail(id string, params *stripe.TestHelpersTreasuryOutboundPaymentFailParams) (*stripe.TreasuryOutboundPayment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Transitions a test mode created OutboundPayment to the posted status. The OutboundPayment must already be in the processing state.
func Post(id string, params *stripe.TestHelpersTreasuryOutboundPaymentPostParams) (*stripe.TreasuryOutboundPayment, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Transitions a test mode created OutboundPayment to the posted status. The OutboundPayment must already be in the processing state.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Post(id string, params *stripe.TestHelpersTreasuryOutboundPaymentPostParams) (*stripe.TreasuryOutboundPayment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Transitions a test mode created OutboundPayment to the returned status. The OutboundPayment must already be in the processing state.
func ReturnOutboundPayment(id string, params *stripe.TestHelpersTreasuryOutboundPaymentReturnOutboundPaymentParams) (*stripe.TreasuryOutboundPayment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Transitions a test mode created OutboundPayment to the returned status. The OutboundPayment must already be in the processing state.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ReturnOutboundPayment(id string, params *stripe.TestHelpersTreasuryOutboundPaymentReturnOutboundPaymentParams) (*stripe.TreasuryOutboundPayment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
