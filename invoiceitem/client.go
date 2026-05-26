//
//
// File generated from our OpenAPI spec
//
//

// Package invoiceitem provides the /v1/invoiceitems APIs
package invoiceitem

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/invoiceitems APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates an item to be added to a draft invoice (up to 250 items per invoice). If no invoice is specified, the item will be on the next invoice created for the customer specified.
func New(params *stripe.InvoiceItemParams) (*stripe.InvoiceItem, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates an item to be added to a draft invoice (up to 250 items per invoice). If no invoice is specified, the item will be on the next invoice created for the customer specified.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.InvoiceItemParams) (*stripe.InvoiceItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the invoice item with the given ID.
func Get(id string, params *stripe.InvoiceItemParams) (*stripe.InvoiceItem, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves the invoice item with the given ID.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.InvoiceItemParams) (*stripe.InvoiceItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the amount or description of an invoice item on an upcoming invoice. Updating an invoice item is only possible before the invoice it's attached to is closed.
func Update(id string, params *stripe.InvoiceItemParams) (*stripe.InvoiceItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the amount or description of an invoice item on an upcoming invoice. Updating an invoice item is only possible before the invoice it's attached to is closed.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.InvoiceItemParams) (*stripe.InvoiceItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deletes an invoice item, removing it from an invoice. Deleting invoice items is only possible when they're not attached to invoices, or if it's attached to a draft invoice.
func Del(id string, params *stripe.InvoiceItemParams) (*stripe.InvoiceItem, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Deletes an invoice item, removing it from an invoice. Deleting invoice items is only possible when they're not attached to invoices, or if it's attached to a draft invoice.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Del(id string, params *stripe.InvoiceItemParams) (*stripe.InvoiceItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of your invoice items. Invoice items are returned sorted by creation date, with the most recently created invoice items appearing first.
func List(params *stripe.InvoiceItemListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Returns a list of your invoice items. Invoice items are returned sorted by creation date, with the most recently created invoice items appearing first.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.InvoiceItemListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for invoice items.
type Iter struct {
	*stripe.Iter
}

// InvoiceItem returns the invoice item which the iterator is currently pointing to.
func (i *Iter) InvoiceItem() *stripe.InvoiceItem { _ = "STUB: not implemented"; return nil }

// InvoiceItemList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) InvoiceItemList() *stripe.InvoiceItemList { _ = "STUB: not implemented"; return nil }

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
