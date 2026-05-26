//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1TestHelpersTestClockService is used to invoke /v1/test_helpers/test_clocks APIs.
type v1TestHelpersTestClockService struct {
	B   Backend
	Key string
}

// Creates a new test clock that can be attached to new customers and quotes.
func (c v1TestHelpersTestClockService) Create(ctx context.Context, params *TestHelpersTestClockCreateParams) (*TestHelpersTestClock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a test clock.
func (c v1TestHelpersTestClockService) Retrieve(ctx context.Context, id string, params *TestHelpersTestClockRetrieveParams) (*TestHelpersTestClock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deletes a test clock.
func (c v1TestHelpersTestClockService) Delete(ctx context.Context, id string, params *TestHelpersTestClockDeleteParams) (*TestHelpersTestClock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Starts advancing a test clock to a specified time in the future. Advancement is done when status changes to Ready.
func (c v1TestHelpersTestClockService) Advance(ctx context.Context, id string, params *TestHelpersTestClockAdvanceParams) (*TestHelpersTestClock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of your test clocks.
func (c v1TestHelpersTestClockService) List(ctx context.Context, listParams *TestHelpersTestClockListParams) *V1List[*TestHelpersTestClock] {
	_ = "STUB: not implemented"
	return nil
}
