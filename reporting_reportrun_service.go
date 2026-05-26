//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1ReportingReportRunService is used to invoke /v1/reporting/report_runs APIs.
type v1ReportingReportRunService struct {
	B   Backend
	Key string
}

// Creates a new object and begin running the report. (Certain report types require a [live-mode API key](https://stripe.com/docs/keys#test-live-modes).)
func (c v1ReportingReportRunService) Create(ctx context.Context, params *ReportingReportRunCreateParams) (*ReportingReportRun, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the details of an existing Report Run.
func (c v1ReportingReportRunService) Retrieve(ctx context.Context, id string, params *ReportingReportRunRetrieveParams) (*ReportingReportRun, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of Report Runs, with the most recent appearing first.
func (c v1ReportingReportRunService) List(ctx context.Context, listParams *ReportingReportRunListParams) *V1List[*ReportingReportRun] {
	_ = "STUB: not implemented"
	return nil
}
