//
//
// File generated from our OpenAPI spec
//
//

// Package configuration provides the /v1/terminal/configurations APIs
package configuration

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/terminal/configurations APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new Configuration object.
func New(params *stripe.TerminalConfigurationParams) (*stripe.TerminalConfiguration, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates a new Configuration object.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.TerminalConfigurationParams) (*stripe.TerminalConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a Configuration object.
func Get(id string, params *stripe.TerminalConfigurationParams) (*stripe.TerminalConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves a Configuration object.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.TerminalConfigurationParams) (*stripe.TerminalConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates a new Configuration object.
func Update(id string, params *stripe.TerminalConfigurationParams) (*stripe.TerminalConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates a new Configuration object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.TerminalConfigurationParams) (*stripe.TerminalConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deletes a Configuration object.
func Del(id string, params *stripe.TerminalConfigurationParams) (*stripe.TerminalConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Deletes a Configuration object.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Del(id string, params *stripe.TerminalConfigurationParams) (*stripe.TerminalConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of Configuration objects.
func List(params *stripe.TerminalConfigurationListParams) *Iter {
	_ = "STUB: not implemented"
	return nil

	// Returns a list of Configuration objects.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) List(listParams *stripe.TerminalConfigurationListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for terminal configurations.
type Iter struct {
	*stripe.Iter
}

// TerminalConfiguration returns the terminal configuration which the iterator is currently pointing to.
func (i *Iter) TerminalConfiguration() *stripe.TerminalConfiguration {
	_ = "STUB: not implemented"
	return nil
}

// TerminalConfigurationList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TerminalConfigurationList() *stripe.TerminalConfigurationList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
