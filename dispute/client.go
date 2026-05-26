//
//
// File generated from our OpenAPI spec
//
//

// Package dispute provides the /v1/disputes APIs
package dispute

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/disputes APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the dispute with the given ID.
func Get(id string, params *stripe.DisputeParams) (*stripe.Dispute, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves the dispute with the given ID.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.DisputeParams) (*stripe.Dispute, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// When you get a dispute, contacting your customer is always the best first step. If that doesn't work, you can submit evidence to help us resolve the dispute in your favor. You can do this in your [dashboard](https://dashboard.stripe.com/disputes), but if you prefer, you can use the API to submit evidence programmatically.
//
// Depending on your dispute type, different evidence fields will give you a better chance of winning your dispute. To figure out which evidence fields to provide, see our [guide to dispute types](https://docs.stripe.com/docs/disputes/categories).
func Update(id string, params *stripe.DisputeParams) (*stripe.Dispute, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// When you get a dispute, contacting your customer is always the best first step. If that doesn't work, you can submit evidence to help us resolve the dispute in your favor. You can do this in your [dashboard](https://dashboard.stripe.com/disputes), but if you prefer, you can use the API to submit evidence programmatically.
//
// Depending on your dispute type, different evidence fields will give you a better chance of winning your dispute. To figure out which evidence fields to provide, see our [guide to dispute types](https://docs.stripe.com/docs/disputes/categories).
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.DisputeParams) (*stripe.Dispute, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Closing the dispute for a charge indicates that you do not have any evidence to submit and are essentially dismissing the dispute, acknowledging it as lost.
//
// The status of the dispute will change from needs_response to lost. Closing a dispute is irreversible.
func Close(id string, params *stripe.DisputeParams) (*stripe.Dispute, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Closing the dispute for a charge indicates that you do not have any evidence to submit and are essentially dismissing the dispute, acknowledging it as lost.
//
// The status of the dispute will change from needs_response to lost. Closing a dispute is irreversible.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Close(id string, params *stripe.DisputeParams) (*stripe.Dispute, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of your disputes.
func List(params *stripe.DisputeListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Returns a list of your disputes.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.DisputeListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for disputes.
type Iter struct {
	*stripe.Iter
}

// Dispute returns the dispute which the iterator is currently pointing to.
func (i *Iter) Dispute() *stripe.Dispute { _ = "STUB: not implemented"; return nil }

// DisputeList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) DisputeList() *stripe.DisputeList { _ = "STUB: not implemented"; return nil }

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
