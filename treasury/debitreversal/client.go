//
//
// File generated from our OpenAPI spec
//
//

// Package debitreversal provides the /v1/treasury/debit_reversals APIs
package debitreversal

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/treasury/debit_reversals APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Reverses a ReceivedDebit and creates a DebitReversal object.
func New(params *stripe.TreasuryDebitReversalParams) (*stripe.TreasuryDebitReversal, error) {
	_ = "STUB: not implemented"
	return nil,

		// Reverses a ReceivedDebit and creates a DebitReversal object.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.TreasuryDebitReversalParams) (*stripe.TreasuryDebitReversal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a DebitReversal object.
func Get(id string, params *stripe.TreasuryDebitReversalParams) (*stripe.TreasuryDebitReversal, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves a DebitReversal object.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.TreasuryDebitReversalParams) (*stripe.TreasuryDebitReversal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of DebitReversals.
func List(params *stripe.TreasuryDebitReversalListParams) *Iter {
	_ = "STUB: not implemented"
	return nil

	// Returns a list of DebitReversals.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) List(listParams *stripe.TreasuryDebitReversalListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for treasury debit reversals.
type Iter struct {
	*stripe.Iter
}

// TreasuryDebitReversal returns the treasury debit reversal which the iterator is currently pointing to.
func (i *Iter) TreasuryDebitReversal() *stripe.TreasuryDebitReversal {
	_ = "STUB: not implemented"
	return nil
}

// TreasuryDebitReversalList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TreasuryDebitReversalList() *stripe.TreasuryDebitReversalList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
