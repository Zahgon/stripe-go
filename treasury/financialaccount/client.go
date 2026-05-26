//
//
// File generated from our OpenAPI spec
//
//

// Package financialaccount provides the /v1/treasury/financial_accounts APIs
package financialaccount

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/treasury/financial_accounts APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new FinancialAccount. Each connected account can have up to three FinancialAccounts by default.
func New(params *stripe.TreasuryFinancialAccountParams) (*stripe.TreasuryFinancialAccount, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates a new FinancialAccount. Each connected account can have up to three FinancialAccounts by default.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.TreasuryFinancialAccountParams) (*stripe.TreasuryFinancialAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the details of a FinancialAccount.
func Get(id string, params *stripe.TreasuryFinancialAccountParams) (*stripe.TreasuryFinancialAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves the details of a FinancialAccount.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.TreasuryFinancialAccountParams) (*stripe.TreasuryFinancialAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the details of a FinancialAccount.
func Update(id string, params *stripe.TreasuryFinancialAccountParams) (*stripe.TreasuryFinancialAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the details of a FinancialAccount.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.TreasuryFinancialAccountParams) (*stripe.TreasuryFinancialAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Closes a FinancialAccount. A FinancialAccount can only be closed if it has a zero balance, has no pending InboundTransfers, and has canceled all attached Issuing cards.
func Close(id string, params *stripe.TreasuryFinancialAccountCloseParams) (*stripe.TreasuryFinancialAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Closes a FinancialAccount. A FinancialAccount can only be closed if it has a zero balance, has no pending InboundTransfers, and has canceled all attached Issuing cards.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Close(id string, params *stripe.TreasuryFinancialAccountCloseParams) (*stripe.TreasuryFinancialAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves Features information associated with the FinancialAccount.
func RetrieveFeatures(id string, params *stripe.TreasuryFinancialAccountRetrieveFeaturesParams) (*stripe.TreasuryFinancialAccountFeatures, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves Features information associated with the FinancialAccount.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) RetrieveFeatures(id string, params *stripe.TreasuryFinancialAccountRetrieveFeaturesParams) (*stripe.TreasuryFinancialAccountFeatures, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the Features associated with a FinancialAccount.
func UpdateFeatures(id string, params *stripe.TreasuryFinancialAccountUpdateFeaturesParams) (*stripe.TreasuryFinancialAccountFeatures, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the Features associated with a FinancialAccount.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) UpdateFeatures(id string, params *stripe.TreasuryFinancialAccountUpdateFeaturesParams) (*stripe.TreasuryFinancialAccountFeatures, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of FinancialAccounts.
func List(params *stripe.TreasuryFinancialAccountListParams) *Iter {
	_ = "STUB: not implemented"
	return nil

	// Returns a list of FinancialAccounts.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) List(listParams *stripe.TreasuryFinancialAccountListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for treasury financial accounts.
type Iter struct {
	*stripe.Iter
}

// TreasuryFinancialAccount returns the treasury financial account which the iterator is currently pointing to.
func (i *Iter) TreasuryFinancialAccount() *stripe.TreasuryFinancialAccount {
	_ = "STUB: not implemented"
	return nil
}

// TreasuryFinancialAccountList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TreasuryFinancialAccountList() *stripe.TreasuryFinancialAccountList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
