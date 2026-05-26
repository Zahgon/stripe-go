//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1CouponService is used to invoke /v1/coupons APIs.
type v1CouponService struct {
	B   Backend
	Key string
}

// You can create coupons easily via the [coupon management](https://dashboard.stripe.com/coupons) page of the Stripe dashboard. Coupon creation is also accessible via the API if you need to create coupons on the fly.
//
// A coupon has either a percent_off or an amount_off and currency. If you set an amount_off, that amount will be subtracted from any invoice's subtotal. For example, an invoice with a subtotal of 100 will have a final total of 0 if a coupon with an amount_off of 200 is applied to it and an invoice with a subtotal of 300 will have a final total of 100 if a coupon with an amount_off of 200 is applied to it.
func (c v1CouponService) Create(ctx context.Context, params *CouponCreateParams) (*Coupon, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the coupon with the given ID.
func (c v1CouponService) Retrieve(ctx context.Context, id string, params *CouponRetrieveParams) (*Coupon, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the metadata of a coupon. Other coupon details (currency, duration, amount_off) are, by design, not editable.
func (c v1CouponService) Update(ctx context.Context, id string, params *CouponUpdateParams) (*Coupon, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// You can delete coupons via the [coupon management](https://dashboard.stripe.com/coupons) page of the Stripe dashboard. However, deleting a coupon does not affect any customers who have already applied the coupon; it means that new customers can't redeem the coupon. You can also delete coupons via the API.
func (c v1CouponService) Delete(ctx context.Context, id string, params *CouponDeleteParams) (*Coupon, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of your coupons.
func (c v1CouponService) List(ctx context.Context, listParams *CouponListParams) *V1List[*Coupon] {
	_ = "STUB: not implemented"
	return nil
}
