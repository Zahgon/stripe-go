//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1BillingAlertService is used to invoke /v1/billing/alerts APIs.
type v1BillingAlertService struct {
	B   Backend
	Key string
}

// Creates a billing alert
func (c v1BillingAlertService) Create(ctx context.Context, params *BillingAlertCreateParams) (*BillingAlert, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a billing alert given an ID
func (c v1BillingAlertService) Retrieve(ctx context.Context, id string, params *BillingAlertRetrieveParams) (*BillingAlert, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Reactivates this alert, allowing it to trigger again.
func (c v1BillingAlertService) Activate(ctx context.Context, id string, params *BillingAlertActivateParams) (*BillingAlert, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Archives this alert, removing it from the list view and APIs. This is non-reversible.
func (c v1BillingAlertService) Archive(ctx context.Context, id string, params *BillingAlertArchiveParams) (*BillingAlert, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deactivates this alert, preventing it from triggering.
func (c v1BillingAlertService) Deactivate(ctx context.Context, id string, params *BillingAlertDeactivateParams) (*BillingAlert, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Lists billing active and inactive alerts
func (c v1BillingAlertService) List(ctx context.Context, listParams *BillingAlertListParams) *V1List[*BillingAlert] {
	_ = "STUB: not implemented"
	return nil
}
