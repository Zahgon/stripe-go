//
//
// File generated from our OpenAPI spec
//
//

// Package balancetransaction provides the /v1/balance_transactions APIs
package balancetransaction

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/balance_transactions APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the balance transaction with the given ID.
//
// Note that this endpoint previously used the path /v1/balance/history/:id.
func Get(id string, params *stripe.BalanceTransactionParams) (*stripe.BalanceTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves the balance transaction with the given ID.
	//
	// Note that this endpoint previously used the path /v1/balance/history/:id.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.BalanceTransactionParams) (*stripe.BalanceTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of transactions that have contributed to the Stripe account balance (e.g., charges, transfers, and so forth). The transactions are returned in sorted order, with the most recent transactions appearing first.
//
// Note that this endpoint was previously called “Balance history” and used the path /v1/balance/history.
func List(params *stripe.BalanceTransactionListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Returns a list of transactions that have contributed to the Stripe account balance (e.g., charges, transfers, and so forth). The transactions are returned in sorted order, with the most recent transactions appearing first.
//
// Note that this endpoint was previously called “Balance history” and used the path /v1/balance/history.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.BalanceTransactionListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for balance transactions.
type Iter struct {
	*stripe.Iter
}

// BalanceTransaction returns the balance transaction which the iterator is currently pointing to.
func (i *Iter) BalanceTransaction() *stripe.BalanceTransaction {
	_ = "STUB: not implemented"
	return nil
}

// BalanceTransactionList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) BalanceTransactionList() *stripe.BalanceTransactionList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
