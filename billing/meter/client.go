//
//
// File generated from our OpenAPI spec
//
//

// Package meter provides the /v1/billing/meters APIs
package meter

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/billing/meters APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a billing meter.
func New(params *stripe.BillingMeterParams) (*stripe.BillingMeter, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates a billing meter.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.BillingMeterParams) (*stripe.BillingMeter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a billing meter given an ID.
func Get(id string, params *stripe.BillingMeterParams) (*stripe.BillingMeter, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves a billing meter given an ID.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.BillingMeterParams) (*stripe.BillingMeter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates a billing meter.
func Update(id string, params *stripe.BillingMeterParams) (*stripe.BillingMeter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates a billing meter.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.BillingMeterParams) (*stripe.BillingMeter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// When a meter is deactivated, no more meter events will be accepted for this meter. You can't attach a deactivated meter to a price.
func Deactivate(id string, params *stripe.BillingMeterDeactivateParams) (*stripe.BillingMeter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// When a meter is deactivated, no more meter events will be accepted for this meter. You can't attach a deactivated meter to a price.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Deactivate(id string, params *stripe.BillingMeterDeactivateParams) (*stripe.BillingMeter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// When a meter is reactivated, events for this meter can be accepted and you can attach the meter to a price.
func Reactivate(id string, params *stripe.BillingMeterReactivateParams) (*stripe.BillingMeter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// When a meter is reactivated, events for this meter can be accepted and you can attach the meter to a price.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Reactivate(id string, params *stripe.BillingMeterReactivateParams) (*stripe.BillingMeter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieve a list of billing meters.
func List(params *stripe.BillingMeterListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Retrieve a list of billing meters.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.BillingMeterListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for billing meters.
type Iter struct {
	*stripe.Iter
}

// BillingMeter returns the billing meter which the iterator is currently pointing to.
func (i *Iter) BillingMeter() *stripe.BillingMeter { _ = "STUB: not implemented"; return nil }

// BillingMeterList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) BillingMeterList() *stripe.BillingMeterList { _ = "STUB: not implemented"; return nil }

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
