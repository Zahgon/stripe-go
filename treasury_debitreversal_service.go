//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1TreasuryDebitReversalService is used to invoke /v1/treasury/debit_reversals APIs.
type v1TreasuryDebitReversalService struct {
	B   Backend
	Key string
}

// Reverses a ReceivedDebit and creates a DebitReversal object.
func (c v1TreasuryDebitReversalService) Create(ctx context.Context, params *TreasuryDebitReversalCreateParams) (*TreasuryDebitReversal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a DebitReversal object.
func (c v1TreasuryDebitReversalService) Retrieve(ctx context.Context, id string, params *TreasuryDebitReversalRetrieveParams) (*TreasuryDebitReversal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of DebitReversals.
func (c v1TreasuryDebitReversalService) List(ctx context.Context, listParams *TreasuryDebitReversalListParams) *V1List[*TreasuryDebitReversal] {
	_ = "STUB: not implemented"
	return nil
}
