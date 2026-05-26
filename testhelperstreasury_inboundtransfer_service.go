//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1TestHelpersTreasuryInboundTransferService is used to invoke /v1/treasury/inbound_transfers APIs.
type v1TestHelpersTreasuryInboundTransferService struct {
	B   Backend
	Key string
}

// Transitions a test mode created InboundTransfer to the failed status. The InboundTransfer must already be in the processing state.
func (c v1TestHelpersTreasuryInboundTransferService) Fail(ctx context.Context, id string, params *TestHelpersTreasuryInboundTransferFailParams) (*TreasuryInboundTransfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Marks the test mode InboundTransfer object as returned and links the InboundTransfer to a ReceivedDebit. The InboundTransfer must already be in the succeeded state.
func (c v1TestHelpersTreasuryInboundTransferService) ReturnInboundTransfer(ctx context.Context, id string, params *TestHelpersTreasuryInboundTransferReturnInboundTransferParams) (*TreasuryInboundTransfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Transitions a test mode created InboundTransfer to the succeeded status. The InboundTransfer must already be in the processing state.
func (c v1TestHelpersTreasuryInboundTransferService) Succeed(ctx context.Context, id string, params *TestHelpersTreasuryInboundTransferSucceedParams) (*TreasuryInboundTransfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
