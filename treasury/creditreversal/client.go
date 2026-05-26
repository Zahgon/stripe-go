//
//
// File generated from our OpenAPI spec
//
//

// Package creditreversal provides the /v1/treasury/credit_reversals APIs
package creditreversal

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/treasury/credit_reversals APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Reverses a ReceivedCredit and creates a CreditReversal object.
func New(params *stripe.TreasuryCreditReversalParams) (*stripe.TreasuryCreditReversal, error) {
	_ = "STUB: not implemented"
	return nil,

		// Reverses a ReceivedCredit and creates a CreditReversal object.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.TreasuryCreditReversalParams) (*stripe.TreasuryCreditReversal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the details of an existing CreditReversal by passing the unique CreditReversal ID from either the CreditReversal creation request or CreditReversal list
func Get(id string, params *stripe.TreasuryCreditReversalParams) (*stripe.TreasuryCreditReversal, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves the details of an existing CreditReversal by passing the unique CreditReversal ID from either the CreditReversal creation request or CreditReversal list
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.TreasuryCreditReversalParams) (*stripe.TreasuryCreditReversal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of CreditReversals.
func List(params *stripe.TreasuryCreditReversalListParams) *Iter {
	_ = "STUB: not implemented"
	return nil

	// Returns a list of CreditReversals.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) List(listParams *stripe.TreasuryCreditReversalListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for treasury credit reversals.
type Iter struct {
	*stripe.Iter
}

// TreasuryCreditReversal returns the treasury credit reversal which the iterator is currently pointing to.
func (i *Iter) TreasuryCreditReversal() *stripe.TreasuryCreditReversal {
	_ = "STUB: not implemented"
	return nil
}

// TreasuryCreditReversalList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TreasuryCreditReversalList() *stripe.TreasuryCreditReversalList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
