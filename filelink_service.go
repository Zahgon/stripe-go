//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1FileLinkService is used to invoke /v1/file_links APIs.
type v1FileLinkService struct {
	B   Backend
	Key string
}

// Creates a new file link object.
func (c v1FileLinkService) Create(ctx context.Context, params *FileLinkCreateParams) (*FileLink, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the file link with the given ID.
func (c v1FileLinkService) Retrieve(ctx context.Context, id string, params *FileLinkRetrieveParams) (*FileLink, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates an existing file link object. Expired links can no longer be updated.
func (c v1FileLinkService) Update(ctx context.Context, id string, params *FileLinkUpdateParams) (*FileLink, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of file links.
func (c v1FileLinkService) List(ctx context.Context, listParams *FileLinkListParams) *V1List[*FileLink] {
	_ = "STUB: not implemented"
	return nil
}
