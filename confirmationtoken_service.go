//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1ConfirmationTokenService is used to invoke /v1/confirmation_tokens APIs.
type v1ConfirmationTokenService struct {
	B   Backend
	Key string
}

// Retrieves an existing ConfirmationToken object
func (c v1ConfirmationTokenService) Retrieve(ctx context.Context, id string, params *ConfirmationTokenRetrieveParams) (*ConfirmationToken, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
