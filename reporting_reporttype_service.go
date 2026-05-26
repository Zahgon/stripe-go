//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1ReportingReportTypeService is used to invoke /v1/reporting/report_types APIs.
type v1ReportingReportTypeService struct {
	B   Backend
	Key string
}

// Retrieves the details of a Report Type. (Certain report types require a [live-mode API key](https://stripe.com/docs/keys#test-live-modes).)
func (c v1ReportingReportTypeService) Retrieve(ctx context.Context, id string, params *ReportingReportTypeRetrieveParams) (*ReportingReportType, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a full list of Report Types.
func (c v1ReportingReportTypeService) List(ctx context.Context, listParams *ReportingReportTypeListParams) *V1List[*ReportingReportType] {
	_ = "STUB: not implemented"
	return nil
}
