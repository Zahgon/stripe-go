//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1SubscriptionItemService is used to invoke /v1/subscription_items APIs.
type v1SubscriptionItemService struct {
	B   Backend
	Key string
}

// Adds a new item to an existing subscription. No existing items will be changed or replaced.
func (c v1SubscriptionItemService) Create(ctx context.Context, params *SubscriptionItemCreateParams) (*SubscriptionItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the subscription item with the given ID.
func (c v1SubscriptionItemService) Retrieve(ctx context.Context, id string, params *SubscriptionItemRetrieveParams) (*SubscriptionItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the plan or quantity of an item on a current subscription.
func (c v1SubscriptionItemService) Update(ctx context.Context, id string, params *SubscriptionItemUpdateParams) (*SubscriptionItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deletes an item from the subscription. Removing a subscription item from a subscription will not cancel the subscription.
func (c v1SubscriptionItemService) Delete(ctx context.Context, id string, params *SubscriptionItemDeleteParams) (*SubscriptionItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of your subscription items for a given subscription.
func (c v1SubscriptionItemService) List(ctx context.Context, listParams *SubscriptionItemListParams) *V1List[*SubscriptionItem] {
	_ = "STUB: not implemented"
	return nil
}
