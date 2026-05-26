//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1PlanService is used to invoke /v1/plans APIs.
type v1PlanService struct {
	B   Backend
	Key string
}

// You can now model subscriptions more flexibly using the [Prices API](https://docs.stripe.com/api#prices). It replaces the Plans API and is backwards compatible to simplify your migration.
func (c v1PlanService) Create(ctx context.Context, params *PlanCreateParams) (*Plan, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the plan with the given ID.
func (c v1PlanService) Retrieve(ctx context.Context, id string, params *PlanRetrieveParams) (*Plan, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specified plan by setting the values of the parameters passed. Any parameters not provided are left unchanged. By design, you cannot change a plan's ID, amount, currency, or billing cycle.
func (c v1PlanService) Update(ctx context.Context, id string, params *PlanUpdateParams) (*Plan, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deleting plans means new subscribers can't be added. Existing subscribers aren't affected.
func (c v1PlanService) Delete(ctx context.Context, id string, params *PlanDeleteParams) (*Plan, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of your plans.
func (c v1PlanService) List(ctx context.Context, listParams *PlanListParams) *V1List[*Plan] {
	_ = "STUB: not implemented"
	return nil
}
