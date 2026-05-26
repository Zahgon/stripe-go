//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1SigmaScheduledQueryRunService is used to invoke /v1/sigma/scheduled_query_runs APIs.
type v1SigmaScheduledQueryRunService struct {
	B   Backend
	Key string
}

// Retrieves the details of an scheduled query run.
func (c v1SigmaScheduledQueryRunService) Retrieve(ctx context.Context, id string, params *SigmaScheduledQueryRunRetrieveParams) (*SigmaScheduledQueryRun, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of scheduled query runs.
func (c v1SigmaScheduledQueryRunService) List(ctx context.Context, listParams *SigmaScheduledQueryRunListParams) *V1List[*SigmaScheduledQueryRun] {
	_ = "STUB: not implemented"
	return nil
}
