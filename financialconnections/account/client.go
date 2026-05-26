//
//
// File generated from our OpenAPI spec
//
//

// Package account provides the /v1/financial_connections/accounts APIs
package account

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/financial_connections/accounts APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the details of an Financial Connections Account.
func GetByID(id string, params *stripe.FinancialConnectionsAccountParams) (*stripe.FinancialConnectionsAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the details of an Financial Connections Account.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) GetByID(id string, params *stripe.FinancialConnectionsAccountParams) (*stripe.FinancialConnectionsAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Disables your access to a Financial Connections Account. You will no longer be able to access data associated with the account (e.g. balances, transactions).
func Disconnect(id string, params *stripe.FinancialConnectionsAccountDisconnectParams) (*stripe.FinancialConnectionsAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Disables your access to a Financial Connections Account. You will no longer be able to access data associated with the account (e.g. balances, transactions).
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Disconnect(id string, params *stripe.FinancialConnectionsAccountDisconnectParams) (*stripe.FinancialConnectionsAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Refreshes the data associated with a Financial Connections Account.
func Refresh(id string, params *stripe.FinancialConnectionsAccountRefreshParams) (*stripe.FinancialConnectionsAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Refreshes the data associated with a Financial Connections Account.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Refresh(id string, params *stripe.FinancialConnectionsAccountRefreshParams) (*stripe.FinancialConnectionsAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Subscribes to periodic refreshes of data associated with a Financial Connections Account. When the account status is active, data is typically refreshed once a day.
func Subscribe(id string, params *stripe.FinancialConnectionsAccountSubscribeParams) (*stripe.FinancialConnectionsAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Subscribes to periodic refreshes of data associated with a Financial Connections Account. When the account status is active, data is typically refreshed once a day.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Subscribe(id string, params *stripe.FinancialConnectionsAccountSubscribeParams) (*stripe.FinancialConnectionsAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Unsubscribes from periodic refreshes of data associated with a Financial Connections Account.
func Unsubscribe(id string, params *stripe.FinancialConnectionsAccountUnsubscribeParams) (*stripe.FinancialConnectionsAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Unsubscribes from periodic refreshes of data associated with a Financial Connections Account.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Unsubscribe(id string, params *stripe.FinancialConnectionsAccountUnsubscribeParams) (*stripe.FinancialConnectionsAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of Financial Connections Account objects.
func List(params *stripe.FinancialConnectionsAccountListParams) *Iter {
	_ = "STUB: not implemented"
	return nil

	// Returns a list of Financial Connections Account objects.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) List(listParams *stripe.FinancialConnectionsAccountListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for financial connections accounts.
type Iter struct {
	*stripe.Iter
}

// FinancialConnectionsAccount returns the financial connections account which the iterator is currently pointing to.
func (i *Iter) FinancialConnectionsAccount() *stripe.FinancialConnectionsAccount {
	_ = "STUB: not implemented"
	return nil
}

// FinancialConnectionsAccountList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) FinancialConnectionsAccountList() *stripe.FinancialConnectionsAccountList {
	_ = "STUB: not implemented"
	return nil
}

// Lists all owners for a given Account
func ListOwners(params *stripe.FinancialConnectionsAccountListOwnersParams) *OwnerIter {
	_ = "STUB: not implemented"
	return nil
}

// Lists all owners for a given Account
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ListOwners(listParams *stripe.FinancialConnectionsAccountListOwnersParams) *OwnerIter {
	_ = "STUB: not implemented"
	return nil
}

// OwnerIter is an iterator for financial connections account owners.
type OwnerIter struct {
	*stripe.Iter
}

// FinancialConnectionsAccountOwner returns the financial connections account owner which the iterator is currently pointing to.
func (i *OwnerIter) FinancialConnectionsAccountOwner() *stripe.FinancialConnectionsAccountOwner {
	_ = "STUB: not implemented"
	return nil
}

// FinancialConnectionsAccountOwnerList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *OwnerIter) FinancialConnectionsAccountOwnerList() *stripe.FinancialConnectionsAccountOwnerList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
