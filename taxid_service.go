//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1TaxIDService is used to invoke /v1/tax_ids APIs.
type v1TaxIDService struct {
	B   Backend
	Key string
}

// Creates a new tax_id object for a customer.
func (c v1TaxIDService) Create(ctx context.Context, params *TaxIDCreateParams) (*TaxID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the tax_id object with the given identifier.
func (c v1TaxIDService) Retrieve(ctx context.Context, id string, params *TaxIDRetrieveParams) (*TaxID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deletes an existing tax_id object.
func (c v1TaxIDService) Delete(ctx context.Context, id string, params *TaxIDDeleteParams) (*TaxID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of tax IDs for a customer.
func (c v1TaxIDService) List(ctx context.Context, listParams *TaxIDListParams) *V1List[*TaxID] {
	_ = "STUB: not implemented"
	return nil
}
