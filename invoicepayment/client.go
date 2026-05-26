//
//
// File generated from our OpenAPI spec
//
//

// Package invoicepayment provides the /v1/invoice_payments APIs
package invoicepayment

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/invoice_payments APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the invoice payment with the given ID.
func Get(id string, params *stripe.InvoicePaymentParams) (*stripe.InvoicePayment, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves the invoice payment with the given ID.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.InvoicePaymentParams) (*stripe.InvoicePayment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// When retrieving an invoice, there is an includable payments property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of payments.
func List(params *stripe.InvoicePaymentListParams) *Iter { _ = "STUB: not implemented"; return nil }

// When retrieving an invoice, there is an includable payments property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of payments.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.InvoicePaymentListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for invoice payments.
type Iter struct {
	*stripe.Iter
}

// InvoicePayment returns the invoice payment which the iterator is currently pointing to.
func (i *Iter) InvoicePayment() *stripe.InvoicePayment { _ = "STUB: not implemented"; return nil }

// InvoicePaymentList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) InvoicePaymentList() *stripe.InvoicePaymentList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
