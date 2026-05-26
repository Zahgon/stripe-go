//
//
// File generated from our OpenAPI spec
//
//

// Package coupon provides the /v1/coupons APIs
package coupon

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/coupons APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// You can create coupons easily via the [coupon management](https://dashboard.stripe.com/coupons) page of the Stripe dashboard. Coupon creation is also accessible via the API if you need to create coupons on the fly.
//
// A coupon has either a percent_off or an amount_off and currency. If you set an amount_off, that amount will be subtracted from any invoice's subtotal. For example, an invoice with a subtotal of 100 will have a final total of 0 if a coupon with an amount_off of 200 is applied to it and an invoice with a subtotal of 300 will have a final total of 100 if a coupon with an amount_off of 200 is applied to it.
func New(params *stripe.CouponParams) (*stripe.Coupon, error) {
	_ = "STUB: not implemented"
	return nil,

		// You can create coupons easily via the [coupon management](https://dashboard.stripe.com/coupons) page of the Stripe dashboard. Coupon creation is also accessible via the API if you need to create coupons on the fly.
		//
		// A coupon has either a percent_off or an amount_off and currency. If you set an amount_off, that amount will be subtracted from any invoice's subtotal. For example, an invoice with a subtotal of 100 will have a final total of 0 if a coupon with an amount_off of 200 is applied to it and an invoice with a subtotal of 300 will have a final total of 100 if a coupon with an amount_off of 200 is applied to it.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.CouponParams) (*stripe.Coupon, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the coupon with the given ID.
func Get(id string, params *stripe.CouponParams) (*stripe.Coupon, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves the coupon with the given ID.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.CouponParams) (*stripe.Coupon, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the metadata of a coupon. Other coupon details (currency, duration, amount_off) are, by design, not editable.
func Update(id string, params *stripe.CouponParams) (*stripe.Coupon, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the metadata of a coupon. Other coupon details (currency, duration, amount_off) are, by design, not editable.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.CouponParams) (*stripe.Coupon, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// You can delete coupons via the [coupon management](https://dashboard.stripe.com/coupons) page of the Stripe dashboard. However, deleting a coupon does not affect any customers who have already applied the coupon; it means that new customers can't redeem the coupon. You can also delete coupons via the API.
func Del(id string, params *stripe.CouponParams) (*stripe.Coupon, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// You can delete coupons via the [coupon management](https://dashboard.stripe.com/coupons) page of the Stripe dashboard. However, deleting a coupon does not affect any customers who have already applied the coupon; it means that new customers can't redeem the coupon. You can also delete coupons via the API.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Del(id string, params *stripe.CouponParams) (*stripe.Coupon, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of your coupons.
func List(params *stripe.CouponListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Returns a list of your coupons.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.CouponListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for coupons.
type Iter struct {
	*stripe.Iter
}

// Coupon returns the coupon which the iterator is currently pointing to.
func (i *Iter) Coupon() *stripe.Coupon { _ = "STUB: not implemented"; return nil }

// CouponList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) CouponList() *stripe.CouponList { _ = "STUB: not implemented"; return nil }

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
