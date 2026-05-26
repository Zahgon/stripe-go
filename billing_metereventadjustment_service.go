//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1BillingMeterEventAdjustmentService is used to invoke /v1/billing/meter_event_adjustments APIs.
type v1BillingMeterEventAdjustmentService struct {
	B   Backend
	Key string
}

// Creates a billing meter event adjustment.
func (c v1BillingMeterEventAdjustmentService) Create(ctx context.Context, params *BillingMeterEventAdjustmentCreateParams) (*BillingMeterEventAdjustment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
