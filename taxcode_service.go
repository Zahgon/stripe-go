//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1TaxCodeService is used to invoke /v1/tax_codes APIs.
type v1TaxCodeService struct {
	B   Backend
	Key string
}

// Retrieves the details of an existing tax code. Supply the unique tax code ID and Stripe will return the corresponding tax code information.
func (c v1TaxCodeService) Retrieve(ctx context.Context, id string, params *TaxCodeRetrieveParams) (*TaxCode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// A list of [all tax codes available](https://stripe.com/docs/tax/tax-categories) to add to Products in order to allow specific tax calculations.
func (c v1TaxCodeService) List(ctx context.Context, listParams *TaxCodeListParams) *V1List[*TaxCode] {
	_ = "STUB: not implemented"
	return nil
}
