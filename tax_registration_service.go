//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1TaxRegistrationService is used to invoke /v1/tax/registrations APIs.
type v1TaxRegistrationService struct {
	B   Backend
	Key string
}

// Creates a new Tax Registration object.
func (c v1TaxRegistrationService) Create(ctx context.Context, params *TaxRegistrationCreateParams) (*TaxRegistration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a Tax Registration object.
func (c v1TaxRegistrationService) Retrieve(ctx context.Context, id string, params *TaxRegistrationRetrieveParams) (*TaxRegistration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates an existing Tax Registration object.
//
// A registration cannot be deleted after it has been created. If you wish to end a registration you may do so by setting expires_at.
func (c v1TaxRegistrationService) Update(ctx context.Context, id string, params *TaxRegistrationUpdateParams) (*TaxRegistration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of Tax Registration objects.
func (c v1TaxRegistrationService) List(ctx context.Context, listParams *TaxRegistrationListParams) *V1List[*TaxRegistration] {
	_ = "STUB: not implemented"
	return nil
}
