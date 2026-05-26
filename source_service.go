//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1SourceService is used to invoke /v1/sources APIs.
type v1SourceService struct {
	B   Backend
	Key string
}

// Creates a new source object.
func (c v1SourceService) Create(ctx context.Context, params *SourceCreateParams) (*Source, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves an existing source object. Supply the unique source ID from a source creation request and Stripe will return the corresponding up-to-date source object information.
func (c v1SourceService) Retrieve(ctx context.Context, id string, params *SourceRetrieveParams) (*Source, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specified source by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
//
// This request accepts the metadata and owner as arguments. It is also possible to update type specific information for selected payment methods. Please refer to our [payment method guides](https://docs.stripe.com/docs/sources) for more detail.
func (c v1SourceService) Update(ctx context.Context, id string, params *SourceUpdateParams) (*Source, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete a specified source for a given customer.
func (c v1SourceService) Detach(ctx context.Context, id string, params *SourceDetachParams) (*Source, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
