//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1FinancialConnectionsTransactionService is used to invoke /v1/financial_connections/transactions APIs.
type v1FinancialConnectionsTransactionService struct {
	B   Backend
	Key string
}

// Retrieves the details of a Financial Connections Transaction
func (c v1FinancialConnectionsTransactionService) Retrieve(ctx context.Context, id string, params *FinancialConnectionsTransactionRetrieveParams) (*FinancialConnectionsTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of Financial Connections Transaction objects.
func (c v1FinancialConnectionsTransactionService) List(ctx context.Context, listParams *FinancialConnectionsTransactionListParams) *V1List[*FinancialConnectionsTransaction] {
	_ = "STUB: not implemented"
	return nil
}
