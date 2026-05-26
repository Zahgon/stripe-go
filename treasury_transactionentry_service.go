//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1TreasuryTransactionEntryService is used to invoke /v1/treasury/transaction_entries APIs.
type v1TreasuryTransactionEntryService struct {
	B   Backend
	Key string
}

// Retrieves a TransactionEntry object.
func (c v1TreasuryTransactionEntryService) Retrieve(ctx context.Context, id string, params *TreasuryTransactionEntryRetrieveParams) (*TreasuryTransactionEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a list of TransactionEntry objects.
func (c v1TreasuryTransactionEntryService) List(ctx context.Context, listParams *TreasuryTransactionEntryListParams) *V1List[*TreasuryTransactionEntry] {
	_ = "STUB: not implemented"
	return nil
}
