//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1BillingCreditBalanceTransactionService is used to invoke /v1/billing/credit_balance_transactions APIs.
type v1BillingCreditBalanceTransactionService struct {
	B   Backend
	Key string
}

// Retrieves a credit balance transaction.
func (c v1BillingCreditBalanceTransactionService) Retrieve(ctx context.Context, id string, params *BillingCreditBalanceTransactionRetrieveParams) (*BillingCreditBalanceTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieve a list of credit balance transactions.
func (c v1BillingCreditBalanceTransactionService) List(ctx context.Context, listParams *BillingCreditBalanceTransactionListParams) *V1List[*BillingCreditBalanceTransaction] {
	_ = "STUB: not implemented"
	return nil
}
