//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1TestHelpersIssuingTransactionService is used to invoke /v1/issuing/transactions APIs.
type v1TestHelpersIssuingTransactionService struct {
	B   Backend
	Key string
}

// Allows the user to capture an arbitrary amount, also known as a forced capture.
func (c v1TestHelpersIssuingTransactionService) CreateForceCapture(ctx context.Context, params *TestHelpersIssuingTransactionCreateForceCaptureParams) (*IssuingTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Allows the user to refund an arbitrary amount, also known as a unlinked refund.
func (c v1TestHelpersIssuingTransactionService) CreateUnlinkedRefund(ctx context.Context, params *TestHelpersIssuingTransactionCreateUnlinkedRefundParams) (*IssuingTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Refund a test-mode Transaction.
func (c v1TestHelpersIssuingTransactionService) Refund(ctx context.Context, id string, params *TestHelpersIssuingTransactionRefundParams) (*IssuingTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
