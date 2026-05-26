//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1IssuingPhysicalBundleService is used to invoke /v1/issuing/physical_bundles APIs.
type v1IssuingPhysicalBundleService struct {
	B   Backend
	Key string
}

// Retrieves a physical bundle object.
func (c v1IssuingPhysicalBundleService) Retrieve(ctx context.Context, id string, params *IssuingPhysicalBundleRetrieveParams) (*IssuingPhysicalBundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of physical bundle objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
func (c v1IssuingPhysicalBundleService) List(ctx context.Context, listParams *IssuingPhysicalBundleListParams) *V1List[*IssuingPhysicalBundle] {
	_ = "STUB: not implemented"
	return nil
}
