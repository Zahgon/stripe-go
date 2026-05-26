//
//
// File generated from our OpenAPI spec
//
//

// Package paymentevaluation provides the /v1/radar/payment_evaluations APIs
package paymentevaluation

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/radar/payment_evaluations APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Request a Radar API fraud risk score from Stripe for a payment before sending it for external processor authorization.
func New(params *stripe.RadarPaymentEvaluationParams) (*stripe.RadarPaymentEvaluation, error) {
	_ = "STUB: not implemented"
	return nil,

		// Request a Radar API fraud risk score from Stripe for a payment before sending it for external processor authorization.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.RadarPaymentEvaluationParams) (*stripe.RadarPaymentEvaluation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
