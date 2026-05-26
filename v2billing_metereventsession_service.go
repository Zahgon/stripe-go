//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v2BillingMeterEventSessionService is used to invoke metereventsession related APIs.
type v2BillingMeterEventSessionService struct {
	B   Backend
	Key string
}

// Creates a meter event session to send usage on the high-throughput meter event stream. Authentication tokens are only valid for 15 minutes, so you will need to create a new meter event session when your token expires.
func (c v2BillingMeterEventSessionService) Create(ctx context.Context, params *V2BillingMeterEventSessionCreateParams) (*V2BillingMeterEventSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
