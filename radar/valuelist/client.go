//
//
// File generated from our OpenAPI spec
//
//

// Package valuelist provides the /v1/radar/value_lists APIs
package valuelist

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/radar/value_lists APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new ValueList object, which can then be referenced in rules.
func New(params *stripe.RadarValueListParams) (*stripe.RadarValueList, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates a new ValueList object, which can then be referenced in rules.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.RadarValueListParams) (*stripe.RadarValueList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a ValueList object.
func Get(id string, params *stripe.RadarValueListParams) (*stripe.RadarValueList, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves a ValueList object.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.RadarValueListParams) (*stripe.RadarValueList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates a ValueList object by setting the values of the parameters passed. Any parameters not provided will be left unchanged. Note that item_type is immutable.
func Update(id string, params *stripe.RadarValueListParams) (*stripe.RadarValueList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates a ValueList object by setting the values of the parameters passed. Any parameters not provided will be left unchanged. Note that item_type is immutable.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.RadarValueListParams) (*stripe.RadarValueList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deletes a ValueList object, also deleting any items contained within the value list. To be deleted, a value list must not be referenced in any rules.
func Del(id string, params *stripe.RadarValueListParams) (*stripe.RadarValueList, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Deletes a ValueList object, also deleting any items contained within the value list. To be deleted, a value list must not be referenced in any rules.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Del(id string, params *stripe.RadarValueListParams) (*stripe.RadarValueList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of ValueList objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
func List(params *stripe.RadarValueListListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Returns a list of ValueList objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.RadarValueListListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for radar value lists.
type Iter struct {
	*stripe.Iter
}

// RadarValueList returns the radar value list which the iterator is currently pointing to.
func (i *Iter) RadarValueList() *stripe.RadarValueList { _ = "STUB: not implemented"; return nil }

// RadarValueListList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) RadarValueListList() *stripe.RadarValueListList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
