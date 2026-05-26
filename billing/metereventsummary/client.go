//
//
// File generated from our OpenAPI spec
//
//

// Package metereventsummary provides the /v1/billing/meters/{id}/event_summaries APIs
package metereventsummary

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/billing/meters/{id}/event_summaries APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieve a list of billing meter event summaries.
func List(params *stripe.BillingMeterEventSummaryListParams) *Iter {
	_ = "STUB: not implemented"
	return nil

	// Retrieve a list of billing meter event summaries.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) List(listParams *stripe.BillingMeterEventSummaryListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for billing meter event summaries.
type Iter struct {
	*stripe.Iter
}

// BillingMeterEventSummary returns the billing meter event summary which the iterator is currently pointing to.
func (i *Iter) BillingMeterEventSummary() *stripe.BillingMeterEventSummary {
	_ = "STUB: not implemented"
	return nil
}

// BillingMeterEventSummaryList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) BillingMeterEventSummaryList() *stripe.BillingMeterEventSummaryList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
