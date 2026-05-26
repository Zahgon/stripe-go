//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1PaymentSourceService is used to invoke /v1/customers/{customer}/sources APIs.
type v1PaymentSourceService struct {
	B   Backend
	Key string
}

// When you create a new credit card, you must specify a customer or recipient on which to create it.
//
// If the card's owner has no default card, then the new card will become the default.
// However, if the owner already has a default, then it will not change.
// To change the default, you should [update the customer](https://docs.stripe.com/api/customers/update) to have a new default_source.
func (c v1PaymentSourceService) Create(ctx context.Context, params *PaymentSourceCreateParams) (*PaymentSource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieve a specified source for a given customer.
func (c v1PaymentSourceService) Retrieve(ctx context.Context, id string, params *PaymentSourceRetrieveParams) (*PaymentSource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update a specified source for a given customer.
func (c v1PaymentSourceService) Update(ctx context.Context, id string, params *PaymentSourceUpdateParams) (*PaymentSource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete a specified source for a given customer.
func (c v1PaymentSourceService) Delete(ctx context.Context, id string, params *PaymentSourceDeleteParams) (*PaymentSource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Verify verifies a source which is used for bank accounts.
// Verify a specified bank account for a given customer.
func (c v1PaymentSourceService) Verify(ctx context.Context, id string, params *PaymentSourceVerifyParams) (*PaymentSource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// List sources for a specified customer.
func (c v1PaymentSourceService) List(ctx context.Context, listParams *PaymentSourceListParams) *V1List[*PaymentSource] {
	_ = "STUB: not implemented"
	return nil
}
