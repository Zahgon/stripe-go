//
//
// File generated from our OpenAPI spec
//
//

// Package productfeature provides the /v1/products/{product}/features APIs
package productfeature

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/products/{product}/features APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a product_feature, which represents a feature attachment to a product
func New(params *stripe.ProductFeatureParams) (*stripe.ProductFeature, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates a product_feature, which represents a feature attachment to a product
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.ProductFeatureParams) (*stripe.ProductFeature, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a product_feature, which represents a feature attachment to a product
func Get(id string, params *stripe.ProductFeatureParams) (*stripe.ProductFeature, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves a product_feature, which represents a feature attachment to a product
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.ProductFeatureParams) (*stripe.ProductFeature, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deletes the feature attachment to a product
func Del(id string, params *stripe.ProductFeatureParams) (*stripe.ProductFeature, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Deletes the feature attachment to a product
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Del(id string, params *stripe.ProductFeatureParams) (*stripe.ProductFeature, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieve a list of features for a product
func List(params *stripe.ProductFeatureListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Retrieve a list of features for a product
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.ProductFeatureListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for product features.
type Iter struct {
	*stripe.Iter
}

// ProductFeature returns the product feature which the iterator is currently pointing to.
func (i *Iter) ProductFeature() *stripe.ProductFeature { _ = "STUB: not implemented"; return nil }

// ProductFeatureList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) ProductFeatureList() *stripe.ProductFeatureList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
