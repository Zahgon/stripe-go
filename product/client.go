//
//
// File generated from our OpenAPI spec
//
//

// Package product provides the /v1/products APIs
package product

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/products APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new product object.
func New(params *stripe.ProductParams) (*stripe.Product, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates a new product object.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.ProductParams) (*stripe.Product, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the details of an existing product. Supply the unique product ID from either a product creation request or the product list, and Stripe will return the corresponding product information.
func Get(id string, params *stripe.ProductParams) (*stripe.Product, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves the details of an existing product. Supply the unique product ID from either a product creation request or the product list, and Stripe will return the corresponding product information.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.ProductParams) (*stripe.Product, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specific product by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
func Update(id string, params *stripe.ProductParams) (*stripe.Product, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specific product by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.ProductParams) (*stripe.Product, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete a product. Deleting a product is only possible if it has no prices associated with it. Additionally, deleting a product with type=good is only possible if it has no SKUs associated with it.
func Del(id string, params *stripe.ProductParams) (*stripe.Product, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Delete a product. Deleting a product is only possible if it has no prices associated with it. Additionally, deleting a product with type=good is only possible if it has no SKUs associated with it.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Del(id string, params *stripe.ProductParams) (*stripe.Product, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of your products. The products are returned sorted by creation date, with the most recently created products appearing first.
func List(params *stripe.ProductListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Returns a list of your products. The products are returned sorted by creation date, with the most recently created products appearing first.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.ProductListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for products.
type Iter struct {
	*stripe.Iter
}

// Product returns the product which the iterator is currently pointing to.
func (i *Iter) Product() *stripe.Product { _ = "STUB: not implemented"; return nil }

// ProductList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) ProductList() *stripe.ProductList { _ = "STUB: not implemented"; return nil }

// Search for products you've previously created using Stripe's [Search Query Language](https://docs.stripe.com/docs/search#search-query-language).
// Don't use search in read-after-write flows where strict consistency is necessary. Under normal operating
// conditions, data is searchable in less than a minute. Occasionally, propagation of new or updated data can be up
// to an hour behind during outages. Search functionality is not available to merchants in India.
func Search(params *stripe.ProductSearchParams) *SearchIter { _ = "STUB: not implemented"; return nil }

// Search for products you've previously created using Stripe's [Search Query Language](https://docs.stripe.com/docs/search#search-query-language).
// Don't use search in read-after-write flows where strict consistency is necessary. Under normal operating
// conditions, data is searchable in less than a minute. Occasionally, propagation of new or updated data can be up
// to an hour behind during outages. Search functionality is not available to merchants in India.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Search(params *stripe.ProductSearchParams) *SearchIter {
	_ = "STUB: not implemented"
	return nil
}

// SearchIter is an iterator for products.
type SearchIter struct {
	*stripe.SearchIter
}

// Product returns the product which the iterator is currently pointing to.
func (i *SearchIter) Product() *stripe.Product { _ = "STUB: not implemented"; return nil }

// ProductSearchResult returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *SearchIter) ProductSearchResult() *stripe.ProductSearchResult {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
