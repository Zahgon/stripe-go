//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1SetupAttemptService is used to invoke /v1/setup_attempts APIs.
type v1SetupAttemptService struct {
	B   Backend
	Key string
}

// Returns a list of SetupAttempts that associate with a provided SetupIntent.
func (c v1SetupAttemptService) List(ctx context.Context, listParams *SetupAttemptListParams) *V1List[*SetupAttempt] {
	_ = "STUB: not implemented"
	return nil
}
