//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1CapabilityService is used to invoke /v1/accounts/{account}/capabilities APIs.
type v1CapabilityService struct {
	B   Backend
	Key string
}

// Retrieves information about the specified Account Capability.
func (c v1CapabilityService) Retrieve(ctx context.Context, id string, params *CapabilityRetrieveParams) (*Capability, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates an existing Account Capability. Request or remove a capability by updating its requested parameter.
func (c v1CapabilityService) Update(ctx context.Context, id string, params *CapabilityUpdateParams) (*Capability, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of capabilities associated with the account. The capabilities are returned sorted by creation date, with the most recent capability appearing first.
func (c v1CapabilityService) List(ctx context.Context, listParams *CapabilityListParams) *V1List[*Capability] {
	_ = "STUB: not implemented"
	return nil
}
