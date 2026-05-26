//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1PaymentIntentAmountDetailsLineItemService is used to invoke /v1/payment_intents/{intent}/amount_details_line_items APIs.
type v1PaymentIntentAmountDetailsLineItemService struct {
	B   Backend
	Key string
}

// Lists all LineItems of a given PaymentIntent.
func (c v1PaymentIntentAmountDetailsLineItemService) List(ctx context.Context, listParams *PaymentIntentAmountDetailsLineItemListParams) *V1List[*PaymentIntentAmountDetailsLineItem] {
	_ = "STUB: not implemented"
	return nil
}
