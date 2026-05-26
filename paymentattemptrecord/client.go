//
//
// File generated from our OpenAPI spec
//
//

// Package paymentattemptrecord provides the /v1/payment_attempt_records APIs
package paymentattemptrecord

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/payment_attempt_records APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves a Payment Attempt Record with the given ID
func Get(id string, params *stripe.PaymentAttemptRecordParams) (*stripe.PaymentAttemptRecord, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves a Payment Attempt Record with the given ID
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.PaymentAttemptRecordParams) (*stripe.PaymentAttemptRecord, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// List all the Payment Attempt Records attached to the specified Payment Record.
func List(params *stripe.PaymentAttemptRecordListParams) *Iter {
	_ = "STUB: not implemented"
	return nil

	// List all the Payment Attempt Records attached to the specified Payment Record.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) List(listParams *stripe.PaymentAttemptRecordListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for payment attempt records.
type Iter struct {
	*stripe.Iter
}

// PaymentAttemptRecord returns the payment attempt record which the iterator is currently pointing to.
func (i *Iter) PaymentAttemptRecord() *stripe.PaymentAttemptRecord {
	_ = "STUB: not implemented"
	return nil
}

// PaymentAttemptRecordList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) PaymentAttemptRecordList() *stripe.PaymentAttemptRecordList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
