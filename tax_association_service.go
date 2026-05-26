//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1TaxAssociationService is used to invoke association related APIs.
type v1TaxAssociationService struct {
	B   Backend
	Key string
}

// Finds a tax association object by PaymentIntent id.
func (c v1TaxAssociationService) Find(ctx context.Context, params *TaxAssociationFindParams) (*TaxAssociation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
