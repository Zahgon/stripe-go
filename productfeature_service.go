//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1ProductFeatureService is used to invoke /v1/products/{product}/features APIs.
type v1ProductFeatureService struct {
	B   Backend
	Key string
}

// Creates a product_feature, which represents a feature attachment to a product
func (c v1ProductFeatureService) Create(ctx context.Context, params *ProductFeatureCreateParams) (*ProductFeature, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a product_feature, which represents a feature attachment to a product
func (c v1ProductFeatureService) Retrieve(ctx context.Context, id string, params *ProductFeatureRetrieveParams) (*ProductFeature, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deletes the feature attachment to a product
func (c v1ProductFeatureService) Delete(ctx context.Context, id string, params *ProductFeatureDeleteParams) (*ProductFeature, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieve a list of features for a product
func (c v1ProductFeatureService) List(ctx context.Context, listParams *ProductFeatureListParams) *V1List[*ProductFeature] {
	_ = "STUB: not implemented"
	return nil
}
