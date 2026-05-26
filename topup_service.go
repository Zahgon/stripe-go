//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1TopupService is used to invoke /v1/topups APIs.
type v1TopupService struct {
	B   Backend
	Key string
}

// Top up the balance of an account
func (c v1TopupService) Create(ctx context.Context, params *TopupCreateParams) (*Topup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the details of a top-up that has previously been created. Supply the unique top-up ID that was returned from your previous request, and Stripe will return the corresponding top-up information.
func (c v1TopupService) Retrieve(ctx context.Context, id string, params *TopupRetrieveParams) (*Topup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the metadata of a top-up. Other top-up details are not editable by design.
func (c v1TopupService) Update(ctx context.Context, id string, params *TopupUpdateParams) (*Topup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Cancels a top-up. Only pending top-ups can be canceled.
func (c v1TopupService) Cancel(ctx context.Context, id string, params *TopupCancelParams) (*Topup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of top-ups.
func (c v1TopupService) List(ctx context.Context, listParams *TopupListParams) *V1List[*Topup] {
	_ = "STUB: not implemented"
	return nil
}
