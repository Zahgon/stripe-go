//
//
// File generated from our OpenAPI spec
//
//

// Package capability provides the /v1/accounts/{account}/capabilities APIs
package capability

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/accounts/{account}/capabilities APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves information about the specified Account Capability.
func Get(id string, params *stripe.CapabilityParams) (*stripe.Capability, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves information about the specified Account Capability.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.CapabilityParams) (*stripe.Capability, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates an existing Account Capability. Request or remove a capability by updating its requested parameter.
func Update(id string, params *stripe.CapabilityParams) (*stripe.Capability, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates an existing Account Capability. Request or remove a capability by updating its requested parameter.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.CapabilityParams) (*stripe.Capability, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of capabilities associated with the account. The capabilities are returned sorted by creation date, with the most recent capability appearing first.
func List(params *stripe.CapabilityListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Returns a list of capabilities associated with the account. The capabilities are returned sorted by creation date, with the most recent capability appearing first.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.CapabilityListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for capabilities.
type Iter struct {
	*stripe.Iter
}

// Capability returns the capability which the iterator is currently pointing to.
func (i *Iter) Capability() *stripe.Capability { _ = "STUB: not implemented"; return nil }

// CapabilityList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) CapabilityList() *stripe.CapabilityList { _ = "STUB: not implemented"; return nil }

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
