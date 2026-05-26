//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1ReviewService is used to invoke /v1/reviews APIs.
type v1ReviewService struct {
	B   Backend
	Key string
}

// Retrieves a Review object.
func (c v1ReviewService) Retrieve(ctx context.Context, id string, params *ReviewRetrieveParams) (*Review, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Approves a Review object, closing it and removing it from the list of reviews.
func (c v1ReviewService) Approve(ctx context.Context, id string, params *ReviewApproveParams) (*Review, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of Review objects that have open set to true. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
func (c v1ReviewService) List(ctx context.Context, listParams *ReviewListParams) *V1List[*Review] {
	_ = "STUB: not implemented"
	return nil
}
