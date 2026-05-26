//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1EntitlementsFeatureService is used to invoke /v1/entitlements/features APIs.
type v1EntitlementsFeatureService struct {
	B   Backend
	Key string
}

// Creates a feature
func (c v1EntitlementsFeatureService) Create(ctx context.Context, params *EntitlementsFeatureCreateParams) (*EntitlementsFeature, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a feature
func (c v1EntitlementsFeatureService) Retrieve(ctx context.Context, id string, params *EntitlementsFeatureRetrieveParams) (*EntitlementsFeature, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update a feature's metadata or permanently deactivate it.
func (c v1EntitlementsFeatureService) Update(ctx context.Context, id string, params *EntitlementsFeatureUpdateParams) (*EntitlementsFeature, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieve a list of features
func (c v1EntitlementsFeatureService) List(ctx context.Context, listParams *EntitlementsFeatureListParams) *V1List[*EntitlementsFeature] {
	_ = "STUB: not implemented"
	return nil
}
