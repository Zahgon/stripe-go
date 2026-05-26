//
//
// File generated from our OpenAPI spec
//
//

// Package activeentitlement provides the /v1/entitlements/active_entitlements APIs
package activeentitlement

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/entitlements/active_entitlements APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieve an active entitlement
func Get(id string, params *stripe.EntitlementsActiveEntitlementParams) (*stripe.EntitlementsActiveEntitlement, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieve an active entitlement
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.EntitlementsActiveEntitlementParams) (*stripe.EntitlementsActiveEntitlement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieve a list of active entitlements for a customer
func List(params *stripe.EntitlementsActiveEntitlementListParams) *Iter {
	_ = "STUB: not implemented"
	return nil

	// Retrieve a list of active entitlements for a customer
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) List(listParams *stripe.EntitlementsActiveEntitlementListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for entitlements active entitlements.
type Iter struct {
	*stripe.Iter
}

// EntitlementsActiveEntitlement returns the entitlements active entitlement which the iterator is currently pointing to.
func (i *Iter) EntitlementsActiveEntitlement() *stripe.EntitlementsActiveEntitlement {
	_ = "STUB: not implemented"
	return nil
}

// EntitlementsActiveEntitlementList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) EntitlementsActiveEntitlementList() *stripe.EntitlementsActiveEntitlementList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
