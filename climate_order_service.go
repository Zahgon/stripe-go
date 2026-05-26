//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1ClimateOrderService is used to invoke /v1/climate/orders APIs.
type v1ClimateOrderService struct {
	B   Backend
	Key string
}

// Creates a Climate order object for a given Climate product. The order will be processed immediately
// after creation and payment will be deducted your Stripe balance.
func (c v1ClimateOrderService) Create(ctx context.Context, params *ClimateOrderCreateParams) (*ClimateOrder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the details of a Climate order object with the given ID.
func (c v1ClimateOrderService) Retrieve(ctx context.Context, id string, params *ClimateOrderRetrieveParams) (*ClimateOrder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specified order by setting the values of the parameters passed.
func (c v1ClimateOrderService) Update(ctx context.Context, id string, params *ClimateOrderUpdateParams) (*ClimateOrder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Cancels a Climate order. You can cancel an order within 24 hours of creation. Stripe refunds the
// reservation amount_subtotal, but not the amount_fees for user-triggered cancellations. Frontier
// might cancel reservations if suppliers fail to deliver. If Frontier cancels the reservation, Stripe
// provides 90 days advance notice and refunds the amount_total.
func (c v1ClimateOrderService) Cancel(ctx context.Context, id string, params *ClimateOrderCancelParams) (*ClimateOrder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Lists all Climate order objects. The orders are returned sorted by creation date, with the
// most recently created orders appearing first.
func (c v1ClimateOrderService) List(ctx context.Context, listParams *ClimateOrderListParams) *V1List[*ClimateOrder] {
	_ = "STUB: not implemented"
	return nil
}
