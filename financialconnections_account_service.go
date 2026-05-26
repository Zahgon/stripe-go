//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1FinancialConnectionsAccountService is used to invoke /v1/financial_connections/accounts APIs.
type v1FinancialConnectionsAccountService struct {
	B   Backend
	Key string
}

// Retrieves the details of an Financial Connections Account.
func (c v1FinancialConnectionsAccountService) GetByID(ctx context.Context, id string, params *FinancialConnectionsAccountRetrieveParams) (*FinancialConnectionsAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Disables your access to a Financial Connections Account. You will no longer be able to access data associated with the account (e.g. balances, transactions).
func (c v1FinancialConnectionsAccountService) Disconnect(ctx context.Context, id string, params *FinancialConnectionsAccountDisconnectParams) (*FinancialConnectionsAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Refreshes the data associated with a Financial Connections Account.
func (c v1FinancialConnectionsAccountService) Refresh(ctx context.Context, id string, params *FinancialConnectionsAccountRefreshParams) (*FinancialConnectionsAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Subscribes to periodic refreshes of data associated with a Financial Connections Account. When the account status is active, data is typically refreshed once a day.
func (c v1FinancialConnectionsAccountService) Subscribe(ctx context.Context, id string, params *FinancialConnectionsAccountSubscribeParams) (*FinancialConnectionsAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Unsubscribes from periodic refreshes of data associated with a Financial Connections Account.
func (c v1FinancialConnectionsAccountService) Unsubscribe(ctx context.Context, id string, params *FinancialConnectionsAccountUnsubscribeParams) (*FinancialConnectionsAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of Financial Connections Account objects.
func (c v1FinancialConnectionsAccountService) List(ctx context.Context, listParams *FinancialConnectionsAccountListParams) *V1List[*FinancialConnectionsAccount] {
	_ = "STUB: not implemented"
	return nil
}

// Lists all owners for a given Account
func (c v1FinancialConnectionsAccountService) ListOwners(ctx context.Context, listParams *FinancialConnectionsAccountListOwnersParams) *V1List[*FinancialConnectionsAccountOwner] {
	_ = "STUB: not implemented"
	return nil
}
