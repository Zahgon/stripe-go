//
//
// File generated from our OpenAPI spec
//
//

// Package price provides the /v1/prices APIs
package price

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/prices APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new [Price for an existing <a href="https://docs.stripe.com/api/products">Product](https://docs.stripe.com/api/prices). The Price can be recurring or one-time.
func New(params *stripe.PriceParams) (*stripe.Price, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates a new [Price for an existing <a href="https://docs.stripe.com/api/products">Product](https://docs.stripe.com/api/prices). The Price can be recurring or one-time.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.PriceParams) (*stripe.Price, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the price with the given ID.
func Get(id string, params *stripe.PriceParams) (*stripe.Price, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves the price with the given ID.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.PriceParams) (*stripe.Price, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specified price by setting the values of the parameters passed. Any parameters not provided are left unchanged.
func Update(id string, params *stripe.PriceParams) (*stripe.Price, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specified price by setting the values of the parameters passed. Any parameters not provided are left unchanged.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.PriceParams) (*stripe.Price, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of your active prices, excluding [inline prices](https://docs.stripe.com/docs/products-prices/pricing-models#inline-pricing). For the list of inactive prices, set active to false.
func List(params *stripe.PriceListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Returns a list of your active prices, excluding [inline prices](https://docs.stripe.com/docs/products-prices/pricing-models#inline-pricing). For the list of inactive prices, set active to false.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.PriceListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for prices.
type Iter struct {
	*stripe.Iter
}

// Price returns the price which the iterator is currently pointing to.
func (i *Iter) Price() *stripe.Price { _ = "STUB: not implemented"; return nil }

// PriceList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) PriceList() *stripe.PriceList { _ = "STUB: not implemented"; return nil }

// Search for prices you've previously created using Stripe's [Search Query Language](https://docs.stripe.com/docs/search#search-query-language).
// Don't use search in read-after-write flows where strict consistency is necessary. Under normal operating
// conditions, data is searchable in less than a minute. Occasionally, propagation of new or updated data can be up
// to an hour behind during outages. Search functionality is not available to merchants in India.
func Search(params *stripe.PriceSearchParams) *SearchIter { _ = "STUB: not implemented"; return nil }

// Search for prices you've previously created using Stripe's [Search Query Language](https://docs.stripe.com/docs/search#search-query-language).
// Don't use search in read-after-write flows where strict consistency is necessary. Under normal operating
// conditions, data is searchable in less than a minute. Occasionally, propagation of new or updated data can be up
// to an hour behind during outages. Search functionality is not available to merchants in India.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Search(params *stripe.PriceSearchParams) *SearchIter {
	_ = "STUB: not implemented"
	return nil
}

// SearchIter is an iterator for prices.
type SearchIter struct {
	*stripe.SearchIter
}

// Price returns the price which the iterator is currently pointing to.
func (i *SearchIter) Price() *stripe.Price { _ = "STUB: not implemented"; return nil }

// PriceSearchResult returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *SearchIter) PriceSearchResult() *stripe.PriceSearchResult {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
