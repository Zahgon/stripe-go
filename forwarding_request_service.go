//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1ForwardingRequestService is used to invoke /v1/forwarding/requests APIs.
type v1ForwardingRequestService struct {
	B   Backend
	Key string
}

// Creates a ForwardingRequest object.
func (c v1ForwardingRequestService) Create(ctx context.Context, params *ForwardingRequestCreateParams) (*ForwardingRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a ForwardingRequest object.
func (c v1ForwardingRequestService) Retrieve(ctx context.Context, id string, params *ForwardingRequestRetrieveParams) (*ForwardingRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Lists all ForwardingRequest objects.
func (c v1ForwardingRequestService) List(ctx context.Context, listParams *ForwardingRequestListParams) *V1List[*ForwardingRequest] {
	_ = "STUB: not implemented"
	return nil
}
