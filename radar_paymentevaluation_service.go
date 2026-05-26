//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1RadarPaymentEvaluationService is used to invoke /v1/radar/payment_evaluations APIs.
type v1RadarPaymentEvaluationService struct {
	B   Backend
	Key string
}

// Request a Radar API fraud risk score from Stripe for a payment before sending it for external processor authorization.
func (c v1RadarPaymentEvaluationService) Create(ctx context.Context, params *RadarPaymentEvaluationCreateParams) (*RadarPaymentEvaluation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
