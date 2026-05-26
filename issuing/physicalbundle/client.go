//
//
// File generated from our OpenAPI spec
//
//

// Package physicalbundle provides the /v1/issuing/physical_bundles APIs
package physicalbundle

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/issuing/physical_bundles APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves a physical bundle object.
func Get(id string, params *stripe.IssuingPhysicalBundleParams) (*stripe.IssuingPhysicalBundle, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves a physical bundle object.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.IssuingPhysicalBundleParams) (*stripe.IssuingPhysicalBundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of physical bundle objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
func List(params *stripe.IssuingPhysicalBundleListParams) *Iter {
	_ = "STUB: not implemented"
	return nil

	// Returns a list of physical bundle objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) List(listParams *stripe.IssuingPhysicalBundleListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for issuing physical bundles.
type Iter struct {
	*stripe.Iter
}

// IssuingPhysicalBundle returns the issuing physical bundle which the iterator is currently pointing to.
func (i *Iter) IssuingPhysicalBundle() *stripe.IssuingPhysicalBundle {
	_ = "STUB: not implemented"
	return nil
}

// IssuingPhysicalBundleList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) IssuingPhysicalBundleList() *stripe.IssuingPhysicalBundleList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
