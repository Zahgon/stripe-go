//
//
// File generated from our OpenAPI spec
//
//

// Package webhookendpoint provides the /v1/webhook_endpoints APIs
package webhookendpoint

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/webhook_endpoints APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// A webhook endpoint must have a url and a list of enabled_events. You may optionally specify the Boolean connect parameter. If set to true, then a Connect webhook endpoint that notifies the specified url about events from all connected accounts is created; otherwise an account webhook endpoint that notifies the specified url only about events from your account is created. You can also create webhook endpoints in the [webhooks settings](https://dashboard.stripe.com/account/webhooks) section of the Dashboard.
func New(params *stripe.WebhookEndpointParams) (*stripe.WebhookEndpoint, error) {
	_ = "STUB: not implemented"
	return nil,

		// A webhook endpoint must have a url and a list of enabled_events. You may optionally specify the Boolean connect parameter. If set to true, then a Connect webhook endpoint that notifies the specified url about events from all connected accounts is created; otherwise an account webhook endpoint that notifies the specified url only about events from your account is created. You can also create webhook endpoints in the [webhooks settings](https://dashboard.stripe.com/account/webhooks) section of the Dashboard.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.WebhookEndpointParams) (*stripe.WebhookEndpoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the webhook endpoint with the given ID.
func Get(id string, params *stripe.WebhookEndpointParams) (*stripe.WebhookEndpoint, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves the webhook endpoint with the given ID.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.WebhookEndpointParams) (*stripe.WebhookEndpoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the webhook endpoint. You may edit the url, the list of enabled_events, and the status of your endpoint.
func Update(id string, params *stripe.WebhookEndpointParams) (*stripe.WebhookEndpoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the webhook endpoint. You may edit the url, the list of enabled_events, and the status of your endpoint.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.WebhookEndpointParams) (*stripe.WebhookEndpoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// You can also delete webhook endpoints via the [webhook endpoint management](https://dashboard.stripe.com/account/webhooks) page of the Stripe dashboard.
func Del(id string, params *stripe.WebhookEndpointParams) (*stripe.WebhookEndpoint, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// You can also delete webhook endpoints via the [webhook endpoint management](https://dashboard.stripe.com/account/webhooks) page of the Stripe dashboard.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Del(id string, params *stripe.WebhookEndpointParams) (*stripe.WebhookEndpoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of your webhook endpoints.
func List(params *stripe.WebhookEndpointListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Returns a list of your webhook endpoints.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.WebhookEndpointListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for webhook endpoints.
type Iter struct {
	*stripe.Iter
}

// WebhookEndpoint returns the webhook endpoint which the iterator is currently pointing to.
func (i *Iter) WebhookEndpoint() *stripe.WebhookEndpoint { _ = "STUB: not implemented"; return nil }

// WebhookEndpointList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) WebhookEndpointList() *stripe.WebhookEndpointList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
