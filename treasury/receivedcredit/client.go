//
//
// File generated from our OpenAPI spec
//
//

// Package receivedcredit provides the /v1/treasury/received_credits APIs
package receivedcredit

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/treasury/received_credits APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the details of an existing ReceivedCredit by passing the unique ReceivedCredit ID from the ReceivedCredit list.
func Get(id string, params *stripe.TreasuryReceivedCreditParams) (*stripe.TreasuryReceivedCredit, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves the details of an existing ReceivedCredit by passing the unique ReceivedCredit ID from the ReceivedCredit list.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.TreasuryReceivedCreditParams) (*stripe.TreasuryReceivedCredit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of ReceivedCredits.
func List(params *stripe.TreasuryReceivedCreditListParams) *Iter {
	_ = "STUB: not implemented"
	return nil

	// Returns a list of ReceivedCredits.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) List(listParams *stripe.TreasuryReceivedCreditListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for treasury received credits.
type Iter struct {
	*stripe.Iter
}

// TreasuryReceivedCredit returns the treasury received credit which the iterator is currently pointing to.
func (i *Iter) TreasuryReceivedCredit() *stripe.TreasuryReceivedCredit {
	_ = "STUB: not implemented"
	return nil
}

// TreasuryReceivedCreditList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TreasuryReceivedCreditList() *stripe.TreasuryReceivedCreditList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
