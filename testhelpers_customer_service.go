//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1TestHelpersCustomerService is used to invoke /v1/customers APIs.
type v1TestHelpersCustomerService struct {
	B   Backend
	Key string
}

// Create an incoming testmode bank transfer
func (c v1TestHelpersCustomerService) FundCashBalance(ctx context.Context, id string, params *TestHelpersCustomerFundCashBalanceParams) (*CustomerCashBalanceTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
