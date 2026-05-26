//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v2CoreEventService is used to invoke event related APIs.
type v2CoreEventService struct {
	B   Backend
	Key string
}

// Retrieves the details of an event.
func (c v2CoreEventService) Retrieve(ctx context.Context, id string, params *V2CoreEventRetrieveParams) (V2CoreEvent, error) {
	_ = "STUB: not implemented"
	return *new(V2CoreEvent), nil
}

// List events, going back up to 30 days.
func (c v2CoreEventService) List(ctx context.Context, listParams *V2CoreEventListParams) *V2List[V2CoreEvent] {
	_ = "STUB: not implemented"
	return nil
}
