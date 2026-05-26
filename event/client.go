//
//
// File generated from our OpenAPI spec
//
//

// Package event provides the /v1/events APIs
package event

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/events APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the details of an event if it was created in the last 30 days. Supply the unique identifier of the event, which you might have received in a webhook.
func Get(id string, params *stripe.EventParams) (*stripe.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves the details of an event if it was created in the last 30 days. Supply the unique identifier of the event, which you might have received in a webhook.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.EventParams) (*stripe.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// List events, going back up to 30 days. Each event data is rendered according to Stripe API version at its creation time, specified in [event object](https://docs.stripe.com/api/events/object) api_version attribute (not according to your current Stripe API version or Stripe-Version header).
func List(params *stripe.EventListParams) *Iter { _ = "STUB: not implemented"; return nil }

// List events, going back up to 30 days. Each event data is rendered according to Stripe API version at its creation time, specified in [event object](https://docs.stripe.com/api/events/object) api_version attribute (not according to your current Stripe API version or Stripe-Version header).
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.EventListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for events.
type Iter struct {
	*stripe.Iter
}

// Event returns the event which the iterator is currently pointing to.
func (i *Iter) Event() *stripe.Event { _ = "STUB: not implemented"; return nil }

// EventList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) EventList() *stripe.EventList { _ = "STUB: not implemented"; return nil }

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
