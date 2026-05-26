//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1TreasuryFinancialAccountService is used to invoke /v1/treasury/financial_accounts APIs.
type v1TreasuryFinancialAccountService struct {
	B   Backend
	Key string
}

// Creates a new FinancialAccount. Each connected account can have up to three FinancialAccounts by default.
func (c v1TreasuryFinancialAccountService) Create(ctx context.Context, params *TreasuryFinancialAccountCreateParams) (*TreasuryFinancialAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the details of a FinancialAccount.
func (c v1TreasuryFinancialAccountService) Retrieve(ctx context.Context, id string, params *TreasuryFinancialAccountRetrieveParams) (*TreasuryFinancialAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the details of a FinancialAccount.
func (c v1TreasuryFinancialAccountService) Update(ctx context.Context, id string, params *TreasuryFinancialAccountUpdateParams) (*TreasuryFinancialAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Closes a FinancialAccount. A FinancialAccount can only be closed if it has a zero balance, has no pending InboundTransfers, and has canceled all attached Issuing cards.
func (c v1TreasuryFinancialAccountService) Close(ctx context.Context, id string, params *TreasuryFinancialAccountCloseParams) (*TreasuryFinancialAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves Features information associated with the FinancialAccount.
func (c v1TreasuryFinancialAccountService) RetrieveFeatures(ctx context.Context, id string, params *TreasuryFinancialAccountRetrieveFeaturesParams) (*TreasuryFinancialAccountFeatures, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the Features associated with a FinancialAccount.
func (c v1TreasuryFinancialAccountService) UpdateFeatures(ctx context.Context, id string, params *TreasuryFinancialAccountUpdateFeaturesParams) (*TreasuryFinancialAccountFeatures, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of FinancialAccounts.
func (c v1TreasuryFinancialAccountService) List(ctx context.Context, listParams *TreasuryFinancialAccountListParams) *V1List[*TreasuryFinancialAccount] {
	_ = "STUB: not implemented"
	return nil
}
