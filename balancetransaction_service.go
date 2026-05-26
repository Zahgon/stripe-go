//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1BalanceTransactionService is used to invoke /v1/balance_transactions APIs.
type v1BalanceTransactionService struct {
	B   Backend
	Key string
}

// Retrieves the balance transaction with the given ID.
//
// Note that this endpoint previously used the path /v1/balance/history/:id.
func (c v1BalanceTransactionService) Retrieve(ctx context.Context, id string, params *BalanceTransactionRetrieveParams) (*BalanceTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of transactions that have contributed to the Stripe account balance (e.g., charges, transfers, and so forth). The transactions are returned in sorted order, with the most recent transactions appearing first.
//
// Note that this endpoint was previously called “Balance history” and used the path /v1/balance/history.
func (c v1BalanceTransactionService) List(ctx context.Context, listParams *BalanceTransactionListParams) *V1List[*BalanceTransaction] {
	_ = "STUB: not implemented"
	return nil
}
