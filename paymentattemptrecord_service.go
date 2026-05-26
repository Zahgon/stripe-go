//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1PaymentAttemptRecordService is used to invoke /v1/payment_attempt_records APIs.
type v1PaymentAttemptRecordService struct {
	B   Backend
	Key string
}

// Retrieves a Payment Attempt Record with the given ID
func (c v1PaymentAttemptRecordService) Retrieve(ctx context.Context, id string, params *PaymentAttemptRecordRetrieveParams) (*PaymentAttemptRecord, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// List all the Payment Attempt Records attached to the specified Payment Record.
func (c v1PaymentAttemptRecordService) List(ctx context.Context, listParams *PaymentAttemptRecordListParams) *V1List[*PaymentAttemptRecord] {
	_ = "STUB: not implemented"
	return nil
}
