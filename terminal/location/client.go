//
//
// File generated from our OpenAPI spec
//
//

// Package location provides the /v1/terminal/locations APIs
package location

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/terminal/locations APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new Location object.
// For further details, including which address fields are required in each country, see the [Manage locations](https://docs.stripe.com/docs/terminal/fleet/locations) guide.
func New(params *stripe.TerminalLocationParams) (*stripe.TerminalLocation, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates a new Location object.
		// For further details, including which address fields are required in each country, see the [Manage locations](https://docs.stripe.com/docs/terminal/fleet/locations) guide.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.TerminalLocationParams) (*stripe.TerminalLocation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a Location object.
func Get(id string, params *stripe.TerminalLocationParams) (*stripe.TerminalLocation, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves a Location object.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.TerminalLocationParams) (*stripe.TerminalLocation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates a Location object by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
func Update(id string, params *stripe.TerminalLocationParams) (*stripe.TerminalLocation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates a Location object by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.TerminalLocationParams) (*stripe.TerminalLocation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deletes a Location object.
func Del(id string, params *stripe.TerminalLocationParams) (*stripe.TerminalLocation, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Deletes a Location object.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Del(id string, params *stripe.TerminalLocationParams) (*stripe.TerminalLocation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of Location objects.
func List(params *stripe.TerminalLocationListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Returns a list of Location objects.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.TerminalLocationListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for terminal locations.
type Iter struct {
	*stripe.Iter
}

// TerminalLocation returns the terminal location which the iterator is currently pointing to.
func (i *Iter) TerminalLocation() *stripe.TerminalLocation { _ = "STUB: not implemented"; return nil }

// TerminalLocationList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TerminalLocationList() *stripe.TerminalLocationList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
