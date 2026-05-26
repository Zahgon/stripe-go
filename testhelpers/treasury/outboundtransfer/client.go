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

// Updates a test mode created OutboundTransfer with tracking details. The OutboundTransfer must not be cancelable, and cannot be in the canceled or failed states.
func Update(id string, params *stripe.TestHelpersTreasuryOutboundTransferParams) (*stripe.TreasuryOutboundTransfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates a test mode created OutboundTransfer with tracking details. The OutboundTransfer must not be cancelable, and cannot be in the canceled or failed states.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.TestHelpersTreasuryOutboundTransferParams) (*stripe.TreasuryOutboundTransfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Transitions a test mode created OutboundTransfer to the failed status. The OutboundTransfer must already be in the processing state.
func Fail(id string, params *stripe.TestHelpersTreasuryOutboundTransferFailParams) (*stripe.TreasuryOutboundTransfer, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Transitions a test mode created OutboundTransfer to the failed status. The OutboundTransfer must already be in the processing state.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Fail(id string, params *stripe.TestHelpersTreasuryOutboundTransferFailParams) (*stripe.TreasuryOutboundTransfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Transitions a test mode created OutboundTransfer to the posted status. The OutboundTransfer must already be in the processing state.
func Post(id string, params *stripe.TestHelpersTreasuryOutboundTransferPostParams) (*stripe.TreasuryOutboundTransfer, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Transitions a test mode created OutboundTransfer to the posted status. The OutboundTransfer must already be in the processing state.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Post(id string, params *stripe.TestHelpersTreasuryOutboundTransferPostParams) (*stripe.TreasuryOutboundTransfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Transitions a test mode created OutboundTransfer to the returned status. The OutboundTransfer must already be in the processing state.
func ReturnOutboundTransfer(id string, params *stripe.TestHelpersTreasuryOutboundTransferReturnOutboundTransferParams) (*stripe.TreasuryOutboundTransfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Transitions a test mode created OutboundTransfer to the returned status. The OutboundTransfer must already be in the processing state.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ReturnOutboundTransfer(id string, params *stripe.TestHelpersTreasuryOutboundTransferReturnOutboundTransferParams) (*stripe.TreasuryOutboundTransfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
