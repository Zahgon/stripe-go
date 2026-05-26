//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1TestHelpersTreasuryOutboundTransferService is used to invoke /v1/treasury/outbound_transfers APIs.
type v1TestHelpersTreasuryOutboundTransferService struct {
	B   Backend
	Key string
}

// Updates a test mode created OutboundTransfer with tracking details. The OutboundTransfer must not be cancelable, and cannot be in the canceled or failed states.
func (c v1TestHelpersTreasuryOutboundTransferService) Update(ctx context.Context, id string, params *TestHelpersTreasuryOutboundTransferUpdateParams) (*TreasuryOutboundTransfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Transitions a test mode created OutboundTransfer to the failed status. The OutboundTransfer must already be in the processing state.
func (c v1TestHelpersTreasuryOutboundTransferService) Fail(ctx context.Context, id string, params *TestHelpersTreasuryOutboundTransferFailParams) (*TreasuryOutboundTransfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Transitions a test mode created OutboundTransfer to the posted status. The OutboundTransfer must already be in the processing state.
func (c v1TestHelpersTreasuryOutboundTransferService) Post(ctx context.Context, id string, params *TestHelpersTreasuryOutboundTransferPostParams) (*TreasuryOutboundTransfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Transitions a test mode created OutboundTransfer to the returned status. The OutboundTransfer must already be in the processing state.
func (c v1TestHelpersTreasuryOutboundTransferService) ReturnOutboundTransfer(ctx context.Context, id string, params *TestHelpersTreasuryOutboundTransferReturnOutboundTransferParams) (*TreasuryOutboundTransfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
