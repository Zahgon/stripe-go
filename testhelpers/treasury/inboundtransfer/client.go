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

// Transitions a test mode created InboundTransfer to the failed status. The InboundTransfer must already be in the processing state.
func Fail(id string, params *stripe.TestHelpersTreasuryInboundTransferFailParams) (*stripe.TreasuryInboundTransfer, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Transitions a test mode created InboundTransfer to the failed status. The InboundTransfer must already be in the processing state.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Fail(id string, params *stripe.TestHelpersTreasuryInboundTransferFailParams) (*stripe.TreasuryInboundTransfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Marks the test mode InboundTransfer object as returned and links the InboundTransfer to a ReceivedDebit. The InboundTransfer must already be in the succeeded state.
func ReturnInboundTransfer(id string, params *stripe.TestHelpersTreasuryInboundTransferReturnInboundTransferParams) (*stripe.TreasuryInboundTransfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Marks the test mode InboundTransfer object as returned and links the InboundTransfer to a ReceivedDebit. The InboundTransfer must already be in the succeeded state.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ReturnInboundTransfer(id string, params *stripe.TestHelpersTreasuryInboundTransferReturnInboundTransferParams) (*stripe.TreasuryInboundTransfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Transitions a test mode created InboundTransfer to the succeeded status. The InboundTransfer must already be in the processing state.
func Succeed(id string, params *stripe.TestHelpersTreasuryInboundTransferSucceedParams) (*stripe.TreasuryInboundTransfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Transitions a test mode created InboundTransfer to the succeeded status. The InboundTransfer must already be in the processing state.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Succeed(id string, params *stripe.TestHelpersTreasuryInboundTransferSucceedParams) (*stripe.TreasuryInboundTransfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
