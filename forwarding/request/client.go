//
//
// File generated from our OpenAPI spec
//
//

// Package request provides the /v1/forwarding/requests APIs
package request

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/forwarding/requests APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a ForwardingRequest object.
func New(params *stripe.ForwardingRequestParams) (*stripe.ForwardingRequest, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates a ForwardingRequest object.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.ForwardingRequestParams) (*stripe.ForwardingRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a ForwardingRequest object.
func Get(id string, params *stripe.ForwardingRequestParams) (*stripe.ForwardingRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves a ForwardingRequest object.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.ForwardingRequestParams) (*stripe.ForwardingRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Lists all ForwardingRequest objects.
func List(params *stripe.ForwardingRequestListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Lists all ForwardingRequest objects.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.ForwardingRequestListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for forwarding requests.
type Iter struct {
	*stripe.Iter
}

// ForwardingRequest returns the forwarding request which the iterator is currently pointing to.
func (i *Iter) ForwardingRequest() *stripe.ForwardingRequest { _ = "STUB: not implemented"; return nil }

// ForwardingRequestList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) ForwardingRequestList() *stripe.ForwardingRequestList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
