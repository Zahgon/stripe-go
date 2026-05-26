//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1TreasuryReceivedDebitService is used to invoke /v1/treasury/received_debits APIs.
type v1TreasuryReceivedDebitService struct {
	B   Backend
	Key string
}

// Retrieves the details of an existing ReceivedDebit by passing the unique ReceivedDebit ID from the ReceivedDebit list
func (c v1TreasuryReceivedDebitService) Retrieve(ctx context.Context, id string, params *TreasuryReceivedDebitRetrieveParams) (*TreasuryReceivedDebit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of ReceivedDebits.
func (c v1TreasuryReceivedDebitService) List(ctx context.Context, listParams *TreasuryReceivedDebitListParams) *V1List[*TreasuryReceivedDebit] {
	_ = "STUB: not implemented"
	return nil
}
