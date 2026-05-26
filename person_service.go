//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1PersonService is used to invoke /v1/accounts/{account}/persons APIs.
type v1PersonService struct {
	B   Backend
	Key string
}

// Creates a new person.
func (c v1PersonService) Create(ctx context.Context, params *PersonCreateParams) (*Person, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves an existing person.
func (c v1PersonService) Retrieve(ctx context.Context, id string, params *PersonRetrieveParams) (*Person, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates an existing person.
func (c v1PersonService) Update(ctx context.Context, id string, params *PersonUpdateParams) (*Person, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deletes an existing person's relationship to the account's legal entity. Any person with a relationship for an account can be deleted through the API, except if the person is the account_opener. If your integration is using the executive parameter, you cannot delete the only verified executive on file.
func (c v1PersonService) Delete(ctx context.Context, id string, params *PersonDeleteParams) (*Person, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of people associated with the account's legal entity. The people are returned sorted by creation date, with the most recent people appearing first.
func (c v1PersonService) List(ctx context.Context, listParams *PersonListParams) *V1List[*Person] {
	_ = "STUB: not implemented"
	return nil
}
