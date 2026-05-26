//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1BillingMeterEventService is used to invoke /v1/billing/meter_events APIs.
type v1BillingMeterEventService struct {
	B   Backend
	Key string
}

// Creates a billing meter event.
func (c v1BillingMeterEventService) Create(ctx context.Context, params *BillingMeterEventCreateParams) (*BillingMeterEvent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
