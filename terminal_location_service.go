//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1TerminalLocationService is used to invoke /v1/terminal/locations APIs.
type v1TerminalLocationService struct {
	B   Backend
	Key string
}

// Creates a new Location object.
// For further details, including which address fields are required in each country, see the [Manage locations](https://docs.stripe.com/docs/terminal/fleet/locations) guide.
func (c v1TerminalLocationService) Create(ctx context.Context, params *TerminalLocationCreateParams) (*TerminalLocation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a Location object.
func (c v1TerminalLocationService) Retrieve(ctx context.Context, id string, params *TerminalLocationRetrieveParams) (*TerminalLocation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates a Location object by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
func (c v1TerminalLocationService) Update(ctx context.Context, id string, params *TerminalLocationUpdateParams) (*TerminalLocation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deletes a Location object.
func (c v1TerminalLocationService) Delete(ctx context.Context, id string, params *TerminalLocationDeleteParams) (*TerminalLocation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of Location objects.
func (c v1TerminalLocationService) List(ctx context.Context, listParams *TerminalLocationListParams) *V1List[*TerminalLocation] {
	_ = "STUB: not implemented"
	return nil
}
