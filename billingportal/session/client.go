//
//
// File generated from our OpenAPI spec
//
//

// Package session provides the /v1/billing_portal/sessions APIs
package session

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/billing_portal/sessions APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a session of the customer portal.
func New(params *stripe.BillingPortalSessionParams) (*stripe.BillingPortalSession, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates a session of the customer portal.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.BillingPortalSessionParams) (*stripe.BillingPortalSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
