//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v2CoreAccountLinkService is used to invoke accountlink related APIs.
type v2CoreAccountLinkService struct {
	B   Backend
	Key string
}

// Creates an AccountLink object that includes a single-use URL that an account can use to access a Stripe-hosted flow for collecting or updating required information.
func (c v2CoreAccountLinkService) Create(ctx context.Context, params *V2CoreAccountLinkCreateParams) (*V2CoreAccountLink, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
