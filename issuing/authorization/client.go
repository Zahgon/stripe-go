//
//
// File generated from our OpenAPI spec
//
//

// Package authorization provides the /v1/issuing/authorizations APIs
package authorization

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/issuing/authorizations APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves an Issuing Authorization object.
func Get(id string, params *stripe.IssuingAuthorizationParams) (*stripe.IssuingAuthorization, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves an Issuing Authorization object.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.IssuingAuthorizationParams) (*stripe.IssuingAuthorization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specified Issuing Authorization object by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
func Update(id string, params *stripe.IssuingAuthorizationParams) (*stripe.IssuingAuthorization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specified Issuing Authorization object by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.IssuingAuthorizationParams) (*stripe.IssuingAuthorization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deprecated: [Deprecated] Approves a pending Issuing Authorization object. This request should be made within the timeout window of the [real-time authorization](https://docs.stripe.com/docs/issuing/controls/real-time-authorizations) flow.
// This method is deprecated. Instead, [respond directly to the webhook request to approve an authorization](https://docs.stripe.com/docs/issuing/controls/real-time-authorizations#authorization-handling).
func Approve(id string, params *stripe.IssuingAuthorizationApproveParams) (*stripe.IssuingAuthorization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deprecated: [Deprecated] Approves a pending Issuing Authorization object. This request should be made within the timeout window of the [real-time authorization](https://docs.stripe.com/docs/issuing/controls/real-time-authorizations) flow.
// This method is deprecated. Instead, [respond directly to the webhook request to approve an authorization](https://docs.stripe.com/docs/issuing/controls/real-time-authorizations#authorization-handling).
func (c Client) Approve(id string, params *stripe.IssuingAuthorizationApproveParams) (*stripe.IssuingAuthorization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deprecated: [Deprecated] Declines a pending Issuing Authorization object. This request should be made within the timeout window of the [real time authorization](https://docs.stripe.com/docs/issuing/controls/real-time-authorizations) flow.
// This method is deprecated. Instead, [respond directly to the webhook request to decline an authorization](https://docs.stripe.com/docs/issuing/controls/real-time-authorizations#authorization-handling).
func Decline(id string, params *stripe.IssuingAuthorizationDeclineParams) (*stripe.IssuingAuthorization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deprecated: [Deprecated] Declines a pending Issuing Authorization object. This request should be made within the timeout window of the [real time authorization](https://docs.stripe.com/docs/issuing/controls/real-time-authorizations) flow.
// This method is deprecated. Instead, [respond directly to the webhook request to decline an authorization](https://docs.stripe.com/docs/issuing/controls/real-time-authorizations#authorization-handling).
func (c Client) Decline(id string, params *stripe.IssuingAuthorizationDeclineParams) (*stripe.IssuingAuthorization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of Issuing Authorization objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
func List(params *stripe.IssuingAuthorizationListParams) *Iter {
	_ = "STUB: not implemented"
	return nil

	// Returns a list of Issuing Authorization objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) List(listParams *stripe.IssuingAuthorizationListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for issuing authorizations.
type Iter struct {
	*stripe.Iter
}

// IssuingAuthorization returns the issuing authorization which the iterator is currently pointing to.
func (i *Iter) IssuingAuthorization() *stripe.IssuingAuthorization {
	_ = "STUB: not implemented"
	return nil
}

// IssuingAuthorizationList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) IssuingAuthorizationList() *stripe.IssuingAuthorizationList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
