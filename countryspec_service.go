//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1CountrySpecService is used to invoke /v1/country_specs APIs.
type v1CountrySpecService struct {
	B   Backend
	Key string
}

// Returns a Country Spec for a given Country code.
func (c v1CountrySpecService) Retrieve(ctx context.Context, id string, params *CountrySpecRetrieveParams) (*CountrySpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Lists all Country Spec objects available in the API.
func (c v1CountrySpecService) List(ctx context.Context, listParams *CountrySpecListParams) *V1List[*CountrySpec] {
	_ = "STUB: not implemented"
	return nil
}
