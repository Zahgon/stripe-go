//
//
// File generated from our OpenAPI spec
//
//

// Package sourcetransaction provides the sourcetransaction related APIs
package sourcetransaction

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /sources/:source_id/transactions APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// List source transactions for a given source.
func List(params *stripe.SourceTransactionListParams) *Iter { _ = "STUB: not implemented"; return nil }

// List source transactions for a given source.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.SourceTransactionListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for source transactions.
type Iter struct {
	*stripe.Iter
}

// SourceTransaction returns the source transaction which the iterator is currently pointing to.
func (i *Iter) SourceTransaction() *stripe.SourceTransaction { _ = "STUB: not implemented"; return nil }

// SourceTransactionList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) SourceTransactionList() *stripe.SourceTransactionList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
