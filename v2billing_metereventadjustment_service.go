//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v2BillingMeterEventAdjustmentService is used to invoke metereventadjustment related APIs.
type v2BillingMeterEventAdjustmentService struct {
	B   Backend
	Key string
}

// Creates a meter event adjustment to cancel a previously sent meter event.
func (c v2BillingMeterEventAdjustmentService) Create(ctx context.Context, params *V2BillingMeterEventAdjustmentCreateParams) (*V2BillingMeterEventAdjustment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
