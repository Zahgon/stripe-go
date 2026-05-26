//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1EntitlementsActiveEntitlementService is used to invoke /v1/entitlements/active_entitlements APIs.
type v1EntitlementsActiveEntitlementService struct {
	B   Backend
	Key string
}

// Retrieve an active entitlement
func (c v1EntitlementsActiveEntitlementService) Retrieve(ctx context.Context, id string, params *EntitlementsActiveEntitlementRetrieveParams) (*EntitlementsActiveEntitlement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieve a list of active entitlements for a customer
func (c v1EntitlementsActiveEntitlementService) List(ctx context.Context, listParams *EntitlementsActiveEntitlementListParams) *V1List[*EntitlementsActiveEntitlement] {
	_ = "STUB: not implemented"
	return nil
}
