//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1TreasuryInboundTransferService is used to invoke /v1/treasury/inbound_transfers APIs.
type v1TreasuryInboundTransferService struct {
	B   Backend
	Key string
}

// Creates an InboundTransfer.
func (c v1TreasuryInboundTransferService) Create(ctx context.Context, params *TreasuryInboundTransferCreateParams) (*TreasuryInboundTransfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the details of an existing InboundTransfer.
func (c v1TreasuryInboundTransferService) Retrieve(ctx context.Context, id string, params *TreasuryInboundTransferRetrieveParams) (*TreasuryInboundTransfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Cancels an InboundTransfer.
func (c v1TreasuryInboundTransferService) Cancel(ctx context.Context, id string, params *TreasuryInboundTransferCancelParams) (*TreasuryInboundTransfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of InboundTransfers sent from the specified FinancialAccount.
func (c v1TreasuryInboundTransferService) List(ctx context.Context, listParams *TreasuryInboundTransferListParams) *V1List[*TreasuryInboundTransfer] {
	_ = "STUB: not implemented"
	return nil
}
