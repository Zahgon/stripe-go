//
//
// File generated from our OpenAPI spec
//
//

// Package plan provides the /v1/plans APIs
package plan

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/plans APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// You can now model subscriptions more flexibly using the [Prices API](https://docs.stripe.com/api#prices). It replaces the Plans API and is backwards compatible to simplify your migration.
func New(params *stripe.PlanParams) (*stripe.Plan, error) {
	_ = "STUB: not implemented"
	return nil,

		// You can now model subscriptions more flexibly using the [Prices API](https://docs.stripe.com/api#prices). It replaces the Plans API and is backwards compatible to simplify your migration.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.PlanParams) (*stripe.Plan, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the plan with the given ID.
func Get(id string, params *stripe.PlanParams) (*stripe.Plan, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves the plan with the given ID.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.PlanParams) (*stripe.Plan, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specified plan by setting the values of the parameters passed. Any parameters not provided are left unchanged. By design, you cannot change a plan's ID, amount, currency, or billing cycle.
func Update(id string, params *stripe.PlanParams) (*stripe.Plan, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specified plan by setting the values of the parameters passed. Any parameters not provided are left unchanged. By design, you cannot change a plan's ID, amount, currency, or billing cycle.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.PlanParams) (*stripe.Plan, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deleting plans means new subscribers can't be added. Existing subscribers aren't affected.
func Del(id string, params *stripe.PlanParams) (*stripe.Plan, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Deleting plans means new subscribers can't be added. Existing subscribers aren't affected.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Del(id string, params *stripe.PlanParams) (*stripe.Plan, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of your plans.
func List(params *stripe.PlanListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Returns a list of your plans.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.PlanListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for plans.
type Iter struct {
	*stripe.Iter
}

// Plan returns the plan which the iterator is currently pointing to.
func (i *Iter) Plan() *stripe.Plan { _ = "STUB: not implemented"; return nil }

// PlanList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) PlanList() *stripe.PlanList { _ = "STUB: not implemented"; return nil }

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
