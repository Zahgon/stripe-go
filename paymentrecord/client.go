//
//
// File generated from our OpenAPI spec
//
//

// Package paymentrecord provides the /v1/payment_records APIs
package paymentrecord

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/payment_records APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves a Payment Record with the given ID
func Get(id string, params *stripe.PaymentRecordParams) (*stripe.PaymentRecord, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves a Payment Record with the given ID
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.PaymentRecordParams) (*stripe.PaymentRecord, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Report a new Payment Record. You may report a Payment Record as it is
//
//	initialized and later report updates through the other report_* methods, or report Payment
//	Records in a terminal state directly, through this method.
func ReportPayment(params *stripe.PaymentRecordReportPaymentParams) (*stripe.PaymentRecord, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Report a new Payment Record. You may report a Payment Record as it is
//
//	initialized and later report updates through the other report_* methods, or report Payment
//	Records in a terminal state directly, through this method.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ReportPayment(params *stripe.PaymentRecordReportPaymentParams) (*stripe.PaymentRecord, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Report a new payment attempt on the specified Payment Record. A new payment
//
//	attempt can only be specified if all other payment attempts are canceled or failed.
func ReportPaymentAttempt(id string, params *stripe.PaymentRecordReportPaymentAttemptParams) (*stripe.PaymentRecord, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Report a new payment attempt on the specified Payment Record. A new payment
//
//	attempt can only be specified if all other payment attempts are canceled or failed.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ReportPaymentAttempt(id string, params *stripe.PaymentRecordReportPaymentAttemptParams) (*stripe.PaymentRecord, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Report that the most recent payment attempt on the specified Payment Record
//
//	was canceled.
func ReportPaymentAttemptCanceled(id string, params *stripe.PaymentRecordReportPaymentAttemptCanceledParams) (*stripe.PaymentRecord, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Report that the most recent payment attempt on the specified Payment Record
//
//	was canceled.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ReportPaymentAttemptCanceled(id string, params *stripe.PaymentRecordReportPaymentAttemptCanceledParams) (*stripe.PaymentRecord, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Report that the most recent payment attempt on the specified Payment Record
//
//	failed or errored.
func ReportPaymentAttemptFailed(id string, params *stripe.PaymentRecordReportPaymentAttemptFailedParams) (*stripe.PaymentRecord, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Report that the most recent payment attempt on the specified Payment Record
//
//	failed or errored.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ReportPaymentAttemptFailed(id string, params *stripe.PaymentRecordReportPaymentAttemptFailedParams) (*stripe.PaymentRecord, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Report that the most recent payment attempt on the specified Payment Record
//
//	was guaranteed.
func ReportPaymentAttemptGuaranteed(id string, params *stripe.PaymentRecordReportPaymentAttemptGuaranteedParams) (*stripe.PaymentRecord, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Report that the most recent payment attempt on the specified Payment Record
//
//	was guaranteed.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ReportPaymentAttemptGuaranteed(id string, params *stripe.PaymentRecordReportPaymentAttemptGuaranteedParams) (*stripe.PaymentRecord, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Report informational updates on the specified Payment Record.
func ReportPaymentAttemptInformational(id string, params *stripe.PaymentRecordReportPaymentAttemptInformationalParams) (*stripe.PaymentRecord, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Report informational updates on the specified Payment Record.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ReportPaymentAttemptInformational(id string, params *stripe.PaymentRecordReportPaymentAttemptInformationalParams) (*stripe.PaymentRecord, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Report that the most recent payment attempt on the specified Payment Record
//
//	was refunded.
func ReportRefund(id string, params *stripe.PaymentRecordReportRefundParams) (*stripe.PaymentRecord, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Report that the most recent payment attempt on the specified Payment Record
//
//	was refunded.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ReportRefund(id string, params *stripe.PaymentRecordReportRefundParams) (*stripe.PaymentRecord, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
