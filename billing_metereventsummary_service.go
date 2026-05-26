//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1BillingMeterEventSummaryService is used to invoke /v1/billing/meters/{id}/event_summaries APIs.
type v1BillingMeterEventSummaryService struct {
	B   Backend
	Key string
}

// Retrieve a list of billing meter event summaries.
func (c v1BillingMeterEventSummaryService) List(ctx context.Context, listParams *BillingMeterEventSummaryListParams) *V1List[*BillingMeterEventSummary] {
	_ = "STUB: not implemented"
	return nil
}
