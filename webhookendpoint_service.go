//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1WebhookEndpointService is used to invoke /v1/webhook_endpoints APIs.
type v1WebhookEndpointService struct {
	B   Backend
	Key string
}

// A webhook endpoint must have a url and a list of enabled_events. You may optionally specify the Boolean connect parameter. If set to true, then a Connect webhook endpoint that notifies the specified url about events from all connected accounts is created; otherwise an account webhook endpoint that notifies the specified url only about events from your account is created. You can also create webhook endpoints in the [webhooks settings](https://dashboard.stripe.com/account/webhooks) section of the Dashboard.
func (c v1WebhookEndpointService) Create(ctx context.Context, params *WebhookEndpointCreateParams) (*WebhookEndpoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the webhook endpoint with the given ID.
func (c v1WebhookEndpointService) Retrieve(ctx context.Context, id string, params *WebhookEndpointRetrieveParams) (*WebhookEndpoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the webhook endpoint. You may edit the url, the list of enabled_events, and the status of your endpoint.
func (c v1WebhookEndpointService) Update(ctx context.Context, id string, params *WebhookEndpointUpdateParams) (*WebhookEndpoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// You can also delete webhook endpoints via the [webhook endpoint management](https://dashboard.stripe.com/account/webhooks) page of the Stripe dashboard.
func (c v1WebhookEndpointService) Delete(ctx context.Context, id string, params *WebhookEndpointDeleteParams) (*WebhookEndpoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of your webhook endpoints.
func (c v1WebhookEndpointService) List(ctx context.Context, listParams *WebhookEndpointListParams) *V1List[*WebhookEndpoint] {
	_ = "STUB: not implemented"
	return nil
}
