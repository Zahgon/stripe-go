//
//
// File generated from our OpenAPI spec
//
//

// Package session provides the /v1/financial_connections/sessions APIs
package session

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/financial_connections/sessions APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// To launch the Financial Connections authorization flow, create a Session. The session's client_secret can be used to launch the flow using Stripe.js.
func New(params *stripe.FinancialConnectionsSessionParams) (*stripe.FinancialConnectionsSession, error) {
	_ = "STUB: not implemented"
	return nil,

		// To launch the Financial Connections authorization flow, create a Session. The session's client_secret can be used to launch the flow using Stripe.js.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.FinancialConnectionsSessionParams) (*stripe.FinancialConnectionsSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the details of a Financial Connections Session
func Get(id string, params *stripe.FinancialConnectionsSessionParams) (*stripe.FinancialConnectionsSession, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves the details of a Financial Connections Session
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.FinancialConnectionsSessionParams) (*stripe.FinancialConnectionsSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
