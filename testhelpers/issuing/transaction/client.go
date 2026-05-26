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

// Allows the user to capture an arbitrary amount, also known as a forced capture.
func CreateForceCapture(params *stripe.TestHelpersIssuingTransactionCreateForceCaptureParams) (*stripe.IssuingTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Allows the user to capture an arbitrary amount, also known as a forced capture.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) CreateForceCapture(params *stripe.TestHelpersIssuingTransactionCreateForceCaptureParams) (*stripe.IssuingTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Allows the user to refund an arbitrary amount, also known as a unlinked refund.
func CreateUnlinkedRefund(params *stripe.TestHelpersIssuingTransactionCreateUnlinkedRefundParams) (*stripe.IssuingTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Allows the user to refund an arbitrary amount, also known as a unlinked refund.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) CreateUnlinkedRefund(params *stripe.TestHelpersIssuingTransactionCreateUnlinkedRefundParams) (*stripe.IssuingTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Refund a test-mode Transaction.
func Refund(id string, params *stripe.TestHelpersIssuingTransactionRefundParams) (*stripe.IssuingTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Refund a test-mode Transaction.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Refund(id string, params *stripe.TestHelpersIssuingTransactionRefundParams) (*stripe.IssuingTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
