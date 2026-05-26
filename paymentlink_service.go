//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1PaymentLinkService is used to invoke /v1/payment_links APIs.
type v1PaymentLinkService struct {
	B   Backend
	Key string
}

// Creates a payment link.
func (c v1PaymentLinkService) Create(ctx context.Context, params *PaymentLinkCreateParams) (*PaymentLink, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieve a payment link.
func (c v1PaymentLinkService) Retrieve(ctx context.Context, id string, params *PaymentLinkRetrieveParams) (*PaymentLink, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates a payment link.
func (c v1PaymentLinkService) Update(ctx context.Context, id string, params *PaymentLinkUpdateParams) (*PaymentLink, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of your payment links.
func (c v1PaymentLinkService) List(ctx context.Context, listParams *PaymentLinkListParams) *V1List[*PaymentLink] {
	_ = "STUB: not implemented"
	return nil
}

// When retrieving a payment link, there is an includable line_items property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
func (c v1PaymentLinkService) ListLineItems(ctx context.Context, listParams *PaymentLinkListLineItemsParams) *V1List[*LineItem] {
	_ = "STUB: not implemented"
	return nil
}
