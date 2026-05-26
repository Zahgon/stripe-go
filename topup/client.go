//
//
// File generated from our OpenAPI spec
//
//

// Package topup provides the /v1/topups APIs
package topup

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/topups APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Top up the balance of an account
func New(params *stripe.TopupParams) (*stripe.Topup, error) {
	_ = "STUB: not implemented"
	return nil,

		// Top up the balance of an account
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.TopupParams) (*stripe.Topup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the details of a top-up that has previously been created. Supply the unique top-up ID that was returned from your previous request, and Stripe will return the corresponding top-up information.
func Get(id string, params *stripe.TopupParams) (*stripe.Topup, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves the details of a top-up that has previously been created. Supply the unique top-up ID that was returned from your previous request, and Stripe will return the corresponding top-up information.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.TopupParams) (*stripe.Topup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the metadata of a top-up. Other top-up details are not editable by design.
func Update(id string, params *stripe.TopupParams) (*stripe.Topup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the metadata of a top-up. Other top-up details are not editable by design.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.TopupParams) (*stripe.Topup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Cancels a top-up. Only pending top-ups can be canceled.
func Cancel(id string, params *stripe.TopupParams) (*stripe.Topup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Cancels a top-up. Only pending top-ups can be canceled.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Cancel(id string, params *stripe.TopupParams) (*stripe.Topup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of top-ups.
func List(params *stripe.TopupListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Returns a list of top-ups.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.TopupListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for topups.
type Iter struct {
	*stripe.Iter
}

// Topup returns the topup which the iterator is currently pointing to.
func (i *Iter) Topup() *stripe.Topup { _ = "STUB: not implemented"; return nil }

// TopupList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TopupList() *stripe.TopupList { _ = "STUB: not implemented"; return nil }

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
