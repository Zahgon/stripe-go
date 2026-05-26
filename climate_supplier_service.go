//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1ClimateSupplierService is used to invoke /v1/climate/suppliers APIs.
type v1ClimateSupplierService struct {
	B   Backend
	Key string
}

// Retrieves a Climate supplier object.
func (c v1ClimateSupplierService) Retrieve(ctx context.Context, id string, params *ClimateSupplierRetrieveParams) (*ClimateSupplier, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Lists all available Climate supplier objects.
func (c v1ClimateSupplierService) List(ctx context.Context, listParams *ClimateSupplierListParams) *V1List[*ClimateSupplier] {
	_ = "STUB: not implemented"
	return nil
}
