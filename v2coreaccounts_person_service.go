//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v2CoreAccountsPersonService is used to invoke person related APIs.
type v2CoreAccountsPersonService struct {
	B   Backend
	Key string
}

// Create a Person. Adds an individual to an Account's identity. You can set relationship attributes and identity information at creation.
func (c v2CoreAccountsPersonService) Create(ctx context.Context, params *V2CoreAccountsPersonCreateParams) (*V2CoreAccountPerson, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a Person associated with an Account.
func (c v2CoreAccountsPersonService) Retrieve(ctx context.Context, id string, params *V2CoreAccountsPersonRetrieveParams) (*V2CoreAccountPerson, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates a Person associated with an Account.
func (c v2CoreAccountsPersonService) Update(ctx context.Context, id string, params *V2CoreAccountsPersonUpdateParams) (*V2CoreAccountPerson, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete a Person associated with an Account.
func (c v2CoreAccountsPersonService) Delete(ctx context.Context, id string, params *V2CoreAccountsPersonDeleteParams) (*V2DeletedObject, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a paginated list of Persons associated with an Account.
func (c v2CoreAccountsPersonService) List(ctx context.Context, listParams *V2CoreAccountsPersonListParams) *V2List[*V2CoreAccountPerson] {
	_ = "STUB: not implemented"
	return nil
}
