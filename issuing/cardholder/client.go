//
//
// File generated from our OpenAPI spec
//
//

// Package cardholder provides the /v1/issuing/cardholders APIs
// For more details, see: https://stripe.com/docs/api/?lang=go#issuing_cardholders
package cardholder

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/issuing/cardholders APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new Issuing Cardholder object that can be issued cards.
func New(params *stripe.IssuingCardholderParams) (*stripe.IssuingCardholder, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates a new Issuing Cardholder object that can be issued cards.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.IssuingCardholderParams) (*stripe.IssuingCardholder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves an Issuing Cardholder object.
func Get(id string, params *stripe.IssuingCardholderParams) (*stripe.IssuingCardholder, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves an Issuing Cardholder object.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.IssuingCardholderParams) (*stripe.IssuingCardholder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specified Issuing Cardholder object by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
func Update(id string, params *stripe.IssuingCardholderParams) (*stripe.IssuingCardholder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specified Issuing Cardholder object by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.IssuingCardholderParams) (*stripe.IssuingCardholder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of Issuing Cardholder objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
func List(params *stripe.IssuingCardholderListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Returns a list of Issuing Cardholder objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.IssuingCardholderListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for issuing cardholders.
type Iter struct {
	*stripe.Iter
}

// IssuingCardholder returns the issuing cardholder which the iterator is currently pointing to.
func (i *Iter) IssuingCardholder() *stripe.IssuingCardholder { _ = "STUB: not implemented"; return nil }

// IssuingCardholderList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) IssuingCardholderList() *stripe.IssuingCardholderList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
