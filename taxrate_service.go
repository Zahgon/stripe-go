//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1TaxRateService is used to invoke /v1/tax_rates APIs.
type v1TaxRateService struct {
	B   Backend
	Key string
}

// Creates a new tax rate.
func (c v1TaxRateService) Create(ctx context.Context, params *TaxRateCreateParams) (*TaxRate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a tax rate with the given ID
func (c v1TaxRateService) Retrieve(ctx context.Context, id string, params *TaxRateRetrieveParams) (*TaxRate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates an existing tax rate.
func (c v1TaxRateService) Update(ctx context.Context, id string, params *TaxRateUpdateParams) (*TaxRate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of your tax rates. Tax rates are returned sorted by creation date, with the most recently created tax rates appearing first.
func (c v1TaxRateService) List(ctx context.Context, listParams *TaxRateListParams) *V1List[*TaxRate] {
	_ = "STUB: not implemented"
	return nil
}
