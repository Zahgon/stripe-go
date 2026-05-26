//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1TreasuryOutboundPaymentService is used to invoke /v1/treasury/outbound_payments APIs.
type v1TreasuryOutboundPaymentService struct {
	B   Backend
	Key string
}

// Creates an OutboundPayment.
func (c v1TreasuryOutboundPaymentService) Create(ctx context.Context, params *TreasuryOutboundPaymentCreateParams) (*TreasuryOutboundPayment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the details of an existing OutboundPayment by passing the unique OutboundPayment ID from either the OutboundPayment creation request or OutboundPayment list.
func (c v1TreasuryOutboundPaymentService) Retrieve(ctx context.Context, id string, params *TreasuryOutboundPaymentRetrieveParams) (*TreasuryOutboundPayment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Cancel an OutboundPayment.
func (c v1TreasuryOutboundPaymentService) Cancel(ctx context.Context, id string, params *TreasuryOutboundPaymentCancelParams) (*TreasuryOutboundPayment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of OutboundPayments sent from the specified FinancialAccount.
func (c v1TreasuryOutboundPaymentService) List(ctx context.Context, listParams *TreasuryOutboundPaymentListParams) *V1List[*TreasuryOutboundPayment] {
	_ = "STUB: not implemented"
	return nil
}
