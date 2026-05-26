//
//
// File generated from our OpenAPI spec
//
//

// Package event provides the event related APIs
package event

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke event related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the details of an event.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2CoreEventParams) (stripe.V2CoreEvent, error) {
	_ = "STUB: not implemented"
	return *new(stripe.V2CoreEvent), nil
}

// List events, going back up to 30 days.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2CoreEventListParams) stripe.Seq2[stripe.V2CoreEvent, error] {
	_ = "STUB: not implemented"
	return nil
}
