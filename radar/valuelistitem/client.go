//
//
// File generated from our OpenAPI spec
//
//

// Package valuelistitem provides the /v1/radar/value_list_items APIs
// For more details, see: https://stripe.com/docs/api/radar/list_items?lang=go
package valuelistitem

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/radar/value_list_items APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new ValueListItem object, which is added to the specified parent value list.
func New(params *stripe.RadarValueListItemParams) (*stripe.RadarValueListItem, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates a new ValueListItem object, which is added to the specified parent value list.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.RadarValueListItemParams) (*stripe.RadarValueListItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a ValueListItem object.
func Get(id string, params *stripe.RadarValueListItemParams) (*stripe.RadarValueListItem, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves a ValueListItem object.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.RadarValueListItemParams) (*stripe.RadarValueListItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deletes a ValueListItem object, removing it from its parent value list.
func Del(id string, params *stripe.RadarValueListItemParams) (*stripe.RadarValueListItem, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Deletes a ValueListItem object, removing it from its parent value list.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Del(id string, params *stripe.RadarValueListItemParams) (*stripe.RadarValueListItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of ValueListItem objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
func List(params *stripe.RadarValueListItemListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Returns a list of ValueListItem objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.RadarValueListItemListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for radar value list items.
type Iter struct {
	*stripe.Iter
}

// RadarValueListItem returns the radar value list item which the iterator is currently pointing to.
func (i *Iter) RadarValueListItem() *stripe.RadarValueListItem {
	_ = "STUB: not implemented"
	return nil
}

// RadarValueListItemList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) RadarValueListItemList() *stripe.RadarValueListItemList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
