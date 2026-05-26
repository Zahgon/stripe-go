//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1BankAccountService is used to invoke bankaccount related APIs.
type v1BankAccountService struct {
	B   Backend
	Key string
}

// Create creates a new bank account
func (c v1BankAccountService) Create(ctx context.Context, params *BankAccountCreateParams) (*BankAccount, error) {
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
func (c v1BankAccountService) Retrieve(ctx context.Context, id string, params *BankAccountRetrieveParams) (*BankAccount, error) {
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
func (c v1BankAccountService) Update(ctx context.Context, id string, params *BankAccountUpdateParams) (*BankAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete a specified external account for a given account.
func (c v1BankAccountService) Delete(ctx context.Context, id string, params *BankAccountDeleteParams) (*BankAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c v1BankAccountService) List(ctx context.Context, listParams *BankAccountListParams) *V1List[*BankAccount] {
	_ = "STUB: not implemented"
	return nil
}

// There's no bank accounts list URL, so we use one sources or external
// accounts. An override on BankAccountListParam's `AppendTo` will add the
// filter `object=bank_account` to make sure that only bank accounts come
// back with the response.
