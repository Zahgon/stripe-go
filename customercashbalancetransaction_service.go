//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1CustomerCashBalanceTransactionService is used to invoke /v1/customers/{customer}/cash_balance_transactions APIs.
type v1CustomerCashBalanceTransactionService struct {
	B   Backend
	Key string
}

// Retrieves a specific cash balance transaction, which updated the customer's [cash balance](https://docs.stripe.com/docs/payments/customer-balance).
func (c v1CustomerCashBalanceTransactionService) Retrieve(ctx context.Context, id string, params *CustomerCashBalanceTransactionRetrieveParams) (*CustomerCashBalanceTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of transactions that modified the customer's [cash balance](https://docs.stripe.com/docs/payments/customer-balance).
func (c v1CustomerCashBalanceTransactionService) List(ctx context.Context, listParams *CustomerCashBalanceTransactionListParams) *V1List[*CustomerCashBalanceTransaction] {
	_ = "STUB: not implemented"
	return nil
}
