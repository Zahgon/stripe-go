//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1PriceService is used to invoke /v1/prices APIs.
type v1PriceService struct {
	B   Backend
	Key string
}

// Creates a new [Price for an existing <a href="https://docs.stripe.com/api/products">Product](https://docs.stripe.com/api/prices). The Price can be recurring or one-time.
func (c v1PriceService) Create(ctx context.Context, params *PriceCreateParams) (*Price, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the price with the given ID.
func (c v1PriceService) Retrieve(ctx context.Context, id string, params *PriceRetrieveParams) (*Price, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specified price by setting the values of the parameters passed. Any parameters not provided are left unchanged.
func (c v1PriceService) Update(ctx context.Context, id string, params *PriceUpdateParams) (*Price, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of your active prices, excluding [inline prices](https://docs.stripe.com/docs/products-prices/pricing-models#inline-pricing). For the list of inactive prices, set active to false.
func (c v1PriceService) List(ctx context.Context, listParams *PriceListParams) *V1List[*Price] {
	_ = "STUB: not implemented"
	return nil
}

// Search for prices you've previously created using Stripe's [Search Query Language](https://docs.stripe.com/docs/search#search-query-language).
// Don't use search in read-after-write flows where strict consistency is necessary. Under normal operating
// conditions, data is searchable in less than a minute. Occasionally, propagation of new or updated data can be up
// to an hour behind during outages. Search functionality is not available to merchants in India.
func (c v1PriceService) Search(ctx context.Context, params *PriceSearchParams) *V1SearchList[*Price] {
	_ = "STUB: not implemented"
	return nil
}
