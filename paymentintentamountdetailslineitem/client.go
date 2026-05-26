//
//
// File generated from our OpenAPI spec
//
//

// Package paymentintentamountdetailslineitem provides the /v1/payment_intents/{intent}/amount_details_line_items APIs
package paymentintentamountdetailslineitem

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/payment_intents/{intent}/amount_details_line_items APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Lists all LineItems of a given PaymentIntent.
func List(params *stripe.PaymentIntentAmountDetailsLineItemListParams) *Iter {
	_ = "STUB: not implemented"
	return nil

	// Lists all LineItems of a given PaymentIntent.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) List(listParams *stripe.PaymentIntentAmountDetailsLineItemListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for payment intent amount details line items.
type Iter struct {
	*stripe.Iter
}

// PaymentIntentAmountDetailsLineItem returns the payment intent amount details line item which the iterator is currently pointing to.
func (i *Iter) PaymentIntentAmountDetailsLineItem() *stripe.PaymentIntentAmountDetailsLineItem {
	_ = "STUB: not implemented"
	return nil
}

// PaymentIntentAmountDetailsLineItemList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) PaymentIntentAmountDetailsLineItemList() *stripe.PaymentIntentAmountDetailsLineItemList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
