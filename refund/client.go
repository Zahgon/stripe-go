//
//
// File generated from our OpenAPI spec
//
//

// Package refund provides the /v1/refunds APIs
package refund

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/refunds APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// When you create a new refund, you must specify a Charge or a PaymentIntent object on which to create it.
//
// Creating a new refund will refund a charge that has previously been created but not yet refunded.
// Funds will be refunded to the credit or debit card that was originally charged.
//
// You can optionally refund only part of a charge.
// You can do so multiple times, until the entire charge has been refunded.
//
// Once entirely refunded, a charge can't be refunded again.
// This method will raise an error when called on an already-refunded charge,
// or when trying to refund more money than is left on a charge.
func New(params *stripe.RefundParams) (*stripe.Refund, error) {
	_ = "STUB: not implemented"
	return nil,

		// When you create a new refund, you must specify a Charge or a PaymentIntent object on which to create it.
		//
		// Creating a new refund will refund a charge that has previously been created but not yet refunded.
		// Funds will be refunded to the credit or debit card that was originally charged.
		//
		// You can optionally refund only part of a charge.
		// You can do so multiple times, until the entire charge has been refunded.
		//
		// Once entirely refunded, a charge can't be refunded again.
		// This method will raise an error when called on an already-refunded charge,
		// or when trying to refund more money than is left on a charge.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.RefundParams) (*stripe.Refund, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the details of an existing refund.
func Get(id string, params *stripe.RefundParams) (*stripe.Refund, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves the details of an existing refund.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.RefundParams) (*stripe.Refund, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the refund that you specify by setting the values of the passed parameters. Any parameters that you don't provide remain unchanged.
//
// This request only accepts metadata as an argument.
func Update(id string, params *stripe.RefundParams) (*stripe.Refund, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the refund that you specify by setting the values of the passed parameters. Any parameters that you don't provide remain unchanged.
//
// This request only accepts metadata as an argument.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.RefundParams) (*stripe.Refund, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Cancels a refund with a status of requires_action.
//
// You can't cancel refunds in other states. Only refunds for payment methods that require customer action can enter the requires_action state.
func Cancel(id string, params *stripe.RefundCancelParams) (*stripe.Refund, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Cancels a refund with a status of requires_action.
//
// You can't cancel refunds in other states. Only refunds for payment methods that require customer action can enter the requires_action state.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Cancel(id string, params *stripe.RefundCancelParams) (*stripe.Refund, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of all refunds you created. We return the refunds in sorted order, with the most recent refunds appearing first. The 10 most recent refunds are always available by default on the Charge object.
func List(params *stripe.RefundListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Returns a list of all refunds you created. We return the refunds in sorted order, with the most recent refunds appearing first. The 10 most recent refunds are always available by default on the Charge object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.RefundListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for refunds.
type Iter struct {
	*stripe.Iter
}

// Refund returns the refund which the iterator is currently pointing to.
func (i *Iter) Refund() *stripe.Refund { _ = "STUB: not implemented"; return nil }

// RefundList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) RefundList() *stripe.RefundList { _ = "STUB: not implemented"; return nil }

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
