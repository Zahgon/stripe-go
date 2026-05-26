//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1BillingMeterService is used to invoke /v1/billing/meters APIs.
type v1BillingMeterService struct {
	B   Backend
	Key string
}

// Creates a billing meter.
func (c v1BillingMeterService) Create(ctx context.Context, params *BillingMeterCreateParams) (*BillingMeter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a billing meter given an ID.
func (c v1BillingMeterService) Retrieve(ctx context.Context, id string, params *BillingMeterRetrieveParams) (*BillingMeter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates a billing meter.
func (c v1BillingMeterService) Update(ctx context.Context, id string, params *BillingMeterUpdateParams) (*BillingMeter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// When a meter is deactivated, no more meter events will be accepted for this meter. You can't attach a deactivated meter to a price.
func (c v1BillingMeterService) Deactivate(ctx context.Context, id string, params *BillingMeterDeactivateParams) (*BillingMeter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// When a meter is reactivated, events for this meter can be accepted and you can attach the meter to a price.
func (c v1BillingMeterService) Reactivate(ctx context.Context, id string, params *BillingMeterReactivateParams) (*BillingMeter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieve a list of billing meters.
func (c v1BillingMeterService) List(ctx context.Context, listParams *BillingMeterListParams) *V1List[*BillingMeter] {
	_ = "STUB: not implemented"
	return nil
}
