//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1ClimateProductService is used to invoke /v1/climate/products APIs.
type v1ClimateProductService struct {
	B   Backend
	Key string
}

// Retrieves the details of a Climate product with the given ID.
func (c v1ClimateProductService) Retrieve(ctx context.Context, id string, params *ClimateProductRetrieveParams) (*ClimateProduct, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Lists all available Climate product objects.
func (c v1ClimateProductService) List(ctx context.Context, listParams *ClimateProductListParams) *V1List[*ClimateProduct] {
	_ = "STUB: not implemented"
	return nil
}
