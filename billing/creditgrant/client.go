//
//
// File generated from our OpenAPI spec
//
//

// Package creditgrant provides the /v1/billing/credit_grants APIs
package creditgrant

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/billing/credit_grants APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a credit grant.
func New(params *stripe.BillingCreditGrantParams) (*stripe.BillingCreditGrant, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates a credit grant.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.BillingCreditGrantParams) (*stripe.BillingCreditGrant, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a credit grant.
func Get(id string, params *stripe.BillingCreditGrantParams) (*stripe.BillingCreditGrant, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves a credit grant.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.BillingCreditGrantParams) (*stripe.BillingCreditGrant, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates a credit grant.
func Update(id string, params *stripe.BillingCreditGrantParams) (*stripe.BillingCreditGrant, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates a credit grant.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.BillingCreditGrantParams) (*stripe.BillingCreditGrant, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Expires a credit grant.
func Expire(id string, params *stripe.BillingCreditGrantExpireParams) (*stripe.BillingCreditGrant, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Expires a credit grant.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Expire(id string, params *stripe.BillingCreditGrantExpireParams) (*stripe.BillingCreditGrant, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Voids a credit grant.
func VoidGrant(id string, params *stripe.BillingCreditGrantVoidGrantParams) (*stripe.BillingCreditGrant, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Voids a credit grant.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) VoidGrant(id string, params *stripe.BillingCreditGrantVoidGrantParams) (*stripe.BillingCreditGrant, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieve a list of credit grants.
func List(params *stripe.BillingCreditGrantListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Retrieve a list of credit grants.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.BillingCreditGrantListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for billing credit grants.
type Iter struct {
	*stripe.Iter
}

// BillingCreditGrant returns the billing credit grant which the iterator is currently pointing to.
func (i *Iter) BillingCreditGrant() *stripe.BillingCreditGrant {
	_ = "STUB: not implemented"
	return nil
}

// BillingCreditGrantList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) BillingCreditGrantList() *stripe.BillingCreditGrantList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
