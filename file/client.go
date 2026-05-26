//
//
// File generated from our OpenAPI spec
//
//

// Package file provides the /v1/files APIs
package file

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/files APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B        stripe.Backend
	BUploads stripe.Backend
	Key      string
}

// To upload a file to Stripe, you need to send a request of type multipart/form-data. Include the file you want to upload in the request, and the parameters for creating a file.
//
// All of Stripe's officially supported Client libraries support sending multipart/form-data.
func New(params *stripe.FileParams) (*stripe.File, error) {
	_ = "STUB: not implemented"
	return nil,

		// To upload a file to Stripe, you need to send a request of type multipart/form-data. Include the file you want to upload in the request, and the parameters for creating a file.
		//
		// All of Stripe's officially supported Client libraries support sending multipart/form-data.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.FileParams) (*stripe.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the details of an existing file object. After you supply a unique file ID, Stripe returns the corresponding file object. Learn how to [access file contents](https://docs.stripe.com/docs/file-upload#download-file-contents).
func Get(id string, params *stripe.FileParams) (*stripe.File, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves the details of an existing file object. After you supply a unique file ID, Stripe returns the corresponding file object. Learn how to [access file contents](https://docs.stripe.com/docs/file-upload#download-file-contents).
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.FileParams) (*stripe.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of the files that your account has access to. Stripe sorts and returns the files by their creation dates, placing the most recently created files at the top.
func List(params *stripe.FileListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Returns a list of the files that your account has access to. Stripe sorts and returns the files by their creation dates, placing the most recently created files at the top.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.FileListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for files.
type Iter struct {
	*stripe.Iter
}

// File returns the file which the iterator is currently pointing to.
func (i *Iter) File() *stripe.File { _ = "STUB: not implemented"; return nil }

// FileList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) FileList() *stripe.FileList { _ = "STUB: not implemented"; return nil }

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
