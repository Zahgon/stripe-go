//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1AccountLinkService is used to invoke /v1/account_links APIs.
type v1AccountLinkService struct {
	B   Backend
	Key string
}

// Creates an AccountLink object that includes a single-use Stripe URL that the platform can redirect their user to in order to take them through the Connect Onboarding flow.
func (c v1AccountLinkService) Create(ctx context.Context, params *AccountLinkCreateParams) (*AccountLink, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
