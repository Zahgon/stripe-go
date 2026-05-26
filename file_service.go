//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1FileService is used to invoke /v1/files APIs.
type v1FileService struct {
	B        Backend
	BUploads Backend
	Key      string
}

// To upload a file to Stripe, you need to send a request of type multipart/form-data. Include the file you want to upload in the request, and the parameters for creating a file.
//
// All of Stripe's officially supported Client libraries support sending multipart/form-data.
func (c v1FileService) Create(ctx context.Context, params *FileCreateParams) (*File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the details of an existing file object. After you supply a unique file ID, Stripe returns the corresponding file object. Learn how to [access file contents](https://docs.stripe.com/docs/file-upload#download-file-contents).
func (c v1FileService) Retrieve(ctx context.Context, id string, params *FileRetrieveParams) (*File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of the files that your account has access to. Stripe sorts and returns the files by their creation dates, placing the most recently created files at the top.
func (c v1FileService) List(ctx context.Context, listParams *FileListParams) *V1List[*File] {
	_ = "STUB: not implemented"
	return nil
}
