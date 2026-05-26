//
//
// File generated from our OpenAPI spec
//
//

// Package person provides the /v1/accounts/{account}/persons APIs
package person

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/accounts/{account}/persons APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new person.
func New(params *stripe.PersonParams) (*stripe.Person, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates a new person.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.PersonParams) (*stripe.Person, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves an existing person.
func Get(id string, params *stripe.PersonParams) (*stripe.Person, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves an existing person.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.PersonParams) (*stripe.Person, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates an existing person.
func Update(id string, params *stripe.PersonParams) (*stripe.Person, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates an existing person.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.PersonParams) (*stripe.Person, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deletes an existing person's relationship to the account's legal entity. Any person with a relationship for an account can be deleted through the API, except if the person is the account_opener. If your integration is using the executive parameter, you cannot delete the only verified executive on file.
func Del(id string, params *stripe.PersonParams) (*stripe.Person, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Deletes an existing person's relationship to the account's legal entity. Any person with a relationship for an account can be deleted through the API, except if the person is the account_opener. If your integration is using the executive parameter, you cannot delete the only verified executive on file.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Del(id string, params *stripe.PersonParams) (*stripe.Person, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of people associated with the account's legal entity. The people are returned sorted by creation date, with the most recent people appearing first.
func List(params *stripe.PersonListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Returns a list of people associated with the account's legal entity. The people are returned sorted by creation date, with the most recent people appearing first.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.PersonListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for persons.
type Iter struct {
	*stripe.Iter
}

// Person returns the person which the iterator is currently pointing to.
func (i *Iter) Person() *stripe.Person { _ = "STUB: not implemented"; return nil }

// PersonList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) PersonList() *stripe.PersonList { _ = "STUB: not implemented"; return nil }

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
