//
//
// File generated from our OpenAPI spec
//
//

// Package review provides the /v1/reviews APIs
package review

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/reviews APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves a Review object.
func Get(id string, params *stripe.ReviewParams) (*stripe.Review, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves a Review object.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.ReviewParams) (*stripe.Review, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Approves a Review object, closing it and removing it from the list of reviews.
func Approve(id string, params *stripe.ReviewApproveParams) (*stripe.Review, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Approves a Review object, closing it and removing it from the list of reviews.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Approve(id string, params *stripe.ReviewApproveParams) (*stripe.Review, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of Review objects that have open set to true. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
func List(params *stripe.ReviewListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Returns a list of Review objects that have open set to true. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.ReviewListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for reviews.
type Iter struct {
	*stripe.Iter
}

// Review returns the review which the iterator is currently pointing to.
func (i *Iter) Review() *stripe.Review { _ = "STUB: not implemented"; return nil }

// ReviewList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) ReviewList() *stripe.ReviewList { _ = "STUB: not implemented"; return nil }

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
