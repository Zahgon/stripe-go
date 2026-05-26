//
//
// File generated from our OpenAPI spec
//
//

// Package configuration provides the /v1/billing_portal/configurations APIs
package configuration

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/billing_portal/configurations APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a configuration that describes the functionality and behavior of a PortalSession
func New(params *stripe.BillingPortalConfigurationParams) (*stripe.BillingPortalConfiguration, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates a configuration that describes the functionality and behavior of a PortalSession
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.BillingPortalConfigurationParams) (*stripe.BillingPortalConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a configuration that describes the functionality of the customer portal.
func Get(id string, params *stripe.BillingPortalConfigurationParams) (*stripe.BillingPortalConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves a configuration that describes the functionality of the customer portal.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.BillingPortalConfigurationParams) (*stripe.BillingPortalConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates a configuration that describes the functionality of the customer portal.
func Update(id string, params *stripe.BillingPortalConfigurationParams) (*stripe.BillingPortalConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates a configuration that describes the functionality of the customer portal.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.BillingPortalConfigurationParams) (*stripe.BillingPortalConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of configurations that describe the functionality of the customer portal.
func List(params *stripe.BillingPortalConfigurationListParams) *Iter {
	_ = "STUB: not implemented"
	return nil

	// Returns a list of configurations that describe the functionality of the customer portal.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) List(listParams *stripe.BillingPortalConfigurationListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for billing portal configurations.
type Iter struct {
	*stripe.Iter
}

// BillingPortalConfiguration returns the billing portal configuration which the iterator is currently pointing to.
func (i *Iter) BillingPortalConfiguration() *stripe.BillingPortalConfiguration {
	_ = "STUB: not implemented"
	return nil
}

// BillingPortalConfigurationList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) BillingPortalConfigurationList() *stripe.BillingPortalConfigurationList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
