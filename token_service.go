//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1TokenService is used to invoke /v1/tokens APIs.
type v1TokenService struct {
	B   Backend
	Key string
}

// Creates a single-use token that represents a bank account's details.
// You can use this token with any v1 API method in place of a bank account dictionary. You can only use this token once. To do so, attach it to a [connected account](https://docs.stripe.com/api#accounts) where [controller.requirement_collection](https://docs.stripe.com/api/accounts/object#account_object-controller-requirement_collection) is application, which includes Custom accounts.
func (c v1TokenService) Create(ctx context.Context, params *TokenCreateParams) (*Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the token with the given ID.
func (c v1TokenService) Retrieve(ctx context.Context, id string, params *TokenRetrieveParams) (*Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
