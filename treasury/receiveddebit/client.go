//
//
// File generated from our OpenAPI spec
//
//

// Package receiveddebit provides the /v1/treasury/received_debits APIs
package receiveddebit

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/treasury/received_debits APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the details of an existing ReceivedDebit by passing the unique ReceivedDebit ID from the ReceivedDebit list
func Get(id string, params *stripe.TreasuryReceivedDebitParams) (*stripe.TreasuryReceivedDebit, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves the details of an existing ReceivedDebit by passing the unique ReceivedDebit ID from the ReceivedDebit list
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.TreasuryReceivedDebitParams) (*stripe.TreasuryReceivedDebit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of ReceivedDebits.
func List(params *stripe.TreasuryReceivedDebitListParams) *Iter {
	_ = "STUB: not implemented"
	return nil

	// Returns a list of ReceivedDebits.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) List(listParams *stripe.TreasuryReceivedDebitListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for treasury received debits.
type Iter struct {
	*stripe.Iter
}

// TreasuryReceivedDebit returns the treasury received debit which the iterator is currently pointing to.
func (i *Iter) TreasuryReceivedDebit() *stripe.TreasuryReceivedDebit {
	_ = "STUB: not implemented"
	return nil
}

// TreasuryReceivedDebitList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TreasuryReceivedDebitList() *stripe.TreasuryReceivedDebitList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
