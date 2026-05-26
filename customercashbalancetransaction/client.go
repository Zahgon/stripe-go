//
//
// File generated from our OpenAPI spec
//
//

// Package customercashbalancetransaction provides the /v1/customers/{customer}/cash_balance_transactions APIs
package customercashbalancetransaction

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/customers/{customer}/cash_balance_transactions APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves a specific cash balance transaction, which updated the customer's [cash balance](https://docs.stripe.com/docs/payments/customer-balance).
func Get(id string, params *stripe.CustomerCashBalanceTransactionParams) (*stripe.CustomerCashBalanceTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves a specific cash balance transaction, which updated the customer's [cash balance](https://docs.stripe.com/docs/payments/customer-balance).
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.CustomerCashBalanceTransactionParams) (*stripe.CustomerCashBalanceTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of transactions that modified the customer's [cash balance](https://docs.stripe.com/docs/payments/customer-balance).
func List(params *stripe.CustomerCashBalanceTransactionListParams) *Iter {
	_ = "STUB: not implemented"
	return nil

	// Returns a list of transactions that modified the customer's [cash balance](https://docs.stripe.com/docs/payments/customer-balance).
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) List(listParams *stripe.CustomerCashBalanceTransactionListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for customer cash balance transactions.
type Iter struct {
	*stripe.Iter
}

// CustomerCashBalanceTransaction returns the customer cash balance transaction which the iterator is currently pointing to.
func (i *Iter) CustomerCashBalanceTransaction() *stripe.CustomerCashBalanceTransaction {
	_ = "STUB: not implemented"
	return nil
}

// CustomerCashBalanceTransactionList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) CustomerCashBalanceTransactionList() *stripe.CustomerCashBalanceTransactionList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
