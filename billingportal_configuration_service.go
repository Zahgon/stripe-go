//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1BillingPortalConfigurationService is used to invoke /v1/billing_portal/configurations APIs.
type v1BillingPortalConfigurationService struct {
	B   Backend
	Key string
}

// Creates a configuration that describes the functionality and behavior of a PortalSession
func (c v1BillingPortalConfigurationService) Create(ctx context.Context, params *BillingPortalConfigurationCreateParams) (*BillingPortalConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a configuration that describes the functionality of the customer portal.
func (c v1BillingPortalConfigurationService) Retrieve(ctx context.Context, id string, params *BillingPortalConfigurationRetrieveParams) (*BillingPortalConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates a configuration that describes the functionality of the customer portal.
func (c v1BillingPortalConfigurationService) Update(ctx context.Context, id string, params *BillingPortalConfigurationUpdateParams) (*BillingPortalConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of configurations that describe the functionality of the customer portal.
func (c v1BillingPortalConfigurationService) List(ctx context.Context, listParams *BillingPortalConfigurationListParams) *V1List[*BillingPortalConfiguration] {
	_ = "STUB: not implemented"
	return nil
}
