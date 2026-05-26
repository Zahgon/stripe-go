//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v2CoreAccountService is used to invoke account related APIs.
type v2CoreAccountService struct {
	B   Backend
	Key string
}

// An Account is a representation of a company, individual or other entity that a user interacts with. Accounts contain identifying information about the entity, and configurations that store the features an account has access to. An account can be configured as any or all of the following configurations: Customer, Merchant and/or Recipient.
func (c v2CoreAccountService) Create(ctx context.Context, params *V2CoreAccountCreateParams) (*V2CoreAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the details of an Account.
func (c v2CoreAccountService) Retrieve(ctx context.Context, id string, params *V2CoreAccountRetrieveParams) (*V2CoreAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the details of an Account.
func (c v2CoreAccountService) Update(ctx context.Context, id string, params *V2CoreAccountUpdateParams) (*V2CoreAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Removes access to the Account and its associated resources. Closed Accounts can no longer be operated on, but limited information can still be retrieved through the API in order to be able to track their history.
func (c v2CoreAccountService) Close(ctx context.Context, id string, params *V2CoreAccountCloseParams) (*V2CoreAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of Accounts.
func (c v2CoreAccountService) List(ctx context.Context, listParams *V2CoreAccountListParams) *V2List[*V2CoreAccount] {
	_ = "STUB: not implemented"
	return nil
}
