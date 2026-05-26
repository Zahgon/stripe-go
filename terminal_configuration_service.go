//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1TerminalConfigurationService is used to invoke /v1/terminal/configurations APIs.
type v1TerminalConfigurationService struct {
	B   Backend
	Key string
}

// Creates a new Configuration object.
func (c v1TerminalConfigurationService) Create(ctx context.Context, params *TerminalConfigurationCreateParams) (*TerminalConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a Configuration object.
func (c v1TerminalConfigurationService) Retrieve(ctx context.Context, id string, params *TerminalConfigurationRetrieveParams) (*TerminalConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates a new Configuration object.
func (c v1TerminalConfigurationService) Update(ctx context.Context, id string, params *TerminalConfigurationUpdateParams) (*TerminalConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deletes a Configuration object.
func (c v1TerminalConfigurationService) Delete(ctx context.Context, id string, params *TerminalConfigurationDeleteParams) (*TerminalConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of Configuration objects.
func (c v1TerminalConfigurationService) List(ctx context.Context, listParams *TerminalConfigurationListParams) *V1List[*TerminalConfiguration] {
	_ = "STUB: not implemented"
	return nil
}
