//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1TaxCalculationService is used to invoke /v1/tax/calculations APIs.
type v1TaxCalculationService struct {
	B   Backend
	Key string
}

// Calculates tax based on the input and returns a Tax Calculation object.
func (c v1TaxCalculationService) Create(ctx context.Context, params *TaxCalculationCreateParams) (*TaxCalculation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a Tax Calculation object, if the calculation hasn't expired.
func (c v1TaxCalculationService) Retrieve(ctx context.Context, id string, params *TaxCalculationRetrieveParams) (*TaxCalculation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the line items of a tax calculation as a collection, if the calculation hasn't expired.
func (c v1TaxCalculationService) ListLineItems(ctx context.Context, listParams *TaxCalculationListLineItemsParams) *V1List[*TaxCalculationLineItem] {
	_ = "STUB: not implemented"
	return nil
}
