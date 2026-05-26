//
//
// File generated from our OpenAPI spec
//
//

// Package transaction provides the /v1/financial_connections/transactions APIs
package transaction

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/financial_connections/transactions APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the details of a Financial Connections Transaction
func Get(id string, params *stripe.FinancialConnectionsTransactionParams) (*stripe.FinancialConnectionsTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves the details of a Financial Connections Transaction
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.FinancialConnectionsTransactionParams) (*stripe.FinancialConnectionsTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of Financial Connections Transaction objects.
func List(params *stripe.FinancialConnectionsTransactionListParams) *Iter {
	_ = "STUB: not implemented"
	return nil

	// Returns a list of Financial Connections Transaction objects.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) List(listParams *stripe.FinancialConnectionsTransactionListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for financial connections transactions.
type Iter struct {
	*stripe.Iter
}

// FinancialConnectionsTransaction returns the financial connections transaction which the iterator is currently pointing to.
func (i *Iter) FinancialConnectionsTransaction() *stripe.FinancialConnectionsTransaction {
	_ = "STUB: not implemented"
	return nil
}

// FinancialConnectionsTransactionList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) FinancialConnectionsTransactionList() *stripe.FinancialConnectionsTransactionList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
