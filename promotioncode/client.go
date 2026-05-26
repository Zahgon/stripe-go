//
//
// File generated from our OpenAPI spec
//
//

// Package promotioncode provides the /v1/promotion_codes APIs
package promotioncode

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/promotion_codes APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// A promotion code points to an underlying promotion. You can optionally restrict the code to a specific customer, redemption limit, and expiration date.
func New(params *stripe.PromotionCodeParams) (*stripe.PromotionCode, error) {
	_ = "STUB: not implemented"
	return nil,

		// A promotion code points to an underlying promotion. You can optionally restrict the code to a specific customer, redemption limit, and expiration date.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.PromotionCodeParams) (*stripe.PromotionCode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the promotion code with the given ID. In order to retrieve a promotion code by the customer-facing code use [list](https://docs.stripe.com/docs/api/promotion_codes/list) with the desired code.
func Get(id string, params *stripe.PromotionCodeParams) (*stripe.PromotionCode, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves the promotion code with the given ID. In order to retrieve a promotion code by the customer-facing code use [list](https://docs.stripe.com/docs/api/promotion_codes/list) with the desired code.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.PromotionCodeParams) (*stripe.PromotionCode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specified promotion code by setting the values of the parameters passed. Most fields are, by design, not editable.
func Update(id string, params *stripe.PromotionCodeParams) (*stripe.PromotionCode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specified promotion code by setting the values of the parameters passed. Most fields are, by design, not editable.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.PromotionCodeParams) (*stripe.PromotionCode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of your promotion codes.
func List(params *stripe.PromotionCodeListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Returns a list of your promotion codes.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.PromotionCodeListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for promotion codes.
type Iter struct {
	*stripe.Iter
}

// PromotionCode returns the promotion code which the iterator is currently pointing to.
func (i *Iter) PromotionCode() *stripe.PromotionCode { _ = "STUB: not implemented"; return nil }

// PromotionCodeList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) PromotionCodeList() *stripe.PromotionCodeList { _ = "STUB: not implemented"; return nil }

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
