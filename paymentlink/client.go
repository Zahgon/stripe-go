//
//
// File generated from our OpenAPI spec
//
//

// Package paymentlink provides the /v1/payment_links APIs
package paymentlink

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/payment_links APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a payment link.
func New(params *stripe.PaymentLinkParams) (*stripe.PaymentLink, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates a payment link.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.PaymentLinkParams) (*stripe.PaymentLink, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieve a payment link.
func Get(id string, params *stripe.PaymentLinkParams) (*stripe.PaymentLink, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieve a payment link.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.PaymentLinkParams) (*stripe.PaymentLink, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates a payment link.
func Update(id string, params *stripe.PaymentLinkParams) (*stripe.PaymentLink, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates a payment link.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.PaymentLinkParams) (*stripe.PaymentLink, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of your payment links.
func List(params *stripe.PaymentLinkListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Returns a list of your payment links.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.PaymentLinkListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for payment links.
type Iter struct {
	*stripe.Iter
}

// PaymentLink returns the payment link which the iterator is currently pointing to.
func (i *Iter) PaymentLink() *stripe.PaymentLink { _ = "STUB: not implemented"; return nil }

// PaymentLinkList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) PaymentLinkList() *stripe.PaymentLinkList { _ = "STUB: not implemented"; return nil }

// When retrieving a payment link, there is an includable line_items property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
func ListLineItems(params *stripe.PaymentLinkListLineItemsParams) *LineItemIter {
	_ = "STUB: not implemented"
	return nil
}

// When retrieving a payment link, there is an includable line_items property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ListLineItems(listParams *stripe.PaymentLinkListLineItemsParams) *LineItemIter {
	_ = "STUB: not implemented"
	return nil
}

// LineItemIter is an iterator for line items.
type LineItemIter struct {
	*stripe.Iter
}

// LineItem returns the line item which the iterator is currently pointing to.
func (i *LineItemIter) LineItem() *stripe.LineItem { _ = "STUB: not implemented"; return nil }

// LineItemList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *LineItemIter) LineItemList() *stripe.LineItemList { _ = "STUB: not implemented"; return nil }

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
