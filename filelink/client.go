//
//
// File generated from our OpenAPI spec
//
//

// Package filelink provides the /v1/file_links APIs
package filelink

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/file_links APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new file link object.
func New(params *stripe.FileLinkParams) (*stripe.FileLink, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates a new file link object.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.FileLinkParams) (*stripe.FileLink, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the file link with the given ID.
func Get(id string, params *stripe.FileLinkParams) (*stripe.FileLink, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves the file link with the given ID.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.FileLinkParams) (*stripe.FileLink, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates an existing file link object. Expired links can no longer be updated.
func Update(id string, params *stripe.FileLinkParams) (*stripe.FileLink, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates an existing file link object. Expired links can no longer be updated.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.FileLinkParams) (*stripe.FileLink, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of file links.
func List(params *stripe.FileLinkListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Returns a list of file links.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.FileLinkListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for file links.
type Iter struct {
	*stripe.Iter
}

// FileLink returns the file link which the iterator is currently pointing to.
func (i *Iter) FileLink() *stripe.FileLink { _ = "STUB: not implemented"; return nil }

// FileLinkList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) FileLinkList() *stripe.FileLinkList { _ = "STUB: not implemented"; return nil }

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
