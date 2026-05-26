//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1CardService is used to invoke card related APIs.
type v1CardService struct {
	B   Backend
	Key string
}

// Create creates a new card
func (c v1CardService) Create(ctx context.Context, params *CardCreateParams) (*Card, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Note that we call this special append method instead of the standard one
// from the form package. We should not use form's because doing so will
// include some parameters that are undesirable here.

// Because card creation uses the custom append above, we have to
// make an explicit call using a form and CallRaw instead of the standard
// Call (which takes a set of parameters).

// Get returns the details of a card.
func (c v1CardService) Retrieve(ctx context.Context, id string, params *CardRetrieveParams) (*Card, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update a specified source for a given customer.
func (c v1CardService) Update(ctx context.Context, id string, params *CardUpdateParams) (*Card, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete a specified source for a given customer.
func (c v1CardService) Delete(ctx context.Context, id string, params *CardDeleteParams) (*Card, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c v1CardService) List(ctx context.Context, listParams *CardListParams) *V1List[*Card] {
	_ = "STUB: not implemented"
	return nil
}

// There's no cards list URL, so we use one sources or external
// accounts. An override on CardListParam's `AppendTo` will add the
// filter `object=card` to make sure that only cards come
// back with the response.
