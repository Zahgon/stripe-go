//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1TestHelpersConfirmationTokenService is used to invoke /v1/confirmation_tokens APIs.
type v1TestHelpersConfirmationTokenService struct {
	B   Backend
	Key string
}

// Creates a test mode Confirmation Token server side for your integration tests.
func (c v1TestHelpersConfirmationTokenService) Create(ctx context.Context, params *TestHelpersConfirmationTokenCreateParams) (*ConfirmationToken, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
