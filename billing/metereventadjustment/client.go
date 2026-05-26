//
//
// File generated from our OpenAPI spec
//
//

// Package metereventadjustment provides the /v1/billing/meter_event_adjustments APIs
package metereventadjustment

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/billing/meter_event_adjustments APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a billing meter event adjustment.
func New(params *stripe.BillingMeterEventAdjustmentParams) (*stripe.BillingMeterEventAdjustment, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates a billing meter event adjustment.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.BillingMeterEventAdjustmentParams) (*stripe.BillingMeterEventAdjustment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
