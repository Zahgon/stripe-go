//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1RadarValueListItemService is used to invoke /v1/radar/value_list_items APIs.
type v1RadarValueListItemService struct {
	B   Backend
	Key string
}

// Creates a new ValueListItem object, which is added to the specified parent value list.
func (c v1RadarValueListItemService) Create(ctx context.Context, params *RadarValueListItemCreateParams) (*RadarValueListItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a ValueListItem object.
func (c v1RadarValueListItemService) Retrieve(ctx context.Context, id string, params *RadarValueListItemRetrieveParams) (*RadarValueListItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deletes a ValueListItem object, removing it from its parent value list.
func (c v1RadarValueListItemService) Delete(ctx context.Context, id string, params *RadarValueListItemDeleteParams) (*RadarValueListItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of ValueListItem objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
func (c v1RadarValueListItemService) List(ctx context.Context, listParams *RadarValueListItemListParams) *V1List[*RadarValueListItem] {
	_ = "STUB: not implemented"
	return nil
}
