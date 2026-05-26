//
//
// File generated from our OpenAPI spec
//
//

// Package cashbalance provides the /v1/customers/{customer}/cash_balance APIs
package cashbalance

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/customers/{customer}/cash_balance APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves a customer's cash balance.
func Get(params *stripe.CashBalanceParams) (*stripe.CashBalance, error) {
	_ = "STUB: not implemented"
	return nil,

		// Retrieves a customer's cash balance.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) Get(params *stripe.CashBalanceParams) (*stripe.CashBalance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Changes the settings on a customer's cash balance.
func Update(params *stripe.CashBalanceParams) (*stripe.CashBalance, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Changes the settings on a customer's cash balance.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Update(params *stripe.CashBalanceParams) (*stripe.CashBalance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
