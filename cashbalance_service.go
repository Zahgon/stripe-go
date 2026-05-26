//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1CashBalanceService is used to invoke /v1/customers/{customer}/cash_balance APIs.
type v1CashBalanceService struct {
	B   Backend
	Key string
}

// Retrieves a customer's cash balance.
func (c v1CashBalanceService) Retrieve(ctx context.Context, params *CashBalanceRetrieveParams) (*CashBalance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Changes the settings on a customer's cash balance.
func (c v1CashBalanceService) Update(ctx context.Context, params *CashBalanceUpdateParams) (*CashBalance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
