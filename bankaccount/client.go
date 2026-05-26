//
//
// File generated from our OpenAPI spec
//
//

// Package bankaccount provides the bankaccount related APIs
package bankaccount

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke bankaccount related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create creates a new bank account
func New(params *stripe.BankAccountParams) (*stripe.BankAccount, error) {
	_ = "STUB: not implemented"
	return nil,

		// Create creates a new bank account
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.BankAccountParams) (*stripe.BankAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Note that we call this special append method instead of the standard one
// from the form package. We should not use form's because doing so will
// include some parameters that are undesirable here.

// Because bank account creation uses the custom append above, we have to
// make an explicit call using a form and CallRaw instead of the standard
// Call (which takes a set of parameters).

// Get returns the details of a bank account.
func Get(id string, params *stripe.BankAccountParams) (*stripe.BankAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Get returns the details of a bank account.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.BankAccountParams) (*stripe.BankAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the metadata, account holder name, account holder type of a bank account belonging to
// a connected account and optionally sets it as the default for its currency. Other bank account
// details are not editable by design.
//
// You can only update bank accounts when [account.controller.requirement_collection is application, which includes <a href="/connect/custom-accounts">Custom accounts](https://docs.stripe.com/api/accounts/object#account_object-controller-requirement_collection).
//
// You can re-enable a disabled bank account by performing an update call without providing any
// arguments or changes.
func Update(id string, params *stripe.BankAccountParams) (*stripe.BankAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the metadata, account holder name, account holder type of a bank account belonging to
// a connected account and optionally sets it as the default for its currency. Other bank account
// details are not editable by design.
//
// You can only update bank accounts when [account.controller.requirement_collection is application, which includes <a href="/connect/custom-accounts">Custom accounts](https://docs.stripe.com/api/accounts/object#account_object-controller-requirement_collection).
//
// You can re-enable a disabled bank account by performing an update call without providing any
// arguments or changes.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.BankAccountParams) (*stripe.BankAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete a specified external account for a given account.
func Del(id string, params *stripe.BankAccountParams) (*stripe.BankAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Delete a specified external account for a given account.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Del(id string, params *stripe.BankAccountParams) (*stripe.BankAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func List(params *stripe.BankAccountListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.BankAccountListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// There's no bank accounts list URL, so we use one sources or external
// accounts. An override on BankAccountListParam's `AppendTo` will add the
// filter `object=bank_account` to make sure that only bank accounts come
// back with the response.

// Iter is an iterator for bank accounts.
type Iter struct {
	*stripe.Iter
}

// BankAccount returns the bank account which the iterator is currently pointing to.
func (i *Iter) BankAccount() *stripe.BankAccount { _ = "STUB: not implemented"; return nil }

// BankAccountList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) BankAccountList() *stripe.BankAccountList { _ = "STUB: not implemented"; return nil }

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
