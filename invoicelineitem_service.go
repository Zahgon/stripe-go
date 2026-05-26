//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1InvoiceLineItemService is used to invoke /v1/invoices/{invoice}/lines APIs.
type v1InvoiceLineItemService struct {
	B   Backend
	Key string
}

// Updates an invoice's line item. Some fields, such as tax_amounts, only live on the invoice line item,
// so they can only be updated through this endpoint. Other fields, such as amount, live on both the invoice
// item and the invoice line item, so updates on this endpoint will propagate to the invoice item as well.
// Updating an invoice's line item is only possible before the invoice is finalized.
func (c v1InvoiceLineItemService) Update(ctx context.Context, id string, params *InvoiceLineItemUpdateParams) (*InvoiceLineItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
