//
//
// File generated from our OpenAPI spec
//
//

// Package card provides the card related APIs
package card

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke card related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create creates a new card
func New(params *stripe.CardParams) (*stripe.Card, error) {
	_ = "STUB: not implemented"
	return nil,

		// Create creates a new card
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.CardParams) (*stripe.Card, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Note that we call this special append method instead of the standard one
// from the form package. We should not use form's because doing so will
// include some parameters that are undesirable here.

// Because card creation uses the custom append above, we have to
// make an explicit call using a form and CallRaw instead of the standard
// Call (which takes a set of parameters).

// Get returns the details of a card.
func Get(id string, params *stripe.CardParams) (*stripe.Card, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Get returns the details of a card.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.CardParams) (*stripe.Card, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update a specified source for a given customer.
func Update(id string, params *stripe.CardParams) (*stripe.Card, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update a specified source for a given customer.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.CardParams) (*stripe.Card, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete a specified source for a given customer.
func Del(id string, params *stripe.CardParams) (*stripe.Card, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Delete a specified source for a given customer.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Del(id string, params *stripe.CardParams) (*stripe.Card, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func List(params *stripe.CardListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.CardListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// There's no cards list URL, so we use one sources or external
// accounts. An override on CardListParam's `AppendTo` will add the
// filter `object=card` to make sure that only cards come
// back with the response.

// Iter is an iterator for cards.
type Iter struct {
	*stripe.Iter
}

// Card returns the card which the iterator is currently pointing to.
func (i *Iter) Card() *stripe.Card { _ = "STUB: not implemented"; return nil }

// CardList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) CardList() *stripe.CardList { _ = "STUB: not implemented"; return nil }

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
