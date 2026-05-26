//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1TreasuryTransactionService is used to invoke /v1/treasury/transactions APIs.
type v1TreasuryTransactionService struct {
	B   Backend
	Key string
}

// Retrieves the details of an existing Transaction.
func (c v1TreasuryTransactionService) Retrieve(ctx context.Context, id string, params *TreasuryTransactionRetrieveParams) (*TreasuryTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a list of Transaction objects.
func (c v1TreasuryTransactionService) List(ctx context.Context, listParams *TreasuryTransactionListParams) *V1List[*TreasuryTransaction] {
	_ = "STUB: not implemented"
	return nil
}
