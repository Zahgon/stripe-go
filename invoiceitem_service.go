//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1InvoiceItemService is used to invoke /v1/invoiceitems APIs.
type v1InvoiceItemService struct {
	B   Backend
	Key string
}

// Creates an item to be added to a draft invoice (up to 250 items per invoice). If no invoice is specified, the item will be on the next invoice created for the customer specified.
func (c v1InvoiceItemService) Create(ctx context.Context, params *InvoiceItemCreateParams) (*InvoiceItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the invoice item with the given ID.
func (c v1InvoiceItemService) Retrieve(ctx context.Context, id string, params *InvoiceItemRetrieveParams) (*InvoiceItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the amount or description of an invoice item on an upcoming invoice. Updating an invoice item is only possible before the invoice it's attached to is closed.
func (c v1InvoiceItemService) Update(ctx context.Context, id string, params *InvoiceItemUpdateParams) (*InvoiceItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deletes an invoice item, removing it from an invoice. Deleting invoice items is only possible when they're not attached to invoices, or if it's attached to a draft invoice.
func (c v1InvoiceItemService) Delete(ctx context.Context, id string, params *InvoiceItemDeleteParams) (*InvoiceItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of your invoice items. Invoice items are returned sorted by creation date, with the most recently created invoice items appearing first.
func (c v1InvoiceItemService) List(ctx context.Context, listParams *InvoiceItemListParams) *V1List[*InvoiceItem] {
	_ = "STUB: not implemented"
	return nil
}
