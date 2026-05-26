//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1BillingPortalSessionService is used to invoke /v1/billing_portal/sessions APIs.
type v1BillingPortalSessionService struct {
	B   Backend
	Key string
}

// Creates a session of the customer portal.
func (c v1BillingPortalSessionService) Create(ctx context.Context, params *BillingPortalSessionCreateParams) (*BillingPortalSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
