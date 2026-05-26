//
//
// File generated from our OpenAPI spec
//
//

// Package alert provides the /v1/billing/alerts APIs
package alert

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/billing/alerts APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a billing alert
func New(params *stripe.BillingAlertParams) (*stripe.BillingAlert, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates a billing alert
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.BillingAlertParams) (*stripe.BillingAlert, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a billing alert given an ID
func Get(id string, params *stripe.BillingAlertParams) (*stripe.BillingAlert, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves a billing alert given an ID
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.BillingAlertParams) (*stripe.BillingAlert, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Reactivates this alert, allowing it to trigger again.
func Activate(id string, params *stripe.BillingAlertActivateParams) (*stripe.BillingAlert, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Reactivates this alert, allowing it to trigger again.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Activate(id string, params *stripe.BillingAlertActivateParams) (*stripe.BillingAlert, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Archives this alert, removing it from the list view and APIs. This is non-reversible.
func Archive(id string, params *stripe.BillingAlertArchiveParams) (*stripe.BillingAlert, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Archives this alert, removing it from the list view and APIs. This is non-reversible.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Archive(id string, params *stripe.BillingAlertArchiveParams) (*stripe.BillingAlert, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deactivates this alert, preventing it from triggering.
func Deactivate(id string, params *stripe.BillingAlertDeactivateParams) (*stripe.BillingAlert, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deactivates this alert, preventing it from triggering.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Deactivate(id string, params *stripe.BillingAlertDeactivateParams) (*stripe.BillingAlert, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Lists billing active and inactive alerts
func List(params *stripe.BillingAlertListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Lists billing active and inactive alerts
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.BillingAlertListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for billing alerts.
type Iter struct {
	*stripe.Iter
}

// BillingAlert returns the billing alert which the iterator is currently pointing to.
func (i *Iter) BillingAlert() *stripe.BillingAlert { _ = "STUB: not implemented"; return nil }

// BillingAlertList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) BillingAlertList() *stripe.BillingAlertList { _ = "STUB: not implemented"; return nil }

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
