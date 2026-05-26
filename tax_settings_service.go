//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1TaxSettingsService is used to invoke /v1/tax/settings APIs.
type v1TaxSettingsService struct {
	B   Backend
	Key string
}

// Retrieves Tax Settings for a merchant.
func (c v1TaxSettingsService) Retrieve(ctx context.Context, params *TaxSettingsRetrieveParams) (*TaxSettings, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates Tax Settings parameters used in tax calculations. All parameters are editable but none can be removed once set.
func (c v1TaxSettingsService) Update(ctx context.Context, params *TaxSettingsUpdateParams) (*TaxSettings, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
