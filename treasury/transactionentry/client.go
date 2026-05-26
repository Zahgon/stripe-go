//
//
// File generated from our OpenAPI spec
//
//

// Package transactionentry provides the /v1/treasury/transaction_entries APIs
package transactionentry

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/treasury/transaction_entries APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves a TransactionEntry object.
func Get(id string, params *stripe.TreasuryTransactionEntryParams) (*stripe.TreasuryTransactionEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves a TransactionEntry object.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.TreasuryTransactionEntryParams) (*stripe.TreasuryTransactionEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a list of TransactionEntry objects.
func List(params *stripe.TreasuryTransactionEntryListParams) *Iter {
	_ = "STUB: not implemented"
	return nil

	// Retrieves a list of TransactionEntry objects.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) List(listParams *stripe.TreasuryTransactionEntryListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for treasury transaction entries.
type Iter struct {
	*stripe.Iter
}

// TreasuryTransactionEntry returns the treasury transaction entry which the iterator is currently pointing to.
func (i *Iter) TreasuryTransactionEntry() *stripe.TreasuryTransactionEntry {
	_ = "STUB: not implemented"
	return nil
}

// TreasuryTransactionEntryList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TreasuryTransactionEntryList() *stripe.TreasuryTransactionEntryList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
