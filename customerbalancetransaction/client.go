//
//
// File generated from our OpenAPI spec
//
//

// Package customerbalancetransaction provides the /v1/customers/{customer}/balance_transactions APIs
package customerbalancetransaction

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/customers/{customer}/balance_transactions APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates an immutable transaction that updates the customer's credit [balance](https://docs.stripe.com/docs/billing/customer/balance).
func New(params *stripe.CustomerBalanceTransactionParams) (*stripe.CustomerBalanceTransaction, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates an immutable transaction that updates the customer's credit [balance](https://docs.stripe.com/docs/billing/customer/balance).
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.CustomerBalanceTransactionParams) (*stripe.CustomerBalanceTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a specific customer balance transaction that updated the customer's [balances](https://docs.stripe.com/docs/billing/customer/balance).
func Get(id string, params *stripe.CustomerBalanceTransactionParams) (*stripe.CustomerBalanceTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves a specific customer balance transaction that updated the customer's [balances](https://docs.stripe.com/docs/billing/customer/balance).
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.CustomerBalanceTransactionParams) (*stripe.CustomerBalanceTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Most credit balance transaction fields are immutable, but you may update its description and metadata.
func Update(id string, params *stripe.CustomerBalanceTransactionParams) (*stripe.CustomerBalanceTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Most credit balance transaction fields are immutable, but you may update its description and metadata.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.CustomerBalanceTransactionParams) (*stripe.CustomerBalanceTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of transactions that updated the customer's [balances](https://docs.stripe.com/docs/billing/customer/balance).
func List(params *stripe.CustomerBalanceTransactionListParams) *Iter {
	_ = "STUB: not implemented"
	return nil

	// Returns a list of transactions that updated the customer's [balances](https://docs.stripe.com/docs/billing/customer/balance).
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) List(listParams *stripe.CustomerBalanceTransactionListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for customer balance transactions.
type Iter struct {
	*stripe.Iter
}

// CustomerBalanceTransaction returns the customer balance transaction which the iterator is currently pointing to.
func (i *Iter) CustomerBalanceTransaction() *stripe.CustomerBalanceTransaction {
	_ = "STUB: not implemented"
	return nil
}

// CustomerBalanceTransactionList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) CustomerBalanceTransactionList() *stripe.CustomerBalanceTransactionList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
