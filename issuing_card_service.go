//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1IssuingCardService is used to invoke /v1/issuing/cards APIs.
type v1IssuingCardService struct {
	B   Backend
	Key string
}

// Creates an Issuing Card object.
func (c v1IssuingCardService) Create(ctx context.Context, params *IssuingCardCreateParams) (*IssuingCard, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves an Issuing Card object.
func (c v1IssuingCardService) Retrieve(ctx context.Context, id string, params *IssuingCardRetrieveParams) (*IssuingCard, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specified Issuing Card object by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
func (c v1IssuingCardService) Update(ctx context.Context, id string, params *IssuingCardUpdateParams) (*IssuingCard, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of Issuing Card objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
func (c v1IssuingCardService) List(ctx context.Context, listParams *IssuingCardListParams) *V1List[*IssuingCard] {
	_ = "STUB: not implemented"
	return nil
}
