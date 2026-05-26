//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1CustomerBalanceTransactionService is used to invoke /v1/customers/{customer}/balance_transactions APIs.
type v1CustomerBalanceTransactionService struct {
	B   Backend
	Key string
}

// Creates an immutable transaction that updates the customer's credit [balance](https://docs.stripe.com/docs/billing/customer/balance).
func (c v1CustomerBalanceTransactionService) Create(ctx context.Context, params *CustomerBalanceTransactionCreateParams) (*CustomerBalanceTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a specific customer balance transaction that updated the customer's [balances](https://docs.stripe.com/docs/billing/customer/balance).
func (c v1CustomerBalanceTransactionService) Retrieve(ctx context.Context, id string, params *CustomerBalanceTransactionRetrieveParams) (*CustomerBalanceTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Most credit balance transaction fields are immutable, but you may update its description and metadata.
func (c v1CustomerBalanceTransactionService) Update(ctx context.Context, id string, params *CustomerBalanceTransactionUpdateParams) (*CustomerBalanceTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of transactions that updated the customer's [balances](https://docs.stripe.com/docs/billing/customer/balance).
func (c v1CustomerBalanceTransactionService) List(ctx context.Context, listParams *CustomerBalanceTransactionListParams) *V1List[*CustomerBalanceTransaction] {
	_ = "STUB: not implemented"
	return nil
}
