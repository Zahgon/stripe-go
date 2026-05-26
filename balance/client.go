//
//
// File generated from our OpenAPI spec
//
//

// Package balance provides the /v1/balance APIs
package balance

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/balance APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the current account balance, based on the authentication that was used to make the request.
//
//	For a sample request, see [Accounting for negative balances](https://docs.stripe.com/docs/connect/account-balances#accounting-for-negative-balances).
func Get(params *stripe.BalanceParams) (*stripe.Balance, error) {
	_ = "STUB: not implemented"
	return nil,

		// Retrieves the current account balance, based on the authentication that was used to make the request.
		//
		//	For a sample request, see [Accounting for negative balances](https://docs.stripe.com/docs/connect/account-balances#accounting-for-negative-balances).
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) Get(params *stripe.BalanceParams) (*stripe.Balance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
