//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1IdentityVerificationReportService is used to invoke /v1/identity/verification_reports APIs.
type v1IdentityVerificationReportService struct {
	B   Backend
	Key string
}

// Retrieves an existing VerificationReport
func (c v1IdentityVerificationReportService) Retrieve(ctx context.Context, id string, params *IdentityVerificationReportRetrieveParams) (*IdentityVerificationReport, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// List all verification reports.
func (c v1IdentityVerificationReportService) List(ctx context.Context, listParams *IdentityVerificationReportListParams) *V1List[*IdentityVerificationReport] {
	_ = "STUB: not implemented"
	return nil
}
