//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1BalanceSettingsService is used to invoke /v1/balance_settings APIs.
type v1BalanceSettingsService struct {
	B   Backend
	Key string
}

// Retrieves balance settings for a given connected account.
//
//	Related guide: [Making API calls for connected accounts](https://docs.stripe.com/connect/authentication)
func (c v1BalanceSettingsService) Retrieve(ctx context.Context, params *BalanceSettingsRetrieveParams) (*BalanceSettings, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates balance settings for a given connected account.
//
//	Related guide: [Making API calls for connected accounts](https://docs.stripe.com/connect/authentication)
func (c v1BalanceSettingsService) Update(ctx context.Context, params *BalanceSettingsUpdateParams) (*BalanceSettings, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
