//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1BillingCreditGrantService is used to invoke /v1/billing/credit_grants APIs.
type v1BillingCreditGrantService struct {
	B   Backend
	Key string
}

// Creates a credit grant.
func (c v1BillingCreditGrantService) Create(ctx context.Context, params *BillingCreditGrantCreateParams) (*BillingCreditGrant, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a credit grant.
func (c v1BillingCreditGrantService) Retrieve(ctx context.Context, id string, params *BillingCreditGrantRetrieveParams) (*BillingCreditGrant, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates a credit grant.
func (c v1BillingCreditGrantService) Update(ctx context.Context, id string, params *BillingCreditGrantUpdateParams) (*BillingCreditGrant, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Expires a credit grant.
func (c v1BillingCreditGrantService) Expire(ctx context.Context, id string, params *BillingCreditGrantExpireParams) (*BillingCreditGrant, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Voids a credit grant.
func (c v1BillingCreditGrantService) VoidGrant(ctx context.Context, id string, params *BillingCreditGrantVoidGrantParams) (*BillingCreditGrant, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieve a list of credit grants.
func (c v1BillingCreditGrantService) List(ctx context.Context, listParams *BillingCreditGrantListParams) *V1List[*BillingCreditGrant] {
	_ = "STUB: not implemented"
	return nil
}
