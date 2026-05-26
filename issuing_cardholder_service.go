//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1IssuingCardholderService is used to invoke /v1/issuing/cardholders APIs.
type v1IssuingCardholderService struct {
	B   Backend
	Key string
}

// Creates a new Issuing Cardholder object that can be issued cards.
func (c v1IssuingCardholderService) Create(ctx context.Context, params *IssuingCardholderCreateParams) (*IssuingCardholder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves an Issuing Cardholder object.
func (c v1IssuingCardholderService) Retrieve(ctx context.Context, id string, params *IssuingCardholderRetrieveParams) (*IssuingCardholder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specified Issuing Cardholder object by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
func (c v1IssuingCardholderService) Update(ctx context.Context, id string, params *IssuingCardholderUpdateParams) (*IssuingCardholder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of Issuing Cardholder objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
func (c v1IssuingCardholderService) List(ctx context.Context, listParams *IssuingCardholderListParams) *V1List[*IssuingCardholder] {
	_ = "STUB: not implemented"
	return nil
}
