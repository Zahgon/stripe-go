//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1SubscriptionScheduleService is used to invoke /v1/subscription_schedules APIs.
type v1SubscriptionScheduleService struct {
	B   Backend
	Key string
}

// Creates a new subscription schedule object. Each customer can have up to 500 active or scheduled subscriptions.
func (c v1SubscriptionScheduleService) Create(ctx context.Context, params *SubscriptionScheduleCreateParams) (*SubscriptionSchedule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the details of an existing subscription schedule. You only need to supply the unique subscription schedule identifier that was returned upon subscription schedule creation.
func (c v1SubscriptionScheduleService) Retrieve(ctx context.Context, id string, params *SubscriptionScheduleRetrieveParams) (*SubscriptionSchedule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates an existing subscription schedule.
func (c v1SubscriptionScheduleService) Update(ctx context.Context, id string, params *SubscriptionScheduleUpdateParams) (*SubscriptionSchedule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Cancels a subscription schedule and its associated subscription immediately (if the subscription schedule has an active subscription). A subscription schedule can only be canceled if its status is not_started or active.
func (c v1SubscriptionScheduleService) Cancel(ctx context.Context, id string, params *SubscriptionScheduleCancelParams) (*SubscriptionSchedule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Releases the subscription schedule immediately, which will stop scheduling of its phases, but leave any existing subscription in place. A schedule can only be released if its status is not_started or active. If the subscription schedule is currently associated with a subscription, releasing it will remove its subscription property and set the subscription's ID to the released_subscription property.
func (c v1SubscriptionScheduleService) Release(ctx context.Context, id string, params *SubscriptionScheduleReleaseParams) (*SubscriptionSchedule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the list of your subscription schedules.
func (c v1SubscriptionScheduleService) List(ctx context.Context, listParams *SubscriptionScheduleListParams) *V1List[*SubscriptionSchedule] {
	_ = "STUB: not implemented"
	return nil
}
