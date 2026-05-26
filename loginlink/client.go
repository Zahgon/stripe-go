//
//
// File generated from our OpenAPI spec
//
//

// Package loginlink provides the /v1/accounts/{account}/login_links APIs
package loginlink

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/accounts/{account}/login_links APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a login link for a connected account to access the Express Dashboard.
//
// You can only create login links for accounts that use the [Express Dashboard](https://docs.stripe.com/connect/express-dashboard) and are connected to your platform.
func New(params *stripe.LoginLinkParams) (*stripe.LoginLink, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates a login link for a connected account to access the Express Dashboard.
		//
		// You can only create login links for accounts that use the [Express Dashboard](https://docs.stripe.com/connect/express-dashboard) and are connected to your platform.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.LoginLinkParams) (*stripe.LoginLink, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
