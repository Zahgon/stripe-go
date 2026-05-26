//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1IssuingTokenService is used to invoke /v1/issuing/tokens APIs.
type v1IssuingTokenService struct {
	B   Backend
	Key string
}

// Retrieves an Issuing Token object.
func (c v1IssuingTokenService) Retrieve(ctx context.Context, id string, params *IssuingTokenRetrieveParams) (*IssuingToken, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Attempts to update the specified Issuing Token object to the status specified.
func (c v1IssuingTokenService) Update(ctx context.Context, id string, params *IssuingTokenUpdateParams) (*IssuingToken, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Lists all Issuing Token objects for a given card.
func (c v1IssuingTokenService) List(ctx context.Context, listParams *IssuingTokenListParams) *V1List[*IssuingToken] {
	_ = "STUB: not implemented"
	return nil
}
