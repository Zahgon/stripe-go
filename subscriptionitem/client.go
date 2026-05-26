//
//
// File generated from our OpenAPI spec
//
//

// Package subscriptionitem provides the /v1/subscription_items APIs
package subscriptionitem

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/subscription_items APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Adds a new item to an existing subscription. No existing items will be changed or replaced.
func New(params *stripe.SubscriptionItemParams) (*stripe.SubscriptionItem, error) {
	_ = "STUB: not implemented"
	return nil,

		// Adds a new item to an existing subscription. No existing items will be changed or replaced.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.SubscriptionItemParams) (*stripe.SubscriptionItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the subscription item with the given ID.
func Get(id string, params *stripe.SubscriptionItemParams) (*stripe.SubscriptionItem, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves the subscription item with the given ID.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.SubscriptionItemParams) (*stripe.SubscriptionItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the plan or quantity of an item on a current subscription.
func Update(id string, params *stripe.SubscriptionItemParams) (*stripe.SubscriptionItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the plan or quantity of an item on a current subscription.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.SubscriptionItemParams) (*stripe.SubscriptionItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deletes an item from the subscription. Removing a subscription item from a subscription will not cancel the subscription.
func Del(id string, params *stripe.SubscriptionItemParams) (*stripe.SubscriptionItem, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Deletes an item from the subscription. Removing a subscription item from a subscription will not cancel the subscription.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Del(id string, params *stripe.SubscriptionItemParams) (*stripe.SubscriptionItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of your subscription items for a given subscription.
func List(params *stripe.SubscriptionItemListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Returns a list of your subscription items for a given subscription.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.SubscriptionItemListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for subscription items.
type Iter struct {
	*stripe.Iter
}

// SubscriptionItem returns the subscription item which the iterator is currently pointing to.
func (i *Iter) SubscriptionItem() *stripe.SubscriptionItem { _ = "STUB: not implemented"; return nil }

// SubscriptionItemList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) SubscriptionItemList() *stripe.SubscriptionItemList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
