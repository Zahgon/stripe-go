//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v2CoreAccountTokenService is used to invoke accounttoken related APIs.
type v2CoreAccountTokenService struct {
	B   Backend
	Key string
}

// Creates an Account Token.
func (c v2CoreAccountTokenService) Create(ctx context.Context, params *V2CoreAccountTokenCreateParams) (*V2CoreAccountToken, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves an Account Token.
func (c v2CoreAccountTokenService) Retrieve(ctx context.Context, id string, params *V2CoreAccountTokenRetrieveParams) (*V2CoreAccountToken, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
