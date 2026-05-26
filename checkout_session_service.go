//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1CheckoutSessionService is used to invoke /v1/checkout/sessions APIs.
type v1CheckoutSessionService struct {
	B   Backend
	Key string
}

// Creates a Checkout Session object.
func (c v1CheckoutSessionService) Create(ctx context.Context, params *CheckoutSessionCreateParams) (*CheckoutSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a Checkout Session object.
func (c v1CheckoutSessionService) Retrieve(ctx context.Context, id string, params *CheckoutSessionRetrieveParams) (*CheckoutSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates a Checkout Session object.
//
// Related guide: [Dynamically update a Checkout Session](https://docs.stripe.com/payments/advanced/dynamic-updates)
func (c v1CheckoutSessionService) Update(ctx context.Context, id string, params *CheckoutSessionUpdateParams) (*CheckoutSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// A Checkout Session can be expired when it is in one of these statuses: open
//
// After it expires, a customer can't complete a Checkout Session and customers loading the Checkout Session see a message saying the Checkout Session is expired.
func (c v1CheckoutSessionService) Expire(ctx context.Context, id string, params *CheckoutSessionExpireParams) (*CheckoutSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of Checkout Sessions.
func (c v1CheckoutSessionService) List(ctx context.Context, listParams *CheckoutSessionListParams) *V1List[*CheckoutSession] {
	_ = "STUB: not implemented"
	return nil
}

// When retrieving a Checkout Session, there is an includable line_items property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
func (c v1CheckoutSessionService) ListLineItems(ctx context.Context, listParams *CheckoutSessionListLineItemsParams) *V1List[*LineItem] {
	_ = "STUB: not implemented"
	return nil
}
