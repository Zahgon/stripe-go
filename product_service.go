//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1ProductService is used to invoke /v1/products APIs.
type v1ProductService struct {
	B   Backend
	Key string
}

// Creates a new product object.
func (c v1ProductService) Create(ctx context.Context, params *ProductCreateParams) (*Product, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the details of an existing product. Supply the unique product ID from either a product creation request or the product list, and Stripe will return the corresponding product information.
func (c v1ProductService) Retrieve(ctx context.Context, id string, params *ProductRetrieveParams) (*Product, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specific product by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
func (c v1ProductService) Update(ctx context.Context, id string, params *ProductUpdateParams) (*Product, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete a product. Deleting a product is only possible if it has no prices associated with it. Additionally, deleting a product with type=good is only possible if it has no SKUs associated with it.
func (c v1ProductService) Delete(ctx context.Context, id string, params *ProductDeleteParams) (*Product, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of your products. The products are returned sorted by creation date, with the most recently created products appearing first.
func (c v1ProductService) List(ctx context.Context, listParams *ProductListParams) *V1List[*Product] {
	_ = "STUB: not implemented"
	return nil
}

// Search for products you've previously created using Stripe's [Search Query Language](https://docs.stripe.com/docs/search#search-query-language).
// Don't use search in read-after-write flows where strict consistency is necessary. Under normal operating
// conditions, data is searchable in less than a minute. Occasionally, propagation of new or updated data can be up
// to an hour behind during outages. Search functionality is not available to merchants in India.
func (c v1ProductService) Search(ctx context.Context, params *ProductSearchParams) *V1SearchList[*Product] {
	_ = "STUB: not implemented"
	return nil
}
