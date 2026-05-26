//
//
// File generated from our OpenAPI spec
//
//

// Package session provides the /v1/checkout/sessions APIs
package session

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/checkout/sessions APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a Checkout Session object.
func New(params *stripe.CheckoutSessionParams) (*stripe.CheckoutSession, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates a Checkout Session object.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.CheckoutSessionParams) (*stripe.CheckoutSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a Checkout Session object.
func Get(id string, params *stripe.CheckoutSessionParams) (*stripe.CheckoutSession, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves a Checkout Session object.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.CheckoutSessionParams) (*stripe.CheckoutSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates a Checkout Session object.
//
// Related guide: [Dynamically update a Checkout Session](https://docs.stripe.com/payments/advanced/dynamic-updates)
func Update(id string, params *stripe.CheckoutSessionParams) (*stripe.CheckoutSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates a Checkout Session object.
//
// Related guide: [Dynamically update a Checkout Session](https://docs.stripe.com/payments/advanced/dynamic-updates)
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.CheckoutSessionParams) (*stripe.CheckoutSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// A Checkout Session can be expired when it is in one of these statuses: open
//
// After it expires, a customer can't complete a Checkout Session and customers loading the Checkout Session see a message saying the Checkout Session is expired.
func Expire(id string, params *stripe.CheckoutSessionExpireParams) (*stripe.CheckoutSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// A Checkout Session can be expired when it is in one of these statuses: open
//
// After it expires, a customer can't complete a Checkout Session and customers loading the Checkout Session see a message saying the Checkout Session is expired.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Expire(id string, params *stripe.CheckoutSessionExpireParams) (*stripe.CheckoutSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of Checkout Sessions.
func List(params *stripe.CheckoutSessionListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Returns a list of Checkout Sessions.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.CheckoutSessionListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for checkout sessions.
type Iter struct {
	*stripe.Iter
}

// CheckoutSession returns the checkout session which the iterator is currently pointing to.
func (i *Iter) CheckoutSession() *stripe.CheckoutSession { _ = "STUB: not implemented"; return nil }

// CheckoutSessionList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) CheckoutSessionList() *stripe.CheckoutSessionList {
	_ = "STUB: not implemented"
	return nil
}

// When retrieving a Checkout Session, there is an includable line_items property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
func ListLineItems(params *stripe.CheckoutSessionListLineItemsParams) *LineItemIter {
	_ = "STUB: not implemented"
	return nil
}

// When retrieving a Checkout Session, there is an includable line_items property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ListLineItems(listParams *stripe.CheckoutSessionListLineItemsParams) *LineItemIter {
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
