//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1ApplicationFeeService is used to invoke /v1/application_fees APIs.
type v1ApplicationFeeService struct {
	B   Backend
	Key string
}

// Retrieves the details of an application fee that your account has collected. The same information is returned when refunding the application fee.
func (c v1ApplicationFeeService) Retrieve(ctx context.Context, id string, params *ApplicationFeeRetrieveParams) (*ApplicationFee, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of application fees you've previously collected. The application fees are returned in sorted order, with the most recent fees appearing first.
func (c v1ApplicationFeeService) List(ctx context.Context, listParams *ApplicationFeeListParams) *V1List[*ApplicationFee] {
	_ = "STUB: not implemented"
	return nil
}
