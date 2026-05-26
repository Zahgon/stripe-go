//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1ShippingRateService is used to invoke /v1/shipping_rates APIs.
type v1ShippingRateService struct {
	B   Backend
	Key string
}

// Creates a new shipping rate object.
func (c v1ShippingRateService) Create(ctx context.Context, params *ShippingRateCreateParams) (*ShippingRate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns the shipping rate object with the given ID.
func (c v1ShippingRateService) Retrieve(ctx context.Context, id string, params *ShippingRateRetrieveParams) (*ShippingRate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates an existing shipping rate object.
func (c v1ShippingRateService) Update(ctx context.Context, id string, params *ShippingRateUpdateParams) (*ShippingRate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of your shipping rates.
func (c v1ShippingRateService) List(ctx context.Context, listParams *ShippingRateListParams) *V1List[*ShippingRate] {
	_ = "STUB: not implemented"
	return nil
}
