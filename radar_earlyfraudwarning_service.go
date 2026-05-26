//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1RadarEarlyFraudWarningService is used to invoke /v1/radar/early_fraud_warnings APIs.
type v1RadarEarlyFraudWarningService struct {
	B   Backend
	Key string
}

// Retrieves the details of an early fraud warning that has previously been created.
//
// Please refer to the [early fraud warning](https://docs.stripe.com/api#early_fraud_warning_object) object reference for more details.
func (c v1RadarEarlyFraudWarningService) Retrieve(ctx context.Context, id string, params *RadarEarlyFraudWarningRetrieveParams) (*RadarEarlyFraudWarning, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of early fraud warnings.
func (c v1RadarEarlyFraudWarningService) List(ctx context.Context, listParams *RadarEarlyFraudWarningListParams) *V1List[*RadarEarlyFraudWarning] {
	_ = "STUB: not implemented"
	return nil
}
