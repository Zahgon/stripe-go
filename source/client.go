//
//
// File generated from our OpenAPI spec
//
//

// Package source provides the /v1/sources APIs
package source

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/sources APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new source object.
func New(params *stripe.SourceParams) (*stripe.Source, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates a new source object.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.SourceParams) (*stripe.Source, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves an existing source object. Supply the unique source ID from a source creation request and Stripe will return the corresponding up-to-date source object information.
func Get(id string, params *stripe.SourceParams) (*stripe.Source, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves an existing source object. Supply the unique source ID from a source creation request and Stripe will return the corresponding up-to-date source object information.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.SourceParams) (*stripe.Source, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specified source by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
//
// This request accepts the metadata and owner as arguments. It is also possible to update type specific information for selected payment methods. Please refer to our [payment method guides](https://docs.stripe.com/docs/sources) for more detail.
func Update(id string, params *stripe.SourceParams) (*stripe.Source, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specified source by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
//
// This request accepts the metadata and owner as arguments. It is also possible to update type specific information for selected payment methods. Please refer to our [payment method guides](https://docs.stripe.com/docs/sources) for more detail.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.SourceParams) (*stripe.Source, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete a specified source for a given customer.
func Detach(id string, params *stripe.SourceDetachParams) (*stripe.Source, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete a specified source for a given customer.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Detach(id string, params *stripe.SourceDetachParams) (*stripe.Source, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
