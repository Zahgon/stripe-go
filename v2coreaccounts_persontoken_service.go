//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v2CoreAccountsPersonTokenService is used to invoke persontoken related APIs.
type v2CoreAccountsPersonTokenService struct {
	B   Backend
	Key string
}

// Creates a Person Token associated with an Account.
func (c v2CoreAccountsPersonTokenService) Create(ctx context.Context, params *V2CoreAccountsPersonTokenCreateParams) (*V2CoreAccountPersonToken, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a Person Token associated with an Account.
func (c v2CoreAccountsPersonTokenService) Retrieve(ctx context.Context, id string, params *V2CoreAccountsPersonTokenRetrieveParams) (*V2CoreAccountPersonToken, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
