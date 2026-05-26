//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1TestHelpersTreasuryOutboundPaymentService is used to invoke /v1/treasury/outbound_payments APIs.
type v1TestHelpersTreasuryOutboundPaymentService struct {
	B   Backend
	Key string
}

// Updates a test mode created OutboundPayment with tracking details. The OutboundPayment must not be cancelable, and cannot be in the canceled or failed states.
func (c v1TestHelpersTreasuryOutboundPaymentService) Update(ctx context.Context, id string, params *TestHelpersTreasuryOutboundPaymentUpdateParams) (*TreasuryOutboundPayment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Transitions a test mode created OutboundPayment to the failed status. The OutboundPayment must already be in the processing state.
func (c v1TestHelpersTreasuryOutboundPaymentService) Fail(ctx context.Context, id string, params *TestHelpersTreasuryOutboundPaymentFailParams) (*TreasuryOutboundPayment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Transitions a test mode created OutboundPayment to the posted status. The OutboundPayment must already be in the processing state.
func (c v1TestHelpersTreasuryOutboundPaymentService) Post(ctx context.Context, id string, params *TestHelpersTreasuryOutboundPaymentPostParams) (*TreasuryOutboundPayment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Transitions a test mode created OutboundPayment to the returned status. The OutboundPayment must already be in the processing state.
func (c v1TestHelpersTreasuryOutboundPaymentService) ReturnOutboundPayment(ctx context.Context, id string, params *TestHelpersTreasuryOutboundPaymentReturnOutboundPaymentParams) (*TreasuryOutboundPayment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
