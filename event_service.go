//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1EventService is used to invoke /v1/events APIs.
type v1EventService struct {
	B   Backend
	Key string
}

// Retrieves the details of an event if it was created in the last 30 days. Supply the unique identifier of the event, which you might have received in a webhook.
func (c v1EventService) Retrieve(ctx context.Context, id string, params *EventRetrieveParams) (*Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// List events, going back up to 30 days. Each event data is rendered according to Stripe API version at its creation time, specified in [event object](https://docs.stripe.com/api/events/object) api_version attribute (not according to your current Stripe API version or Stripe-Version header).
func (c v1EventService) List(ctx context.Context, listParams *EventListParams) *V1List[*Event] {
	_ = "STUB: not implemented"
	return nil
}
