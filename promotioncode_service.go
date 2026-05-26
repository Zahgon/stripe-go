//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1PromotionCodeService is used to invoke /v1/promotion_codes APIs.
type v1PromotionCodeService struct {
	B   Backend
	Key string
}

// A promotion code points to an underlying promotion. You can optionally restrict the code to a specific customer, redemption limit, and expiration date.
func (c v1PromotionCodeService) Create(ctx context.Context, params *PromotionCodeCreateParams) (*PromotionCode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the promotion code with the given ID. In order to retrieve a promotion code by the customer-facing code use [list](https://docs.stripe.com/docs/api/promotion_codes/list) with the desired code.
func (c v1PromotionCodeService) Retrieve(ctx context.Context, id string, params *PromotionCodeRetrieveParams) (*PromotionCode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specified promotion code by setting the values of the parameters passed. Most fields are, by design, not editable.
func (c v1PromotionCodeService) Update(ctx context.Context, id string, params *PromotionCodeUpdateParams) (*PromotionCode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of your promotion codes.
func (c v1PromotionCodeService) List(ctx context.Context, listParams *PromotionCodeListParams) *V1List[*PromotionCode] {
	_ = "STUB: not implemented"
	return nil
}
