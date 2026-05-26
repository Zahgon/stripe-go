//
//
// File generated from our OpenAPI spec
//
//

// Package eventdestination provides the eventdestination related APIs
package eventdestination

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke eventdestination related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a new event destination.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2CoreEventDestinationParams) (*stripe.V2CoreEventDestination, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the details of an event destination.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2CoreEventDestinationParams) (*stripe.V2CoreEventDestination, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update the details of an event destination.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.V2CoreEventDestinationParams) (*stripe.V2CoreEventDestination, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete an event destination.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Del(id string, params *stripe.V2CoreEventDestinationParams) (*stripe.V2DeletedObject, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Disable an event destination.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Disable(id string, params *stripe.V2CoreEventDestinationDisableParams) (*stripe.V2CoreEventDestination, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Enable an event destination.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Enable(id string, params *stripe.V2CoreEventDestinationEnableParams) (*stripe.V2CoreEventDestination, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Send a `ping` event to an event destination.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Ping(id string, params *stripe.V2CoreEventDestinationPingParams) (stripe.V2CoreEvent, error) {
	_ = "STUB: not implemented"
	return *new(stripe.V2CoreEvent), nil
}

// Lists all event destinations.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2CoreEventDestinationListParams) stripe.Seq2[*stripe.V2CoreEventDestination, error] {
	_ = "STUB: not implemented"
	return nil
}
