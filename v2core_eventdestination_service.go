//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v2CoreEventDestinationService is used to invoke eventdestination related APIs.
type v2CoreEventDestinationService struct {
	B   Backend
	Key string
}

// Create a new event destination.
func (c v2CoreEventDestinationService) Create(ctx context.Context, params *V2CoreEventDestinationCreateParams) (*V2CoreEventDestination, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the details of an event destination.
func (c v2CoreEventDestinationService) Retrieve(ctx context.Context, id string, params *V2CoreEventDestinationRetrieveParams) (*V2CoreEventDestination, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update the details of an event destination.
func (c v2CoreEventDestinationService) Update(ctx context.Context, id string, params *V2CoreEventDestinationUpdateParams) (*V2CoreEventDestination, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete an event destination.
func (c v2CoreEventDestinationService) Delete(ctx context.Context, id string, params *V2CoreEventDestinationDeleteParams) (*V2DeletedObject, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Disable an event destination.
func (c v2CoreEventDestinationService) Disable(ctx context.Context, id string, params *V2CoreEventDestinationDisableParams) (*V2CoreEventDestination, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Enable an event destination.
func (c v2CoreEventDestinationService) Enable(ctx context.Context, id string, params *V2CoreEventDestinationEnableParams) (*V2CoreEventDestination, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Send a `ping` event to an event destination.
func (c v2CoreEventDestinationService) Ping(ctx context.Context, id string, params *V2CoreEventDestinationPingParams) (V2CoreEvent, error) {
	_ = "STUB: not implemented"
	return *new(V2CoreEvent), nil
}

// Lists all event destinations.
func (c v2CoreEventDestinationService) List(ctx context.Context, listParams *V2CoreEventDestinationListParams) *V2List[*V2CoreEventDestination] {
	_ = "STUB: not implemented"
	return nil
}
