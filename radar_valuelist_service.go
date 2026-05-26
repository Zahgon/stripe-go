//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1RadarValueListService is used to invoke /v1/radar/value_lists APIs.
type v1RadarValueListService struct {
	B   Backend
	Key string
}

// Creates a new ValueList object, which can then be referenced in rules.
func (c v1RadarValueListService) Create(ctx context.Context, params *RadarValueListCreateParams) (*RadarValueList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a ValueList object.
func (c v1RadarValueListService) Retrieve(ctx context.Context, id string, params *RadarValueListRetrieveParams) (*RadarValueList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates a ValueList object by setting the values of the parameters passed. Any parameters not provided will be left unchanged. Note that item_type is immutable.
func (c v1RadarValueListService) Update(ctx context.Context, id string, params *RadarValueListUpdateParams) (*RadarValueList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deletes a ValueList object, also deleting any items contained within the value list. To be deleted, a value list must not be referenced in any rules.
func (c v1RadarValueListService) Delete(ctx context.Context, id string, params *RadarValueListDeleteParams) (*RadarValueList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of ValueList objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
func (c v1RadarValueListService) List(ctx context.Context, listParams *RadarValueListListParams) *V1List[*RadarValueList] {
	_ = "STUB: not implemented"
	return nil
}
