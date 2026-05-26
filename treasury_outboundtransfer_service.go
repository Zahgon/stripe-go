//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1TreasuryOutboundTransferService is used to invoke /v1/treasury/outbound_transfers APIs.
type v1TreasuryOutboundTransferService struct {
	B   Backend
	Key string
}

// Creates an OutboundTransfer.
func (c v1TreasuryOutboundTransferService) Create(ctx context.Context, params *TreasuryOutboundTransferCreateParams) (*TreasuryOutboundTransfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the details of an existing OutboundTransfer by passing the unique OutboundTransfer ID from either the OutboundTransfer creation request or OutboundTransfer list.
func (c v1TreasuryOutboundTransferService) Retrieve(ctx context.Context, id string, params *TreasuryOutboundTransferRetrieveParams) (*TreasuryOutboundTransfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// An OutboundTransfer can be canceled if the funds have not yet been paid out.
func (c v1TreasuryOutboundTransferService) Cancel(ctx context.Context, id string, params *TreasuryOutboundTransferCancelParams) (*TreasuryOutboundTransfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of OutboundTransfers sent from the specified FinancialAccount.
func (c v1TreasuryOutboundTransferService) List(ctx context.Context, listParams *TreasuryOutboundTransferListParams) *V1List[*TreasuryOutboundTransfer] {
	_ = "STUB: not implemented"
	return nil
}
